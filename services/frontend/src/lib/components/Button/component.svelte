<script lang="ts">
  import type { ButtonType } from "./types";
  import { Button as BitsUiButton } from "bits-ui";
  import "./component.sass";

  interface Props {
    href?: string | undefined;
    buttonType?: ButtonType;
    disabled?: boolean | undefined;
    label?: string;
    linebreak?: boolean;
    download?: boolean | undefined;
    onclick?: () => void;
    children?: import("svelte").Snippet;
  }

  let {
    href = undefined,
    buttonType = "default",
    label = "",
    download = undefined,
    linebreak = false,
    disabled,
    children,
    onclick = () => {},
  }: Props = $props();
</script>

<div class="dcs-theme">
  <BitsUiButton.Root
    {download}
    {onclick}
    {disabled}
    {href}
    class="button button-type-{buttonType} {linebreak ? '' : 'no-wrap'}"
    >{label}{@render children?.()}</BitsUiButton.Root
  >
</div>

<style lang="sass">
  :global(button)
    // background-color: var(--button-background-color)
  :global(button.no-wrap)
    white-space: nowrap

</style>
