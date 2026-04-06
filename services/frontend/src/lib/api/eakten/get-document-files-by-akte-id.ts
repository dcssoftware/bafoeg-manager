import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { DocumentsHttpResponse } from "./types";

export async function GetDocumentFilesByAkteID(eaktenID: string): Promise<DocumentsHttpResponse | undefined> {
  try {
    const queryParams = new URLSearchParams();

    const response = await fetch(`/api/v1/eakten/${eaktenID}/documents?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch eakten data');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;
  } catch (error) {

    console.error('Failed to fetch eakten data', error);
    return undefined
  }
}