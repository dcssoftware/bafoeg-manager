
export async function isApplicationUpdatableByApplicationID(applicationID: string): Promise<boolean> {
  try {
    const response = await fetch(`/api/v1/applications/${applicationID}/is-updatable`);
    if (!response.ok) {
      return false
    }
    return true;
  } catch (error) {
    console.error('Failed to check if application is updatable', error);
    return false;
  }
}