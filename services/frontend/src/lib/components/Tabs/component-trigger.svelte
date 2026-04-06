<script lang="ts">
  import { Tabs as TabsPrimitive } from "bits-ui";

  interface Props {
    value: string;
    onclick?: () => void;
    onkeydown?: () => void;
    onfocus?: () => void;
    children?: any;
  }

  let {
    children,
    value = $bindable(),
    onclick,
    onkeydown,
    onfocus,
  }: Props = $props();
</script>

<!-- 
cn(
    "ring-offset-background 
    focus-visible:ring-ring 
    data-[state=active]:bg-background 
    data-[state=active]:text-foreground 
    inline-flex 
    items-center 
    justify-center 
    whitespace-nowrap 
    rounded-md 
    px-3 
    py-1 
    text-sm 
    font-medium 
    transition-all 
    focus-visible:outline-none 
    focus-visible:ring-2 
    focus-visible:ring-offset-2 
    disabled:pointer-events-none 
    disabled:opacity-50 
    data-[state=active]:shadow",
    className
  )
-->

<div class="dcs-theme">
  <TabsPrimitive.Trigger
    class="tab-trigger"
    {value}
    {onclick}
    {onkeydown}
    {onfocus}
  >
    {@render children?.()}
  </TabsPrimitive.Trigger>
</div>

<style lang="sass">
  .dcs-theme
    :global(.tab-trigger)
      border: none
      background-color: var(--background-color-tertiary)
      font-size: 1em
      padding: .5rem 1rem
      cursor: pointer
    :global(.tab-trigger[data-state=active])
      background-color: var(--color-green)
      color: var(--font-color-white)
      border-radius: 5px
</style>
