<script lang="ts">
  import { ContentLayout } from "$lib/components/ContentLayout/";
  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();
</script>

<div class="layout">
  <ContentLayout mode="center-blog">
    {@render children?.()}
  </ContentLayout>
</div>

<style lang="sass">
    .layout
        :global(p)
          line-height: 1.5

        :global(a) 
          text-decoration: none
          color: var(--anchor-font-color)
        
        :global(h1)
          border-bottom: 1px solid var(--header-font-border-color)
          padding-bottom: 1rem
          margin-bottom: 2rem

</style>
