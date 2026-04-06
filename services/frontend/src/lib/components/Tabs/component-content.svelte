<script lang="ts">
  import { Tabs as TabsPrimitive } from "bits-ui";

  interface Props {
    children?: import("svelte").Snippet;
    value: string;
  }

  let { children, value = $bindable() }: Props = $props();
</script>

<!-- 
{cn(
    "ring-offset-background focus-visible:ring-ring mt-2 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2",
    className
  )}
-->

<div class="dcs-theme">
  <TabsPrimitive.Content class="tab-content" {value}>
    {@render children?.()}
  </TabsPrimitive.Content>
</div>

<style lang="sass">
  .dcs-theme
    :global(.tab-content)
      background-color: var(--background-color-tertiary)
      color: var(--foreground-color-primary)
      padding: 2rem
      margin-top: 1rem

</style>
