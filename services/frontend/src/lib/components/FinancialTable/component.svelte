<script lang="ts">
  import { convertMoney } from "$lib/functions/convert-money";

  interface Props {
    table: FinancialTable;
  }
  let { table }: Props = $props();

  let ergebnis: number = $state(0);

  $effect(() => {
    let erg: number = 0;
    for (let row of table.values) {
      erg += row.value;
    }
    ergebnis = erg;
  });

  function convertMoneyDisplay(value: number): string {
    let erg = convertMoney(value);
    if (value > 0) {
      erg = "+" + erg;
    }
    return erg;
  }
</script>

<div class="table">
  <h4>{table.title}</h4>
  <table>
    <tbody>
      {#each table.values as tablerow}
        <tr>
          <th>{tablerow.header}</th>
          <td>{convertMoneyDisplay(tablerow.value)}</td>
        </tr>
      {/each}
      <tr>
        <th>{table.ergebnis_title ?? "Ergebnis"}</th>
        <td>{convertMoney(ergebnis)}</td>
      </tr>
    </tbody>
  </table>
</div>

<style lang="sass">
  .table
    background-color: var(--background-color-secondary)
    padding: 1rem
    border-radius: 5px
    display: flex
    flex-direction: column
    table
      width: 100%
      margin-top: auto
      th, td
        padding: .5rem 0
        text-align: left
        border-bottom: 1px solid var(--font-color)
      td
        text-align: right
      tr:last-of-type
        th, td
          border-bottom: none
</style>
