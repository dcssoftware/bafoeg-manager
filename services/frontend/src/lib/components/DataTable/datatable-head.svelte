<script lang="ts">
  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();
</script>

<thead>{@render children?.()}</thead>

<style lang="sass">
    thead
        background-color: var(--datatable-head-background-color, var(--table-head-background-color))
        color: var(--datatable-head-color, var(--table-head-color))
        font-weight: bold
        text-align: left
        border-bottom: 1px solid #e0e0e0
        :global(td)
            background-color: var(--datatable-head-background-color, var(--table-head-background-color)) !important
            color: var(--datatable-head-color, var(--table-head-color))

</style>
