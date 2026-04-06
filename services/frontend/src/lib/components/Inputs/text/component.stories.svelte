<script module lang="ts">
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { default as Component } from "./component.svelte";
  import { _ } from "svelte-i18n";

  import "@fontsource/inter";
  import "$lib/theme/index.sass";
  import { initApp } from "$lib/init-app";

  initApp();

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: "BAföG/Inputs/Text",
    component: Component,
    tags: ["autodocs"],
    args: {},
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
<Story name="Default" args={{}} />
