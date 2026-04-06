<script lang="ts" module>
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { Header } from ".";
  import { initApp } from "$lib/init-app";
  import type { SelfInformation } from "$lib/api/user/types/self-information";

  import "@fontsource/inter";
  import "$lib/theme/index.sass";

  initApp();

  const args: SelfInformation = {
    id: "1234",
    username: "MAXMUSTERMANN",
    displayName: "Max Mustermann",
    email: "max.mustermann@bafög.de",
    profilePicture: "https://cataas.com/cat",
    permissions: ["admin"],
  };

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: "BAföG/Header/Header",
    component: Header,
    tags: ["autodocs"],
    args: {
      user: args,
    },
    parameters: {
      user: {
        control: {
          type: "object",
        },
        description: "The user object",
      },
    },
  });
</script>

<!-- More on writing stories with args: https://storybook.js.org/docs/writing-stories/args -->
<Story name="Primary" />
