<script lang="ts">
  import type { SchoolModelType } from "$lib/api/schools";
  import { getSchoolByID } from "$lib/api/schools/get-school-by-id";
  import { _ } from "svelte-i18n";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "../DataTable";

  interface Props {
    schoolID: string;
    onSingleSelect?(id: string): void;
  }
  let { schoolID, onSingleSelect = (id: string) => {} }: Props = $props();

  let schoolDataPromise: Promise<SchoolModelType | undefined> | undefined =
    $state(undefined);

  function loadData() {
    schoolDataPromise = getSchoolByID(schoolID);
  }

  $effect(() => {
    loadData();
  });
</script>

{#await schoolDataPromise}
  <span>{$_("states.loading")}</span>
{:then school}
  {#if school !== undefined}
    <div class="datatable">
      {#if true}
        <DataTable>
          <DataTableHead>
            <DataTableRow>
              <DataTableColumn>#</DataTableColumn>
              <DataTableColumn>Abschluss</DataTableColumn>
              <DataTableColumn>FOS Berufsabschluss Benötigt</DataTableColumn>
              <DataTableColumn
                >BOS Berufsqualifizierender Abschluss</DataTableColumn
              >
              <DataTableColumn
                >Fachschule / Berufsschule Benötigt</DataTableColumn
              >
            </DataTableRow>
          </DataTableHead>
          <DataTableBody>
            {#each school.degree as degree, i}
              {@const number = i + 1}
              <DataTableRow onClick={() => onSingleSelect(degree.id)}>
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
  {/if}
{/await}
