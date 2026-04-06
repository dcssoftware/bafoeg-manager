import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationRevisionModelResponseType } from "./types/application-revision-model-type";

export async function getApplicationsRevisionsByApplicationID(page: number, applicationID?: string): Promise<ApplicationRevisionModelResponseType | undefined> {
    try {
        const queryParams = new URLSearchParams();
        queryParams.append('page', page.toString());

        const response = await fetch(`/api/v1/applications/${applicationID}/revision?${queryParams.toString()}`);
        if (!response.ok) {
            throw new Error('Failed to fetch application revisions by application ID');
        }

        const rawData = await response.json()
        const data = convertStructTypes(rawData)
        return data;
    } catch (error) {

        console.error('Failed to fetch application revisions by application ID', error);
        return undefined
    }
}
