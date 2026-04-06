export interface Payment {
  id: string;
  applicantID: string
  applicationID: string;
  amount: number;
  statusID: string;
  statusIdentifier: string;
  description: string;
  iban: string;
  bic: string;
  direction: "incoming" | "outgoing";
  executed: Date;
  created: Date;
}