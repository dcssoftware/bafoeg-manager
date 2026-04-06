export interface ApplicantsBySchool {
  count: number;
  maxCount: number;
  applicants: ApplicantBySchool[];
}

export interface ApplicantBySchool {
  id: string;
  firstname: string;
  lastname: string;
  classLevel: string;
  statusIdentifier: string;
  degree: ApplicantBySchoolDegree;
  address: ApplicantBySchoolAddress
}

export interface ApplicantBySchoolDegree {
  id: string;
  name: string;
}

export interface ApplicantBySchoolAddress {
  zipCode: string;
  city: string;
  country: string;
}