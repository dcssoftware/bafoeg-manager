<script lang="ts">
  import { convertMoney } from "$lib/functions/convert-money";

  interface Props {
    ID: string;
    outgoingAmount?: number | null;
    packbackAmount?: number | null;
  }

  let { ID, outgoingAmount, packbackAmount }: Props = $props();
</script>

<div class="balance-sheet">
  <div class="col">
    <span>Antragsteller ID</span>
    <span class="id">{ID}</span>
  </div>
  <div class="spacer"></div>
  <div class="col">
    <span>Bisher ausgezahlten Betrag</span>
    <span class="amount">{convertMoney(outgoingAmount ?? 0)}</span>
  </div>
  <div class="col">
    <span>Zurückzahlungspflichtiger Betrag</span>
    <span class="amount">{convertMoney(packbackAmount ?? 0)}</span>
  </div>
</div>

<style lang="sass">
  .balance-sheet
    background-color: var(--background-color-tertiary)
    padding: 2rem
    display: flex
    flex-direction: row
    gap: 2rem

    .spacer
      margin-right: auto

    .col
      display: flex
      flex-direction: column
      gap: 1rem
      font-size: 1.2rem
      span.amount
        font-size: 2rem
        font-weight: bold
        color: var(--color-green)
        text-align: right
</style>
