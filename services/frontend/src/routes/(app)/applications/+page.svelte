<script lang="ts">
  import { _ } from "svelte-i18n";
  import { goto } from "$app/navigation";
  import {
    DataTable,
    DataTableHead,
    DataTableRow,
    DataTableColumn,
    DataTableBody,
  } from "$lib/components/DataTable";
  import { getApplications } from "$lib/api/applications/get-applications";
  import type { ApplicationModelResponseType } from "$lib/api/applications/types/application-model-type";
  import { UserPicker } from "$lib/components/UserPicker";
  import { OrganizationRegionPicker } from "$lib/components/OrganizationRegionPicker";
  import { Pagination } from "$lib/components/Pagination";
  import { getRemainingTimeColor } from "$lib/functions/dynamic-color/remaining-time-color";
  import { getApplicationsMetrics } from "$lib/api/applications";
  import type { ApplicationMetricsModelResponseType } from "$lib/api/applications/types/application-metrics-model";
  import { userState } from "$lib/states/user";
  import { FilterTextInput } from "$lib/components/Inputs";
  import { Button } from "$lib/components/Button";
  import { OrganizationBehördePicker } from "$lib/components/OrganizationBehordePicker";
  import { OrganizationAbteilungPicker } from "$lib/components/OrganizationAbteilungPicker";

  let applicationsPromise:
    | Promise<ApplicationModelResponseType | undefined>
    | undefined = $state(undefined);

  let applicationsMetricsPromise:
    | Promise<ApplicationMetricsModelResponseType | undefined>
    | undefined = $state(undefined);

  let filterShowFinishedApplications: boolean = $state(false);
  let filterSearchTerm: string = $state("");

  let selectedUserID: string | undefined = $state($userState?.id);
  let selectedOrganizationRegionID: string | undefined = $state(undefined);
  let selectedOrganizationBehördeID: string | undefined = $state(undefined);
  let selectedOrganizationAbteilungID: string | undefined = $state(undefined);

  let applicationPageNumber: number = $state(1);

  function setFilter(
    userID?: string,
    orgaRegionID?: string,
    orgaBehördeID?: string,
    orgaAbteilungID?: string,
    text?: string
  ) {
    selectedOrganizationRegionID = orgaRegionID;
    selectedOrganizationBehördeID = orgaBehördeID;
    selectedOrganizationAbteilungID = orgaAbteilungID;

    const url = new URL(window.location.href);
    url.searchParams.set("userID", userID ?? "");
    url.searchParams.set("orgaRegionID", orgaRegionID ?? "");
    url.searchParams.set("orgaBehördeID", orgaBehördeID ?? "");
    url.searchParams.set("orgaAbteilungID", orgaAbteilungID ?? "");
    url.searchParams.set("text", text ?? "");
    window.history.replaceState({}, "", url.toString());
  }

  function changeFilterUser(userID: string | null) {
    selectedUserID = userID ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      selectedOrganizationAbteilungID,
      filterSearchTerm
    );
  }

  function changeFilterOrganizationRegion(region: string | null) {
    selectedOrganizationRegionID = region ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      undefined,
      undefined,
      filterSearchTerm
    );
  }

  function changeFilterOrganizationBehörde(behörde: string | null) {
    selectedOrganizationBehördeID = behörde ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      undefined,
      filterSearchTerm
    );
  }

  function changeFilterOrganizationAbteilung(abteilung: string | null) {
    selectedOrganizationAbteilungID = abteilung ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      selectedOrganizationAbteilungID,
      filterSearchTerm
    );
  }

  function changeFilterText(text: string) {
    filterSearchTerm = text;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      selectedOrganizationAbteilungID,
      filterSearchTerm
    );
  }

  async function loadData(pageNumber: number) {
    applicationsPromise = getApplications(
      pageNumber,
      selectedUserID,
      filterShowFinishedApplications,
      filterSearchTerm
    );
    applicationsMetricsPromise = getApplicationsMetrics(
      selectedUserID ?? "",
      filterShowFinishedApplications
    );
  }

  $effect(() => {
    loadData(applicationPageNumber);
  });
</script>

<div class="header">
  <h1>
    {$_("page.application.overview.header")}
  </h1>
</div>

<div class="statistics-wrapper">
  {#await applicationsMetricsPromise}
    <span>{$_("states.loading")}</span>
  {:then applicationsMetrics}
    {#if applicationsMetrics !== undefined}
      {@const i18npath = "page.application.overview.statistics"}
      <div class="statistics">
        <div class="statistic-box open-applications">
          <span>{$_(`${i18npath}.total`)}</span>
          <span class="number">{applicationsMetrics.total}</span>
        </div>
        <div class="statistic-box new-applications-today">
          <span>{$_(`${i18npath}.new-today`)}</span>
          <span class="number">{applicationsMetrics.newToday}</span>
        </div>
        <div class="statistic-box in-progress">
          <span>{$_(`${i18npath}.in-progress`)}</span>
          <span class="number">{applicationsMetrics.inProgress}</span>
        </div>
        <div class="statistic-box in-progress">
          <span>{$_(`${i18npath}.user-assigned`)}</span>
          <span class="number"
            >{applicationsMetrics.userAssigned != 0
              ? applicationsMetrics.userAssigned
              : "-"}</span
          >
        </div>
      </div>
    {/if}
  {/await}
</div>

<div class="controlls">
  <div class="filter-components">
    <label>
      <span>Region</span>
      <OrganizationRegionPicker
        selectedUserID={selectedOrganizationRegionID ?? ""}
        onChange={(id: string | null) => changeFilterOrganizationRegion(id)}
      />
    </label>
    <label>
      <span>Behörde</span>
      <OrganizationBehördePicker
        regionID={selectedOrganizationRegionID ?? ""}
        selectedUserID={selectedOrganizationBehördeID ?? ""}
        onChange={(id: string | null) => changeFilterOrganizationBehörde(id)}
      />
    </label>
    <label>
      <span>Abteilung</span>
      <OrganizationAbteilungPicker
        regionID={selectedOrganizationRegionID ?? ""}
        behördeID={selectedOrganizationBehördeID ?? ""}
        selectedUserID={selectedOrganizationAbteilungID ?? ""}
        onChange={(id: string | null) => changeFilterOrganizationAbteilung(id)}
      />
    </label>
    <label>
      <span>Sachbearbeiter</span>
      <UserPicker
        selectedUserID={selectedUserID ?? ""}
        onChange={(userID: string | null) => changeFilterUser(userID)}
      />
    </label>
    <label style="margin-left: auto;">
      <input type="checkbox" bind:checked={filterShowFinishedApplications} />
      <span>Bearbeitete anzeigen</span>
    </label>
    <label>
      <Button>Tabellen Spalten</Button>
    </label>
  </div>

  <div class="filter-components">
    <label style="flex-grow: 1;">
      <span>Antrag-Suchkriterium</span>
      <FilterTextInput
        {filterSearchTerm}
        onFilterChange={(value: string) => changeFilterText(value)}
      />
    </label>
    <!-- <SearchTextInput /> -->
  </div>
</div>

{#await applicationsPromise}
  <span>{$_("page.application.overview.loading")}</span>
{:then applicationsWithMetaData}
  {#if applicationsWithMetaData != undefined && applicationsWithMetaData.maxCount === 0}
    <span>{$_("page.application.overview.no-applications-found")}</span>
  {:else if applicationsWithMetaData != undefined}
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
            <DataTableColumn
              >{$_("page.application.overview.datatable.name")}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.educational-institution"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.educational-class-level"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.educational-institution-city"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.clerk"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.status"
              )}</DataTableColumn
            >
            <DataTableColumn
              >{$_(
                "page.application.overview.datatable.remaining-time"
              )}</DataTableColumn
            >
          </DataTableRow>
        </DataTableHead>
        <DataTableBody>
          {#each applicationsWithMetaData.application as row, i}
            {@const count = applicationPageNumber * 10 - 10 + (i + 1)}
            {@const remainingDays = row.processingTime.remainingTimeInDays}
            {@const remainingDaysInPercent =
              row.processingTime.remainingTimeInPercent}
            {@const remainingDaysInPercentString =
              row.processingTime.remainingTimeInPercent.toLocaleString(
                "de-DE",
                {
                  unit: "percent",
                  maximumFractionDigits: 2,
                  minimumFractionDigits: 2,
                }
              )}
            {@const remainingTimeClass = getRemainingTimeColor(
              remainingDaysInPercent
            )}

            <DataTableRow onClick={() => goto("/applications/" + row.id)}>
              <DataTableColumn align="Right">{count}</DataTableColumn>
              <DataTableColumn title={row.id}
                >{row.id.split("-")[0]}</DataTableColumn
              >
              <DataTableColumn
                >{row.applicant.firstname ?? ""}
                {row.applicant.lastname}</DataTableColumn
              >
              <DataTableColumn>{row.school.name}</DataTableColumn>
              <DataTableColumn align="Right">{row.classLevel}</DataTableColumn>
              <DataTableColumn
                >{row.school.address.city} ({row.school.address
                  .country})</DataTableColumn
              >
              <DataTableColumn>
                {#if row.assignedUser !== null}
                  {row.assignedUser.displayName}
                {:else}
                  {$_("components.user-picker.no-user-assigned")}
                {/if}
              </DataTableColumn>
              <DataTableColumn
                >{$_(
                  "states.application-status." + row.status.identifier
                )}</DataTableColumn
              >
              <DataTableColumn align="Right" classList={remainingTimeClass}
                >{remainingDays}
                {$_("states.time.days")} ({remainingDaysInPercentString} %)</DataTableColumn
              >
            </DataTableRow>
          {/each}
        </DataTableBody>
      </DataTable>

      <div class="pagination">
        <Pagination
          itemsMaxCount={applicationsWithMetaData.maxCount}
          itemsPerPage={25}
          bind:selectedPage={applicationPageNumber}
        />
      </div>
    </div>
  {:else}
    <span>Internal Server Error</span>
  {/if}
{/await}

<style lang="sass">
  .statistics-wrapper
    min-height: 5rem 
    margin-bottom: 3rem
  .statistics
    display: grid
    grid-template-columns: repeat(4, 1fr)
    gap: 1rem
    .statistic-box
      border-radius: 5px
      background-color: var(--background-color-tertiary)
      display: flex
      justify-content: space-between
      padding: 1rem
      align-items: center

      .number
        background-color: var(--color-green)
        color: var(--color-white)
        height: 3rem
        width: 4rem
        border-radius: 5px
        align-content: center
        text-align: center
        font-size: 1.5rem
        font-weight: bold

  .filter-components
    display: flex
    width: 100%
    gap: 2rem
    label
      display: block
  .controlls
    display: flex
    gap: 1rem
    flex-direction: column
  .datatable
    margin-top: 3rem
</style>
