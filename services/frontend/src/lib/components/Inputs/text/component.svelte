<script lang="ts">
  import { generateHash } from "$lib/functions/random/gen-hash";

  interface Props {
    label?: string | undefined;
    readOnly?: boolean;
    value?: string | null;
  }

  let {
    label = undefined,
    readOnly = false,
    value = $bindable(),
  }: Props = $props();

  const componentID = generateHash(8);
</script>

<div class="element">
  <label for={"input-label-" + componentID}>{label}</label>
  <input
    class:isEditNode={!readOnly}
    bind:value
    id={"input-label-" + componentID}
    readonly={readOnly}
    placeholder="no value given"
  />
</div>

<style lang="sass">
    .element
      display: flex
      flex-direction: column
      gap: .5rem 
      width: 100%

    input
      border: none
      outline: 0
      background: transparent
      font-color: var(--font-color)
      border-bottom: 1px solid transparent
      cursor: default
      font-size: 1rem
      width: 100%

    .isEditNode
      border-bottom-color: var(--font-color)
      

</style>
