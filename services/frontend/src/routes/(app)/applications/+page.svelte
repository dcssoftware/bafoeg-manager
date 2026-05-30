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
  import { onDestroy } from "svelte";

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

  let showTableRowOptions: boolean = $state(false);
  let showTableRowOptionElement: HTMLElement | undefined = $state();
  let showTableRowOptionParentElement: HTMLElement | undefined = $state();
  let selectedTableRowOptions: {
    id: string;
    displayName: string;
    status: boolean;
  }[] = $state([
    {
      id: "counter",
      displayName: $_("page.application.overview.datatable.count"),
      status: true,
    },
    {
      id: "digest",
      displayName: $_("page.application.overview.datatable.digest-short"),
      status: true,
    },
    {
      id: "applicant_name",
      displayName: $_("page.application.overview.datatable.name"),
      status: true,
    },
    {
      id: "education_name",
      displayName: $_(
        "page.application.overview.datatable.educational-institution",
      ),
      status: true,
    },
    {
      id: "grade",
      displayName: $_(
        "page.application.overview.datatable.educational-class-level",
      ),
      status: true,
    },
    {
      id: "cty",
      displayName: $_(
        "page.application.overview.datatable.educational-institution-city",
      ),
      status: true,
    },
    {
      id: "assignee",
      displayName: $_("page.application.overview.datatable.clerk"),
      status: true,
    },
    {
      id: "status",
      displayName: $_("page.application.overview.datatable.status"),
      status: true,
    },
    {
      id: "remaining_time",
      displayName: $_("page.application.overview.datatable.remaining-time"),
      status: true,
    },
  ]);

  let applicationPageNumber: number = $state(1);

  function setFilter(
    userID?: string,
    orgaRegionID?: string,
    orgaBehördeID?: string,
    orgaAbteilungID?: string,
    text?: string,
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
      filterSearchTerm,
    );
  }

  function changeFilterOrganizationRegion(region: string | null) {
    selectedOrganizationRegionID = region ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      undefined,
      undefined,
      filterSearchTerm,
    );
  }

  function changeFilterOrganizationBehörde(behörde: string | null) {
    selectedOrganizationBehördeID = behörde ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      undefined,
      filterSearchTerm,
    );
  }

  function changeFilterOrganizationAbteilung(abteilung: string | null) {
    selectedOrganizationAbteilungID = abteilung ?? undefined;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      selectedOrganizationAbteilungID,
      filterSearchTerm,
    );
  }

  function changeFilterText(text: string) {
    filterSearchTerm = text;
    setFilter(
      selectedUserID,
      selectedOrganizationRegionID,
      selectedOrganizationBehördeID,
      selectedOrganizationAbteilungID,
      filterSearchTerm,
    );
  }

  async function loadData(pageNumber: number) {
    applicationsPromise = getApplications(
      pageNumber,
      selectedUserID,
      filterShowFinishedApplications,
      filterSearchTerm,
    );
    applicationsMetricsPromise = getApplicationsMetrics(
      selectedUserID ?? "",
      filterShowFinishedApplications,
    );
  }

  $effect(() => {
    loadData(applicationPageNumber);
  });

  $effect(() => {
    if (showTableRowOptionElement && showTableRowOptionParentElement) {
      const handleClick = (e: MouseEvent) => {
        const dialogDimensions =
          showTableRowOptionParentElement!.getBoundingClientRect();
        if (
          dialogDimensions !== null &&
          (e.clientX < dialogDimensions.left ||
            e.clientX > dialogDimensions.right ||
            e.clientY < dialogDimensions.top ||
            e.clientY > dialogDimensions.bottom)
        ) {
          showTableRowOptions = false;
        }
      };
      document.addEventListener("click", handleClick);
      return () => {
        document?.removeEventListener("click", handleClick);
      };
    }
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
    <div class="extra-menu" bind:this={showTableRowOptionParentElement}>
      <label>
        <Button onclick={() => (showTableRowOptions = !showTableRowOptions)}
          >Tabellen Spalten</Button
        >
      </label>
      <div
        class="menu"
        bind:this={showTableRowOptionElement}
        class:show-menu={showTableRowOptions}
      >
        <ul>
          {#each selectedTableRowOptions as item}
            <li>
              <label>
                <input type="checkbox" bind:checked={item.status} />
                <span>{item.displayName}</span>
              </label>
            </li>
          {/each}
        </ul>
      </div>
    </div>
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
            {#each selectedTableRowOptions as column}
              {#if column.status}
                <DataTableColumn>{column.displayName}</DataTableColumn>
              {/if}
            {/each}
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
                },
              )}
            {@const remainingTimeClass = getRemainingTimeColor(
              remainingDaysInPercent,
            )}

            <DataTableRow onClick={() => goto("/applications/" + row.id)}>
              {#if selectedTableRowOptions.find((el) => el.id === "counter")?.status}
                <DataTableColumn align="Right">{count}</DataTableColumn>
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "digest")?.status}
                <DataTableColumn>{row.id.split("-")[0]}</DataTableColumn>
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "applicant_name")?.status}
                <DataTableColumn
                  >{row.applicant.firstname ?? ""}
                  {row.applicant.lastname}</DataTableColumn
                >
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "education_name")?.status}
                <DataTableColumn>{row.school.name}</DataTableColumn>
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "grade")?.status}
                <DataTableColumn align="Right">{row.classLevel}</DataTableColumn
                >
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "cty")?.status}
                <DataTableColumn
                  >{row.school.address.city} ({row.school.address
                    .country})</DataTableColumn
                >
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "assignee")?.status}
                <DataTableColumn
                  >{#if row.assignedUser !== null}
                    {row.assignedUser.displayName}
                  {:else}
                    {$_("components.user-picker.no-user-assigned")}
                  {/if}
                </DataTableColumn>
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "status")?.status}
                <DataTableColumn
                  >{$_(
                    "states.application-status." + row.status.identifier,
                  )}</DataTableColumn
                >
              {/if}
              {#if selectedTableRowOptions.find((el) => el.id === "remaining_time")?.status}
                <DataTableColumn align="Right"
                  >{$_("states.time.days")} ({remainingDaysInPercentString} %)</DataTableColumn
                >
              {/if}
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
    .extra-menu
      position: relative
      .menu
        position: absolute
        background-color: var(--background-color-tertiary)
        display: none
        &.show-menu
          display: block
        ul
          list-style-type: none
          padding: 1rem
          margin: 0
          display: flex
          flex-direction: column
          gap: 0.5rem
  .controlls
    display: flex
    gap: 1rem
    flex-direction: column
  .datatable
    margin-top: 3rem
</style>
