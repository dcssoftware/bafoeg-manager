<script lang="ts">
  import { Button } from "$lib/components/Button";
  import { DocumentIndicator } from "$lib/components/Document-Indicator";
  import { convertFileSize } from "$lib/functions/convert-file-size";
  import type { FileUploadFile } from "./types";

  interface Props {
    onUpload: (file: FileUploadFile) => Promise<void>;
    onUploadFileFinished: () => void;
    verifyFileUnique: (file: File) => Promise<boolean>;
  }

  let { onUpload = onUploadDefault, onUploadFileFinished }: Props = $props();

  let isLoading: boolean = $state(false);
  let uploadedFiles: FileUploadFile[] = $state([]);

  function preventDefault(event: Event) {
    event.preventDefault();
    event.stopPropagation();
  }

  async function handleDrop(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();
    const files = event.dataTransfer?.files;
    if (files && files.length > 0) {
      const fileInput = document.querySelector(
        'input[type="file"]'
      ) as HTMLInputElement;
      if (fileInput) {
        fileInput.files = files;
      }
    }
    if (files && files.length > 0) {
      for (const file of files) {
        uploadedFiles.push({ file: file, status: "pending", error: undefined });
      }
    }
  }

  async function handleFileInput(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
      for (const file of target.files) {
        uploadedFiles.push({ file: file, status: "pending", error: undefined });
      }
    }
  }

  function removeFile(fileIndex: number, fileName: string) {
    uploadedFiles = uploadedFiles.filter((fileObj, index) => {
      const isIndex = index == fileIndex;
      const name = fileName == fileObj.file.name;
      return !(isIndex && name);
    });
  }

  async function triggerUpload() {
    isLoading = true;
    for (const file of uploadedFiles) {
      try {
        await onUpload(file);
        file.status = "done";
      } catch (error) {
        file.status = "error";
      }
    }

    await new Promise((resolve) => {
      const timeoutID = setTimeout(async () => {
        await onUploadFileFinished();
        resolve(timeoutID);
      }, 7500);
    });

    uploadedFiles = [];
  }

  async function onUploadDefault(file: FileUploadFile) {
    throw new Error("Upload failed");
  }
</script>

<div class="pdf-upload">
  {#if isLoading}
    {#if uploadedFiles != null && uploadedFiles.length > 0}
      <div>
        <table>
          <thead>
            <tr>
              <th>File Name</th>
              <th>File Size</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {#each Array.from(uploadedFiles) as fileObj}
              <tr>
                <td>{fileObj.file.name}</td>
                <td>{convertFileSize(fileObj.file.size)}</td>
                <td>Status: {fileObj.status}</td>
              </tr>
            {/each}
          </tbody>
        </table>
        <div>
          <span
            >Nach dem Upload bitte kurz warten, die Seite lädt die Daten neu</span
          >
        </div>
      </div>
    {:else}
      <p>No files uploaded yet.</p>
    {/if}
  {:else}
    <div class="upload-queue">
      {#if uploadedFiles != null && uploadedFiles.length > 0}
        <div class="inside-list">
          <ul>
            {#each Array.from(uploadedFiles) as fileObj, fileIndex}
              <li>
                <DocumentIndicator
                  documentID={""}
                  documentName={fileObj.file.name}
                  documentType={fileObj.file.type}
                  documentSize={fileObj.file.size}
                  documentClickable={false}
                  documentCreated={new Date(fileObj.file.lastModified)}
                />
                <div class="button">
                  <Button
                    onclick={() => removeFile(fileIndex, fileObj.file.name)}
                    buttonType="danger">Remove</Button
                  >
                </div>
              </li>
            {/each}
          </ul>
        </div>
      {:else}
        <span>No files uploaded yet.</span>
      {/if}
    </div>
    <label
      ondragenter={preventDefault}
      ondragover={preventDefault}
      ondragleave={preventDefault}
      ondrop={handleDrop}
    >
      <input
        type="file"
        multiple
        oninput={handleFileInput}
        accept="application/pdf"
      />
      <div>
        <span>Drop Files here</span>
      </div>
    </label>
    <div>
      <Button onclick={triggerUpload}>Upload</Button>
    </div>
  {/if}
</div>

<style lang="sass">
  .upload-queue
    div.inside-list
      user-select: none
      overflow-x: scroll
      background-color: var(--background-color-tertiary)
      ul
        display: flex
        flex-direction: row
        gap: 1rem
        list-style-type: none
        padding: 1rem 
        li
          min-width: 300px
          width: 300px
          overflow: hidden
          display: flex
          flex-direction: column
          align-content: stretch
          justify-content: center
          padding: 2rem

          div.button
            :global(button)
              width: 100%
              margin-top: max(auto, 1rem)
              text-align: center

          &:hover
            background-color: var(--background-color-secondary)

  .pdf-upload
      input
        display: none
      label
        height: 30vh
        width: 100%
        display: flex
        justify-content: center
        align-items: center
        border: 2px dashed var(--font-color)
        border-radius: 10px
        cursor: pointer
        user-select: none
        margin: 2rem 0 1rem
</style>
