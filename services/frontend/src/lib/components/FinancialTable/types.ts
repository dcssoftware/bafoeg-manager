interface FinancialTable {
  title: string;
  values: FinancialTableRow[];
  ergebnis_title?: string | undefined;
}
interface FinancialTableRow {
  header: string;
  value: number;
  isResult?: boolean;
}