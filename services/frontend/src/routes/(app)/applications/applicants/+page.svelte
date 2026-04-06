<script lang="ts">
  import { _ } from "svelte-i18n";
  import { ApplicantSelect } from "$lib/components/ApplicantSelect";
  import { goto } from "$app/navigation";
  import { Button } from "$lib/components/Button";
  import { Dialog } from "$lib/components/Dialog";
  import { TextInput } from "$lib/components/Inputs";
  import type { CreateApplicantModel } from "$lib/api/applicants/types/create-applicant";
  import { createApplicant } from "$lib/api/applicants/create-applicant";

  let isOpenCreateApplicantDialog: boolean = $state(false);
  let someHash: string = $state("");

  let createApplicantModel: CreateApplicantModel = {
    firstname: "",
    lastname: "",

    street: "",
    houseNumber: "",
    zipCode: "",
    city: "",
    country: "",
  };

  async function reloadPage() {
    await loadData(Math.random().toString(36).substring(2, 15));
  }

  async function submitCreateApplicant() {
    try {
      await createApplicant(createApplicantModel);
    } catch (error) {
      alert("Fehler beim Erstellen des Antragstellers");
    }
  }

  function loadData(hash: string) {
    someHash = hash;
  }
</script>

{#key someHash}
  {#snippet CreateApplicantDialog()}
    <div class="form create-applicant">
      <TextInput
        label={"Vorname"}
        bind:value={createApplicantModel.firstname}
      />
      <TextInput
        label={"Nachname"}
        bind:value={createApplicantModel.lastname}
      />
      <TextInput label={"Straße"} bind:value={createApplicantModel.street} />
      <TextInput
        label={"Hausnummer"}
        bind:value={createApplicantModel.houseNumber}
      />
      <TextInput label={"PLZ"} bind:value={createApplicantModel.zipCode} />
      <TextInput label={"Stadt"} bind:value={createApplicantModel.city} />
      <TextInput label={"Land"} bind:value={createApplicantModel.country} />
      <Button onclick={() => submitCreateApplicant()}>Nutzer erstellen</Button>
    </div>
  {/snippet}

  <Dialog
    bind:isOpen={isOpenCreateApplicantDialog}
    dialogContent={CreateApplicantDialog}
  ></Dialog>

  <div class="header">
    <h1>
      {$_("page.applicant.overview.header")}
    </h1>
  </div>

  <Button onclick={() => (isOpenCreateApplicantDialog = true)}
    >Neuer User</Button
  >
  <ApplicantSelect
    onSingleSelect={(id: string) => goto("/applications/applicants/" + id)}
  />
{/key}
