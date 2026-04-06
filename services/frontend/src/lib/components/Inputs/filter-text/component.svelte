<script lang="ts">
  import { debounce } from "$lib/functions/debounce";

  interface Props {
    filterSearchTerm?: string;
    placeholder?: string;
    onFilterChange?: (value: string) => void;
  }

  let {
    filterSearchTerm = $bindable(),
    placeholder = "Suche...",
    onFilterChange,
  }: Props = $props();

  function handleFilterChange(event: Event) {
    const input = event.target as HTMLInputElement;
    if (onFilterChange) {
      onFilterChange(input.value);
    }
  }
</script>

<input
  type="text"
  {placeholder}
  bind:value={filterSearchTerm}
  onkeyup={debounce((event: Event) => handleFilterChange(event), 500)}
/>

<style lang="sass">
  input[type="text"]
    border: none
    background-color: var(--background-color-tertiary)
    color: var(--font-color)
    font-size: 1rem
    padding: .75rem 1rem
    border-radius: 5px
    width: 100%
    box-sizing: border-box
</style>
