import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { EakteHttpResponse } from "./types";

export async function getEakten(page: number): Promise<EakteHttpResponse | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('page', page.toString());

    const response = await fetch(`/api/v1/eakten?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch eakten data');
    }

    const rawData = await response.json()
    const data = convertStructTypes(rawData)
    return data;
  } catch (error) {

    console.error('Failed to fetch eakten data', error);
    return undefined
  }
}