<script module lang="ts">
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { Sidebar } from ".";
  import { _ } from "svelte-i18n";

  import "@fontsource/inter";
  import "$lib/theme/index.sass";
  import { get } from "svelte/store";
  import { getNavigationItems } from "$lib/navigation";
  import { initApp } from "$lib/init-app";

  initApp();

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: "BAföG/Sidebar/Sidebar",
    component: Sidebar,
    tags: ["autodocs"],
    args: {
      navigation: getNavigationItems(get(_)),
    },
    parameters: {
      navigation: {
        control: {
          type: "object",
        },
      },
    },
  });
</script>

<!-- More on writing stories with args: https://storybook.js.org/docs/writing-stories/args -->
<Story name="Primary" />
