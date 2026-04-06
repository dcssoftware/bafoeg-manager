package http

import "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service"

type RAGHandler struct {
	service *service.RAGService
}

func NewRAGHandler(service *service.RAGService) *RAGHandler {
	return &RAGHandler{
		service: service,
	}
}
