<script lang="ts">
  import { _ } from "svelte-i18n";
  import {
    DataTable,
    DataTableBody,
    DataTableColumn,
    DataTableHead,
    DataTableRow,
  } from "$lib/components/DataTable";
  import { getSchools } from "$lib/api/schools/get-schools";
  import { FilterTextInput } from "$lib/components/Inputs";
  import { Pagination } from "$lib/components/Pagination";
  import type { SchoolShortResponseModelType } from "$lib/api/schools/types/school-short-model-type";

  interface Props {
    onSingleSelect?(id: string): void;
  }
  let { onSingleSelect = (id: string) => {} }: Props = $props();

  let schoolDataPromise:
    | Promise<SchoolShortResponseModelType | undefined>
    | undefined = $state(undefined);

  let filterSearchTerm: string = $state("");
  let applicationPageNumber: number = $state(1);

  async function loadData() {
    schoolDataPromise = getSchools(1, filterSearchTerm);
  }

  $effect(() => {
    loadData();
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

<!-- <Button>{$_("page.school-management.create-school")}</Button> -->

{#await schoolDataPromise then schoolsModel}
  <div class="datatable">
    <DataTable>
      <DataTableHead>
        <DataTableRow>
          <DataTableColumn>Nr</DataTableColumn>
          <DataTableColumn>ID</DataTableColumn>
          <DataTableColumn>Name</DataTableColumn>
          <DataTableColumn>Type</DataTableColumn>
          <DataTableColumn>City</DataTableColumn>
        </DataTableRow>
      </DataTableHead>
      <DataTableBody>
        {#if schoolsModel !== undefined}
          {#each schoolsModel.schools as school, i}
            {@const number = i + 1}
            {@const schoolID = school.id.split("-")[0]}
            <DataTableRow onClick={() => onSingleSelect(school.id)}>
              <DataTableColumn>{number}</DataTableColumn>
              <DataTableColumn title={school.id}>{schoolID}</DataTableColumn>
              <DataTableColumn>{school.name}</DataTableColumn>
              <DataTableColumn>{school.type.name}</DataTableColumn>
              <DataTableColumn>{school.city} ({school.country})</DataTableColumn
              >
            </DataTableRow>
          {/each}
        {/if}
      </DataTableBody>
    </DataTable>
  </div>

  <div class="pagination">
    <Pagination
      itemsMaxCount={schoolsModel?.maxCount ?? 0}
      itemsPerPage={schoolsModel?.count ?? 0}
      bind:selectedPage={applicationPageNumber}
    />
  </div>
{/await}

<style lang="sass">
  .datatable
    margin-top: 3rem
</style>
