<script lang="ts">
  import { goto } from "$app/navigation";
  import { getEakten } from "$lib/api/eakten";
  import type { EakteHttpResponse } from "$lib/api/eakten/types";
  import { uploadEaktenFiles } from "$lib/api/eakten/upload-file";
  import { Button } from "$lib/components/Button";
  import {
    DataTable,
    DataTableBody,
    DataTableColumn,
    DataTableHead,
    DataTableRow,
  } from "$lib/components/DataTable";
  import { Dialog } from "$lib/components/Dialog";
  import { FileInput } from "$lib/components/Inputs/file-input";
  import type { FileUploadFile } from "$lib/components/Inputs/file-input/types";
  import { Pagination } from "$lib/components/Pagination";
  import { dateToFormatStringLong } from "$lib/functions/date";
  import { _ } from "svelte-i18n";

  let dialogUploadFileIsOpen: boolean = $state(false);
  let fileUploadRefreshHash: string = $state("");
  let applicationPageNumber: number = $state(1);
  let eaktenPromise: Promise<EakteHttpResponse | undefined> | undefined =
    $state(undefined);

  async function uploadFile(file: FileUploadFile) {
    const result = await uploadEaktenFiles(file.file);
  }

  async function uploadFileFinished() {
    fileUploadRefreshHash = Math.random().toString(36).substring(2, 15);
    dialogUploadFileIsOpen = false;
  }

  async function loadData() {
    eaktenPromise = getEakten(applicationPageNumber);
  }

  $effect(() => {
    loadData();
  });
</script>

<Dialog bind:isOpen={dialogUploadFileIsOpen} dialogContent={uploadFileDialog} />

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
  <h1>E-Akte</h1>
  <div class="actions">
    <Button onclick={() => (dialogUploadFileIsOpen = true)}
      >E-Akte hochladen</Button
    >
  </div>
</div>

{#await eaktenPromise}
  <span>Load...</span>
{:then eakten}
  <div class="datatable">
    <DataTable>
      <DataTableHead>
        <DataTableRow>
          <DataTableColumn>Count</DataTableColumn>
          <DataTableColumn>ID</DataTableColumn>
          <DataTableColumn>Aktenzeichen</DataTableColumn>
          <DataTableColumn>Type</DataTableColumn>
          <DataTableColumn>Vertraulichkeit</DataTableColumn>
          <DataTableColumn>Datum</DataTableColumn>
          <DataTableColumn>Auto-Mapping</DataTableColumn>
        </DataTableRow>
      </DataTableHead>
      <DataTableBody>
        {#each eakten?.eakten as eakte, i}
          {@const counterMin =
            (applicationPageNumber - 1) * (eakten?.count ?? 0)}
          <DataTableRow onClick={() => goto("/eakte/" + eakte?.id)}>
            <DataTableColumn align="Right"
              >{counterMin + (i + 1)}</DataTableColumn
            >
            <DataTableColumn title={eakte?.id}>{eakte?.id}</DataTableColumn>
            <DataTableColumn>{eakte?.aktenzeichen}</DataTableColumn>
            <DataTableColumn>{eakte?.type.identifier}</DataTableColumn>
            <DataTableColumn>{eakte?.vertraulichkeit}</DataTableColumn>
            <DataTableColumn
              >{dateToFormatStringLong(eakte.created)}</DataTableColumn
            >
            <DataTableColumn></DataTableColumn>
          </DataTableRow>
        {/each}
      </DataTableBody>
    </DataTable>

    <div class="pagination">
      <Pagination
        itemsMaxCount={eakten?.maxCount}
        itemsPerPage={eakten?.count}
        bind:selectedPage={applicationPageNumber}
      />
    </div>
  </div>
{/await}

<style lang="sass">
  .header
    display: flex
    justify-content: space-between
    align-items: center
</style>
