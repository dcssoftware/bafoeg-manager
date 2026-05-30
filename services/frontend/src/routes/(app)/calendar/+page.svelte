<script lang="ts">
  import { Calendar, Editor, WillowDark } from "@svar-ui/svelte-calendar";
  import { Locale } from "@svar-ui/svelte-core";
  import { de } from "@svar-ui/calendar-locales";
  import { de as deCore } from "@svar-ui/core-locales";

  const words = { ...de, ...deCore };
  const events = [
    {
      id: 1,
      start: new Date(2026, 4, 5, 9, 0),
      end: new Date(2026, 4, 5, 10, 0),
      text: "Standup",
    },
  ];
  let api = $state();
  const date = new Date();
</script>

<div class="root">

  <h1>⚠️ Noch in Entwicklung // Work in Process</h1>

  <WillowDark>
    <Locale {words}>
      <Calendar {events} {date} />
      {#if api}
        <Editor {api} />
      {/if}
    </Locale>
  </WillowDark>
</div>

<style lang="sass">
  .root
    max-width: 100%
</style>
