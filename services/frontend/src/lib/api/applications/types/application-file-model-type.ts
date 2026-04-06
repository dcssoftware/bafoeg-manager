export interface ApplicationFileModelResponseType {
  count: number;
  maxCount: number;
  files: ApplicationFileModelType[];
}

export interface ApplicationFileModelType {
  source: string;
  id: string;
  applicationID: string;
  eakte: ApplicationFileModelEakteModelType;
  file: ApplicationFileModelFileModelType;
}

export interface ApplicationFileModelEakteModelType {
  akteID: string;
  vertraulichkeit: string;
  created: Date;
}

export interface ApplicationFileModelFileModelType {
  fileID: string;
  name: string;
  size: number;
  type: string;
  created: Date;
}
