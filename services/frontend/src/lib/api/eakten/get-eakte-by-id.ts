import type { EakteModel } from "./types";

export async function getEakteByID(eaktenID: string): Promise<EakteModel> {
  try {
    const response = await fetch(`/api/v1/eakten/${eaktenID}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Error fetching eakte by ID: ${response.statusText}`);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Failed to fetch eakte by ID:", error);
    throw error;
  }
}
