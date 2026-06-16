<script lang="ts">
  import { _ } from "svelte-i18n";
  import {
    DataTable,
    DataTableBody,
    DataTableColumn,
    DataTableHead,
    DataTableRow,
  } from "$lib/components/DataTable";
  import { Pagination } from "$lib/components/Pagination";
  import type { Payment } from "$lib/api/payments/types";
  import { convertMoney } from "$lib/functions/convert-money";
  import { prettifyIban } from "$lib/functions/prettify-iban";
  import { dateToFormatStringLong } from "$lib/functions/date";
  import { Dialog } from "$lib/components/Dialog";
  import { Button } from "$lib/components/Button";
    import PaymentFlow from "./payment-flow.svelte";

  interface Props {
    paymentData: Payment[] | undefined;
  }

  let { paymentData }: Props = $props();

  let isOpen: boolean = $state(false);
</script>

{#snippet createPaymentDialog()}
  <div>
    <label>
      <span>IBAN</span>
      <input type="text" name="" id="" />
    </label>
    <label>
      <span>BIC</span>
      <input type="text" name="" id="" />
    </label>
    <label>
      <span>Betrag</span>
      <input type="text" name="" id="" />
    </label>
    <label>
      <span>Verwendungszweck</span>
      <input type="text" name="" id="" />
    </label>
    <Button>Überweisung erstellen</Button>
  </div>
{/snippet}

<Dialog bind:isOpen dialogContent={createPaymentDialog} />

<div class="header-with-actions">
  <h2>Zahlungsverlauf</h2>
  <div class="actions">
    <Button onclick={() => (isOpen = !isOpen)}>Überweisung hinzufügen</Button>
  </div>
</div>
{#if paymentData !== undefined}
  <DataTable>
    <DataTableHead>
      <DataTableRow>
        <DataTableColumn
          >{$_("page.application.overview.datatable.count")}</DataTableColumn
        >
        <DataTableColumn>Datum</DataTableColumn>
        <DataTableColumn>Art</DataTableColumn>
        <DataTableColumn>Status</DataTableColumn>
        <DataTableColumn>Betrag</DataTableColumn>
        <DataTableColumn>Konto</DataTableColumn>
        <DataTableColumn>Ausgeführt am</DataTableColumn>
        <DataTableColumn></DataTableColumn>
      </DataTableRow>
    </DataTableHead>
    <DataTableBody>
      {#each paymentData as payment, i}
        <DataTableRow>
          <DataTableColumn>{i + 1}</DataTableColumn>
          <DataTableColumn
            >{dateToFormatStringLong(payment.created)}</DataTableColumn
          >
          <DataTableColumn>{payment.direction}</DataTableColumn>
          <DataTableColumn>{payment.statusIdentifier}</DataTableColumn>
          <DataTableColumn>{convertMoney(payment.amount)}</DataTableColumn>
          <DataTableColumn>{prettifyIban(payment.iban)}</DataTableColumn>
          <DataTableColumn
            >{dateToFormatStringLong(payment.executed)}</DataTableColumn
          >
          <DataTableColumn>Abbrechen</DataTableColumn>
        </DataTableRow>
      {/each}
    </DataTableBody>
  </DataTable>
  <Pagination itemsPerPage={15} itemsMaxCount={25} selectedPage={1} />
{/if}
