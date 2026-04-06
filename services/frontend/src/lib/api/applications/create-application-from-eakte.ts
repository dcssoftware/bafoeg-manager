export async function createApplicationFromEakte(applicantID: string, eakteID: string, schoolDegreeID: string, educationStartDate: Date, educationEndDate: Date, classLevel: string, labelIDs: string[] = []): Promise<boolean> {
  try {
    const requestBody: ApplicationCreateFromEakteInsertModel = {
      applicantID,
      eakteID,
      schoolDegreeID,
      educationStartDate: educationStartDate.toISOString(),
      educationEndDate: educationEndDate.toISOString(),
      classLevel,
      labelIDs
    };

    const response = await fetch('/api/v1/applications/eakte-new-application', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody)
    });

    if (!response.ok) {
      throw new Error('Failed to create application from E-Akte');
    }

    return true;
  } catch (error) {
    console.error('Error creating application from E-Akte:', error);
    return false;
  }
}