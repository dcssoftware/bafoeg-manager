export interface RagInformationModelResponseType {
  aiModelServerAddress: string;
  aiModelServerPort: number;
  aiModelServerSecure: boolean;

  embeddingModelname: string;
  requestingModelname: string;
}