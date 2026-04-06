<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import { getOrganizationBehördenByRegion } from "$lib/api/organization/get-behoerde-by-region-id";
  import type { BehördeResponseModelType } from "$lib/api/organization/models/behoerde";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "$lib/components/DataTable";
  import { _ } from "svelte-i18n";

  const regionID = page.params.region ?? "";
  let behördenPromise:
    | Promise<BehördeResponseModelType | undefined>
    | undefined = $state(undefined);

  async function loadData() {
    console.log("Loading data for region ID:", regionID);
    behördenPromise = getOrganizationBehördenByRegion(1, regionID);
  }

  $effect(() => {
    loadData();
  });
</script>

<h1>Behörde</h1>

{#await behördenPromise then behörden}
  <div class="datatable">
    <DataTable>
      <DataTableHead>
        <DataTableRow>
          <DataTableColumn
            >{$_("page.application.overview.datatable.count")}</DataTableColumn
          >
          <DataTableColumn
            >{$_(
              "page.application.overview.datatable.digest-short"
            )}</DataTableColumn
          >
          <DataTableColumn>Name</DataTableColumn>
        </DataTableRow>
      </DataTableHead>
      <DataTableBody>
        {#each behörden?.behoerden as behörde, i}
          <DataTableRow
            onClick={() => goto(`/organization/${regionID}/${behörde.id}`)}
          >
            <DataTableColumn align="Right">{i + 1}</DataTableColumn>
            <DataTableColumn title={"1"}>{behörde.id}</DataTableColumn>
            <DataTableColumn>{behörde.name}</DataTableColumn>
          </DataTableRow>
        {/each}
      </DataTableBody>
    </DataTable>

    <!-- <div class="pagination">
      <Pagination itemsMaxCount={3} itemsPerPage={25} />
    </div> -->
  </div>
{/await}
