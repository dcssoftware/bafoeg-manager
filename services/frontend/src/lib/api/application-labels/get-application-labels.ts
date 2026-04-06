import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationLabelsResponse } from "./types/labels";

export async function getApplicationLabels(page: number): Promise<ApplicationLabelsResponse | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('page', page.toString());
    const response = await fetch(`/api/v1/applications/labels?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch application label data');
    }
    const rawData = await response.json()
    const data = convertStructTypes(rawData)

    return data;
  } catch (error) {

    console.error('Failed to fetch application label data', error);
    return undefined
  }
}