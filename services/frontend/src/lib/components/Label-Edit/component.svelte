<script lang="ts">
  import { Button } from "$lib/components/Button";
  import { Label } from "$lib/components/Label";
  import type { LabelEditModelType } from ".";

  interface Props {
    color: LabelEditModelType;
  }

  let { color = $bindable() }: Props = $props();

  function handleDelete() {}
  function handleSave() {}
</script>

<div class="label">
  <div class="header">
    <span class="label-text">{color.name}</span>
  </div>
  <div class="picker"></div>
  <div class="preview">
    <Label {color}></Label>
  </div>
  <div class="actions">
    <Button onclick={handleDelete}>Delete</Button>
    <Button onclick={handleSave}>Save</Button>
  </div>
</div>

<style lang="sass">
  .label
    background-color: var(--background-color-secondary)
    display: flex
    flex-direction: column
    gap: 1rem
    padding: 1rem
    user-select: none
    .picker
      display: flex
      flex-direction: row
      align-items: center
      gap: 1rem
    .actions
      display: flex
      flex-direction: row
      gap: 0.5rem
</style>
