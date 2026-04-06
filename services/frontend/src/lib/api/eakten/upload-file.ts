import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { EaktenFileUploadStatus } from "./types/eakten-file-upload-status-type";

export async function uploadEaktenFiles(file: File): Promise<EaktenFileUploadStatus | undefined> {
  try {
    const formData = new FormData();
    formData.append("file", file);

    const response = await fetch(`/api/v1/eakten/documents/upload`, {
      method: "POST",
      body: formData
    });

    if (response.status !== 200 && response.status !== 201) {
      throw new Error('Failed to upload eakten file');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;
  } catch (error) {
    console.error('Failed to upload eakten file', error);
    throw error;
  }
}
