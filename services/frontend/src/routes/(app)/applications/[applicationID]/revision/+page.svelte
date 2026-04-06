<script lang="ts">
  import { _ } from "svelte-i18n";
  import { page } from "$app/state";
  import { getApplicationsRevisionsByApplicationID } from "$lib/api/applications/get-application-revision-by-application-id";
  import type { ApplicationRevisionModelResponseType } from "$lib/api/applications/types/application-revision-model-type";
  import { IconHalfFile } from "$lib/components/Icons";
  import { Pagination } from "$lib/components/Pagination";
  import { dateToFormatStringShort } from "$lib/functions/date";

  let applicationRevisionsPromise:
    | Promise<ApplicationRevisionModelResponseType | undefined>
    | undefined = $state(undefined);

  let applicationPageNumber: number = $state(1);
  const applicationID = $state(page.params.applicationID);

  function loadData(
    applicationID: string | undefined,
    applicationPageNumber: number
  ) {
    applicationRevisionsPromise = getApplicationsRevisionsByApplicationID(
      applicationPageNumber,
      applicationID
    );
  }

  $effect(() => {
    loadData(applicationID, applicationPageNumber);
  });
</script>

<h1>{$_("page.application.revisions.header")}</h1>

{#await applicationRevisionsPromise then data}
  {#if data !== undefined}
    <div class="revision-list">
      {#each data.revisions as row, i}
        <!-- {@const count = applicationPageNumber * 10 - 10 + (i + 1)} -->

        <div class="revision">
          <span>{dateToFormatStringShort(row.created)}</span>
          <span>{row.header}</span> <br />
          <span class="spacer update_source">Unknown</span>
          <span class="digest">{row.id.split("-")[0]}</span>
          <a href="/applications/{applicationID}/diff/abc/abc">
            <span class="icon"><IconHalfFile /></span>
          </a>
        </div>
      {/each}
    </div>

    <div class="pagination">
      <Pagination
        itemsMaxCount={data.maxCount}
        itemsPerPage={25}
        bind:selectedPage={applicationPageNumber}
      />
    </div>
  {:else}
    <div class="no-entries">
      <span>{$_("page.application.revisions.no-entries-found")}</span>
      <br />
      <br />
      <span
        >Leere Revisionen sind nicht vorgesehen. Bitte den Fehler an die IT und
        an den Programm-Entwickler melden.</span
      >
      <br />
      <span
        >Empty revisions are not a desired state. Please report that bug to your
        IT-Administrator and the developer of this application.</span
      >
    </div>
  {/if}
{/await}

<style lang="sass">
  .revision-list 
    display: flex
    flex-direction: column
    gap: 1rem
    width: 100%

  .revision 
    display: flex
    align-items: flex-start
    gap: 1rem
    width: 100%
    padding: 1rem
    background-color: var(--background-color-tertiary)

    span
      padding: 0.5rem

    .spacer
      margin-left: auto
    .update_source
      background-color: var(--color-green-70)
      color: var(--color-white)
      border-radius: 5px
      padding: 0.5rem
      font-size: 0.8rem
      font-weight: bold
    .digest
      border: 1px solid var(--font-color)
      color: var(--font-color)
      border-radius: 5px
      padding: 0.5rem
      font-size: 0.8rem

    .icon
      :global(svg)
        width: 1rem
        height: 1rem
        fill: var(--font-color)
</style>
