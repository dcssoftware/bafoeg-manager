<script lang="ts">
  import type { FlexboxAlignItems } from ".";

  interface Props {
    wrap?: boolean;
    alignHorizintal?: FlexboxAlignItems;
    children?: import('svelte').Snippet;
  }

  let { wrap = false, alignHorizintal = "start", children }: Props = $props();
</script>

<div class="flexbox h-align-{alignHorizintal}" class:wrap>
  {@render children?.()}
</div>

<style lang="sass">
    .flexbox
      display: flex
      gap: var(--flexbox-gap)
      width: 100%

      &.wrap
        flex-wrap: wrap
      
      &.h-align-start
        align-items: flex-start
      
      &.h-align-center
        align-items: center

      &.h-align-end
        align-items: flex-end

</style>
