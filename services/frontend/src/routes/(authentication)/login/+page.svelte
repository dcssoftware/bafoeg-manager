<script lang="ts">
  import { goto } from "$app/navigation";
  import { LogoAuthentik, LogoEntraId, logoPlain } from "$lib/assets";
  import { IconCaretRight } from "$lib/components/Icons";
  import { generateHash } from "$lib/functions/random/gen-hash";
  import { _ } from "svelte-i18n";
</script>

<svelte:head>
  <title>{$_("page.titles.login")}</title>
</svelte:head>

<div class="content">
  <div class="branding">
    <img src={logoPlain} alt="Logo BAföG Manager" />
  </div>
  <div class="login-provider">
    <ul>
      <li>
        <div
          class="provider authentic"
          onclick={() => goto("/api/v1/oauth/redirect?hash=" + generateHash(8))}
        >
          <img src={LogoAuthentik} alt="Authentik Logo" />
          <span>Login mit Authentik</span>
          <div class="icon">
            <IconCaretRight />
          </div>
        </div>
      </li>
      <li>
        <div
          class="provider authentic"
          onclick={() => goto("/api/v1/oauth/redirect?hash=" + generateHash(8))}
        >
          <img src={LogoEntraId} alt="EntraID Logo" />
          <span>Login mit EntraID</span>
          <div class="icon">
            <IconCaretRight />
          </div>
        </div>
      </li>
    </ul>
  </div>
</div>

<style lang="sass">
  .content
    .branding
      width: 100%
      display: flex
      justify-content: center
      img
        height: 10rem
    .login-provider
      margin-top: 4rem
      ul
        display: flex
        flex-direction: column
        gap: 1rem
        list-style-type: none
        padding: 0
        li
          background-color: var(--background-color-secondary)
          border-radius: 5px
          padding: 1rem
          &:hover 
            background-color: var(--background-color)
            cursor: pointer
          span
            font-size: 1.5rem
          img
            height: 2rem        
          .provider
            padding: .5rem 1rem
            display: flex
            align-items: center
            gap: 1rem

            .icon
              margin-left: auto
              :global(svg)
                width: 2rem
                height: 2rem
                fill: var(--font-color)
</style>
