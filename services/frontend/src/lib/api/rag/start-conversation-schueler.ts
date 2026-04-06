import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { RagConversationIDModelType } from "./type/rag-conversation-id";

export async function startConversationSchueler(): Promise<RagConversationIDModelType | undefined> {
  try {
    const response = await fetch(`/api/v1/rag/bafoeg/schueler`, {
      method: 'PUT'
    });
    if (!response.ok) {
      throw new Error('Failed to start new RAG Schüler conversation');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)

    return data
  } catch (error) {

    console.error('Failed to start new RAG Schüler conversation', error);
    return undefined
  }
}