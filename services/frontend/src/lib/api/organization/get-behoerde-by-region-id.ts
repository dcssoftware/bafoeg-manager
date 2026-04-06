import type { BehördeResponseModelType } from "./models/behoerde";

export async function getOrganizationBehördenByRegion(page: number, regionId: string): Promise<BehördeResponseModelType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append("page", page.toString());

    const response = await fetch(`/api/v1/organization/behoerden/${regionId}?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch organization regions data');
    }

    const data: BehördeResponseModelType = await response.json()
    return data;

  } catch (error) {

    console.error('Failed to fetch organization regions data', error);
    return undefined
  }
}