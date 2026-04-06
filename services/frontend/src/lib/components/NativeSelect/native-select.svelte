<script lang="ts">
  import { createBubbler } from "svelte/legacy";

  interface Props {
    data?: { label: string; value: string }[];
    value?: string | null;
    placeholder?: string;
    onChange?: () => void;
    required?: boolean;
  }

  let {
    data = [],
    value = $bindable(null),
    placeholder = "",
    required = false,
    onChange,
  }: Props = $props();
</script>

<select bind:value {placeholder} onchange={onChange} {required}>
  {#each data as dataRow}
    <option value={dataRow.value}>{dataRow.label}</option>
  {/each}
</select>
