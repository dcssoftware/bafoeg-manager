<script lang="ts">
  import { Tabs as TabsPrimitive } from "bits-ui";

  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();
</script>

<!--
{cn(
    "bg-muted text-muted-foreground inline-flex h-9 items-center justify-center rounded-lg p-1",
    className
  )}
-->

<div class="dcs-theme">
  <div class="scrollbar">
    <TabsPrimitive.List class="tabs-list">
      {@render children?.()}
    </TabsPrimitive.List>
  </div>
</div>

<style lang="sass">
  .dcs-theme
    .scrollbar
      white-space: nowrap
      overflow: scroll

      :global(.tabs-list)
        background-color: var(--background-color-tertiary)
        display: inline-flex
        flex-wrap: wrap
        border-radius: 5px
</style>
