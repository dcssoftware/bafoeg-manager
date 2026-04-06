import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { Licenses } from "./types/licenses";

export async function getLicenses(): Promise<Licenses | undefined> {
  try {
    const response = await fetch(`/api/asset/licenses`);
    if (!response.ok) {
      throw new Error('Failed to fetch licenses data');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;
  } catch (error) {

    console.error('Failed to fetch licenses data', error);
    return undefined
  }
}
