export async function requestRagSchueler(conversationID: string | undefined, userInput: string): Promise<ReadableStreamDefaultReader<Uint8Array<ArrayBuffer>> | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('prompt', userInput);

    if (conversationID !== undefined) {
      queryParams.append('conversationID', conversationID);
    }

    const response = await fetch(`/api/v1/rag/bafoeg/schueler/request?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to request RAG Schüler data');
    }

    const reader = response.body?.getReader()

    return reader
  } catch (error) {

    console.error('Failed to request RAG Schüler data', error);
    return undefined
  }
}