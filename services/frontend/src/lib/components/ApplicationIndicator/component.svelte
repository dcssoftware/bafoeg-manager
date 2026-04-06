<script lang="ts">
  interface Props {
    indicator: number;
  }
  let { indicator }: Props = $props();
</script>

<div class="field-id">
  <span>{indicator}</span>
</div>

<style lang="sass">
  .field-id
    background-color: var(--application-inputbox-author-identifier-color-applicant)
    padding: .5rem
    width: calc((.5rem * 2 ) + 2ch)
    display: flex
    align-items: stretch
    justify-content: center
    border-radius: 5px

    span
      display: inline-block
      color: var(--color-white)
</style>
