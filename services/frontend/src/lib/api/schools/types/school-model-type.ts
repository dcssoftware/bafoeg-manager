import type { SchoolDegreeModelType } from "./school-degrees-model-type";

export interface SchoolModelType {
  id: string;
  name: string;
  type: SchoolModelType;
  degree: SchoolDegreeModelType[];
  street: string;
  houseNumber: string;
  city: string;
  zipCode: string;
  country: string;
}

