<script lang="ts">
  import { allowedLanguages, getLocale } from "$lib/i18n";
  import { setLocale } from "$lib/i18n/locale";
  import { NativeSelect } from "$lib/components/NativeSelect";
  import { _ } from "svelte-i18n";

  let currentLanguage = $state(getLocale(allowedLanguages));

  function setLanguage() {
    setLocale(currentLanguage ?? "", allowedLanguages);
    location.reload();
  }
</script>

<h1>{$_("page.settings.device_header")}</h1>
<span>{$_("page.settings.device_description")}:</span>

<div class="menu">
  <NativeSelect
    data={[
      { label: "English", value: "en" },
      { label: "Deutsch", value: "de" },
    ]}
    placeholder={$_("page.settings.language.placeholder")}
    bind:value={currentLanguage}
    required
    onChange={setLanguage}
  />
  <!-- description={$_("page.settings.language.description")} -->
  <!-- label={$_("page.settings.language.label")} -->
</div>

<style lang="sass">
  span
    margin-top: .75rem
    display: block
  .menu
    margin-top: 2rem
    display: flex
    flex-direction: column
    gap: 1.5rem
</style>
