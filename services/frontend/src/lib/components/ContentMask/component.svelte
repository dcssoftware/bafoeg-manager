<script lang="ts">
  interface Props {
    header?: string;
    children?: import("svelte").Snippet;
  }

  let { header, children }: Props = $props();
</script>

<div class="content-mask">
  {#if header !== undefined}
    <div class="header">
      <span>{header}</span>
    </div>
  {/if}
  <div class="content">
    {@render children?.()}
  </div>
</div>

<style lang="sass">
    .content-mask
        background-color: var(--application-form-content-background-color)
        border: 2px solid var(--application-form-header-border-color)
        border-radius: 10px
        .header
          padding: 1rem 2rem
          font-size: 1.3rem
          user-select: none
          background-color: var(--application-form-header-background-color)
          color: var(--application-form-header-color)
        .content
          padding: 2rem
</style>
