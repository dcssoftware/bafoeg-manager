<script lang="ts">
  import { _ } from "svelte-i18n";
  import {
    DataTable,
    DataTableBody,
    DataTableColumn,
    DataTableHead,
    DataTableRow,
  } from "$lib/components/DataTable";
  import { getUsers } from "$lib/api/user/get-users";
  import type { User } from "$lib/api/user/types/user-model";
  import { goto } from "$app/navigation";

  let userDataPromise: Promise<User[] | undefined> | undefined =
    $state(undefined);

  async function loadData() {
    userDataPromise = getUsers();
  }

  $effect(() => {
    loadData();
  });
</script>

<div class="header">
  <h1>User-Management</h1>
</div>

<div class="datatable">
  <DataTable>
    <DataTableHead>
      <DataTableRow>
        <DataTableColumn>ID</DataTableColumn>
        <DataTableColumn>Username</DataTableColumn>
        <DataTableColumn>Display Name</DataTableColumn>
        <DataTableColumn>Email</DataTableColumn>
      </DataTableRow>
    </DataTableHead>
    <DataTableBody>
      {#await userDataPromise then userData}
        {#if userData != undefined}
          {#each userData as data}
            <DataTableRow onClick={() => goto(`/user-management/${data.id}`)}>
              <DataTableColumn>{data.id}</DataTableColumn>
              <DataTableColumn>{data.username}</DataTableColumn>
              <DataTableColumn>{data.displayName}</DataTableColumn>
              <DataTableColumn>{data.email}</DataTableColumn>
            </DataTableRow>
          {/each}
        {/if}
      {/await}
    </DataTableBody>
  </DataTable>
</div>

<style lang="sass">
  .datatable
    margin-top: 3rem
</style>
