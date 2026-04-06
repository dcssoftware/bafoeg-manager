import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationModelResponseType } from ".";

export async function getApplicationByID(applicationID: string): Promise<ApplicationModelResponseType | undefined> {
  try {
    const response = await fetch(`/api/v1/applications/${applicationID}`);
    if (!response.ok) {
      throw new Error('Failed to fetch applications data by applicant ID');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;

  } catch (error) {
    console.error('Failed to fetch applications data by applicant ID', error);
    return undefined
  }
}