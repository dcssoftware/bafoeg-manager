export function dateToFormatStringShort(date: Date): string {
  const options: Intl.DateTimeFormatOptions = {
    year: "2-digit",
    month: "2-digit",
    day: "2-digit",
  };
  return new Intl.DateTimeFormat("de-DE", options).format(date);
}