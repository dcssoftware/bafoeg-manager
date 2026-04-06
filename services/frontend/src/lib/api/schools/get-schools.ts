import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { SchoolShortResponseModelType } from "./types/school-short-model-type";

export async function getSchools(page: number, filter: string): Promise<SchoolShortResponseModelType | undefined> {
    try {

        const queryParams = new URLSearchParams();
        queryParams.append('page', page.toString());
        if (filter) {
            queryParams.append('filterResult', filter);
        }

        const response = await fetch(`/api/v1/schools?${queryParams.toString()}`);
        if (!response.ok) {
            throw new Error('Failed to fetch school data');
        }

        const rawData = await response.json()
        const data = convertStructTypes(rawData)

        return data;
    } catch (error) {

        console.error('Failed to fetch school data', error);
        return undefined
    }
}
