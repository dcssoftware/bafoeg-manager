<script lang="ts">
  import type { ApplicationLabelsResponse } from "$lib/api/application-labels";
  import { getApplicationLabels } from "$lib/api/application-labels/get-application-labels";
  import { Label } from "$lib/components/Label";

  let applicationLabelsPromise:
    | Promise<ApplicationLabelsResponse | undefined>
    | undefined = $state(undefined);

  function loadData() {
    applicationLabelsPromise = getApplicationLabels(1);
  }

  $effect(() => {
    loadData();
  });
</script>

<h1>Labels</h1>

<div class="information-not-editable">
  <h2>Aktuell gibt es nur statische Labels</h2>
  <span
    >Während der Beta des Programms wird es nur vorgegebene Labels geben. Später
    können hier aber auch eigene Labels erstellt, verwaltet und bei BAföG
    Anträgen gesetzt, sowie danach gefiltert werden</span
  >
</div>

{#await applicationLabelsPromise then applicationLabels}
  {#if applicationLabels !== undefined}
    <div class="labels-section">
      <h2>Verfügbare Labels:</h2>
      <div class="labels">
        {#each applicationLabels.labels as label}
          <div class="label">
            <Label color={label.style}>{label.name}</Label>
          </div>
        {/each}
      </div>
    </div>
  {:else}
    <span>Error loading labels</span>
  {/if}
{:catch error}
  <span>Error: {error.message}</span>
{/await}

<style lang="sass">
  .information-not-editable
    background-color: var(--color-blue-40)
    width: 100%
    padding: 1rem
    border-radius: 5px
    margin-bottom: 3rem
    h2
      margin: 0
      padding-bottom: 2rem
  .labels
    display: flex
    flex-direction: column
    gap: 10px
</style>
