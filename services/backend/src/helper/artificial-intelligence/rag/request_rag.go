package rag

import (
	"context"
	"errors"
	"fmt"
	"strings"

	artificialintelligence "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func RequestRAG(provider *artificialintelligence.AIProvider, prompt string, pgDatabasename string, messages []models.ConversationMessage, streamFunc func(ctx context.Context, reasoningChunk []byte, chunk []byte) error) (*llms.ContentResponse, error) {
	if provider == nil {
		return nil, errors.New("ai provider is nil")
	}

	requestModel, _, requestModelErr := provider.DefaultRequestModelInstance()
	if requestModelErr != nil {
		return nil, requestModelErr
	}
	if requestModel == nil {
		return nil, errors.New("default request model is nil")
	}
	if provider.OllamaEmbedder == nil {
		return nil, errors.New("default embedding provider is nil")
	}

	embedder, embedderErr := embeddings.NewEmbedder(
		provider.OllamaEmbedder,
	)
	if embedderErr != nil {
		return nil, embedderErr
	}

	var dbMetaDataFilter map[string]any

	pgvectorStore, pgvectorStoreErr := CreatePGVectorConnection(
		pgvector.WithEmbeddingTableName(pgDatabasename),
		pgvector.WithEmbedder(embedder),
		pgvector.WithCollectionMetadata(dbMetaDataFilter),
	)
	if pgvectorStoreErr != nil {
		return nil, pgvectorStoreErr
	}

	var maxDocuments int = 10
	schemaDocuments, schemaDocumentsErr := pgvectorStore.SimilaritySearch(context.Background(), prompt, maxDocuments)
	if schemaDocumentsErr != nil {
		return nil, schemaDocumentsErr
	}

	_ = schemaDocuments

	var messageContent []llms.MessageContent = []llms.MessageContent{}

	for _, message := range messages {
		messageContent = append(messageContent, llms.MessageContent{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextPart(message.Message),
			},
		})
	}

	// messageContent = append(messageContent, llms.MessageContent{
	// 	Role: llms.ChatMessageTypeHuman,
	// 	Parts: []llms.ContentPart{
	// 		llms.TextPart(buildRagPrompt(schemaDocuments, prompt)),
	// 	},
	// })

	messageContent = append(messageContent, llms.MessageContent{
		Role: llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{
			llms.TextPart(prompt),
		},
	})

	response, responseErr := requestModel.GenerateContent(
		context.Background(),
		messageContent,
		llms.WithTemperature(1),
		llms.WithModel(provider.DefaultRequestModelName()),
		llms.WithInterleaveThinking(true),
		llms.WithMaxTokens(8000),
		llms.WithStreamingReasoningFunc(streamFunc),
		llms.WithStreamThinking(true),
		llms.WithReturnThinking(true),
		llms.WithThinkingMode(llms.ThinkingModeMedium),
	)

	return response, responseErr
}

func buildRagPrompt(schemaDocuments []schema.Document, prompt string) string {
	var ctxBuilder strings.Builder
	ctxBuilder.WriteString("Du bist ein Assistent im BAföG Genehmigungsverfahren. Du antwortest NUR mit Daten aus deinem Context!")
	ctxBuilder.WriteString("Antworte prägnant, nutze lieber den Kontakt als eigene Trainingsdaten und liefere möglichst viele Quellenangaben. ")
	ctxBuilder.WriteString("Nutze den folgenden Kontext um die Frage zu beantworten.")
	ctxBuilder.WriteString("\n\n")
	ctxBuilder.WriteString("Kontext:\n\n")

	// godump.Dump(schemaDocuments)

	/*
		todo lieber in html mäßige xml chunks unterteilen, damit die quellenangaben besser erkannt werden können
	*/
	for _, r := range schemaDocuments {
		documentname := ""
		documentpage := r.Metadata["page"]
		ctxBuilder.WriteString("<chunk>\n")
		ctxBuilder.WriteString(fmt.Sprintf("<documentname name=\"%s\" \\>\n", documentname))
		ctxBuilder.WriteString(fmt.Sprintf("<documentpage page=\"%d\" \\>\n", documentpage))
		ctxBuilder.WriteString("<content>\n")
		ctxBuilder.WriteString(r.PageContent)
		ctxBuilder.WriteString("</content>\n")
		ctxBuilder.WriteString("</chunk>\n\n")
	}

	ctxBuilder.WriteString("Frage: \n\n")
	ctxBuilder.WriteString(prompt)
	ctxBuilder.WriteString("\n\n")
	ctxBuilder.WriteString("Antworte möglichst genau! Sollte die Antwort nicht im Kontext stehen, antworte mit 'Dazu liegen mir derzeit keine Informationen im Kontext vor'.\n")

	return ctxBuilder.String()
}
