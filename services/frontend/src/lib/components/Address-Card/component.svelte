<script lang="ts">
  import type { Component } from "svelte";
  import type { PersonAddressCardProps, SchoolAddressCardProps } from "./types";
  import { IconExternalLink, IconPen } from "$lib/components/Icons";

  interface Props {
    header?: string | undefined;
    externalLink?: string | undefined;
    data: PersonAddressCardProps | SchoolAddressCardProps;
    isUpdatable?: boolean;
    IconComponent: Component;
    onChange?: (() => void) | undefined;
  }

  let {
    header = "DEFAULT TEXT",
    data,
    IconComponent,
    externalLink,
    isUpdatable = false,
    onChange = undefined,
  }: Props = $props();

  // Helper to check if data is a person
  function isPerson(
    data: PersonAddressCardProps | SchoolAddressCardProps
  ): data is PersonAddressCardProps {
    return "firstname" in data && "lastname" in data;
  }
</script>

<div class="address-card">
  {#if externalLink !== undefined}
    <div class="external-link">
      {#if isUpdatable && onChange !== undefined}
        <a onclick={onChange}>
          <IconPen />
        </a>
      {/if}
      <a href={externalLink}>
        <IconExternalLink />
      </a>
    </div>
  {/if}
  <div class="content">
    <div class="form-icon">
      <IconComponent />
    </div>
    <div class="address-section">
      <div class="header">
        <span>{header}</span>
      </div>
      <div class="name">
        {#if isPerson(data)}
          <span>{data.firstname} {data.lastname}</span>
        {:else}
          <span>{data.name}</span>
        {/if}
      </div>
      <div class="address">
        <div>
          <span class="street">{data.street}</span>
          <span class="housenumber">{data.houseNumber}</span>
        </div>
        <div>
          <span class="zip">{data.postalCode}</span>
          <span class="city">{data.city}</span>
        </div>
        <div>
          <span class="country">{data.country}</span>
        </div>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
  .address-card
    background-color: var(--background-color-tertiary)
    padding: 2rem 3rem
    position: relative
    .external-link
      position: absolute
      top: 2rem
      right: 3rem
      display: flex
      gap: 1rem
      :global(svg)
        $size: 1.2rem
        width: $size
        height: $size
        fill: var(--font-color)
        cursor: pointer

    .content
      flex-grow: 1
      align-items: center
      display: flex
      gap: 2rem
    .form-icon
      :global(svg)
        $icon-size: 50px
        fill: var(--font-color)
        height: $icon-size
        width: $icon-size
    .address-section
      display: flex
      flex-direction: column
      gap: 1rem
      .name
        font-size: 1.7rem
        font-weight: bold
      .address
        display: flex
        flex-direction: column
        gap: .25rem
        font-size: 1.2rem
        &>*
          display: flex
          gap: 1ch
</style>
