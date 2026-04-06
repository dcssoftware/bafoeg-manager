<script module>
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { HeaderProfile } from ".";
  import { initApp } from "$lib/init-app";

  import "@fontsource/inter";
  import "$lib/theme/index.sass";

  initApp();

  const args = {
    id: "1234",
    username: "MAXMUSTERMANN",
    display_name: "Max Mustermann",
    email: "max.mustermann@bafög.de",
    profile_picture: "https://cataas.com/cat",
    permissions: ["admin"],
  };

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: "BAföG/Header/Profile",
    component: HeaderProfile,
    tags: ["autodocs"],
    args: {
      userdata: args,
    },
    parameters: {
      userdata: {
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
