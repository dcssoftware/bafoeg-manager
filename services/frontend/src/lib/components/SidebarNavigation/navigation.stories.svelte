<script module lang="ts">
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { SidebarNavigation } from ".";

  import "@fontsource/inter";
  import "$lib/theme/index.sass";
  import { getNavigationItems } from "$lib/navigation";
  import { get } from "svelte/store";
  import { _ } from "svelte-i18n";
  import { initApp } from "$lib/init-app";

  initApp();

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: "BAföG/Sidebar/Navigation",
    component: SidebarNavigation,
    tags: ["autodocs"],
    args: {
      urlInput: { pathname: "/" } as URL,
      navigation: getNavigationItems(get(_)),
    },
    parameters: {
      navigation: {
        control: {
          type: "object",
        },
      },
      urlInput: {
        control: {
          type: "text",
        },
      },
    },
  });
</script>

<!-- More on writing stories with args: https://storybook.js.org/docs/writing-stories/args -->
<Story name="Primary" />
