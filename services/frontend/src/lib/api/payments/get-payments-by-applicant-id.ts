
import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { Payment } from "./types";

export async function getPaymentsByApplicantId(page: number, applicantID: string,): Promise<Payment[] | undefined> {
  try {
    const queryParams = new URLSearchParams();
    queryParams.append('page', page.toString());

    const response = await fetch(`/api/v1/applications/applicants/${applicantID}/payments?${queryParams.toString()}`);
    if (!response.ok) {
      throw new Error('Failed to fetch payments data by applicant id');
    }

    const data = convertStructTypes(await response.json())

    return data;
  } catch (error) {
    console.error('Failed to fetch payments data by applicant id', error);
    return undefined
  }
}