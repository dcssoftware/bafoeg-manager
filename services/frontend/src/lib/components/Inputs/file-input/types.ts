export interface FileUploadFile {
  file: File;
  status: string;
  error: "virus" | "file-already-exists" | "file-too-large" | "unsupported-format" | undefined;
}