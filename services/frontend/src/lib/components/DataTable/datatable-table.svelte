<script lang="ts">
  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();
</script>

<table>{@render children?.()}</table>

<style lang="sass">
    table 
        border-collapse: collapse
        width: 100%

        :global(td)
            padding: 10px
</style>
