<script lang="ts">
  import {
    updateApplicationAssignedUser,
    type ApplicationModelResponseType,
  } from "$lib/api/applications";
  import { dateToFormatStringLong } from "$lib/functions/date";
  import { Label } from "$lib/components/Label";
  import { _ } from "svelte-i18n";
  import { IconPlus } from "$lib/components/Icons";
  import { convertMoney } from "$lib/functions/convert-money";
  import { UserPicker } from "$lib/components/UserPicker";
  import { generateHash } from "$lib/functions/random/gen-hash";

  const i18nPath = "page.application.editor.basic-infos";

  interface Props {
    application: ApplicationModelResponseType;
    applicationReloadHash: string;
  }

  let { application, applicationReloadHash = $bindable() }: Props = $props();

  async function changeAssignedUser(userID: string) {
    await updateApplicationAssignedUser(application.id, userID);

    applicationReloadHash = generateHash(8);
  }

  function getDueDayFormat(days: number): string {
    if (days === 0) {
      return $_(i18nPath + ".processing-time.due-today");
    } else if (days === 1) {
      return $_(i18nPath + ".processing-time.due-in-one-day");
    } else if (days === -1) {
      return $_(i18nPath + ".processing-time.overdue");
    } else if (days < 0) {
      return (
        Math.abs(days) +
        " " +
        $_(i18nPath + ".processing-time.overdue-since-days")
      );
    } else {
      return days + " " + $_(i18nPath + ".processing-time.due-in-days");
    }
  }
</script>

<div class="application-info">
  <div class="row">
    <div class="requested-education">
      <span class="">{$_(i18nPath + ".requested-education")}</span>
      <div class="degree-status">
        <span class="title">{application.school.degree.name}</span>
        <span
          class="status status-identifier-{application.status.identifier
            .toLowerCase()
            .replaceAll('_', '-')}">{application.status.name}</span
        >
      </div>
    </div>
    <span class="current-vermoegenswerter-anspruch">{convertMoney(951)}</span>
  </div>
  <div class="row">
    <div>
      <div class="assigned-user">
        <span></span>
        <UserPicker
          selectedUserID={application.assignedUser?.id}
          onChange={changeAssignedUser}
        />
      </div>
    </div>
  </div>
  <div class="row">
    <div class="information">
      <div class="col aktuelle-antragsnummer">
        <span class="">{$_(i18nPath + ".current-application-id")}</span>
        <span class="title">{application.id}</span>
      </div>
      <!-- <div class="col vorherige-antragsnummer">
        <span class=""
          >{$_(i18nPath + ".previous-application.previous-bafoeg-amt")}</span
        >
        <span class="title">Bildungsbehörde Musterstadt</span>
      </div>
      <div class="col vorherige-antragsnummer">
        <span class=""
          >{$_(
            i18nPath + ".previous-application.previous-application-id"
          )}</span
        >
        <span class="title">-</span>
      </div> -->
      <div class="col spacer vorherige-antragsnummer">
        <span class="">{$_(i18nPath + ".current-application.created")}</span>
        <span class="title">{dateToFormatStringLong(application.created)}</span>
      </div>
      <div class="col vorherige-antragsnummer">
        <span class=""
          >{$_(i18nPath + ".current-application.last-updated")}</span
        >
        <span class="title">{dateToFormatStringLong(application.updated)}</span>
      </div>
      <div class="col vorherige-antragsnummer">
        <span class="">{$_(i18nPath + ".processing-time.due-at")}</span>
        <span
          class="title align-right"
          style:background-color={application.processingTime
            .remainingTimeInDays <= 0
            ? "var(--color-red)"
            : "inherit"}
          >{getDueDayFormat(
            application.processingTime.remainingTimeInDays
          )}</span
        >
      </div>
    </div>
  </div>
  <div class="row">
    <div class="labels">
      {#each application.labels as label}
        <Label color={label.color}>{label.name}</Label>
      {/each}
      <div class="add-label" title="currently not implemented">
        <IconPlus />
        <span>Label hinzufügen</span>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
  .application-info
    background-color: var(--background-color-tertiary)
    padding: 2rem 3rem
    display: flex
    flex-direction: column
    gap: 2rem

    .row 
      display: flex
      align-items: center
      gap: 2rem

      div.labels
        display: flex
        gap: 1rem

        .add-label
          display: flex
          align-items: center
          gap: 0.5rem
          user-select: none
          cursor: not-allowed

          :global(svg)
            width: 1.5rem
            height: 1.5rem
            fill: var(--font-color)

          span
            color: var(--font-color)

        // span
        //   background-color: var(--color-green-30);
        //   padding: 0.25rem .5rem
        //   border-radius: 1rem
        //   color: var(--font-color)
        //   font-size: 1rem
        //   font-weight: bold

      div.information
        display: flex
        width: 100%
        align-items: stretch
        gap: 1rem

        .spacer
          margin-left: auto !important

        .col
          display: inline-flex
          flex-direction: column
          gap: .5rem

          span
            display: inline
          
          span.align-right
            text-align: right

          span.title
            font-size: 1.2rem
            font-weight: bold

      div.requested-education
        display: flex
        flex-direction: column

        .degree-status
          display: flex
          align-items: center
          gap: 1rem

        span.title
          font-size: 1.7rem
          font-weight: bold
        
        span.status
            border-radius: 5px
            background-color: var(--color-grey-70)
            padding: 0.25rem 0.5rem
            user-select: none

            &.status-identifier-in-progress
              color: var(--font-color-white)
              background-color: var(--color-blue)
            &.status-identifier-approved
              color: var(--font-color-white)
              background-color: var(--color-green-70)
            &.status-identifier-denied
              color: var(--font-color-white)
              background-color: var(--color-red)
            &.status-identifier-response-awaited
              color: var(--font-color-white)
              background-color: var(--color-yellow)
            

      span.current-vermoegenswerter-anspruch
        color: var(--color-green)
        font-size: 1.7rem
        font-weight: bold
        margin-left: auto

</style>
