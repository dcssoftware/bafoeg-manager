import type { RegionResponseModelType } from "./models/region";

export async function getOrganizationRegions(page: number): Promise<RegionResponseModelType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append("page", page.toString());

    const response = await fetch(`/api/v1/organization/regions?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch organization regions data');
    }

    const data: RegionResponseModelType = await response.json()
    return data;

  } catch (error) {

    console.error('Failed to fetch organization regions data', error);
    return undefined
  }
}