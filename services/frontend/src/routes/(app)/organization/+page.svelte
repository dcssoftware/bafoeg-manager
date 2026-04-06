<script lang="ts">
  import { goto } from "$app/navigation";
  import { getOrganizationRegions } from "$lib/api/organization/get-regions";
  import type { RegionResponseModelType } from "$lib/api/organization/models/region";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "$lib/components/DataTable";
  import { Pagination } from "$lib/components/Pagination";
  import { _ } from "svelte-i18n";

  let selectedPage: number = $state(1);
  let regionsPromise: Promise<RegionResponseModelType | undefined> | undefined =
    $state(undefined);

  async function loadData() {
    regionsPromise = getOrganizationRegions(1);
  }

  $effect(() => {
    loadData();
  });
</script>

<h1>Region</h1>

{#await regionsPromise then regions}
  {#if regions !== undefined}
    <div class="datatable">
      <DataTable>
        <DataTableHead>
          <DataTableRow>
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.count"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.digest-short"
              )}</DataTableColumn
            >
            <DataTableColumn>Identifier</DataTableColumn>
            <DataTableColumn>Name</DataTableColumn>
          </DataTableRow>
        </DataTableHead>
        <DataTableBody>
          {#each regions.regions as region, i}
            <DataTableRow onClick={() => goto("/organization/" + region.id)}>
              <DataTableColumn align="Right">{i + 1}</DataTableColumn>
              <DataTableColumn title={region.id}>{region.id}</DataTableColumn>
              <DataTableColumn>{region.identifier}</DataTableColumn>
              <DataTableColumn>{region.name}</DataTableColumn>
            </DataTableRow>
          {/each}
        </DataTableBody>
      </DataTable>
    </div>
    <div class="pagination">
      <Pagination
        itemsMaxCount={regions.maxCount}
        itemsPerPage={25}
        bind:selectedPage
      />
    </div>
  {/if}
{/await}
