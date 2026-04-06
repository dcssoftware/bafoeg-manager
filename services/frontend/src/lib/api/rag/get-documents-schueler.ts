import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { RAGDocumentsSchülerStudierendenModelResponseType } from "./type/rag-documents-schueler-studierenden-model-type";

export async function getDocumentsSchüler(filterText: string): Promise<RAGDocumentsSchülerStudierendenModelResponseType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('filterResult', filterText);

    const response = await fetch(`/api/v1/rag/bafoeg/schueler?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch rag document data for schüler');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;

  } catch (error) {
    console.error('Failed to fetch rag document data for schüler', error);
    return undefined
  }
}