export interface ApplicationLabelsResponse {
  labels: ApplicationLabelType[];
}

export interface ApplicationLabelType {
  id: string;
  name: string;
  style: ApplicationLabelStyleType;
}

export interface ApplicationLabelStyleType {
  id: string;
  name: string;

  colorDark: string;
  bgColorDark: string;
  borderColorDark: string;

  colorLight: string;
  bgColorLight: string;
  borderColorLight: string;
}