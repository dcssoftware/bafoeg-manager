<script lang="ts">
  import { ContentLayout } from "$lib/components/ContentLayout";

  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();
</script>

<ContentLayout mode="center-form">
  {@render children?.()}
</ContentLayout>

<style lang="sass">
    .spacing
        padding: 2rem
</style>
