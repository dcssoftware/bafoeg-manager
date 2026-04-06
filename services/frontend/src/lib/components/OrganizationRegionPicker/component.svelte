<script lang="ts">
  import { _ } from "svelte-i18n";
  import { Select } from "bits-ui";
  import "./component.sass";
  import { IconCaretDown, IconUserSlah } from "$lib/components/Icons";
  import { defaultUserID, defaultEventChangeSelectedUserID } from "./component";
  import { onMount } from "svelte";
  import { getOrganizationRegions } from "$lib/api/organization/get-regions";
  import type { RegionResponseModelType } from "$lib/api/organization/models/region";

  // Workaround for Bits UI not supporting null/empty values
  // Using a special string value to represent "no user selected"
  const NO_USER_SELECTED_VALUE = "__no_user_selected__";

  interface Props {
    selectedUserID: string | null;
    onChange?: (userID: string | null) => void;
  }

  let {
    onChange = defaultEventChangeSelectedUserID,
    selectedUserID = $bindable<string | null>(defaultUserID),
  }: Props = $props();

  // Internal state that Bits UI can work with (always a string)
  let internalSelectedValue = $state<string>(NO_USER_SELECTED_VALUE);

  let regionsPromise: Promise<RegionResponseModelType | undefined> | undefined =
    $state(undefined);
  let isOpen: boolean | undefined = $state(false);

  // Sync internal value with external prop
  $effect(() => {
    internalSelectedValue = selectedUserID || NO_USER_SELECTED_VALUE;
  });

  export async function loadData() {
    regionsPromise = getOrganizationRegions(1);
  }

  async function handleUserOnChange(value: string) {
    // Convert Bits UI value back to external representation
    const newSelectedUserID = value === NO_USER_SELECTED_VALUE ? null : value;
    selectedUserID = newSelectedUserID;

    // Call external onChange with proper null value
    onChange(newSelectedUserID);
  }

  onMount(() => {
    if (selectedUserID === null) {
      selectedUserID = NO_USER_SELECTED_VALUE;
    }
  });

  $effect(() => {
    loadData();
  });
</script>

{#await regionsPromise then promiseItems}
  {#if promiseItems}
    {@const regions = [
      {
        id: NO_USER_SELECTED_VALUE,
        identifier: "NO_REGION",
        name: "No Region Selected",
      },
      ...(promiseItems.regions ?? []),
    ]}
    {@const selected = regions.find(
      (region) => region.id == internalSelectedValue
    )}
    <div title={selected?.id} class="dcs-theme">
      <Select.Root
        type="single"
        value={internalSelectedValue}
        onValueChange={handleUserOnChange}
        bind:open={isOpen}
      >
        <Select.Trigger class="user-picker-trigger {isOpen ? 'open' : ''}">
          <!-- {selectedLabel ? selectedLabel : placeholder} -->

          <div class="user-picker-trigger-content">
            <div class="name">
              <span>{selected?.name} </span>
            </div>
            <div class="icon">
              <IconCaretDown />
            </div>
          </div>
        </Select.Trigger>
        <Select.Portal>
          <Select.Content class="dcs-theme user-picker-content">
            <div class="user-picker-scrollbar">
              <Select.ScrollUpButton>up</Select.ScrollUpButton>
              <Select.Viewport class="user-picker-viewport">
                {#each regions as user}
                  <Select.Item
                    value={user.id}
                    label={user.name}
                    disabled={false}
                    class="user-picker-item"
                  >
                    {#snippet children({ selected })}
                      <div class="row">
                        <div class="name">
                          <span>
                            {user.name}
                          </span>
                        </div>
                        <div class="checkbox">
                          <span>
                            {selected ? "✅" : ""}
                          </span>
                        </div>
                      </div>
                    {/snippet}
                  </Select.Item>
                {/each}
              </Select.Viewport>
              <Select.ScrollDownButton>down</Select.ScrollDownButton>
            </div>
          </Select.Content>
        </Select.Portal>
      </Select.Root>
    </div>
  {:else}
    <div class="dcs-theme">
      <div class="user-picker-trigger-content">
        <div class="image">
          <IconUserSlah />
        </div>
        <div class="name">
          <span>{$_("components.user-picker.no-user-selected")}</span>
        </div>
        <div class="icon">
          <IconCaretDown />
        </div>
      </div>
    </div>
  {/if}
{/await}

<style lang="sass">

  :global(.dcs-theme .user-picker-trigger-trigger)
    border: 1px solid transparent
    
  :global(.dcs-theme .user-picker-trigger-trigger.open)
    border-color: var(--font-color)

  // .user-input-search
  //   border: none
  //   outline: transparent
  //   padding: 0.5rem 1rem
  //   margin-bottom: 2rem
  :global(.dcs-theme.user-picker-content)
    max-height: 450px
    overflow-y: scroll
    background-color: var(--background-color-secondary)
    border: 1px solid var(--font-color)

  :global(.dcs-theme .user-picker-item)
    cursor: pointer

  :global(.dcs-theme .user-picker-item[data-selected])
    background-color: var(--color-blue-20)
  :global(.dcs-theme .user-picker-item[data-highlighted])
    background-color: var(--color-yellow-20) !important

</style>
