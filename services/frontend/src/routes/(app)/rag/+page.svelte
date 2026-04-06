<script lang="ts">
  import { getRagInformation } from "$lib/api/rag/get-rag-information";
  import type { RagInformationModelResponseType } from "$lib/api/rag/type/rag-information-model-type";
  import { IconChild, IconStudent } from "$lib/components/Icons";
  import { userState } from "$lib/states/user";

  let ragInformationPromise:
    | Promise<RagInformationModelResponseType | undefined>
    | undefined = $state(undefined);

  async function loadData() {
    ragInformationPromise = getRagInformation();
  }

  $effect(() => {
    loadData();
  });
</script>

<h1>RAG Datenbasis</h1>
<p>
  Das RAG (Retrieval-Augmented Generation) ist die Wissensdatenbank für die
  KI-Unterstützung im BAföG-Antragsprozess. Abgelegte Daten werden für die KI
  aufbereitet und anschließend genutzt, um Antworten zu generieren. Sollten neue
  Gesetze, Beschlüsse oder Richtlinien erlassen werden, fügen Sie diese bitte
  der Datenbasis hinzu.
</p>

<div class="model-information">
  {#await ragInformationPromise}
    <span>Load...</span>
  {:then ragInformation}
    <div class="ai-model-information">
      <b>Developer Information</b>
      <table>
        <tbody>
          <tr>
            <td>Server Address</td>
            <td>{ragInformation?.aiModelServerAddress}</td>
          </tr>
          <tr>
            <td>Server Port</td>
            <td>{ragInformation?.aiModelServerPort}</td>
          </tr>
          <tr>
            <td>Server Secure</td>
            <td>{ragInformation?.aiModelServerSecure}</td>
          </tr>
          <tr>
            <td>Embedding Modelname</td>
            <td>{ragInformation?.embeddingModelname}</td>
          </tr>
          <tr>
            <td>Requesting Modelname</td>
            <td>{ragInformation?.requestingModelname}</td>
          </tr>
        </tbody>
      </table>
    </div>
  {/await}
</div>

<div class="selection">
  {#if $userState?.permissions.includes("read:rag-management-studierenden-files")}
    <a href="/rag/studierenden-bafoeg-regeln">
      <IconStudent />
      <span>Wissensdatenbank Studierenden BAföG</span>
    </a>
  {/if}
  {#if $userState?.permissions.includes("read:rag-management-schueler-files")}
    <a href="/rag/schueler-bafoeg-regeln">
      <IconChild />
      <span>Wissensdatenbank Schüler BAföG</span>
    </a>
  {/if}
</div>

<style lang="sass">

  .ai-model-information
    background-color: var(--background-color-tertiary)
    display: inline-block
    padding: 1rem
    b
      font-size: 1.3rem 
    table
      margin-top: .5rem
      tr:hover
        cursor: default
        background-color: var(--background-color-secondary)
      td
        padding: .25rem

  .selection
    display: flex
    justify-content: center
    gap: 1rem
    margin-top: 2rem

    a
      padding: 1rem 2rem
      text-decoration: none
      font-weight: bold
      display: flex
      flex-direction: column
      align-items: center
      gap: 0.5rem
      user-select: none

      :global(svg)
        $size: 10rem
        fill: var(--font-color)
        width: $size
        height: $size

</style>
