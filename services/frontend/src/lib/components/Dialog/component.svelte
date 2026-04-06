<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    isOpen: boolean;
    dialogContent: Snippet;
    onClose?: () => void;
  }

  let {
    isOpen = $bindable(false),
    dialogContent,
    onClose = () => {},
  }: Props = $props();

  let dialogComponent: HTMLDialogElement | undefined = $state(undefined);

  function closeDialog() {
    isOpen = false;
    onClose();
  }

  $effect(() => {
    if (dialogComponent) {
      const handleClick = (e: MouseEvent) => {
        const dialogDimensions = dialogComponent!.getBoundingClientRect();
        if (
          e.clientX < dialogDimensions.left ||
          e.clientX > dialogDimensions.right ||
          e.clientY < dialogDimensions.top ||
          e.clientY > dialogDimensions.bottom
        ) {
          closeDialog();
        }
      };

      dialogComponent.addEventListener("click", handleClick);

      return () => {
        dialogComponent?.removeEventListener("click", handleClick);
      };
    }
  });

  $effect(() => {
    if (dialogComponent && isOpen) {
      dialogComponent?.showModal();
    } else if (dialogComponent) {
      dialogComponent?.close();
    }
  });
</script>

<dialog bind:this={dialogComponent} class="dialog">
  <div class="actions">
    <button onclick={closeDialog}>Schließen</button>
  </div>
  <div class="content">
    {@render dialogContent()}
  </div>
</dialog>

<style lang="sass">
  dialog
    background-color: var(--background-color-secondary)
    border: none
    border-radius: 10px
    padding: 1rem

    .actions
      display: flex
      justify-content: flex-end
      button
        background-color: var(--color-red)
        color: var(--color-white)
        border: none
        padding: 0.5rem 1rem
        border-radius: 5px
        cursor: pointer

    .content
      margin-top: 1rem
      min-width: 50vw
      min-height: 20vh
      height: var(--dialog-height, auto)
      min-height: var(--dialog-minheight, inherit)
      width: var(--dialog-width, 50vw)

    &::backdrop
      background-color: rgba(0, 0, 0, 0.5)
      backdrop-filter: blur(2px)
      cursor: pointer
</style>
