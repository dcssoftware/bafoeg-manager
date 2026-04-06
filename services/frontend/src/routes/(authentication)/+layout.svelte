<script lang="ts">
  import "$lib/theme/index.sass";
  import "$lib/theme/variables.scss";
  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();
</script>

<div class="background">
  <div class="content">
    {@render children?.()}
  </div>
</div>

<style lang="sass">
    :global(body)
        margin: 0
        padding: 0
    .background
        width: 100%
        height: 100%
        display: block
        position: fixed
        // background-image: linear-gradient( 73.1deg,  #5f7eb0 8%, #22334f 86.9%) !important
        background-color: var(--background-color)
    .content
        position: absolute
        top: 50%
        left: 50%
        transform: translate(-50%, -50%)
        // background: var(--background-color-white)
        background: var(--background-color-tertiary)
        color: var(--font-color)
        // color: var(--font-color-black)
        display: inline-block
        padding: 2rem
        border-radius: 10px
        min-width: 35vw
</style>
