
export async function updateApplicationAssignedSchoolDegree(applicationID: string, schoolDegreeID: string): Promise<boolean> {
  try {
    const response = await fetch(`/api/v1/applications/${applicationID}/change-assigned-school-degree/${schoolDegreeID}`, {
      method: 'PATCH',
    });
    if (!response.ok) {
      throw new Error('Failed to update application assigned school degree');
    }
    return true;
  } catch (error) {
    console.error('Failed to update application assigned school degree', error);
    return false;
  }
}