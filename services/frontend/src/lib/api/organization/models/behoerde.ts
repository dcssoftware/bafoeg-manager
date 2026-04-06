export interface BehördeResponseModelType {
  count: number;
  maxCount: number;
  behoerden: BehördeModelType[];
}

export interface BehördeModelType {
  id: string;
  name: string;
  region_id: string;
}