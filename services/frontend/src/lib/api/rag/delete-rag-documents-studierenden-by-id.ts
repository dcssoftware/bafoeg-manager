export async function deleteRagDocumentsStudierendenByID(fileID: string): Promise<boolean> {
	try {
		const response = await fetch(`/api/v1/rag/bafoeg/studierenden/${fileID}`, {
			method: 'DELETE',
		});
		if (!response.ok) {
			throw new Error('Failed to delete RAG documents for Studierenden by ID');
		}
		return true;
	} catch (error) {
		console.error('Error deleting RAG documents for Studierenden by ID:', error);
		throw error;
	}
}

