import type { ApplicationFileModelType } from "$lib/api/applications/types/application-file-model-type";

export function sortByFileType(files: ApplicationFileModelType[]) {
  const priority = {
    pdf: 1,
    png: 2,
    jpg: 2,
    jpeg: 2,
    webp: 2,
  } as const;

  return files.sort((a, b) => {
    const extA = a.fileType.toLowerCase();
    const extB = b.fileType.toLowerCase();

    const prioA = priority[extA as keyof typeof priority] || 3;
    const prioB = priority[extB as keyof typeof priority] || 3;

    if (prioA === prioB) {
      return extA.localeCompare(extB);
    }

    return prioA - prioB;
  });
}