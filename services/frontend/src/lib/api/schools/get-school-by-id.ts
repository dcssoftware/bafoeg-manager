import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { SchoolModelType } from ".";

export async function getSchoolByID(schoolID: string): Promise<SchoolModelType | undefined> {
    try {
        const response = await fetch(`/api/v1/schools/${schoolID}`);
        if (!response.ok) {
            throw new Error('Failed to fetch school data by id');
        }

        const rawData = await response.json()
        const data = convertStructTypes(rawData)

        return data;
    } catch (error) {

        console.error('Failed to fetch school data by id', error);
        return undefined
    }
}
