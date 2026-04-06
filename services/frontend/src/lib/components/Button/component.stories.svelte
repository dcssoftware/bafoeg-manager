<script module lang="ts">
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { Button } from ".";

  import "@fontsource/inter";
  import "$lib/theme/index.sass";

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: "BAföG/Native-Elements/Button",
    component: Button,
    tags: ["autodocs"],
    args: {
      label: "Example Button2",
      buttonType: "danger",
      href: "https://google.com",
    },
    parameters: {
      label: {
        control: {
          type: "text",
        },
        value: "Example Button",
      },
      buttonType: {
        control: {
          type: "select",
          options: ["default", "submit", "danger"],
        },
        value: "submit",
      },
      href: {
        control: {
          type: "text",
        },
        value: "htttps://google.com",
      },
    },
  });
</script>

<!-- More on writing stories with args: https://storybook.js.org/docs/writing-stories/args -->
<Story name="Primary" args={{ buttonType: "default" }} />

<Story name="Submit" args={{ buttonType: "submit" }} />

<Story name="Danger" args={{ buttonType: "danger" }} />
