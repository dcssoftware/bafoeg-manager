package rag

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	ai "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/dcssoftware/bafoeg-manager/src/helper/encoding"
	"github.com/dcssoftware/bafoeg-manager/src/helper/files/pdf"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

type MetaDataEntry struct {
	Key   string
	Value any
}

func CreateVectorsPDF(pdfFile []byte, pgDatabasename string, metadata []MetaDataEntry) error {

	// ollamaLLMEmbedding, ollamaLLMEmbeddingErr := CreateOllamaConnection(
	// 	ollama.WithModel(configuration.OllamaAPI.EmbeddingModelname),
	// )
	// if ollamaLLMEmbeddingErr != nil {
	// 	return ollamaLLMEmbeddingErr
	// }

	ollamaLLMRequesting, ollamaLLMRequestingErr := ai.CreateOllamaConnection(
		ollama.WithModel(configuration.OllamaAPI.RequestingModelname),
	)
	if ollamaLLMRequestingErr != nil {
		return ollamaLLMRequestingErr
	}

	// Build document chunks from the PDF
	docs, err := BuildDocumentsFromPDF(
		ollamaLLMRequesting,
		pdfFile,
	)
	if err != nil {
		return err
	}

	// add custom meta data to vector
	for _, doc := range docs {
		for _, data := range metadata {
			doc.Metadata[data.Key] = data.Value
		}
	}

	// Store the embeddings for the chunks
	_, err = StoreEmbeddings(docs, pgDatabasename)
	if err != nil {
		return err
	}

	return nil
}

// BuildDocumentsFromPDF loads, splits, and sanitizes the PDF into documents ready for embedding.
func BuildDocumentsFromPDF(requestingOllamaLLM *ollama.LLM, pdfFile []byte) ([]schema.Document, error) {

	var pdfFileReader io.ReaderAt
	pdfFileReader = bytes.NewReader(pdfFile)

	// Compute the PDF size in bytes
	size, sizeErr := pdf.GetPDFSize(pdfFileReader)
	if sizeErr != nil {
		return nil, sizeErr
	}

	pdfDocumentLoad := documentloaders.NewPDF(pdfFileReader, size)

	documentPages, documentPagesErr := pdf.ReadPlainText(pdfFileReader)
	if documentPagesErr != nil {
		return nil, documentPagesErr
	}

	var rawDocumentContentPages strings.Builder
	for _, page := range documentPages {
		rawDocumentContentPages.WriteString(page.PageContent + "\n\n")
	}

	// Use conservative chunking to stay within common embedding model limits
	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(configuration.OllamaAPI.TextSplitterChunkSize),
		textsplitter.WithChunkOverlap(configuration.OllamaAPI.TextSplitterChunkSize),
	)

	pdfSchema, pdfSchemaErr := pdfDocumentLoad.LoadAndSplit(context.Background(), splitter)
	if pdfSchemaErr != nil {
		return nil, pdfSchemaErr
	}

	filtered := encoding.SanitizeDocuments(pdfSchema)
	if len(filtered) == 0 {
		return nil, fmt.Errorf("no non-empty text chunks found in PDF; cannot create embeddings")
	}

	// enrich contextual retrieval by adding context to that chunk
	var enrichedChunks []schema.Document

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, configuration.Conjobs.RagVectorProcessor.ParallelJobs)

	for _, chunk := range filtered {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(c schema.Document) {
			defer wg.Done()
			defer func() { <-semaphore }()

			newChunk := CreateChunkContextFromDocument(requestingOllamaLLM, rawDocumentContentPages.String(), c)
			enrichedChunks = append(enrichedChunks, newChunk)
		}(chunk)
	}
	wg.Wait()

	return enrichedChunks, nil
}

// StoreEmbeddings connects to Ollama + pgvector and persists the provided documents as embeddings.
func StoreEmbeddings(docs []schema.Document, pgDatabasename string) ([]string, error) {
	ollamaLLM, ollamaLLMErr := ai.CreateOllamaConnection(
		ollama.WithModel(configuration.OllamaAPI.EmbeddingModelname),
	)
	if ollamaLLMErr != nil {
		return nil, ollamaLLMErr
	}

	embedder, embedderErr := embeddings.NewEmbedder(ollamaLLM)
	if embedderErr != nil {
		return nil, embedderErr
	}

	pgvectorStore, pgvectorStoreErr := CreatePGVectorConnection(
		pgvector.WithEmbeddingTableName(pgDatabasename),
		pgvector.WithEmbedder(embedder),
	)
	if pgvectorStoreErr != nil {
		return nil, pgvectorStoreErr
	}

	return pgvectorStore.AddDocuments(
		context.Background(),
		docs,
		vectorstores.WithEmbedder(embedder),
	)
}

func CreateChunkContextFromDocument(requestingOllamaLLM *ollama.LLM, document string, chunk schema.Document) schema.Document {

	prompt := fmt.Sprintf(`
<document> 
	%s
</document> 
Hier ist ein unterteilter Chunk aus dem Dokument, den wir mit Kontext erweitern möchten:
<chunk> 
	%s
</chunk> 
Bitte generiere einen kurzen, prägnanten Kontext, um diesen Abschnitt innerhalb des Gesamtdokuments mit Themenbezug einzuordnen, 
falls du dir unsicher bist, versuche es erst auf BAföG zu beziehen oder antworte mit "", 
um so die Suche nach diesem Abschnitt zu verbessern. 
Beantworte die Frage nur mit dem prägnanten Kontext und nichts anderem. 
`,
		document,
		chunk.PageContent,
	)
	_ = prompt

	extraContext, extraContextErr := requestingOllamaLLM.Call(
		context.Background(),
		prompt,
		llms.WithTemperature(0),
		// llms.WithThinking(llms.DefaultThinkingConfig()),
	)
	if extraContextErr != nil {
		logger.Error(
			nil,
			"cannot create context to chunk",
			extraContextErr,
			"",
		)
		return chunk
	}

	newPageContent := fmt.Sprintf("Kontext: \n%s\n\n Inhalt: \n%s", extraContext, chunk.PageContent)
	chunk.PageContent = newPageContent

	return chunk
}
