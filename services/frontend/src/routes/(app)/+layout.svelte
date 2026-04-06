<script lang="ts">
  // import code
  import { _ } from "svelte-i18n";

  import { goto } from "$app/navigation";

  import Cookies from "js-cookie";

  import { AppLayout } from "$lib/components/AppLayout";
  import { getSelfInformation } from "$lib/api/user/get-self-information";
  import { refreshJwtToken } from "$lib/api/authentication/refresh-jwt-token";
  import { initApp } from "$lib/init-app";
  import { addPreventStrgSHandler } from "$lib/functions/shortcuts";

  // install fonts + styles
  import "@fontsource-variable/inter";
  import "@fontsource/roboto";
  import "$lib/theme/index.sass";
  import { refreshJwtTokenIfNeeded } from "$lib/api/authentication/refresh-jwt.token-if-needed";
  import { userState } from "$lib/states/user";
  interface Props {
    children?: import("svelte").Snippet;
  }

  let { children }: Props = $props();

  initApp();

  let preMount: boolean = $state(true);
  let pageIsLoading: boolean = $state(true);

  $effect(() => {
    onMountFunc();
  });

  async function onMountFunc() {
    setInterval(() => {
      refreshJwtTokenIfNeeded();
    }, 1000 * 5);
    preMount = false;

    // remove hash from url
    const refreshHahName = "refresh-hash";
    const url = new URL(window.location.href);

    if (url.searchParams.has(refreshHahName)) {
      url.searchParams.delete(refreshHahName);
      history.replaceState(null, "", url.toString());
    }

    // check if jwt exists
    let jwtToken = Cookies.get("jwt");

    if (jwtToken === undefined || jwtToken === "") {
      // check jwt refresh token
      const newToken = await refreshJwtToken();

      if (newToken === null || newToken === "") {
        window.location.href = "/login";
        return;
      }
    }

    let selfData = await getSelfInformation();

    if (selfData === undefined || selfData == null) {
      goto("/login");
    }

    addPreventStrgSHandler();

    userState.set(selfData);
    pageIsLoading = false;
  }
</script>

{#if pageIsLoading}
  <div>
    <p>Loading...</p>
  </div>
{:else if !pageIsLoading && $userState != null}
  {#if !preMount}
    <AppLayout user={$userState}>
      {@render children?.()}
    </AppLayout>
  {/if}
{:else}
  <div>
    <p>Unexpexted Error</p>
  </div>
{/if}

<style lang="sass">

</style>
