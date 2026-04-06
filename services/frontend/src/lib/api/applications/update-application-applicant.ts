// /applications/:applicationID/change-assigned-applicant/:applicantID
export async function updateApplicationApplicant(applicationID: string, newApplicantID: string): Promise<boolean> {
  try {
    const response = await fetch(`/api/v1/applications/${applicationID}/change-assigned-applicant/${newApplicantID}`, {
      method: 'PATCH',
    });
    if (response.ok) {
      return true;
    }
    return false
  } catch (error) {
    console.error('Failed to update application applicant', error);
    return false
  }
}
