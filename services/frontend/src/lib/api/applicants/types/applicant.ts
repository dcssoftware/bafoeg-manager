import type { ApplicantAddress, ApplicantContact, ApplicantTrainingsAddress } from ".";

export interface ApplicantResponse {
  maxCount: number;
  count: number;
  applicants: Applicant[];
}

export interface Applicant {
  id: string;
  firstname: string;
  lastname: string;

  balanceOutgoing: number | null;
  balancePayback: number | null;

  address: ApplicantAddress;
  trainingsAddress?: ApplicantTrainingsAddress;
  contact: ApplicantContact;
}