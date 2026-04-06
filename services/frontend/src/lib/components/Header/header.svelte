<script lang="ts">
  import { _ } from "svelte-i18n";
  import { HeaderProfile } from "$lib/components/HeaderProfile";
  import type { SelfInformation } from "$lib/api/user/types/self-information";
  import { logo } from "$lib/assets";

  interface Props {
    user: SelfInformation;
  }

  let { user }: Props = $props();
</script>

<div class="header">
  <a href="/" class="branding">
    <img src={logo} alt="logo" height="45" />
  </a>

  <div class="header-right"></div>
  <div class="profile">
    <HeaderProfile userdata={user} />
  </div>
</div>

<style lang="sass">
.header 
    user-select: none
    display: flex
    align-items: center
    gap: 1rem
    padding: .5rem 1rem
    color: var(--header-font-color)
    background-color: var(--header-background-color)
    a 
      text-decoration: none

    .header-right 
        margin-left: auto
        display: flex
        align-items: center
        gap: 1rem
    
    a.branding
        text-decoration: none
        color: var(--header-font-color)
        display: flex
        align-items: center
        gap: 1rem
</style>
