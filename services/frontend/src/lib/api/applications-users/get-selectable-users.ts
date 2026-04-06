import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationAssignableUserModelType } from ".";

export async function getApplicationsSelectableUsers(): Promise<ApplicationAssignableUserModelType[] | undefined> {
    try {
        const response = await fetch(`/api/v1/applications/assignable-users/`);
        if (!response.ok) {
            throw new Error('Failed to fetch application assignable users data');
        }

        const rawData = await response.json()
        const data = convertStructTypes(rawData)

        return data;
    } catch (error) {

        console.error('Failed to fetch application assignable users data', error);
        return undefined
    }
}
