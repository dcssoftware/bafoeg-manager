<script lang="ts">
  import type { ApplicationLabelStyleType } from "$lib/api/application-labels";
  import type { ApplicationLabelColorModelType } from "$lib/api/applications";
  import type { LabelColor } from ".";

  interface Props {
    color:
      | LabelColor
      | ApplicationLabelStyleType
      | ApplicationLabelColorModelType;
    children?: import("svelte").Snippet;
  }

  let { color, children }: Props = $props();
</script>

<div
  class="preview"
  style="background-color: light-dark({color.bgColorLight}, {color.bgColorDark}); color: light-dark({color.colorLight}, {color.colorDark}); border-color: light-dark({color.borderColorLight}, {color.borderColorDark});"
>
  <span>{@render children?.()}</span>
</div>

<style lang="sass">
  .preview 
    display: inline-block
    border-radius: 5px
    padding: 0.3rem .75rem
    font-size: .9rem
    user-select: none
    border: 1px solid
</style>
