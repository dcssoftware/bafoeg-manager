<script lang="ts">
  import { _ } from "svelte-i18n";
  import { IconHome, IconSchool, IconStudent } from "$lib/components/Icons";
  import { AddressCard } from "$lib/components/Address-Card";
  import type {
    ApplicationApplicantModelType,
    ApplicationSchoolModelType,
  } from "$lib/api/applications";
  import { Dialog } from "$lib/components/Dialog";
  import DatatableTable from "$lib/components/DataTable/datatable-table.svelte";
  import DatatableHead from "$lib/components/DataTable/datatable-head.svelte";
  import DatatableColumn from "$lib/components/DataTable/datatable-column.svelte";
  import DatatableBody from "$lib/components/DataTable/datatable-body.svelte";
  import DatatableRow from "$lib/components/DataTable/datatable-row.svelte";
  import { FilterTextInput } from "$lib/components/Inputs";
  import { Pagination } from "$lib/components/Pagination";
  import type { SchoolShortResponseModelType } from "$lib/api/schools/types/school-short-model-type";
  import { getSchools } from "$lib/api/schools/get-schools";
  import { getSchoolDegreesBySchoolID } from "$lib/api/schools/get-school-degrees-by-school-id";
  import type { SchoolDegreesReponseModelType } from "$lib/api/schools/types/school-degrees-model-type";
  import { generateHash } from "$lib/functions/random/gen-hash";
  import { updateApplicationAssignedSchoolDegree } from "$lib/api/applications/update-application-school-degree";
  import { page } from "$app/state";
  import { isApplicationUpdatableByApplicationID } from "$lib/api/applications/is-application-updatable-by-application-id";
  import { Button } from "$lib/components/Button";
  import { updateApplicationApplicant } from "$lib/api/applications/update-application-applicant";
  import { getApplicants, type ApplicantResponse } from "$lib/api/applicants";

  interface Props {
    applicant: ApplicationApplicantModelType;
    school: ApplicationSchoolModelType;
    allowWrap: boolean;
    refreshApplicationHash?: string | undefined;
  }

  let {
    applicant,
    school,
    allowWrap = false,
    refreshApplicationHash = $bindable(),
  }: Props = $props();

  let showSchoolsSelectorDialog: boolean = $state(false);
  let showApplicantsSelectorDialog: boolean = $state(false);
  let filterSchoolSearchTerm: string = $state("");
  let filterApplicantSearchTerm: string = $state("");
  let selectedSchoolId: string | undefined = $state(undefined);
  let selectableSchoolPageNumber: number = $state(1);
  let selectableApplicantPageNumber: number = $state(1);
  let selectableSchoolDegreePageNumber: number = $state(1);
  const applicationID = page.params.applicationID ?? "";
  let applicationIsUpdatable: boolean = $state(false);

  let schoolDegreesDataPromise:
    | Promise<SchoolDegreesReponseModelType | undefined>
    | undefined = $state(undefined);

  let schoolDataPromise:
    | Promise<SchoolShortResponseModelType | undefined>
    | undefined = $state(undefined);

  let applicantsDataPromise:
    | Promise<ApplicantResponse | undefined>
    | undefined = $state(undefined);

  async function changeApplicant(applicantID: string) {
    await updateApplicationApplicant(applicationID, applicantID);
    refreshApplicationHash = generateHash(8);
  }

  async function changeSchoolDegree(degreeID: string) {
    await updateApplicationAssignedSchoolDegree(applicationID, degreeID);
    refreshApplicationHash = generateHash(8);
  }

  function changeSchoolSetSchoolID(schoolID: string) {
    filterSchoolSearchTerm = "";
    selectedSchoolId = schoolID;
  }

  async function loadData(
    filterSchoolSearchTerm: string,
    filterApplicantSearchTerm: string,
  ) {
    schoolDataPromise = getSchools(
      selectableSchoolPageNumber,
      filterSchoolSearchTerm,
    );

    if (selectedSchoolId !== undefined) {
      schoolDegreesDataPromise = getSchoolDegreesBySchoolID(
        selectableSchoolDegreePageNumber,
        selectedSchoolId,
      );
    }

    applicantsDataPromise = getApplicants(
      selectableApplicantPageNumber,
      filterApplicantSearchTerm,
    );

    applicationIsUpdatable =
      await isApplicationUpdatableByApplicationID(applicationID);
  }

  $effect(() => {
    loadData(filterSchoolSearchTerm, filterApplicantSearchTerm);
  });
</script>

<Dialog
  bind:isOpen={showSchoolsSelectorDialog}
  dialogContent={SchoolSelectionDialog}
/>

<Dialog
  bind:isOpen={showApplicantsSelectorDialog}
  dialogContent={ApplicantSelectionDialog}
/>

{#snippet SchoolSelectionDialog()}
  {#key showSchoolsSelectorDialog}
    {#if selectedSchoolId === undefined}
      <div class="school-selection">
        <div class="header">
          <h2>Schule auswählen</h2>
        </div>
        <div class="search">
          <FilterTextInput
            onFilterChange={(value) => (filterSchoolSearchTerm = value)}
            placeholder="Schule suchen..."
          />
        </div>
        <br />
        {#await schoolDataPromise then schoolData}
          <div class="select">
            <DatatableTable>
              <DatatableHead>
                <DatatableColumn>#</DatatableColumn>
                <DatatableColumn>Name</DatatableColumn>
                <DatatableColumn>Typ</DatatableColumn>
                <DatatableColumn>Stadt</DatatableColumn>
              </DatatableHead>
              <DatatableBody>
                {#each schoolData?.schools as school, i}
                  <DatatableRow
                    onClick={() => changeSchoolSetSchoolID(school.id)}
                  >
                    <DatatableColumn>{i + 1}</DatatableColumn>
                    <DatatableColumn>{school.name}</DatatableColumn>
                    <DatatableColumn>{school.type.name}</DatatableColumn>
                    <DatatableColumn
                      >{school.city} ({school.country})</DatatableColumn
                    >
                  </DatatableRow>
                {/each}
              </DatatableBody>
            </DatatableTable>
          </div>
          <div class="pagination">
            <Pagination
              itemsPerPage={25}
              itemsMaxCount={schoolData?.maxCount}
              bind:selectedPage={selectableSchoolPageNumber}
            />
          </div>
        {/await}
        <div>
          <Button
            onclick={() => {
              selectedSchoolId = undefined;
            }}>Zurück</Button
          >
        </div>
      </div>
    {:else}
      <div class="degree select">
        <div class="header">
          <h2>Abschluss auswählen</h2>
        </div>
        {#await schoolDegreesDataPromise then schoolDegrees}
          <div class="select">
            <DatatableTable>
              <DatatableHead>
                <DatatableColumn>#</DatatableColumn>
                <DatatableColumn>Abschluss</DatatableColumn>
                <DatatableColumn>FOS Berufsabschluss Benötigt</DatatableColumn>
                <DatatableColumn
                  >BOS Berufsqualifizierender Abschluss</DatatableColumn
                >
                <DatatableColumn
                  >Fachschule Berufsschule Benötigt</DatatableColumn
                >
              </DatatableHead>
              <DatatableBody>
                {#each schoolDegrees?.degrees as degree, i}
                  <DatatableRow onClick={() => changeSchoolDegree(degree.id)}>
                    <DatatableColumn>{i + 1}</DatatableColumn>
                    <DatatableColumn>{degree.name}</DatatableColumn>
                    <DatatableColumn
                      >{degree.fachschuleBerufsschuleRequired
                        ? "✅"
                        : "❌"}</DatatableColumn
                    >
                    <DatatableColumn
                      >{degree.bosBerufsqualifizierenderAbschluss
                        ? "✅"
                        : "❌"}</DatatableColumn
                    >
                    <DatatableColumn
                      >{degree.fachschuleBerufsschuleRequired
                        ? "✅"
                        : "❌"}</DatatableColumn
                    >
                  </DatatableRow>
                {/each}
              </DatatableBody>
            </DatatableTable>
          </div>
          <div class="pagination">
            <Pagination
              itemsPerPage={25}
              itemsMaxCount={schoolDegrees?.maxCount}
              bind:selectedPage={selectableSchoolDegreePageNumber}
            />
          </div>
        {/await}
      </div>
    {/if}
  {/key}
{/snippet}

{#snippet ApplicantSelectionDialog()}
  {#key showApplicantsSelectorDialog}
    <div>
      <div class="header">
        <h2>Antragsteller auswählen</h2>
      </div>
      <div class="search">
        <FilterTextInput
          onFilterChange={(value) => (filterApplicantSearchTerm = value)}
          placeholder="Antragsteller suchen..."
        />
      </div>
      <br />
      {#await applicantsDataPromise then applicantsData}
        <div class="select">
          <DatatableTable>
            <DatatableHead>
              <DatatableColumn>#</DatatableColumn>
              <DatatableColumn>Vorname</DatatableColumn>
              <DatatableColumn>Nachname</DatatableColumn>
            </DatatableHead>
            <DatatableBody>
              {#each applicantsData?.applicants as applicant, i}
                <DatatableRow onClick={() => changeApplicant(applicant.id)}>
                  <DatatableColumn>{i + 1}</DatatableColumn>
                  <DatatableColumn>{applicant.firstname}</DatatableColumn>
                  <DatatableColumn>{applicant.lastname}</DatatableColumn>
                </DatatableRow>
              {/each}
            </DatatableBody>
          </DatatableTable>
        </div>
        <div class="pagination">
          <Pagination
            itemsPerPage={25}
            itemsMaxCount={applicantsData?.maxCount}
            bind:selectedPage={selectableApplicantPageNumber}
          />
        </div>
      {/await}
    </div>
  {/key}
{/snippet}

<div class="address-cards address-cards-{allowWrap ? 'wrap' : 'nowrap'}">
  <AddressCard
    IconComponent={IconHome}
    data={{
      firstname: applicant.firstname,
      lastname: applicant.lastname,
      street: applicant.address.street,
      houseNumber: applicant.address.houseNumber,
      postalCode: applicant.address.zipCode,
      city: applicant.address.city,
      country: applicant.address.country,
    }}
    isUpdatable={applicationIsUpdatable}
    externalLink="/applications/applicants/{applicant.id}"
    header={$_("page.application.editor.basic-infos.address.permanent")}
    onChange={() => {
      showApplicantsSelectorDialog = true;
    }}
  />
  {#if applicant.trainingsAddress != null}
    <AddressCard
      IconComponent={IconStudent}
      data={{
        firstname: applicant.firstname,
        lastname: applicant.lastname,
        street: applicant.trainingsAddress.street,
        houseNumber: applicant.trainingsAddress.houseNumber,
        postalCode: applicant.trainingsAddress.zipCode,
        city: applicant.trainingsAddress.city,
        country: applicant.trainingsAddress.country,
      }}
      header={$_("page.application.editor.basic-infos.address.trainings")}
    />
  {/if}
  <AddressCard
    IconComponent={IconSchool}
    data={{
      name: school.name,
      street: school.address.street,
      houseNumber: school.address.houseNumber,
      postalCode: school.address.zipCode,
      city: school.address.city,
      country: school.address.country,
    }}
    isUpdatable={applicationIsUpdatable}
    onChange={() => {
      showSchoolsSelectorDialog = true;
    }}
    externalLink="/school-management/{school.id}"
    header={$_("page.application.editor.basic-infos.address.school")}
  />
</div>

<style lang="sass">
  .address-cards
    display: grid
    grid-template-columns: 1fr 1fr 1fr
    gap: 2rem
    .student-education-address, .student-permanent-address, .school-address
      background-color: var(--background-color-tertiary)
      flex-grow: 1
      align-items: center
      padding: 2rem 3rem
      display: flex
      gap: 2rem
      .form-icon
        :global(svg)
          $icon-size: 50px
          fill: var(--font-color)
          height: $icon-size
          width: $icon-size
  @media (min-width: 1800px) 
   .address-cards-wrap
    grid-template-columns: 1fr 1fr 1fr
  @media (min-width: 2500px) 
   .address-cards-wrap
    grid-template-columns: 1fr

</style>
