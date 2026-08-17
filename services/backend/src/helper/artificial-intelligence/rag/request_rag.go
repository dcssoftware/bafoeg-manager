package rag

import (
	"context"
	"fmt"
	"strings"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	ai "github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag/models"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func RequestRAG(ctx context.Context, prompt string, pgDatabasename string, messages []models.ConversationMessage, streamfunc func(ctx context.Context, chunk []byte) error) (response string, err error) {

	ollamaLLMEmbedding, ollamaLLMEmbeddingErr := ai.CreateOllamaConnection(
		ollama.WithModel(configuration.OllamaAPI.EmbeddingModelname),
	)
	if ollamaLLMEmbeddingErr != nil {
		return "", ollamaLLMEmbeddingErr
	}

	ollamaLLMRequesting, ollamaLLMRequestingErr := ai.CreateOllamaConnection(
		ollama.WithModel(configuration.OllamaAPI.RequestingModelname),
	)
	if ollamaLLMRequestingErr != nil {
		return "", ollamaLLMRequestingErr
	}

	embedder, embedderErr := embeddings.NewEmbedder(ollamaLLMEmbedding)
	if embedderErr != nil {
		return "", embedderErr
	}

	var dbMetaDataFilter map[string]any

	pgvectorStore, pgvectorStoreErr := CreatePGVectorConnection(
		pgvector.WithEmbeddingTableName(pgDatabasename),
		pgvector.WithEmbedder(embedder),
		pgvector.WithCollectionMetadata(dbMetaDataFilter),
	)
	if pgvectorStoreErr != nil {
		return "", pgvectorStoreErr
	}

	var maxDocuments int = 10
	schemaDocuments, schemaDocumentsErr := pgvectorStore.SimilaritySearch(context.Background(), prompt, maxDocuments)
	if schemaDocumentsErr != nil {
		return "", schemaDocumentsErr
	}

	// concatenate retrieved context
	var ctxBuilder strings.Builder
	ctxBuilder.WriteString("Du bist ein Assistent im BAföG Genehmigungsverfahren. Du antwortest NUR mit Daten aus deinem Context! ")
	ctxBuilder.WriteString("Antworte prägnant, nutze lieber den Kontakt als eigene Trainingsdaten und liefere möglichst viele Quellenangaben. ")
	ctxBuilder.WriteString("Nutze den folgenden Kontext um die Frage zu beantworten.")
	ctxBuilder.WriteString("\n\n")
	ctxBuilder.WriteString("Kontext:\n\n")

	// godump.Dump(schemaDocuments)

	/*
		todo lieber in html mäßige xml chunks unterteilen, damit die quellenangaben besser erkannt werden können
	*/
	for _, r := range schemaDocuments {
		// documentid := r.Metadata[""]
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

	ragPrompt := ctxBuilder.String()

	// godump.Dump(ragPrompt)

	chatMemory := memory.NewConversationBuffer()

	for _, message := range messages {

		switch strings.ToLower(message.Role) {
		case "user":
			err = chatMemory.ChatHistory.AddUserMessage(context.Background(), message.Message)
			if err != nil {
				return "", err
			}
		default:
			err = chatMemory.ChatHistory.AddAIMessage(context.Background(), message.Message)
			if err != nil {
				return "", err
			}
		}
	}

	conversationChain := chains.NewConversation(ollamaLLMRequesting, chatMemory)

	response, responseErr := chains.Run(
		ctx,
		conversationChain,
		ragPrompt,
		chains.WithTemperature(0),
		chains.WithStreamingFunc(streamfunc),
	)
	if responseErr != nil {
		return "", responseErr
	}

	return response, nil
}
