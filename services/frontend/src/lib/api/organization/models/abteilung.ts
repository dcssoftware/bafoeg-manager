export interface AbteilungenResponseModelType {
  count: number;
  maxCount: number;
  abteilungen: AbteilungModelType[];
}

export interface AbteilungModelType {
  id: string;
  name: string;
  behoerde_id: string;
}