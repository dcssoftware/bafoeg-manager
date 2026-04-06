<script lang="ts">
  import { _ } from "svelte-i18n";
  import {
    getApplicationFilesByApplicationID,
    type ApplicationStatusModelType,
  } from "$lib/api/applications";
  import type {
    ApplicationFileModelResponseType,
    ApplicationFileModelType,
  } from "$lib/api/applications/types/application-file-model-type";
  import { uploadApplicationFiles } from "$lib/api/applications/upload-application-files";
  import { Button } from "$lib/components/Button";
  import { Dialog } from "$lib/components/Dialog";
  import { DocumentIndicator } from "$lib/components/Document-Indicator";
  import { FileInput } from "$lib/components/Inputs/file-input";
  import type { FileUploadFile } from "$lib/components/Inputs/file-input/types";
  import { Pagination } from "$lib/components/Pagination";
  import { FileViewer } from "$lib/components/FileViewer";

  interface Props {
    applicationID: string | undefined;
    applicationStatus: ApplicationStatusModelType;
  }

  let { applicationID, applicationStatus }: Props = $props();

  let fileUploadRefreshHash: string = $state("");
  let dialogShowFileIsOpen: boolean = $state(false);
  let dialogUploadFileIsOpen: boolean = $state(false);
  // let showFileURL: string = $state("");
  // let fileTypeIndicator: string | undefined = $state(undefined);
  let showFileObj: ApplicationFileModelType | undefined = $state(undefined);

  let filePageNumber: number = $state(1);

  let applicationFilesPromise:
    | Promise<ApplicationFileModelResponseType | undefined>
    | undefined = $state(undefined);

  function showFile(document: ApplicationFileModelType) {
    showFileObj = document;
    dialogShowFileIsOpen = true;
  }

  function getDocumentShowURL(
    document: ApplicationFileModelType | undefined,
  ): string {
    if (document?.source === "APPLICATION_DOCUMENT") {
      return `/api/v1/applications/${applicationID}/files/${document?.id}`;
    } else if (document?.source === "EAKTE_DOCUMENT") {
      return `/api/v1/eakten/documents/${document?.id}`;
    }
    return "";
  }

  function openFileUpload() {
    dialogUploadFileIsOpen = true;
  }

  async function uploadFile(file: FileUploadFile) {
    if (applicationID) {
      await uploadApplicationFiles(applicationID, file.file);
    } else {
      throw new Error("Application ID is not defined");
    }
  }

  async function uploadFileFinished() {
    fileUploadRefreshHash = Math.random().toString(36).substring(2, 15);
    loadData(Math.random().toString(36).substring(2, 15));
    dialogUploadFileIsOpen = false;
  }

  function applicationUploadTitle(state: string) {
    const base = "illegal-state-warning.upload-document-after-completion.";
    switch (state) {
      case "APPROVED":
        return $_(base + "approved");
      case "DENIED":
        return $_(base + "denied");
      default:
        return $_("page.application.file-upload.popover-title");
    }
  }

  function loadData(hash: string = "") {
    applicationFilesPromise = getApplicationFilesByApplicationID(
      filePageNumber,
      applicationID,
    );
  }

  $effect(() => {
    loadData();
  });
</script>

<Dialog
  bind:isOpen={dialogShowFileIsOpen}
  dialogContent={showFileDialog}
  --dialog-width="80vw"
  --dialog-height="80vh"
/>

<Dialog bind:isOpen={dialogUploadFileIsOpen} dialogContent={uploadFileDialog} />

{#snippet showFileDialog()}
  {@const url = getDocumentShowURL(showFileObj)}
  <FileViewer
    showFileURL={url}
    fileTypeIndicator={showFileObj?.file?.type}
    metadata={{}}
    closeDialog={() => (dialogShowFileIsOpen = false)}
  />
{/snippet}

{#snippet uploadFileDialog()}
  {#key fileUploadRefreshHash}
    <FileInput
      onUpload={uploadFile}
      onUploadFileFinished={uploadFileFinished}
      verifyFileUnique={() => {
        return Promise.resolve(true);
      }}
    />
  {/key}
{/snippet}

<div class="header">
  <h2>Dokumente</h2>
  <div
    class="upload-documents"
    title={applicationUploadTitle(applicationStatus.identifier)}
  >
    <Button
      disabled={applicationStatus.identifier === "APPROVED" ||
        applicationStatus.identifier === "DENIED"}
      onclick={openFileUpload}>Upload File</Button
    >
  </div>
</div>

{#await applicationFilesPromise}
  <span>{$_("states.loading")}</span>
{:then files}
  {#if files !== undefined}
    <div class="documents">
      {#if files.count > 0}
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <!-- {#each sortByFileType(files.files) as document} -->
        {#each files.files as document}
          <DocumentIndicator
            documentID={document.id}
            documentName={document.file.name}
            documentSize={document.file.size}
            documentType={document?.file?.type}
            documentCreated={document?.file?.created}
            documentClickable={true}
            showFile={() => showFile(document)}
          />
        {/each}
      {:else}
        <div class="no-documents-label">
          <span>Keine Dokumente gefunden.</span>
        </div>
      {/if}
    </div>
    <div class="pagination">
      <Pagination
        itemsMaxCount={files.maxCount}
        itemsPerPage={25}
        bind:selectedPage={filePageNumber}
      />
    </div>
  {/if}
{/await}

<style lang="sass">
  .header
    display: flex
    justify-content: space-between
    margin-bottom: 1rem

  .documents 
    display: flex
    flex-wrap: wrap
    align-items: stretch
    gap: 1rem

    .no-documents-label
      width: 100%
      text-align: center
      margin: 5rem
      color: var(--color-gray-70)
</style>
