<script lang="ts">
  import { _ } from "svelte-i18n";
  import { page } from "$app/state";
  import type { SchoolModelType } from "$lib/api/schools";
  import { getSchoolByID } from "$lib/api/schools/get-school-by-id";
  import { AddressCard } from "$lib/components/Address-Card";
  import { IconSchool } from "$lib/components/Icons";
  import {
    DataTable,
    DataTableBody,
    DataTableColumn,
    DataTableHead,
    DataTableRow,
  } from "$lib/components/DataTable";
  import { goto } from "$app/navigation";
  import type { ApplicantsBySchool } from "$lib/api/applicants/types/applicant-by-school";
  import { Pagination } from "$lib/components/Pagination";
  import { getApplicantsBySchoolID } from "$lib/api/applicants";
  import { error } from "@sveltejs/kit";

  let schoolDataPromise: Promise<SchoolModelType | undefined> | undefined =
    $state(undefined);
  let schoolApplicantsPromise:
    | Promise<ApplicantsBySchool | undefined>
    | undefined = $state(undefined);
  let applicantsPageNumber: number = $state(1);

  function loadData() {
    // if no schoolID is provided, throw 404
    if (!page.params.schoolID) {
      throw error(404, "School not found");
    }

    schoolDataPromise = getSchoolByID(page.params.schoolID);
    schoolApplicantsPromise = getApplicantsBySchoolID(
      applicantsPageNumber,
      page.params.schoolID,
      true
    );
  }

  $effect(() => {
    loadData();
  });
</script>

<div class="header">
  <h1>{$_("page.school-management.school.header")}</h1>
</div>

{#await schoolDataPromise}
  <span>{$_("states.loading")}</span>
{:then school}
  {#if school !== undefined}
    <div class="school-data">
      <div>
        <AddressCard
          IconComponent={IconSchool}
          data={{
            name: school.name,
            street: school.street,
            houseNumber: school.houseNumber,
            postalCode: school.zipCode,
            city: school.city,
            country: school.country,
          }}
          header={$_("page.application.editor.basic-infos.address.school")}
        />
      </div>

      <div>
        <h2>{$_("page.school-management.school.degree-courses")}</h2>
        <div class="datatable">
          {#if true}
            {@const i18npath =
              "page.school-management.school.datatables.degree-courses"}
            <DataTable>
              <DataTableHead>
                <DataTableRow>
                  <DataTableColumn>{$_(`${i18npath}.nr`)}</DataTableColumn>
                  <DataTableColumn
                    >{$_(`${i18npath}.degress-course`)}</DataTableColumn
                  >
                  <DataTableColumn
                    >{$_(`${i18npath}.fos-required`)}</DataTableColumn
                  >
                  <DataTableColumn
                    >{$_(`${i18npath}.bos-qualifying`)}</DataTableColumn
                  >
                  <DataTableColumn
                    >{$_(`${i18npath}.work-degree-required`)}</DataTableColumn
                  >
                </DataTableRow>
              </DataTableHead>
              <DataTableBody>
                {#each school.degree as degree, i}
                  {@const number = i + 1}
                  <DataTableRow>
                    <DataTableColumn>{number}</DataTableColumn>
                    <DataTableColumn>{degree.name}</DataTableColumn>
                    <DataTableColumn align="Center"
                      >{degree.fosBerufsabschlussRequired
                        ? "✅"
                        : "❌"}</DataTableColumn
                    >
                    <DataTableColumn align="Center"
                      >{degree.bosBerufsqualifizierenderAbschluss
                        ? "✅"
                        : "❌"}</DataTableColumn
                    >
                    <DataTableColumn align="Center"
                      >{degree.fachschuleBerufsschuleRequired
                        ? "✅"
                        : "❌"}</DataTableColumn
                    >
                  </DataTableRow>
                {/each}
              </DataTableBody>
            </DataTable>
          {/if}
        </div>
      </div>

      <div class="">
        <h2>
          {$_(
            "page.school-management.school.current-students-with-application"
          )}
        </h2>
        {#await schoolApplicantsPromise then schoolApplicants}
          {#if schoolApplicants !== undefined}
            {@const i18npath =
              "page.school-management.school.datatables.current-students"}
            <div class="datatable">
              <DataTable>
                <DataTableHead>
                  <DataTableRow>
                    <DataTableColumn>{$_(`${i18npath}.nr`)}</DataTableColumn>
                    <DataTableColumn>{$_(`${i18npath}.id`)}</DataTableColumn>
                    <DataTableColumn>{$_(`${i18npath}.name`)}</DataTableColumn>
                    <DataTableColumn>{$_(`${i18npath}.city`)}</DataTableColumn>
                    <DataTableColumn>{$_(`${i18npath}.degree`)}</DataTableColumn
                    >
                    <DataTableColumn>{$_(`${i18npath}.class`)}</DataTableColumn>
                    <DataTableColumn>{$_(`${i18npath}.status`)}</DataTableColumn
                    >
                  </DataTableRow>
                </DataTableHead>
                <DataTableBody>
                  {#each schoolApplicants.applicants as applicant, i}
                    {@const number = i + 1}
                    {@const schoolID = applicant.id.split("-")[0]}
                    <DataTableRow
                      onClick={() =>
                        goto("/applications/applicants/" + applicant.id)}
                    >
                      <DataTableColumn>{number}</DataTableColumn>
                      <DataTableColumn>{schoolID}</DataTableColumn>
                      <DataTableColumn
                        >{applicant.firstname}
                        {applicant.lastname}</DataTableColumn
                      >
                      <DataTableColumn
                        >{applicant.address.zipCode}
                        {applicant.address.city}</DataTableColumn
                      >
                      <DataTableColumn>{applicant.degree.name}</DataTableColumn>
                      <DataTableColumn>{applicant.classLevel}</DataTableColumn>
                      <DataTableColumn
                        >{$_(
                          `states.application-status.${applicant.statusIdentifier}`
                        )}</DataTableColumn
                      >
                    </DataTableRow>
                  {/each}
                </DataTableBody>
              </DataTable>
            </div>
            <div class="pagination">
              <Pagination
                itemsMaxCount={schoolApplicants.maxCount}
                itemsPerPage={25}
                bind:selectedPage={applicantsPageNumber}
              />
            </div>
          {/if}
        {/await}
      </div>
    </div>
  {:else}
    <span>{$_("page.school-management.school-not-found")}</span>
  {/if}
{/await}

<style lang="sass">
  .school-data
    display: flex;
    flex-direction: column;
    gap: 2rem;

</style>
