package models

type RagInformationModel struct {
	AIModelServerAddress string `json:"aiModelServerAddress"`
	AIModelServerPort    uint   `json:"aiModelServerPort"`
	AIModelServerSecure  bool   `json:"aiModelServerSecure"`

	EmbeddingModelname  string `json:"embeddingModelname"`
	RequestingModelname string `json:"requestingModelname"`
}
