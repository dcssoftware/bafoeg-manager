<script lang="ts">
  import { CalendarDateTime } from "@internationalized/date";
  import { DateRangeField, type DateRange } from "bits-ui";

  let myValue = $state<DateRange>({
    start: new CalendarDateTime(2024, 8, 3, 12, 30),
    end: new CalendarDateTime(2024, 8, 4, 12, 30),
  });
</script>

<div class="dcs-theme">
  <DateRangeField.Root
    disabled
    class="date-range-field"
    locale="de-DE"
    value={myValue}
    granularity="day"
  >
    <DateRangeField.Label></DateRangeField.Label>
    <div class="input-components">
      {#each ["start", "end"] as const as type}
        <DateRangeField.Input {type} class="input">
          {#snippet children({ segments })}
            {#each segments as { part, value }}
              <DateRangeField.Segment {part}>
                {value}
              </DateRangeField.Segment>
            {/each}
          {/snippet}
        </DateRangeField.Input>
      {/each}
    </div>
  </DateRangeField.Root>
</div>

<style lang="sass">
  .dcs-theme
    :global(.input-components)
      background-color: var(--background-color-secondary)
      padding: .5rem 1rem
      border-radius: 5px
      display: flex
      gap: 1rem
      margin-top: 1rem
    :global(.input)
      border: none

</style>
