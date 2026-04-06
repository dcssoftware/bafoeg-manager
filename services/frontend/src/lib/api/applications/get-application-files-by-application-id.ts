import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationFileModelResponseType } from "./types/application-file-model-type";

export async function getApplicationFilesByApplicationID(page: number, applicationID?: string): Promise<ApplicationFileModelResponseType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('page', page.toString());

    const response = await fetch(`/api/v1/applications/${applicationID}/files?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch application files by application ID');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;
  } catch (error) {

    console.error('Failed to fetch application files by application ID', error);
    return undefined
  }
}