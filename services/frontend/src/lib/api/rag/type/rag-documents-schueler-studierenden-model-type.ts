export interface RAGDocumentsSchülerStudierendenModelResponseType {
  documents: RAGDocumentsSchülerStudierendenModelResponseType[];
  count: number;
  maxCount: number;
}

export interface RAGDocumentsSchülerStudierendenModelResponseType {
  id: string;
  fileName: string;
  fileType: string;
  fileSize: string;
  status: RAGDocumentsSchülerStudierendenStatusModelResponseType;
  created: Date;
  createdFrom: RAGDocumentsSchülerStudierendenCreatedFromModelResponseType;
}

export interface RAGDocumentsSchülerStudierendenCreatedFromModelResponseType {
  id: string;
  displayName: string;
  username: string;
}

export interface RAGDocumentsSchülerStudierendenStatusModelResponseType {
  id: string;
  identifier: string;
  name: string;
}
