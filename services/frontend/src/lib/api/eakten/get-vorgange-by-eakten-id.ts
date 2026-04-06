import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { VorgangHttpResponse } from "./types";

export async function getVorgängeByEaktenID(eaktenID: string): Promise<VorgangHttpResponse[] | undefined> {
  try {
    const queryParams = new URLSearchParams();

    const response = await fetch(`/api/v1/eakten/${eaktenID}/vorgang?${queryParams.toString()}`);
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