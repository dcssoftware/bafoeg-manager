export async function updateApplicationAssignedUser(applicationID: string, userID: string): Promise<boolean | undefined> {
  try {
    const response = await fetch(`/api/v1/applications/${applicationID}/change-assigned-user/${userID}`, {
      method: 'PATCH',
    });
    if (!response.ok) {
      throw new Error('Failed to update application assigned user by application ID');
    }

    return true;
  } catch (error) {
    console.error('Failed to update application assigned user by application ID', error);
    return false;
  }
}