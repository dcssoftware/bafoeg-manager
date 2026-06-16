<script lang="ts">
  import { AddressCard } from "$lib/components/Address-Card";
  import { _ } from "svelte-i18n";
  import Applications from "./components/applications.svelte";
  import PreviousPayment from "./components/previous-payment.svelte";
  import { IconHome, IconStudent } from "$lib/components/Icons";
  import * as Tabs from "$lib/components/Tabs";
  import BalanceSheet from "./components/balance-sheet.svelte";
  import { getApplicantByID, type Applicant } from "$lib/api/applicants";
  import { page } from "$app/state";
  import { getPaymentsByApplicantId } from "$lib/api/payments/get-payments-by-applicant-id";
  import type { Payment } from "$lib/api/payments/types";
  import ApplicantActions from "./components/applicant-actions.svelte";
  import { error } from "@sveltejs/kit";
    import PaymentFlow from "./components/payment-flow.svelte";

  let applicantPromise: Promise<Applicant | undefined> | undefined =
    $state(undefined);

  let paymentHistoryByApplicantPromise:
    | Promise<Payment[] | undefined>
    | undefined = $state(undefined);

  const applicantID = page.params.applicantID;

  async function loadData(applicantID: string | undefined) {
    if (!applicantID) {
      throw error(400, "Applicant ID is undefined");
    }

    applicantPromise = getApplicantByID(applicantID);
    paymentHistoryByApplicantPromise = getPaymentsByApplicantId(1, applicantID);
  }

  $effect(() => {
    loadData(applicantID);
  });
</script>

<h1>Antragsteller Übersicht</h1>

{#await applicantPromise}
  <span>{$_("states.loading")}</span>
{:then applicant}
  {#if applicant !== undefined}
    <div class="content">
      <div class="row">
        <div class="basic-information">
          <AddressCard
            IconComponent={IconHome}
            data={{
              firstname: applicant.firstname,
              lastname: applicant.lastname,
              street: applicant.address.street,
              houseNumber: applicant.address.houseNumber,
              postalCode: applicant.address.zipCode,
              city: applicant.address.city,
              country: applicant.address.country,
            }}
            header={$_("page.application.applicant.address.permanent")}
          />
          {#if applicant.trainingsAddress != null}
            <AddressCard
              IconComponent={IconStudent}
              data={{
                firstname: applicant.firstname,
                lastname: applicant.lastname,
                street: applicant.trainingsAddress.street,
                houseNumber: applicant.trainingsAddress.houseNumber,
                postalCode: applicant.trainingsAddress.zipCode,
                city: applicant.trainingsAddress.city,
                country: applicant.trainingsAddress.country,
              }}
              header={$_("page.application.applicant.address.last-training")}
            />
          {/if}
        </div>
      </div>

      <div class="row">
        <BalanceSheet
          ID={applicant.id}
          outgoingAmount={applicant.balanceOutgoing}
          packbackAmount={applicant.balancePayback}
        />
      </div>

      <div class="row">
        <Tabs.Root value="applications" class="">
          <Tabs.List>
            <Tabs.Trigger value="applications"
              >{$_(
                "page.application.applicant.tabs.tab-headers.applications"
              )}</Tabs.Trigger
            >
            <Tabs.Trigger value="payment-history"
              >{$_(
                "page.application.applicant.tabs.tab-headers.payment-history"
              )}</Tabs.Trigger
            >
            <Tabs.Trigger value="payment-diagram">
              Payment Flow
            </Tabs.Trigger>
            <Tabs.Trigger value="applicant-actions"
              >Benutzer Aktionen</Tabs.Trigger
            >
          </Tabs.List>
          <Tabs.Content value="applications">
            <Applications />
          </Tabs.Content>
          <Tabs.Content value="payment-history">
            {#await paymentHistoryByApplicantPromise}
              <span>{$_("states.loading")}</span>
            {:then payments}
              <PreviousPayment paymentData={payments} />
            {/await}
          </Tabs.Content>
          <Tabs.Content value="payment-diagram">
              <PaymentFlow />
          </Tabs.Content>
          <Tabs.Content value="applicant-actions">
            <ApplicantActions />
          </Tabs.Content>
        </Tabs.Root>
      </div>
    </div>
  {:else}
    <span>{$_("page.applicant.not-found")}</span>
  {/if}
{/await}

<style lang="sass">
  .content
    display: flex
    flex-direction: column
    gap: 2rem
    .row
      .basic-information
        display: grid
        grid-template-columns: repeat(2, 1fr)
        gap: 1rem
        width: 100%

        .address
          display: flex
          flex-direction: column
          gap: 1rem
          background-color: var(--background-color-tertiary)
          min-width: 50%
</style>
