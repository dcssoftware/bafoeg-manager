<script lang="ts">
  import { goto } from "$app/navigation";
  import { Button } from "../Button";

  interface Props {
    showFileURL: string;
    fileTypeIndicator: string | undefined;
    closeDialog(): void;
    metadata?: Record<string, string>;
  }

  let {
    showFileURL,
    fileTypeIndicator,
    metadata,
    closeDialog = () => {},
  }: Props = $props();
</script>

<div class="file-viewer">
  <div class="sidebar-information">
    <h2>Dateiname</h2>
    <table>
      <tbody>
        {#each Object.entries(metadata ?? {}) as [key, value]}
          <tr>
            <td>{key}</td>
            <td>{value}</td>
          </tr>
        {/each}
      </tbody>
    </table>
    <!-- <Button onclick={closeDialog}>Ganze Message zeigen</Button> -->
  </div>
  <div class="content">
    {#if ["pdf", "application/pdf"].includes(fileTypeIndicator ?? "")}
      <iframe
        title="Requested Document"
        src={showFileURL}
        width="100%"
        height="100%"
        style="box-sizing: border-box;"
      ></iframe>
    {:else if ["png", "image/png", "jpg", "image/jpg", "jpeg", "image/jpeg", "webp", "image/webp"].includes(fileTypeIndicator ?? "")}
      <div class="show-picture">
        <img src={showFileURL} alt="Requested Document" />
      </div>
    {:else}
      <div>
        <p>Unknown File Type</p>
        <Button download onclick={() => goto(showFileURL)}>Download File</Button
        >
      </div>
    {/if}
  </div>
</div>

<style lang="sass">
  .file-viewer
    height: 100%
    display: flex
    flex-direction: row
    gap: 1rem
    .sidebar-information
      min-width: 350px
      max-width: 35%
      border-right: 1px solid var(--font-color)
      padding-right: 1rem
      box-sizing: border-box
    .content
      flex-grow: 1
      height: 100%
      display: flex
      flex-direction: column
      user-select: none
      iframe
        border: none
        flex-grow: 1

      .show-picture
        overflow: hidden
        flex-grow: 1
        img
          width: 100%
          height: 100%
          object-fit: contain
</style>
