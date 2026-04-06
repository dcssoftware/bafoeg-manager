<script lang="ts">
  import { page } from "$app/stores";
  import { _ } from "svelte-i18n";
  import { slide } from "svelte/transition";
  import { changeNavigationActiveElement } from "./match-url";
  import {
    getNavigationItems,
    type NavigationItemsType,
  } from "$lib/navigation/navigation-list";
  import { userState } from "$lib/states/user";

  let urlPageStore: URL | undefined | null = $state(undefined);

  page.subscribe((value) => {
    urlPageStore = value?.url;
  });

  function updateUrl(
    fromStore: URL | undefined | null,
    fromInput: URL | undefined | null
  ) {
    if (fromStore !== undefined && fromStore !== null) {
      return fromStore;
    } else if (fromInput !== undefined && fromInput !== null) {
      return fromInput;
    }
  }

  interface Props {
    urlInput: URL | undefined | null;
    navigation?: NavigationItemsType;
  }

  let { urlInput, navigation = getNavigationItems($_) }: Props = $props();

  let activeElementIndex: number = $state(0);

  let url = $derived(updateUrl(urlPageStore, urlInput));
  $effect(() => {
    activeElementIndex = changeNavigationActiveElement(url as URL, navigation);
  });
</script>

<ul class="navigation">
  {#each navigation as nav, counter}
    {@const activeElement = counter === activeElementIndex}
    {#if nav.permission == undefined || $userState?.permissions.includes(nav.permission)}
      <li>
        <a class:active-element={activeElement} href={nav.href}>
          <div class="navigation-header">
            <div class="icon">
              <nav.icon />
            </div>
            <span>{nav.name}</span>
          </div>
        </a>
        {#if activeElement}
          <div class="subelements" transition:slide>
            {#if nav.subelements}
              {#each nav.subelements as sub}
                {#if sub.permission == undefined || $userState?.permissions.includes(sub.permission)}
                  <div class="subelement">
                    <a href={sub.href}>{sub.name}</a>
                  </div>
                {/if}
              {/each}
            {/if}
          </div>
        {/if}
      </li>
    {/if}
  {/each}
  <li style="margin-top: auto;"></li>
</ul>

<style lang="sass">

    ul.navigation
        list-style-type: none
        display: flex
        flex-direction: column
        user-select: none
        // gap: .5rem
        margin: 0
        padding: 0
        li
            a 
                text-decoration: none
                color: var(--sidebar-font-color)
                padding: .5rem 1rem
                display: block
                &:hover
                  background-color: var(--sidebar-background-color-hover)
                &.active-element
                  background-color: var(--sidebar-background-color-active)

                .navigation-header
                  display: flex
                  gap: 1rem
                  align-items: center
                  .icon
                    --icon-side: 1.2rem
                    height: var(--icon-side)
                    width: var(--icon-side)
                    :global(svg)
                      fill: var(--sidebar-font-color)
            
            .subelements
              display: flex
              flex-direction: column
              background-color: var(--sidebar-background-color-submenu)

              a
                padding-left: 1.5rem
                &:hover
                  background-color: var(--sidebar-background-color-submenu-hover)

</style>
