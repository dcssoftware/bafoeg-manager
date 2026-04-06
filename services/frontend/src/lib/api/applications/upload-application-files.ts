import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationFileUploadStatus } from "./types/application-file-upload-status-type";

export async function uploadApplicationFiles(applicationID: string, file: File): Promise<ApplicationFileUploadStatus | undefined> {
  try {
    const formData = new FormData();
    formData.append("file", file);

    const response = await fetch(`/api/v1/applications/${applicationID}/uploadFiles`, {
      method: "POST",
      body: formData
    });

    if (response.status !== 200 && response.status !== 201) {
      throw new Error('Failed to upload application file');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;
  } catch (error) {
    console.error('Failed to upload application file', error);
    throw error;
  }
}
