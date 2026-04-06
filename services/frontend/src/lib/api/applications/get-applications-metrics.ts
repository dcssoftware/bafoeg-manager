
import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { ApplicationMetricsModelResponseType } from "./types/application-metrics-model";

export async function getApplicationsMetrics(userID: string, showFinishedApplications: boolean = false): Promise<ApplicationMetricsModelResponseType | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('userID', userID);
    queryParams.append('showAllApplications', String(showFinishedApplications));

    const response = await fetch(`/api/v1/applications/metrics?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch applications metrics');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;

  } catch (error) {
    console.error('Failed to fetch applications metrics', error);
    return undefined
  }
}