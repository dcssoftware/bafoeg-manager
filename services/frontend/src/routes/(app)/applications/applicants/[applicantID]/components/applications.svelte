<script lang="ts">
  import { _ } from "svelte-i18n";
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import {
    DataTable,
    DataTableBody,
    DataTableColumn,
    DataTableHead,
    DataTableRow,
  } from "$lib/components/DataTable";
  import { Pagination } from "$lib/components/Pagination";
  import { getApplicationsByApplicantID } from "$lib/api/applications/get-applications-by-applicant-id";
  import type { ApplicationModelResponseType } from "$lib/api/applications/types/application-model-type";

  let applicationsPromise:
    | Promise<ApplicationModelResponseType | undefined>
    | undefined = $state(undefined);

  let applicationPageNumber: number = $state(1);

  async function loadData(pageNumber: number) {
    applicationsPromise = getApplicationsByApplicantID(
      pageNumber,
      page.params.applicantID
    );
  }

  $effect(() => {
    loadData(applicationPageNumber);
  });
</script>

<h2>Gestellte Anträge</h2>

{#await applicationsPromise then applications}
  {#if applications !== undefined}
    <DataTable>
      <DataTableHead>
        <DataTableRow>
          <DataTableColumn
            >{$_("page.applicant.application.datatable.count")}</DataTableColumn
          >
          <DataTableColumn
            >{$_(
              "page.applicant.application.datatable.school"
            )}</DataTableColumn
          >
          <DataTableColumn
            >{$_(
              "page.applicant.application.datatable.degree"
            )}</DataTableColumn
          >
          <DataTableColumn
            >{$_(
              "page.applicant.application.datatable.class-level"
            )}</DataTableColumn
          >
          <DataTableColumn
            >{$_(
              "page.applicant.application.datatable.assigned-user"
            )}</DataTableColumn
          >
          <DataTableColumn
            >{$_(
              "page.applicant.application.datatable.status"
            )}</DataTableColumn
          >
        </DataTableRow>
      </DataTableHead>
      <DataTableBody>
        {#each applications.application as row, i}
          {@const count = applicationPageNumber * 10 - 10 + (i + 1)}
          <DataTableRow onClick={() => goto(`/applications/${row.id}`)}>
            <DataTableColumn>{count}</DataTableColumn>
            <DataTableColumn>{row.school.name}</DataTableColumn>
            <DataTableColumn>{row.school.degree.name}</DataTableColumn>
            <DataTableColumn>{row.classLevel}</DataTableColumn>
            <DataTableColumn>
              {#if row.assignedUser !== null}
                {row.assignedUser.displayName}
              {:else}
                {$_("components.user-picker.no-user-assigned")}
              {/if}
            </DataTableColumn>
            <DataTableColumn>{row.status.name}</DataTableColumn>
          </DataTableRow>
        {/each}
      </DataTableBody>
    </DataTable>
    <div class="pagination">
      <Pagination
        itemsMaxCount={applications.maxCount}
        itemsPerPage={25}
        bind:selectedPage={applicationPageNumber}
      />
    </div>
  {/if}
{/await}
