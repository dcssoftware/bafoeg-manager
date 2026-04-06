package service

import (
	"context"
	"fmt"
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	"github.com/dcssoftware/bafoeg-manager/src/helper/artificial-intelligence/rag"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	"github.com/google/uuid"
)

func (s *CronjobRagVectorProcessorService) ProcessFiles(ctx context.Context) error {

	documentsSchüler, documentsSchülerErr := s.storage.GetProcessableFilesSchülerFromDatabase(nil)
	if documentsSchülerErr != nil {
		return documentsSchülerErr
	}

	for _, document := range documentsSchüler {

		logger.Info(
			uuid.NewString(),
			fmt.Sprintf("Process (Schüler)  %s (%s)", document.Filename, document.ID),
			"",
		)

		if ctx.Err() != nil {
			return ctx.Err()
		}

		documentContent, documentContentErr := s.storageS3.DownloadDocumentSchülerByID(document.ID.String())
		if documentContentErr != nil {
			s.storage.SetProcessedFileStatusSchüler(nil, document.ID, documentContentErr.Error())
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := rag.CreateVectorsPDF(documentContent, configuration.OllamaAPI.DatabaseTablenameRAGSchueler, []rag.MetaDataEntry{
			{
				Key:   "originDocumentID",
				Value: document.ID,
			},
			{
				Key:   "createdAt",
				Value: time.Now().String(),
			},
		})
		if err != nil {
			s.storage.SetProcessedFileStatusSchüler(nil, document.ID, err.Error())
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		s.storage.SetProcessedFileStatusSchüler(nil, document.ID, "")

		logger.Debug(
			uuid.NewString(),
			fmt.Sprintf(
				"CRONJOB Processed RAG file: %s (%s) from: %s (@%s) status: successful\n",
				document.Filename,
				document.ID,
				document.CreatedFrom.DisplayName,
				document.CreatedFrom.Username,
			),
			nil,
		)
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	documentsStudierenden, documentsStudierendenErr := s.storage.GetProcessableFilesStudierendenFromDatabase(nil)
	if documentsStudierendenErr != nil {
		return documentsStudierendenErr
	}

	for _, document := range documentsStudierenden {

		logger.Info(
			uuid.NewString(),
			fmt.Sprintf("Process (Studierenden)  %s (%s)", document.Filename, document.ID),
			"",
		)

		if ctx.Err() != nil {
			return ctx.Err()
		}

		documentContent, documentContentErr := s.storageS3.DownloadDocumentStudierendenByID(document.ID.String())
		if documentContentErr != nil {
			s.storage.SetProcessedFileStatusStudierenden(nil, document.ID, documentContentErr.Error())
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := rag.CreateVectorsPDF(documentContent, configuration.OllamaAPI.DatabaseTablenameRAGStudierenden, []rag.MetaDataEntry{
			{
				Key:   "originDocumentID",
				Value: document.ID,
			},
			{
				Key:   "createdAt",
				Value: time.Now().String(),
			},
		})
		if err != nil {
			s.storage.SetProcessedFileStatusStudierenden(nil, document.ID, err.Error())
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		s.storage.SetProcessedFileStatusStudierenden(nil, document.ID, "")

		logger.Debug(
			uuid.NewString(),
			fmt.Sprintf(
				"CRONJOB Processed RAG file: %s (%s) from: %s (@%s) status: successful\n",
				document.Filename,
				document.ID,
				document.CreatedFrom.DisplayName,
				document.CreatedFrom.Username,
			),
			nil,
		)
	}

	return nil
}
