import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicantResponse } from ".";

export async function getApplicants(page: number, filter: string): Promise<ApplicantResponse | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('page', page.toString());
    if (filter) {
      queryParams.append('filterResult', filter);
    }

    const response = await fetch(`/api/v1/applications/applicants/?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch applicants data');
    }

    const data = convertStructTypes(await response.json())

    return data;
  } catch (error) {

    console.error('Failed to fetch applicants data', error);
    return undefined
  }
}
