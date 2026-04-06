<script lang="ts">
  let { onClick = undefined, children } = $props();
</script>

<tr class:notClickable={onClick === undefined} onclick={onClick}
  >{@render children?.()}</tr
>

<style lang="sass">
    tr:nth-child(even) 
        background-color: var(--datatable-row-even-background-color, var(--table-row-even-background-color))

    tr:nth-child(odd) 
        background-color: var(--datatable-row-odd-background-color, var(--table-row-odd-background-color))

    tr:hover
        background-color: var(--datatable-row-hover-background-color, var(--table-row-hover-background-color))
        cursor: pointer
    tr.notClickable:hover
        cursor: default !important

</style>
