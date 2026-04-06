import { filesize } from "filesize";

export function convertFileSize(size: number): string {
  return filesize(size, { standard: "jedec" });
}  