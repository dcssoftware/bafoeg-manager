import type { SchoolTypeModelType } from "./school-model-type";

export interface SchoolShortResponseModelType {
  schools: SchoolShortModelType[];
  maxCount: number;
  count: number;
}

export interface SchoolShortModelType {
  id: string;
  name: string;
  type: SchoolTypeModelType;
  street: string;
  houseNumber: string;
  city: string;
  zipCode: string;
  country: string;
}

