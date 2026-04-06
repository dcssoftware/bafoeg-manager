import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationModelResponseType } from "$lib/api/applications/types/application-model-type";

export async function getApplicationsByApplicantID(page: number, applicantID?: string): Promise<ApplicationModelResponseType | undefined> {
    try {
        const queryParams = new URLSearchParams();
        queryParams.append('page', page.toString());

        const response = await fetch(`/api/v1/applications/by-applicant-id/${applicantID}?${queryParams.toString()}`);
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
