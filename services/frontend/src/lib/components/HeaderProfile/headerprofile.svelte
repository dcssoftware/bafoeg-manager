<script lang="ts">
  import type { SelfInformation } from "$lib/api/user/types/self-information";
  import { logoutSession } from "$lib/functions/logout/logout";
  import { Exit, Gear } from "radix-icons-svelte";

  let showProfileMenu: boolean = $state(false);

  interface Props {
    userdata: SelfInformation;
  }

  let { userdata }: Props = $props();

  function clickOutside(
    element: HTMLElement,
    callbackFunction: CallableFunction
  ) {
    function onClick(event: MouseEvent) {
      if (!element.contains(event.target as Node)) {
        callbackFunction();
      }
    }

    document.body.addEventListener("click", onClick);

    return {
      update(newCallbackFunction: CallableFunction) {
        callbackFunction = newCallbackFunction;
      },
      destroy() {
        document.body.removeEventListener("click", onClick);
      },
    };
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  class="container"
  use:clickOutside={() => {
    showProfileMenu = false;
  }}
  onclick={() => {
    showProfileMenu = !showProfileMenu;
  }}
>
  <div class="account">
    <img
      src={userdata.profilePicture}
      alt=""
      loading="lazy"
      class="profilepicture"
    />
    <div
      class="account-name"
      title="{userdata.displayName} ({userdata.username})"
    >
      <span class="displayname">{userdata.displayName} </span>
      <span class="username">
        ({userdata.username})
      </span>
    </div>
  </div>
  <div class="menu" class:showMenu={showProfileMenu}>
    <div class="greeting">
      <div class="greeting-text">
        <span>Guten Tag,</span>
      </div>
      <div class="name">
        <span>{userdata.displayName}</span>
      </div>
    </div>
    <hr />
    <div class="options">
      <!-- svelte-ignore a11y_click_events_have_key_events -->
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <a class="option settings" href="/settings">
        <span>
          <Gear />
        </span>
        <span>Settings</span>
      </a>
      <!-- svelte-ignore a11y_click_events_have_key_events -->
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <div class="option logout" onclick={logoutSession}>
        <span>
          <Exit color="#f00" />
        </span>
        <span>Logout</span>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
.container
  position: relative
.menu
  position: absolute
  right: 0
  display: none
  padding: 1rem 2rem
  border-radius: 5px
  min-width: 150px
  background-color: var(--sidebar-background-color-secondary)

  &.showMenu
    display: block

  .greeting
    .greeting-text
      font-size: .8rem
  
  hr
    border-color: var(--sidebar-font-color-secondary)

  .options
    display: flex
    gap: 1rem
    justify-content: center
    
    .option, a.option
      font-size: .8rem
      text-decoration: none


  .option
    display: flex
    flex-direction: column
    align-items: center
    cursor: pointer

    span
      color: var(--sidebar-font-color-secondary)

.account
    // border: 1px solid #000
    cursor: pointer
    border-radius: 5px
    display: flex
    align-items: center
    gap: .5rem
    padding: .25rem .5rem

    .profilepicture
        $size: 30px
        aspect-ratio: 1 / 1
        border-radius: 100%
        display: block
        height: $size
        width: $size
        background-color: #f60
        object-fit: cover

    .account-name
        display: flex
        flex-direction: column
        gap: .1rem

        .displayname
          font-size: 1rem
          text-overflow: dots 
        .username
          font-size: .8rem
          color: var(--sidebar-font-color-secondary)

    :global(.menu button:hover)
        background-color: var(--sidebar-background-color-hover)
    :global(.menu button svg)
        color: var(--sidebar-svg-color)

</style>
