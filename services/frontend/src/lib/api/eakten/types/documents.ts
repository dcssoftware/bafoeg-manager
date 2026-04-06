export interface DocumentsHttpResponse {
  maxCount: number;
  count: number;
  documents: DocumentsModel[];
}

export interface DocumentsModel {
  id: string;
  source_xdomea_file: boolean;
  source: DocumentsModelSource;
  files: DocumentsModelFile;
  created: Date;
}

export interface DocumentsModelSource {
  id: string;
  identifier: string;
  name: string;
}

export interface DocumentsModelFile {
  id: string;
  name: string;
  type: string;
  size: number;
  created: Date;
}

