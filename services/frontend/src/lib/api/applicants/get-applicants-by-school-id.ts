import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicantsBySchool } from "./types/applicant-by-school";

export async function getApplicantsBySchoolID(page: number, schoolID: string, isActive: boolean): Promise<ApplicantsBySchool | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('page', page.toString());
    queryParams.append('isActive', isActive ? 'true' : 'false');

    const response = await fetch(`/api/v1/applications/applicants/by-school/${schoolID}?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch applicants data by school id');
    }

    const data = convertStructTypes(await response.json())

    return data;
  } catch (error) {
    console.error('Failed to fetch applicants data by school id', error);
    return undefined
  }
}