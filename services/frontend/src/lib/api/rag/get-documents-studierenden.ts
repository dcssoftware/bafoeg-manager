import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { RAGDocumentsSchülerStudierendenModelResponseType } from "./type/rag-documents-schueler-studierenden-model-type";

export async function getDocumentsStudierenden(filterText: string): Promise<RAGDocumentsSchülerStudierendenModelResponseType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('filterResult', filterText);

    const response = await fetch(`/api/v1/rag/bafoeg/studierenden?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch rag document data for studierenden');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;

  } catch (error) {
    console.error('Failed to fetch rag document data for studierenden', error);
    return undefined
  }
}