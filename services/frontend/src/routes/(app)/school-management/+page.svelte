<script lang="ts">
  import { _ } from "svelte-i18n";
  import { SchoolSelect } from "$lib/components/SchoolSelect";
  import { goto } from "$app/navigation";
</script>

<div class="header">
  <h1>{$_("page.school-management.header")}</h1>
</div>

<SchoolSelect
  onSingleSelect={(id: string) => goto("/school-management/" + id)}
/>
