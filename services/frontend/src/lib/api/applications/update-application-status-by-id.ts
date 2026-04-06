export async function updateApplicationStatusByID(applicationID: string, status: string): Promise<undefined> {
  try {
    const response = await fetch(`/api/v1/applications/${applicationID}/change-status/${status}`, {
      method: 'PATCH',
    });
    if (!response.ok) {
      throw new Error('Failed to update application status by application ID');
    }

    return undefined;
  } catch (error) {
    console.error('Failed to update application status by application ID', error);
    return undefined
  }
}