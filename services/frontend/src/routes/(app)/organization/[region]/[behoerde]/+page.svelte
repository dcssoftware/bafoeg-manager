<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import { getOrganizationAbteilungByBehördenID } from "$lib/api/organization/get-abteilung-by-behoerde-id";
  import type { AbteilungenResponseModelType } from "$lib/api/organization/models/abteilung";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "$lib/components/DataTable";
  import { Pagination } from "$lib/components/Pagination";
  import { _ } from "svelte-i18n";

  const regionID = page.params.region ?? "";
  const behördeID = page.params.behoerde ?? "";
  let selectedPage: number = $state(1);

  let abteilungenPromise:
    | Promise<AbteilungenResponseModelType | undefined>
    | undefined = $state(undefined);

  async function loadData() {
    abteilungenPromise = getOrganizationAbteilungByBehördenID(
      selectedPage,
      behördeID
    );
  }

  $effect(() => {
    loadData();
  });
</script>

<h1>Abteilungen</h1>

{#await abteilungenPromise then abteilungen}
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
        {#each abteilungen?.abteilungen as abteilung, i}
          <DataTableRow
            onClick={() =>
              goto(`/organization/${regionID}/${behördeID}/${abteilung.id}`)}
          >
            <DataTableColumn align="Right">{i + 1}</DataTableColumn>
            <DataTableColumn title={"1"}>{abteilung.id}</DataTableColumn>
            <DataTableColumn>{abteilung.name}</DataTableColumn>
          </DataTableRow>
        {/each}
      </DataTableBody>
    </DataTable>

    <div class="pagination">
      <Pagination
        itemsMaxCount={abteilungen?.count ?? 0}
        itemsPerPage={25}
        bind:selectedPage
      />
    </div>
  </div>
{/await}
