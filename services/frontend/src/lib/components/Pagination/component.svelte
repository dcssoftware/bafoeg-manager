<script lang="ts">
  import { Pagination } from "bits-ui";

  interface Props {
    itemsPerPage?: number;
    itemsMaxCount?: number;
    selectedPage: number;
  }

  let {
    itemsMaxCount = 0,
    itemsPerPage = 10,
    selectedPage = $bindable(),
  }: Props = $props();
</script>

<div class="pagination">
  <Pagination.Root
    count={itemsMaxCount}
    perPage={itemsPerPage}
    bind:page={selectedPage}
  >
    {#snippet children({ pages, range })}
      <div class="pagination-items">
        <Pagination.PrevButton class="pagination-item">
          <span>Prev</span>
        </Pagination.PrevButton>
        <div class="pagination-items">
          {#each pages as page (page.key)}
            {#if page.type === "ellipsis"}
              <div class="">...</div>
            {:else}
              <Pagination.Page {page} class="pagination-item">
                {page.value}
              </Pagination.Page>
            {/if}
          {/each}
        </div>
        <Pagination.NextButton class="pagination-item">
          <span>Next</span>
        </Pagination.NextButton>
      </div>
      <p class="info">
        Showing {range.start} - {range.end}
      </p>
    {/snippet}
  </Pagination.Root>
</div>

<style lang="sass">

  .pagination
    display: flex
    justify-content: center
    margin-top: 5rem
    gap: 0.5rem
  .pagination
    :global(.pagination-items)  
      display: flex
      gap: 1rem
      align-items: center
    :global(.pagination-item)
      border: none
      font-size: 1.2rem
      padding: 0.5rem 1rem
      background-color: var(--button-background-color, var(--background-color-tertiary))
      border-radius: 5px
    :global(.info)
      text-align: center
</style>
