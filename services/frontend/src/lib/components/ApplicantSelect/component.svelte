<script lang="ts">
  import { type ApplicantResponse, getApplicants } from "$lib/api/applicants";
  import { _ } from "svelte-i18n";
  import { Pagination } from "$lib/components/Pagination";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "../DataTable";
  import { FilterTextInput } from "../Inputs";

  interface Props {
    onSingleSelect?(id: string): void;
  }

  let { onSingleSelect: onSelect = (id: string) => {} }: Props = $props();

  let applicantsDataPromise:
    | Promise<ApplicantResponse | undefined>
    | undefined = $state(undefined);

  let applicationPageNumber: number = $state(1);
  let filterSearchTerm: string = $state("");

  async function loadData(filterSearchTerm: string) {
    applicantsDataPromise = getApplicants(
      applicationPageNumber,
      filterSearchTerm
    );
  }

  $effect(() => {
    loadData(filterSearchTerm);
  });
</script>

<div class="controlls">
  <div class="filter-components">
    <label>
      <FilterTextInput
        {filterSearchTerm}
        onFilterChange={(value) => (filterSearchTerm = value)}
      />
    </label>
  </div>
</div>

{#await applicantsDataPromise then applicantData}
  {#if applicantData != undefined}
    <div class="datatable">
      <DataTable>
        <DataTableHead>
          <DataTableRow>
            <DataTableColumn
              >{$_("page.applicant.overview.datatable.count")}</DataTableColumn
            >
            <DataTableColumn
              >{$_("page.applicant.overview.datatable.id")}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.applicant.overview.datatable.firstname"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.applicant.overview.datatable.lastname"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.applicant.overview.datatable.streety"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_("page.applicant.overview.datatable.city")}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.applicant.overview.datatable.country"
              )}</DataTableColumn
            >
          </DataTableRow>
        </DataTableHead>
        <DataTableBody>
          {#each applicantData.applicants as data, i}
            {@const number = i + 1}
            {@const applicantID = data.id.split("-")[0]}
            <DataTableRow onClick={() => onSelect(data.id)}>
              <DataTableColumn>{number}</DataTableColumn>
              <DataTableColumn>{applicantID}</DataTableColumn>
              <DataTableColumn>{data.firstname}</DataTableColumn>
              <DataTableColumn>{data.lastname}</DataTableColumn>
              <!-- <DataTableColumn>Maidenname is not implemented yet</DataTableColumn> -->
              <DataTableColumn
                >{data.address.street}
                {data.address.houseNumber}</DataTableColumn
              >
              <DataTableColumn
                >{data.address.zipCode} {data.address.city}</DataTableColumn
              >
              <DataTableColumn>{data.address.country}</DataTableColumn>
            </DataTableRow>
          {/each}
        </DataTableBody>
      </DataTable>
      <div class="pagination">
        <Pagination
          itemsMaxCount={applicantData.maxCount}
          itemsPerPage={25}
          bind:selectedPage={applicationPageNumber}
        />
      </div>
    </div>
  {/if}
{/await}

<style lang="sass">
  .controlls
    margin-top: 5rem
  .datatable
    margin-top: 3rem
</style>
