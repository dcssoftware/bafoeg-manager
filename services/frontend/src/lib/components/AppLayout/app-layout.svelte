<script lang="ts">
  import { Header } from "$lib/components/Header";
  import { Sidebar } from "$lib/components/Sidebar";
  import type { SelfInformation } from "$lib/api/user/types/self-information";
  import { onMount } from "svelte";

  interface Props {
    user: SelfInformation;
    children?: import("svelte").Snippet;
  }

  let { user, children }: Props = $props();

  let appHeaderOffsetHeight: number | undefined = $state();
  let sidebarOffsetWidth: number | undefined = $state();
  let siteContentElement: HTMLElement | undefined = $state();
  let sidebarAnchorElement: HTMLElement | undefined = $state();
  let headerAnchorElement: HTMLElement | undefined = $state();

  onMount(() => {
    setSidebarGapTop(appHeaderOffsetHeight);
    setSidebarAnchorWidth(sidebarOffsetWidth);
    setSidebarHeight(appHeaderOffsetHeight);
    setContentGapTop(appHeaderOffsetHeight);
  });

  function setSidebarAnchorWidth(height: number | undefined) {
    if (sidebarAnchorElement === null || sidebarAnchorElement === undefined) {
      return;
    }
    height = height ?? 0;
    sidebarAnchorElement.style.setProperty(
      "--sidebar-anchor-width",
      `${height}px`
    );
  }

  function setSidebarHeight(navHeight: number | undefined) {
    if (siteContentElement === null || siteContentElement === undefined) {
      return;
    }
    navHeight = navHeight ?? 0;
    siteContentElement.style.setProperty(
      "--sidebar-height",
      `calc(100vh - ${navHeight}px)`
    );
  }

  function setSidebarGapTop(height: number | undefined) {
    if (headerAnchorElement === null || headerAnchorElement === undefined) {
      return;
    }
    height = height ?? 0;
    headerAnchorElement.style.setProperty("--navbar-gap-top", `${height}px`);
  }

  function setContentGapTop(height: number | undefined) {
    if (siteContentElement === null || siteContentElement === undefined) {
      return;
    }
    height = height ?? 0;
    siteContentElement.style.setProperty("--sidebar-gap-top", `${height}px`);
  }
  $effect(() => {
    setSidebarGapTop(appHeaderOffsetHeight);
  });
  $effect(() => {
    setSidebarAnchorWidth(sidebarOffsetWidth);
  });
  $effect(() => {
    setSidebarHeight(appHeaderOffsetHeight);
  });
  $effect(() => {
    setContentGapTop(appHeaderOffsetHeight);
  });
</script>

<div class="app-layout">
  <div class="navbar-anchor" bind:this={headerAnchorElement}>
    <nav bind:offsetHeight={appHeaderOffsetHeight}>
      <Header {user} />
    </nav>
  </div>
  <div class="site-content" bind:this={siteContentElement}>
    <div class="sidebar-anchor" bind:this={sidebarAnchorElement}>
      <sidebar bind:offsetWidth={sidebarOffsetWidth}>
        <Sidebar />
      </sidebar>
    </div>
    <main>
      {@render children?.()}
    </main>
  </div>
</div>

<style lang="sass">
  .app-layout
    min-height: 100vh
    display: flex
    flex-direction: column

    .block 
      display: block
      width: 100%
      text-align: center
      // background-color: var(--color-red)
      &.block-production-red
        color: var(--color-white)
        background-color: var(--color-red)

    .site-content
      display: flex
      flex-grow: 1
      overflow-y: scroll
      align-items: stretch

  .navbar-anchor
    height: var(--navbar-gap-top)
    width: 100%
    display: block
  
  nav
    position: fixed
    top: 0
    left: 0
    width: 100%
    z-index: 150

  .sidebar-anchor
    width: var(--sidebar-anchor-width)

  sidebar
      display: block !important
      width: 0px
      z-index: 100
      position: fixed
      top: calc(-1px + var(--sidebar-gap-top)) // -1px because otherwise it missmatches with the navbar by one pixel
      width: var(--sidebar-width) !important
      height: calc(1px + var(--sidebar-height)) !important

  main
    margin: 2rem 3rem
    overflow-y: visible
    width: calc(100% - var(--sidebar-width))
  
  @media print
    .navbar-anchor
      display: none 
    .sidebar-anchor
      display: none
    main
      margin: 0
      width: 100%
</style>
