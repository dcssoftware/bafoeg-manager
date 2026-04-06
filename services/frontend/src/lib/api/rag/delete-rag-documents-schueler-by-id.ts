export async function deleteRagDocumentsSchuelerByID(fileID: string): Promise<boolean> {
  try {
    const response = await fetch(`/api/v1/rag/bafoeg/schueler/${fileID}`, {
      method: 'DELETE',
    });
    if (!response.ok) {
      throw new Error('Failed to delete RAG documents for Schüler by ID');
    }
    return true;
  } catch (error) {
    console.error('Error deleting RAG documents for Schüler by ID:', error);
    throw error;
  }
}
