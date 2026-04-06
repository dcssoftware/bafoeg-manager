export function prettifyIban(iban: string): string {
  const cleanIban = iban.replace(/\s+/g, '').toUpperCase();
  return cleanIban.replace(/(.{4})/g, '$1 ').trim();
}