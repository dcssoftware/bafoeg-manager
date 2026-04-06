<script lang="ts">
  import { _ } from "svelte-i18n";
  import { GithubLogo } from "radix-icons-svelte";
  import { LinkedinLogo } from "radix-icons-svelte";
  import { getLicenses, type Licenses } from "$lib/api/general";

  let licensesPromise: Promise<Licenses | undefined> | undefined =
    $state(undefined);

  async function loadData() {
    licensesPromise = getLicenses();
  }

  $effect(() => {
    loadData();
  });
</script>

<h1>{$_("page.legal.licenses.header-software-licenses")}</h1>

<h2>{$_("page.legal.licenses.header-participating-people")}</h2>

<p>{$_("page.legal.licenses.text-participating-people")}</p>

<ul class="contibutors">
  <li>
    <span>Tim Riedl</span>
    <a href="https://www.linkedin.com/in/riedl-tim/" target="_blank"
      ><LinkedinLogo /></a
    >
    <a href="https://www.github.com/uvulpos" target="_blank"><GithubLogo /></a>
  </li>
  <li>
    <span>Mathias Rachid</span>
  </li>
  <li>
    <span>Tetiana Kryvokon</span>
    <a href="https://www.linkedin.com/in/kryvokon" target="_blank"
      ><LinkedinLogo /></a
    >
  </li>
</ul>

<h2>
  {$_("page.legal.licenses.header-software-licenses")}
</h2>

<div style="display: flex; gap: 1rem; flex-wrap: wrap;">
  <div>
    <h3>{$_("page.legal.licenses.header-frontend")}</h3>

    {#await licensesPromise}
      <span>loading...</span>
    {:then licenses}
      {#if licenses != undefined && licenses.frontend != undefined}
        <ol>
          {#each licenses.frontend as license}
            <li>{license}</li>
          {/each}
        </ol>
      {/if}
    {/await}
  </div>

  <div>
    <h3>{$_("page.legal.licenses.header-backend")}</h3>

    {#await licensesPromise}
      <span>loading...</span>
    {:then licenses}
      {#if licenses != undefined && licenses.backend != undefined}
        <ol>
          {#each licenses.backend as license}
            <li>{license}</li>
          {/each}
        </ol>
      {/if}
    {/await}
  </div>
</div>

<style>
  ul.contibutors {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
    list-style-position: inside;
  }
</style>
