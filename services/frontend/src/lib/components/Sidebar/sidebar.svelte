<script lang="ts">
  import { _ } from "svelte-i18n";
  import { SidebarNavigation } from "$lib/components/SidebarNavigation";
  import { getNavigationItems } from "$lib/navigation/navigation-list";
  import type { NavigationItemsType } from "$lib/navigation";

  import "$lib/theme/index.sass";

  interface Props {
    navigation?: NavigationItemsType;
    url?: URL;
  }

  let {
    navigation = getNavigationItems($_),
    url = { pathname: "/" } as URL,
  }: Props = $props();
</script>

<ul class="sidebar-content">
  <li>
    <SidebarNavigation {navigation} urlInput={url} />
  </li>
  <li class="margin-top"></li>
  <li class="subscription-notice">
    <div class="subscription-notice">
      <span class="title">Beta Version</span>
      <span>Bitte keine echten, personenbezogenen Daten verarbeiten</span>
    </div>
  </li>
  <li class="legal-section">
    <hr />
    <div class="legal">
      <a href="/legal/imprint">{$_("page.navigation.imprint")}</a>
      <a href="/legal/gdpr-notice">{$_("page.navigation.gdpr-notice")}</a>
    </div>
  </li>
</ul>

<style lang="sass">
  ul.sidebar-content
    color: var(--sidebar-font-color)
    background-color: var(--sidebar-background-color)
    box-sizing: border-box
    font-size: 1em
    min-height: max(100%, 70vh)

    list-style-type: none
    display: flex
    flex-direction: column
    margin: 0
    padding: 3rem 0 1rem
    // height: 100%
    &>li
      &.margin-top
        margin-top: auto
      &.subscription-notice
        margin-bottom: 1rem
        div.subscription-notice
          display: flex
          flex-direction: column
          gap: 0.5rem
          padding: 1rem
          margin: 0 1rem
          background-color: var(--color-green-80)
          border-radius: 5px

          .title
            font-weight: 700
            font-size: 1.2rem
      &.legal-section, &.account-section
        padding: 0 1rem
      .legal
        font-size: .8rem
        display: flex
        align-items: center
        justify-content: space-around
        user-select: none

        a 
          text-decoration: none
          color: var(--sidebar-font-color-secondary)
          padding: .5rem
          transition: 200ms opacity ease-in-out
          &:hover
            opacity: .6
</style>
