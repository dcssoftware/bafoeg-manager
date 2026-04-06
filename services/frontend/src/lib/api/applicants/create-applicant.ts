import type { CreateApplicantModel } from "./types/create-applicant";

export async function createApplicant(model: CreateApplicantModel): Promise<void> {
  try {
    fetch(`/api/v1/applications/applicants`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(model),
    })
  } catch (error) {
    console.error("Error creating applicant:", error);
    throw error;
  }
}