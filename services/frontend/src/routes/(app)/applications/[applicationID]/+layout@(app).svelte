<script lang="ts">
  interface Props {
    children?: import("svelte").Snippet;
  }
  let { children }: Props = $props();
</script>

<div class="content">
  {@render children?.()}
</div>

<style lang="sass">

</style>
