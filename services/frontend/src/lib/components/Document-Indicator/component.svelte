<script lang="ts">
  import { convertFileSize } from "$lib/functions/convert-file-size";
  import { dateToFormatStringLong } from "$lib/functions/date";

  interface Props {
    documentID: string;
    documentName: string;
    documentType: string;
    documentSize: number;
    documentCreated: Date;
    documentClickable: boolean;
    showFile?: (documentID: string, documentType: string) => void;
  }

  let {
    documentID,
    documentName,
    documentType,
    documentSize,
    documentCreated,
    documentClickable = false,
    showFile = showFileDefault,
  }: Props = $props();

  let documentNameType = $derived(translateFileType(documentType));

  function translateFileType(documentType: string) {
    switch (documentType) {
      case "application/pdf":
        return { name: "PDF", indicator: "pdf" };
      case "image/png":
        return { name: "PNG", indicator: "png" };
      case "image/jpeg":
      case "image/jpg":
        return { name: "JPG", indicator: "jpg" };
      case "image/webp":
        return { name: "WebP", indicator: "webp" };
      case "application/zip":
        return { name: "Zip", indicator: "zip" };
      default:
        return { name: documentType, indicator: "unknown" };
    }
  }

  function showFileDefault(documentID: string, documentType: string) {}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_missing_attribute -->
<a
  class="document"
  class:documentClickable
  onclick={() => showFile(documentID, documentNameType.indicator)}
>
  <div class="document-wrapper" title={documentName}>
    <div class="logo logo-type-{documentNameType.indicator}">
      <span>{documentNameType.name}</span>
    </div>
    <div class="document-info">
      <span class="document-name">{documentName}</span>
      <span class="document-size">{convertFileSize(documentSize)}</span>
      <span class="document-uploaded-at"
        >{dateToFormatStringLong(documentCreated)}</span
      >
    </div>
  </div>
</a>

<style lang="sass">
  .document
      border-radius: 10px
      padding: 1rem
      width: 300px
      max-width: 100%
      box-sizing: border-box

      &.documentClickable
        cursor: pointer


      &:hover
        background-color: var(--background-color-secondary)
        .document-wrapper
          .logo
            background-color: var(--color-red)

      .document-wrapper
        display: flex
        flex-direction: column
        align-items: center
        gap: 1rem

        text-overflow: ellipsis
        overflow: hidden

        .logo
          height: 250px
          width: 250px
          max-width: 100%
          color: var(--color-white)
          background-color: var(--color-red)
          display: flex
          align-items: center
          border-radius: 10px
          white-space: nowrap
          justify-content: center;

          &.logo-type-pdf
            background-color: var(--color-red)
          &.logo-type-png, &.logo-type-jpg, &.logo-type-webp
            background-color: var(--color-blue)
          &.logo-type-zip
            background-color: var(--color-green)
          &.logo-type-unknown
            background-color: var(--color-grey)

          span
            font-size: 1.5rem

        .document-info
          display: flex
          flex-direction: column
          align-items: center
          gap: 0.5rem
          text-align: center
          width: 250px

          max-width: 100% 
          text-overflow: ellipsis

          span
            display: block
            width: 100%
            text-align: left

          .document-name
            font-weight: bold
            font-size: 1.2rem
            line-height: 1.8rem
            overflow: hidden
            text-overflow: ellipsis
            height: 3.5rem

          .document-size,
          .document-uploaded-at
            text-align: center

</style>
