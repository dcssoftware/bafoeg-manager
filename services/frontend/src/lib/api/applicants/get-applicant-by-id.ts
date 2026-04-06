import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { Applicant } from "./types";

export async function getApplicantByID(applicantID: string): Promise<Applicant | undefined> {
  try {
    const queryParams = new URLSearchParams();

    const response = await fetch(`/api/v1/applications/applicants/${applicantID}?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch applicants data by id');
    }

    const data = convertStructTypes(await response.json())

    return data;
  } catch (error) {
    console.error('Failed to fetch applicants data by id', error);
    return undefined
  }
}