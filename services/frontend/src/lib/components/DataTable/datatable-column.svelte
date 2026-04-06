<script lang="ts">
  import type { TextAlign } from ".";

  interface Props {
    title?: string;
    align?: TextAlign;
    classList?: string;
    children?: import("svelte").Snippet;
  }

  let {
    title,
    children,
    align = "Left",
    classList = $bindable(),
  }: Props = $props();
</script>

<td {title} class="{classList} align{align}">{@render children?.()}</td>

<style lang="sass">
  .alignCenter
    text-align: center
  .alignRight
    text-align: right
</style>
