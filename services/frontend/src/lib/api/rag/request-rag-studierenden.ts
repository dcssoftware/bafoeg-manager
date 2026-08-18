export async function requestRagStudierenden(
  conversationID: string | undefined,
  userInput: string,
  abortSignal?: AbortSignal,
): Promise<ReadableStreamDefaultReader<Uint8Array<ArrayBuffer>> | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('prompt', userInput);

    if (conversationID !== undefined) {
      queryParams.append('conversationID', conversationID);
    }

    const response = await fetch(`/api/v1/rag/bafoeg/studierenden/request?${queryParams.toString()}`, {
      signal: abortSignal,
    });
    if (!response.ok) {
      throw new Error('Failed to request RAG Studierenden data');
    }

    const reader = response.body?.getReader()

    return reader
  } catch (error) {
    if (error instanceof DOMException && error.name === 'AbortError') {
      return undefined;
    }

    console.error('Failed to request RAG Studierenden data', error);
    return undefined
  }
}