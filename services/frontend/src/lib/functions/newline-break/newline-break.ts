export function newline2Break(input: string): string {
  return input.replaceAll("\n", "<br>")
}