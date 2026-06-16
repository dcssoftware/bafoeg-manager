<script lang="ts">
  import {
    GetDocumentFilesByAkteID,
    GetDocumentFilesByVorgangsID,
    getVorgängeByEaktenID,
  } from "$lib/api/eakten";
  import type {
    DocumentsHttpResponse,
    EakteModel,
    VorgangHttpResponse,
  } from "$lib/api/eakten/types";
  import { DocumentIndicator } from "$lib/components/Document-Indicator";
  import {
    IconPen,
    IconSchool,
    IconStudent,
  } from "$lib/components/Icons";
  import { Pagination } from "$lib/components/Pagination";
  import { dateToFormatStringShort } from "$lib/functions/date";
  import { onMount } from "svelte";
  import { Button } from "$lib/components/Button";
  import { Dialog } from "$lib/components/Dialog";
  import { ApplicantSelect } from "$lib/components/ApplicantSelect";
  import { SchoolSelect } from "$lib/components/SchoolSelect";
  import DatatableTable from "$lib/components/DataTable/datatable-table.svelte";
  import DatatableHead from "$lib/components/DataTable/datatable-head.svelte";
  import { DataTableRow, DataTableColumn } from "$lib/components/DataTable";
  import DatatableBody from "$lib/components/DataTable/datatable-body.svelte";
  import { getApplicantByID } from "$lib/api/applicants";
  import { getSchoolByID } from "$lib/api/schools/get-school-by-id";
  import { SchoolDegreeSelect } from "$lib/components/SchoolDegreeSelect";
  import { createApplicationFromEakte } from "$lib/api/applications";
  import { page } from "$app/state";
  import { getEaktenApplicationMapping } from "$lib/api/eakten/get-eakten-application-mapping";
  import { FileViewer } from "$lib/components/FileViewer";
  import { AddressCard } from "$lib/components/Address-Card";
  import { _ } from "svelte-i18n";
  import { getEakteByID } from "$lib/api/eakten/get-eakte-by-id";

  interface CreateAntragType {
    applicantID?: string;
    schoolID?: string;
    degreeID?: string;
    className?: string;
    schoolyearStart?: Date;
    schoolyearEnd?: Date;
  }

  interface showFileData {
    showFileURL: string;
    fileTypeIndicator: string;
  }

  let filterVorgang: string = $state("");
  const eaktenID = page.params.eaktenID ?? "";
  let vorgängePromise: Promise<VorgangHttpResponse[] | undefined> | undefined =
    $state(undefined);
  let eaktePromise: Promise<EakteModel | undefined> | undefined =
    $state(undefined);
  let documentFilesPromise:
    | Promise<DocumentsHttpResponse | undefined>
    | undefined = $state(undefined);
  let eaktenApplicationMappingPromise:
    | Promise<EaktenApplicationMappingModel | null>
    | undefined = $state(undefined);
  let dialogShowFileIsOpen: boolean = $state(false);
  let classNameInput: string = $state("");
  let schoolyearStartInput: string = $state("");
  let schoolyearEndInput: string = $state("");
  let showFileData: showFileData | undefined = $state(undefined);

  let isNewAntragDialogOpen: boolean = $state(false);
  let dialogNewApplicationPage: number = $state(1);
  let createAntrag: CreateAntragType = $state({});

  function setVorgangFilter(value: string | null) {
    const url = new URL(window.location.href);

    if (value !== null && value !== "") {
      url.searchParams.set("vorgang", value);
      filterVorgang = value;
    } else {
      url.searchParams.delete("vorgang");
      filterVorgang = "";
    }

    window.history.replaceState({}, "", url.toString());
  }

  function newApplicationSetApplicant(id: string) {
    createAntrag.applicantID = id;
    dialogNewApplicationPage++;
  }

  function newApplicationSetSchool(id: string) {
    createAntrag.schoolID = id;
    dialogNewApplicationPage++;
  }

  function newApplicationSetDegree(id: string) {
    createAntrag.degreeID = id;
    dialogNewApplicationPage++;
  }

  function newApplicationSetMetaData() {
    createAntrag.className = classNameInput;
    createAntrag.schoolyearStart = new Date(schoolyearStartInput);
    createAntrag.schoolyearEnd = new Date(schoolyearEndInput);
    dialogNewApplicationPage++;
  }

  async function submitNewApplication() {
    let success = await createApplicationFromEakte(
      createAntrag.applicantID ?? "",
      eaktenID,
      createAntrag.degreeID ?? "",
      createAntrag.schoolyearStart ?? new Date(),
      createAntrag.schoolyearEnd ?? new Date(),
      createAntrag.className ?? "",
      [],
    );

    if (!success) {
      // alert("Antrag konnte nicht erstellt werden. Internal error.");
      return;
    }

    closeNewAntragDialog();
  }

  function showFile(documentID: string, indicator: string) {
    showFileData = {
      showFileURL: `/api/v1/eakten/documents/${documentID}`,
      fileTypeIndicator: indicator,
    };
  }

  async function loadData() {
    const url = new URL(window.location.href);
    const urlFilterVorgangValue = url.searchParams.get("vorgang");
    if (urlFilterVorgangValue != null || urlFilterVorgangValue === "") {
      filterVorgang = urlFilterVorgangValue;
    }

    eaktenApplicationMappingPromise = getEaktenApplicationMapping(eaktenID);
    eaktePromise = getEakteByID(eaktenID);
    vorgängePromise = getVorgängeByEaktenID(eaktenID);
  }

  async function loadVorgangFiles() {
    if (filterVorgang === "" || filterVorgang == null) {
      documentFilesPromise = GetDocumentFilesByAkteID(eaktenID);
    } else {
      documentFilesPromise = GetDocumentFilesByVorgangsID(
        eaktenID,
        filterVorgang,
      );
    }
  }

  function closeNewAntragDialog() {
    isNewAntragDialogOpen = false;
    dialogNewApplicationPage = 1;
    createAntrag = {};
  }

  function closeShowFileDialog() {
    dialogShowFileIsOpen = false;
  }

  $effect(() => {
    loadVorgangFiles();
  });

  onMount(() => loadData());
</script>

<Dialog
  bind:isOpen={isNewAntragDialogOpen}
  dialogContent={createAntragView}
  onClose={() => closeNewAntragDialog()}
/>

<Dialog
  bind:isOpen={dialogShowFileIsOpen}
  dialogContent={showFileDialog}
  --dialog-width="80vw"
  --dialog-height="80vh"
/>

{#snippet showFileDialog()}
  <FileViewer
    showFileURL={showFileData?.showFileURL ?? ""}
    fileTypeIndicator={showFileData?.fileTypeIndicator}
    closeDialog={closeShowFileDialog}
  />
{/snippet}

{#snippet createAntragView()}
  <h2>Antrag erstellen</h2>

  {#if dialogNewApplicationPage === 1}
    <div class="antragsteller">
      <div>
        <Button disabled>Neuer Antragsteller</Button>
      </div>
      <div>
        <ApplicantSelect
          onSingleSelect={(id: string) => newApplicationSetApplicant(id)}
        />
      </div>
    </div>
  {:else if dialogNewApplicationPage === 2}
    <div class="school">
      <div>
        <SchoolSelect
          onSingleSelect={(id: string) => newApplicationSetSchool(id)}
        />
      </div>
    </div>
  {:else if dialogNewApplicationPage === 3}
    <div class="degree">
      <div>
        <SchoolDegreeSelect
          schoolID={createAntrag.schoolID ?? ""}
          onSingleSelect={(id: string) => newApplicationSetDegree(id)}
        />
      </div>
    </div>
  {:else if dialogNewApplicationPage === 4}
    <div>
      <table>
        <tbody>
          <tr>
            <td>Klasse</td>
            <td><input bind:value={classNameInput} type="text" /></td>
          </tr>
          <tr>
            <td>Schuljahr Start</td>
            <td><input bind:value={schoolyearStartInput} type="date" /></td>
          </tr>
          <tr>
            <td>Schuljahr Ende</td>
            <td><input bind:value={schoolyearEndInput} type="date" /></td>
          </tr>
        </tbody>
      </table>
      <div>
        <Button onclick={() => newApplicationSetMetaData()}>Weiter</Button>
      </div>
    </div>
  {:else if dialogNewApplicationPage === 5}
    <div>
      <div>
        <DatatableTable>
          <DatatableHead>
            <DataTableRow>
              <DataTableColumn>Feld</DataTableColumn>
              <DataTableColumn>Wert</DataTableColumn>
            </DataTableRow>
          </DatatableHead>
          <DatatableBody>
            {#await getApplicantByID(createAntrag.applicantID ?? "") then antragsteller}
              <DataTableRow>
                <DataTableColumn>Antragsteller</DataTableColumn>
                <DataTableColumn
                  >{antragsteller?.firstname}
                  {antragsteller?.lastname}</DataTableColumn
                >
              </DataTableRow>
            {/await}
            {#await getSchoolByID(createAntrag.schoolID ?? "") then schule}
              <DataTableRow>
                <DataTableColumn>Schule</DataTableColumn>
                <DataTableColumn>{schule?.name}</DataTableColumn>
              </DataTableRow>
              <DataTableRow>
                <DataTableColumn>Abschluss</DataTableColumn>
                <DataTableColumn
                  >{schule?.degree.find((d) => d.id === createAntrag.degreeID)
                    ?.name}</DataTableColumn
                >
              </DataTableRow>
            {/await}
            <DataTableRow>
              <DataTableColumn>Klasse</DataTableColumn>
              <DataTableColumn>{createAntrag.className}</DataTableColumn>
            </DataTableRow>
            <DataTableRow>
              <DataTableColumn>Schuljahr Start</DataTableColumn>
              <DataTableColumn
                >{createAntrag.schoolyearStart?.toUTCString()}</DataTableColumn
              >
            </DataTableRow>
            <DataTableRow>
              <DataTableColumn>Schuljahr Ende</DataTableColumn>
              <DataTableColumn
                >{createAntrag.schoolyearEnd?.toUTCString()}</DataTableColumn
              >
            </DataTableRow>
          </DatatableBody>
        </DatatableTable>
      </div>
      <div>
        <Button onclick={submitNewApplication}>Antrag erstellen</Button>
      </div>
    </div>
  {:else}
    <div class="error">
      <span>Unbekannte Seite im Dialog.</span>
    </div>
  {/if}

  <div>
    <Button onclick={() => dialogNewApplicationPage--}>Zurück</Button>
  </div>
{/snippet}

{#await eaktePromise then eakte}
  <h1>Mein Data Import</h1>
  <div class="data-import">
    {#await eaktenApplicationMappingPromise then eaktenmapping}
      {#if eaktenmapping === null || eaktenmapping === undefined}
        <div>
          <Button onclick={() => (isNewAntragDialogOpen = true)}
            >Neuen Antrag erstellen</Button
          >
        </div>
      {:else}
        <div class="application-mapping">
          <div class="mapper">
            <div class="found-application">
              <AddressCard
                IconComponent={IconPen}
                data={{
                  firstname: eaktenmapping.applicant.firstname,
                  lastname: eaktenmapping.applicant.lastname,
                }}
                externalLink={`/applications/${eaktenmapping.application.id}`}
                header={$_(
                  "page.application.editor.basic-infos.address.permanent",
                )}
              />
              <AddressCard
                IconComponent={IconStudent}
                data={{
                  firstname: eaktenmapping.applicant.firstname,
                  lastname: eaktenmapping.applicant.lastname,
                }}
                externalLink="/applications/applicants/{eaktenmapping.applicant
                  .id}"
                header={$_(
                  "page.application.editor.basic-infos.address.permanent",
                )}
              />
              <AddressCard
                IconComponent={IconSchool}
                data={{
                  name: eaktenmapping.school.name,
                }}
                externalLink="/school-management/{eaktenmapping.school.id}"
                header={$_(
                  "page.application.editor.basic-infos.address.school",
                )}
              />
            </div>
          </div>
        </div>
      {/if}
    {/await}
    <div class="import">
      <div class="selector">
        <!-- svelte-ignore a11y_missing_attribute -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <a
          onclick={() => {
            setVorgangFilter(null);
          }}
          class="item eakte"
        >
          <span> E-Akte: </span>
          <span> {eakte?.aktenzeichen} </span>
        </a>
        <!-- svelte-ignore a11y_missing_attribute -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        {#await vorgängePromise then vorgänge}
          {#each vorgänge as vorgang}
            {@const isSelected = filterVorgang === vorgang.id}
            <a
              onclick={() => {
                setVorgangFilter(vorgang.id);
              }}
              class="item vorgang"
            >
              <span>
                {#if isSelected}
                  👉🏻
                {/if}
                Vorgang {dateToFormatStringShort(vorgang.created)}:
              </span>
              <span> {vorgang.vorgangszeichen} </span>
            </a>
          {/each}
        {/await}
      </div>
      <div class="files">
        <h2>Importierte Dateien</h2>

        {#await documentFilesPromise then documentFiles}
          {#if documentFiles !== undefined}
            <div class="documents">
              <div class="files">
                {#if documentFiles.maxCount > 0}
                  {#each documentFiles.documents as file, index}
                    <DocumentIndicator
                      documentID="1"
                      documentName={file.files.name}
                      documentType={file.files.type}
                      documentSize={file.files.size}
                      documentCreated={file.files.created}
                      documentClickable={true}
                      showFile={() => {
                        showFile(`${file.id}`, file.files.type);
                        dialogShowFileIsOpen = true;
                      }}
                    />
                  {/each}
                {:else}
                  <span>Keine Dateien gefunden.</span>
                {/if}
              </div>
              <Pagination
                itemsPerPage={25}
                itemsMaxCount={documentFiles.maxCount}
                selectedPage={1}
              />
            </div>
          {/if}
        {/await}
      </div>
    </div>
  </div>
{/await}

<style lang="sass">

  .data-import
    display: flex
    flex-direction: column
    gap: 2rem

    .application-mapping
      .mapper
        .found-application
          display: grid 
          grid-template-columns: 1fr 1fr 1fr
          gap: 1rem
          margin-bottom: 1rem
          .user
          .application
          :global(svg)
            fill: var(--font-color)

    .import
      display: flex
      gap: 2rem

      .selector
        background-color: var(--background-color-tertiary)
        padding: 1rem
        display: flex
        flex-direction: column
        gap: 1rem
        min-width: 300px

        .item
          padding: 0.5rem
          color: var(--font-color-white)
          cursor: pointer
          display: flex
          flex-direction: column
          gap: 0.5rem
          user-select: none

          span:first-child
            font-weight: bold
            font-size: 1.2rem


          &.eakte
            background-color: var(--color-blue)

          &.vorgang
            background-color: var(--color-green)
            margin-left: 2rem

      .files
        background-color: var(--background-color-tertiary)
        padding: 1rem
        flex-grow: 1

        .documents
          .files 
            display: flex
            flex-wrap: wrap
            align-items: stretch
            gap: 1rem


</style>
