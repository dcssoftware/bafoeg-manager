import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { RagInformationModelResponseType } from "./type/rag-information-model-type";

export async function getRagInformation(): Promise<RagInformationModelResponseType | undefined> {
  try {
    const response = await fetch(`/api/v1/rag/information`);
    if (!response.ok) {
      throw new Error('Failed to fetch rag information data');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;

  } catch (error) {
    console.error('Failed to fetch rag information data', error);
    return undefined
  }
}