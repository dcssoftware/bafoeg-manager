export interface SchoolDegreesReponseModelType {
  degrees: SchoolDegreeModelType[];
  count: number;
  maxCount: number;
}

export interface SchoolDegreeModelType {
  id: string;
  name: string;
  schoolID?: string;
  fosBerufsabschlussRequired: boolean;
  bosBerufsqualifizierenderAbschluss: boolean;
  fachschuleBerufsschuleRequired: boolean;
}