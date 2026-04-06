<script lang="ts">
  import { _ } from "svelte-i18n";
  import { logoPlain } from "$lib/assets";

  interface Props {
    value: string;
  }

  let { value = $bindable() }: Props = $props();
</script>

<div class="header">
  <a href="/" class="branding">
    <img src={logoPlain} alt="logo" height="45" />
  </a>
  <div class="name">
    <span>BAföG Manager</span>
  </div>
  <div class="spacer"></div>
  <div class="header-value">
    <span>{value}</span>
  </div>
</div>

<style lang="sass">
.header 
    user-select: none
    display: flex
    align-items: center
    gap: 1rem
    padding: .5rem 1rem
    color: var(--font-color)
    border-bottom: 1px solid var(--font-color)
    a 
      text-decoration: none

    a.branding
        text-decoration: none
        color: var(--font-color)
        display: flex
        align-items: center
        gap: 1rem
        height: 45px
    .spacer
      margin-left: auto
    .name
        font-size: 1.5rem
        font-weight: bold
        color: var(--font-color)
        text-decoration: none
</style>
