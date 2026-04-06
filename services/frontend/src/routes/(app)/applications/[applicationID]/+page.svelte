<script lang="ts">
  import { _ } from "svelte-i18n";
  import { page } from "$app/state";
  import AddressCards from "./components/basic/address-cards.svelte";
  import ApplicationInfo from "./components/basic/application-info.svelte";
  import * as Tabs from "$lib/components/Tabs";
  import FinancialsOwn from "./components/tabs/financials-own.svelte";
  import FinancialsParents from "./components/tabs/financials-parents.svelte";
  import InformationParents from "./components/tabs/information-parents.svelte";
  import { IconClockRotate } from "$lib/components/Icons";
  import { getApplicationByID } from "$lib/api/applications/get-application-by-id";
  import { type ApplicationModelResponseType } from "$lib/api/applications";
  import AiResponse from "./components/tabs/ai-response.svelte";
  import Cv from "./components/tabs/cv.svelte";
  import Overview from "./components/tabs/overview.svelte";
  import UploadedDocuments from "./components/tabs/uploaded-documents.svelte";
  import { generateHash } from "$lib/functions/random/gen-hash";
  import ActionPanel from "./components/basic/action-panel.svelte";

  const applicationID = page.params.applicationID;
  let baseURL: string = `/applications/${applicationID}/`;

  let tabValue = $state("overview");
  let applicationReloadHash = $state(generateHash(8));

  let applicationPromise:
    | Promise<ApplicationModelResponseType | undefined>
    | undefined = $state(undefined);

  function loadData(hash: string) {
    const url = new URL(window.location.href);
    const urlTabValue = url.searchParams.get("tab");
    if (urlTabValue != null || urlTabValue === "") {
      tabValue = urlTabValue;
    }
    applicationPromise = getApplicationByID(applicationID ?? "");
  }

  function gotoAiReport() {
    tabValue = "ai-response";
  }

  function updateTabSwitchURL(value: string) {
    const url = new URL(window.location.href);
    url.searchParams.set("tab", value);
    window.history.replaceState({}, "", url.toString());
  }

  $effect(() => {
    loadData(applicationReloadHash);
  });
</script>

<div class="component">
  <div class="header">
    <h1>{$_("page.application.editor.header")}</h1>
  </div>

  {#await applicationPromise}
    <span>{$_("states.loading")}</span>
  {:then application}
    {#if application !== undefined}
      {#if application.status.identifier === "APPROVED"}
        <div class="already-done-banner approved">
          <span>Bereits genehmigt</span>
        </div>
      {:else if application.status.identifier === "DENIED"}
        <div class="already-done-banner denied">
          <span>Bereits abgelehnt</span>
        </div>
      {/if}

      <div class="layout">
        <div class="sidebar">
          <AddressCards
            allowWrap={true}
            applicant={application.applicant}
            school={application.school}
            bind:refreshApplicationHash={applicationReloadHash}
          />
          <ActionPanel
            applicationID={application.id}
            currentStatus={application.status}
            reloadPage={() => loadData(generateHash(8))}
          />
        </div>
        <div class="content">
          <div class="base-information">
            <ApplicationInfo bind:applicationReloadHash {application} />

            <div class="versions">
              <div class="current-version">
                <span
                  >{$_(
                    "page.application.editor.basic-infos.versions.current-version"
                  )}</span
                >
                <span>04f34f49</span>
              </div>
              <a href={baseURL + "revision"} class="show-revisions">
                <span
                  >{$_(
                    "page.application.editor.basic-infos.versions.previous-version"
                  )}</span
                >
                <IconClockRotate />
              </a>
            </div>

            <div class="cards">
              <div class="cards-web">
                <div class="tabs">
                  <Tabs.Root
                    bind:value={tabValue}
                    onValueChange={updateTabSwitchURL}
                    class=""
                  >
                    <Tabs.List>
                      <Tabs.Trigger value="overview">Übersicht</Tabs.Trigger>
                      <!-- <Tabs.Trigger value="cv">Lebenslauf</Tabs.Trigger>
                      <Tabs.Trigger value="own-finance"
                        >{$_(
                          "page.application.editor.tabs.tab-headers.own-financial"
                        )}</Tabs.Trigger
                      >
                      <Tabs.Trigger value="parents"
                        >{$_(
                          "page.application.editor.tabs.tab-headers.parents"
                        )}</Tabs.Trigger
                      >
                      <Tabs.Trigger value="parent-finances"
                        >{$_(
                          "page.application.editor.tabs.tab-headers.parents-financial"
                        )}</Tabs.Trigger
                      > -->
                      <Tabs.Trigger value="uploaded-documents"
                        >Hochgeladene Dokumente</Tabs.Trigger
                      >
                      <!-- <Tabs.Trigger value="ai-response">KI-Analyse</Tabs.Trigger
                      > -->
                    </Tabs.List>
                    <Tabs.Content value="overview">
                      <Overview {gotoAiReport} />
                    </Tabs.Content>
                    <!-- <Tabs.Content value="cv">
                      <Cv />
                    </Tabs.Content>
                    <Tabs.Content value="own-finance">
                      <FinancialsOwn />
                    </Tabs.Content>
                    <Tabs.Content value="parents">
                      <InformationParents />
                    </Tabs.Content>
                    <Tabs.Content value="parent-finances">
                      <FinancialsParents />
                    </Tabs.Content> -->
                    <Tabs.Content value="uploaded-documents">
                      <UploadedDocuments
                        {applicationID}
                        applicationStatus={application.status}
                      />
                    </Tabs.Content>
                    <!-- <Tabs.Content value="ai-response">
                      <AiResponse />
                    </Tabs.Content> -->
                  </Tabs.Root>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    {/if}
  {/await}
</div>

<style lang="sass">
  .already-done-banner
    background-color: var(--color-blue)
    color: var(--color-white)
    font-size: 1.2rem;
    padding: 1rem
    margin: 2rem 0

    &.approved
      background-color: var(--color-green)
    &.denied
      background-color: var(--color-red)

  .layout
    display: flex
    flex-direction: column
    gap: 2rem

  @media (min-width: 2500px) 
    .layout
      flex-direction: row-reverse
      .content
        flex-grow: 1
      .sidebar
        min-width: 700px
        width: 700px

  .sidebar
    display: flex
    flex-direction: column
    gap: 2rem

    .application-actions
      background-color: var(--background-color-tertiary)
      padding: 2rem 3rem
      display: flex
      flex-direction: column
      gap: 1rem

      :global(button)
        width: 100%

  .base-information
    display: flex
    flex-direction: column
    gap: 2rem

  .versions
    display: flex
    flex-direction: row
    justify-content: space-between
    align-items: center
    background-color: var(--background-color-tertiary)
    padding: 1rem 2rem

    a
      text-decoration: none
      &:focus, &:active, &:visited
        color: var(--font-color)

    .show-revisions
      display: flex
      align-items: center
      gap: .5rem
      :global(svg)
        $size: 1.5rem
        height: $size
        width: $size
        fill: var(--font-color)

  .application-dates
    background-color: var(--background-color-tertiary)
    padding: 2rem 3rem

  .print-section
    margin-bottom: 20px

</style>
