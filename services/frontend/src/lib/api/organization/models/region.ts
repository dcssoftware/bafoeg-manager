export interface RegionResponseModelType {
  count: number;
  maxCount: number;
  regions: RegionModelType[];
}

export interface RegionModelType {
  id: string;
  identifier: string;
  name: string;
}