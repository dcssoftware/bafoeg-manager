<script lang="ts">
  import {
    IconCartFlatbed,
    IconHeadset,
    IconStudent,
  } from "$lib/components/Icons";
  import { userState } from "$lib/states/user";
  import { _ } from "svelte-i18n";

  async function loadData() {}

  $effect(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>{$_("page.titles.home")}</title>
</svelte:head>

<div class="banner">
  <div class="alert-warning warning-beta">
    <strong>Beta Version</strong>
    <span>
      Diese Anwendung befindet sich derzeit noch in der Beta-Phase und kann
      Fehler sowie Unvollständigkeiten enthalten. <br /><br /> Bitte melden Sie etwaige
      Probleme oder Feedback über die Support-Optionen unten.
    </span>
  </div>
</div>

<div class="welcome">
  <h1>Willkommen, {$userState?.displayName}</h1>

  <div class="cards">
    <div class="card">
      <div class="head">
        <IconStudent />
        <span>Anträge</span>
      </div>
      <div class="suboptions">
        <ul>
          <li>
            <a href="/applications">
              <span>Deine Anträge</span>
              <span>99+</span>
            </a>
          </li>
          <li>
            <a href="/applications">
              <span>Deine Abteilungs- Anträge</span>
              <span>99+</span>
            </a>
          </li>
          <li>
            <a href="/applications">
              <span>Alle Anträge</span>
              <span>99+</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
    <div class="card">
      <div class="head">
        <IconCartFlatbed />
        <span>E-Akte</span>
      </div>
      <div class="suboptions">
        <ul>
          <li>
            <a href="/eakte">
              <span>E-Akte Übersicht</span>
            </a>
          </li>
          <li>
            <a href="/eakte">
              <span>E-Akten zuweisen</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
    <div class="card">
      <div class="head">
        <IconHeadset />
        <span>Support</span>
      </div>
      <div class="suboptions">
        <ul>
          <li>
            <a
              href="mailto:support@bafoegdigitalmanager.de?subject=%5BFeature-Request%5D"
            >
              <span>Feature Request</span>
            </a>
          </li>
          <li>
            <a
              href="mailto:support@bafoegdigitalmanager.de?subject=%5BBug-Report%5D"
            >
              <span>Bug Report</span>
            </a>
          </li>
          <li>
            <a
              href="mailto:support@bafoegdigitalmanager.de?subject=%5BSupport%5D"
            >
              <span>Need Help?</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
  .banner
    margin-bottom: 5rem
    .alert-warning
      background-color: var(--color-green-80)
      color: var(--font-color)
      display: flex
      flex-direction: column
      gap: 1rem
      padding: 2rem
      border-radius: 5px
      font-size: 1rem

      strong
        font-size: 1.5rem
        font-weight: bold

  .cards
    display: flex
    margin-top: 3rem
    gap: 2rem
    flex-wrap: wrap

    .card
      background-color: var(--background-color-tertiary);
      padding: 1rem
      border-radius: 10px
      display: flex
      flex-direction: column
      min-width: 350px

      .head
        font-size: 1.5rem
        font-weight: bold
        display: flex
        align-items: center
        gap: 1rem
        :global(svg)
          width: 2.5rem
          height: 2.5rem
          fill: var(--font-color)
      
      .suboptions
        ul
          list-style-type: none
          padding: 0
          display: flex
          flex-direction: column
          user-select: none

          li
            &:not(:last-child)
              border-bottom: 2px solid var(--border-color)
              
            a
              color: var(--font-color)
              text-decoration: none
              padding: .5rem .25rem
              display: block
              &:hover
                background-color: var(--background-color-secondary)
                
              &[href=""]
                cursor: not-allowed
                &:hover
                  background-color: transparent

</style>
