<script lang="ts">
  import type { ApplicationStatusModelType } from "$lib/api/applications";
  import { updateApplicationStatusByID } from "$lib/api/applications/update-application-status-by-id";
  import { Button } from "$lib/components/Button";
  import { Dialog } from "$lib/components/Dialog";
  import { IconCalendar, IconCheck, IconXmark } from "$lib/components/Icons";

  interface Props {
    applicationID: string;
    currentStatus: ApplicationStatusModelType;
    reloadPage: () => void;
  }

  let { applicationID, currentStatus, reloadPage }: Props = $props();

  let isOpenDialogApprove = $state(false);
  let isOpenDialogDeny = $state(false);

  function allowStatusApproveRejection(status: ApplicationStatusModelType) {
    switch (status.identifier) {
      case "IN_PROGRESS":
      case "RESPONSE_AWAITED":
        return true;
      default:
        return false;
    }
  }

  async function handleApproval() {
    if (!allowStatusApproveRejection(currentStatus)) return;
    await updateApplicationStatusByID(applicationID, "APPROVED");
    reloadPage();
  }

  async function handleRejection() {
    if (!allowStatusApproveRejection(currentStatus)) return;
    await updateApplicationStatusByID(applicationID, "DENIED");
    reloadPage();
  }
</script>

{#snippet confirmApproval()}
  <div>
    <span>Möchtest du wirklich genehmigen?</span>
    <Button buttonType="submit" onclick={handleApproval}>Genehmigen</Button>
  </div>
{/snippet}

{#snippet confirmRejection()}
  <div>
    <span>Möchtest du wirklich ablehnen?</span>
    <Button buttonType="danger" onclick={handleRejection}>Ablehnen</Button>
  </div>
{/snippet}

<Dialog bind:isOpen={isOpenDialogApprove} dialogContent={confirmApproval}
></Dialog>
<Dialog bind:isOpen={isOpenDialogDeny} dialogContent={confirmRejection}
></Dialog>

<div class="actions">
  <Button
    buttonType="submit"
    disabled={!allowStatusApproveRejection(currentStatus)}
    onclick={() => (isOpenDialogApprove = true)}
  >
    <div class="content">
      <IconCheck />
      <span>Genehmigen</span>
    </div>
  </Button>
  <Button
    buttonType="danger"
    disabled={!allowStatusApproveRejection(currentStatus)}
    onclick={() => (isOpenDialogDeny = true)}
  >
    <div class="content">
      <IconXmark />
      <span>Ablehnen</span>
    </div>
  </Button>
  <Button buttonType="default" disabled>
    <div class="content">
      <IconCalendar />
      <span>Änderung anfragen (erst später verfügbar)</span>
    </div>
  </Button>
  <Button buttonType="default" disabled>
    <div class="content">
      <IconCalendar />
      <span>Absprachetermin vereinbaren (erst später verfügbar)</span>
    </div>
  </Button>
</div>

<style lang="sass">
  .actions
    display: flex
    flex-direction: column
    gap: 1rem
    padding: 1rem
    background-color: var(--background-color-tertiary)
    color: var(--color-blue);

    .content
      display: flex
      align-items: center
      gap: 0.5rem

      :global(svg)
        $size: 2rem
        height: $size
        width: $size
        fill: var(--color-white)

    :global(button)
      width: 100%
</style>
