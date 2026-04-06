export async function uploadRAGrelevantDocumentsSchüler(file: File) {
  const formData = new FormData();
  formData.append("file", file);


  const response = await fetch("/api/v1/rag/bafoeg/schueler", {
    method: "POST",
    body: formData,
  });

  if (!response.ok) {
    throw new Error("Failed to upload file");
  }
}