export async function getEaktenApplicationMapping(eakteAkteID: string): Promise<EaktenApplicationMappingModel | null> {
  try {
    const response = await fetch(
      `/api/v1/eakten/${eakteAkteID}/get-application-mapping`
    );
    if (!response.ok) {
      throw new Error("Failed to fetch E-Akte application mapping");
    }

    const data = await response.json() satisfies EaktenApplicationMappingModel;
    return data;

  } catch (error) {

    console.error("Error fetching E-Akte application mapping:", error);
    return null;
  }
}