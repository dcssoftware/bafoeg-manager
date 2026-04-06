import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { SchoolDegreesReponseModelType } from "./types/school-degrees-model-type";

export async function getSchoolDegreesBySchoolID(page: number, schoolID: string): Promise<SchoolDegreesReponseModelType | undefined> {
    try {
        const queryParams = new URLSearchParams();
        queryParams.append('page', page.toString());

        const response = await fetch(`/api/v1/schools/${schoolID}/degrees?${queryParams.toString()}`);
        if (!response.ok) {
            throw new Error('Failed to fetch school degree data by school id');
        }

        const rawData = await response.json()
        const data = convertStructTypes(rawData)

        return data;
    } catch (error) {

        console.error('Failed to fetch school degree data by school id', error);
        return undefined
    }
}
