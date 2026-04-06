import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationModelResponseType } from "$lib/api/applications/types/application-model-type";
import { stringToBooleanString } from "$lib/functions/string-to-boolean-string";

export async function getApplications(page: number, userID?: string, showFinishedApplications: boolean = false, filterResult: string = ""): Promise<ApplicationModelResponseType | undefined> {
    try {
        const queryParams = new URLSearchParams();
        queryParams.append('page', page.toString());
        if (userID) {
            queryParams.append('userID', userID);
        }
        queryParams.append('showAllApplications', stringToBooleanString(showFinishedApplications));
        queryParams.append('filterResult', filterResult);

        const response = await fetch(`/api/v1/applications?${queryParams.toString()}`);
        if (!response.ok) {
            throw new Error('Failed to fetch applications data');
        }

        const rawData = await response.json()
        const data = convertStructTypes(rawData)
        return data;
    } catch (error) {

        console.error('Failed to fetch applications data', error);
        return undefined
    }
}
