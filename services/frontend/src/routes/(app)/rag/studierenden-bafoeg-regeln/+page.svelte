<script lang="ts">
  import { goto } from "$app/navigation";
  import { _ } from "svelte-i18n";
  import { getDocumentsStudierenden } from "$lib/api/rag/get-documents-studierenden";
  import type { RAGDocumentsSchülerStudierendenModelResponseType } from "$lib/api/rag/type/rag-documents-schueler-studierenden-model-type";
  import { uploadRAGrelevantDocumentsStudierenden } from "$lib/api/rag/upload-rag-relevant-documents-studierenden";
  import { Button } from "$lib/components/Button";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "$lib/components/DataTable";
  import { Dialog } from "$lib/components/Dialog";
  import { FilterTextInput } from "$lib/components/Inputs";
  import { FileInput } from "$lib/components/Inputs/file-input";
  import type { FileUploadFile } from "$lib/components/Inputs/file-input/types";
  import { Pagination } from "$lib/components/Pagination";
  import { dateToFormatStringLong } from "$lib/functions/date";
  import { userState } from "$lib/states/user";
  import { LLMChat } from "$lib/components/LLM-Chat";
  import { deleteRagDocumentsStudierendenByID } from "$lib/api/rag/delete-rag-documents-studierenden-by-id";

  let applicationPageNumber: number = $state(1);
  let fileUploadRefreshHash: string = $state("");
  let dialogUploadFileIsOpen: boolean = $state(false);
  let dialogUseRagIsOpen: boolean = $state(false);
  let filterSearchTerm: string = $state("");
  let dialogShowFileIsOpen: boolean = $state(false);
  let ragConversationID: string | undefined = $state(undefined);

  let showFileURL: string = $state("");
  let fileTypeIndicator: string | undefined = $state(undefined);

  let ragDocumentPromise:
    | Promise<RAGDocumentsSchülerStudierendenModelResponseType | undefined>
    | undefined = $state(undefined);

  async function reloadPage() {
    await loadData(Math.random().toString(36).substring(2, 15));
  }

  async function uploadFileFinished() {
    fileUploadRefreshHash = Math.random().toString(36).substring(2, 15);
    loadData(Math.random().toString(36).substring(2, 15));
    dialogUploadFileIsOpen = false;
  }

  async function uploadFile(file: FileUploadFile) {
    await uploadRAGrelevantDocumentsStudierenden(file.file);
  }

  function showFile(fileID: string, filetype: string) {
    showFileURL = `/api/v1/rag/bafoeg/studierenden/${fileID}`;
    fileTypeIndicator = filetype;
    dialogShowFileIsOpen = true;
  }

  async function deleteDocument(fileID: string) {
    const result = await deleteRagDocumentsStudierendenByID(fileID);
    if (!result) {
      alert("Fehler beim Löschen des Dokuments");
      return;
    }
    reloadPage();
  }

  async function loadData(hash: string = "") {
    ragDocumentPromise = getDocumentsStudierenden(filterSearchTerm);
  }

  $effect(() => {
    loadData();
  });
</script>

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

{#snippet useRagDialog()}
  <LLMChat ragDataSource={"studierenden"} bind:ragConversationID />
{/snippet}

{#snippet showFileDialog()}
  <div class="file-viewer" style="height: 100%;">
    {#if fileTypeIndicator === "application/pdf"}
      <iframe
        title="Requested Document"
        src={showFileURL}
        width="100%"
        height="100%"
        style="box-sizing: border-box;"
      ></iframe>
    {:else}
      <div>
        <p>Unknown File Type</p>
        <Button download onclick={() => goto(showFileURL)}>Download File</Button
        >
      </div>
    {/if}
  </div>
{/snippet}

<Dialog
  bind:isOpen={dialogShowFileIsOpen}
  dialogContent={showFileDialog}
  --dialog-width="80vw"
  --dialog-height="80vh"
/>

<Dialog bind:isOpen={dialogUploadFileIsOpen} dialogContent={uploadFileDialog} />

<Dialog
  bind:isOpen={dialogUseRagIsOpen}
  dialogContent={useRagDialog}
  --dialog-width="80vw"
  --dialog-minheight="0vh"
/>

<h1>Studierenden BAföG</h1>
<div class="actions">
  <FilterTextInput
    {filterSearchTerm}
    onFilterChange={(value) => (filterSearchTerm = value)}
  />
  {#if $userState?.permissions.includes("upload:rag-management-studierenden-files")}
    <Button onclick={() => (dialogUploadFileIsOpen = true)}
      >RAG-Dokument hochladen</Button
    >
  {/if}
  {#if $userState?.permissions.includes("upload:rag-management-studierenden-files")}
    <Button onclick={() => (dialogUseRagIsOpen = true)}>RAG testen</Button>
  {/if}
</div>

{#await ragDocumentPromise}
  <span>{$_("states.loading")}</span>
{:then ragDocuments}
  {#if ragDocuments != undefined && ragDocuments.maxCount === 0}
    <span>Keine Anträge gefunden</span>
  {:else if ragDocuments != undefined}
    <div class="datatable">
      <DataTable>
        <DataTableHead>
          <DataTableRow>
            <DataTableColumn>Dokumentenname</DataTableColumn>
            <DataTableColumn>Hochgeladen am</DataTableColumn>
            <DataTableColumn>Hochgeladen von</DataTableColumn>
            <DataTableColumn>Aktion</DataTableColumn>
          </DataTableRow>
        </DataTableHead>
        <DataTableBody>
          {#each ragDocuments.documents as document, i}
            <DataTableRow>
              <DataTableColumn>{document.fileName}</DataTableColumn>
              <DataTableColumn
                >{dateToFormatStringLong(document.created)}</DataTableColumn
              >
              <DataTableColumn
                >{document.createdFrom.displayName}</DataTableColumn
              >
              <DataTableColumn>
                <div class="actions">
                  <Button
                    onclick={() => showFile(document.id, document.fileType)}
                    >Ansehen</Button
                  >
                  <Button
                    buttonType="danger"
                    onclick={() => deleteDocument(document.id)}>Löschen</Button
                  >
                  <!-- <Button buttonType="danger">Löschen</Button> -->
                </div>
              </DataTableColumn>
            </DataTableRow>
          {/each}
        </DataTableBody>
      </DataTable>
    </div>

    <div class="pagination">
      <Pagination
        itemsMaxCount={ragDocuments.maxCount}
        itemsPerPage={25}
        bind:selectedPage={applicationPageNumber}
      />
    </div>
  {/if}
{/await}

<style lang="sass">
  .datatable
    margin-top: 2rem

  .actions
    display: flex
    justify-content: flex-end
    gap: 0.5rem
</style>
