<script lang="ts">
  import type { ContentLayoutModes } from ".";

  interface Props {
    mode?: ContentLayoutModes;
    children?: import("svelte").Snippet;
  }

  let { mode = "default", children }: Props = $props();
</script>

<div class="layout mode-{mode}">
  {@render children?.()}
</div>

<style lang="sass">
    .mode-default
        padding: 2rem
    .mode-center-blog
        margin: auto
        padding: max(2rem)
        max-width: 1000px
    .mode-center-form
        margin: auto
        padding: max(2rem)
        max-width: min(1800px, 90%)
    @media print
      .layout
        max-width: 100%
        padding: 0
        margin: 0
</style>
