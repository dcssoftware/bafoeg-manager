export interface EakteHttpResponse {
  maxCount: number;
  count: number;
  eakten: EakteModel[];
}

export interface EakteModel {
  id: string;
  aktenzeichen: string;
  type: EakteModelType;
  vertraulichkeit: string;
  created: Date;
}

export interface EakteModelType {
  id: string;
  name: string;
  identifier: string;
}