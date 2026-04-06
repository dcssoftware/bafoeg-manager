import type { AbteilungenResponseModelType } from "./models/abteilung";

export async function getOrganizationAbteilungByBehördenID(page: number, behoerdeID: string): Promise<AbteilungenResponseModelType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append("page", page.toString());

    const response = await fetch(`/api/v1/organization/abteilungen/${behoerdeID}?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch organization abteilungen data');
    }

    const data: AbteilungenResponseModelType = await response.json()
    return data;

  } catch (error) {

    console.error('Failed to fetch organization regions data', error);
    return undefined
  }
}

