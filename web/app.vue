<script lang="ts" setup>
import "~/assets/styles/app.scss";

import { pocketbase } from "~/api";
import { useAccountStore } from "~/slices/account/store";
import { mapRecordToAccount } from "~/slices/account/utils";

// NOTE: Initially created to avoid flicker when checking auth (seems fine...)
const loadingAuth = ref(false);

const accountStore = useAccountStore();

onMounted(() => {
  if (!pocketbase.authStore.isValid) {
    pocketbase.authStore.clear();
    return;
  }

  // Initiate refresh token workflow if valid
  pocketbase
    .collection("accounts")
    .authRefresh()
    .catch(() => {
      pocketbase.authStore.clear();
    });
});

pocketbase.authStore.onChange(async (token, model) => {
  if (!model || !pocketbase.authStore.isValid) {
    // Only redirect to login page IF user was previously authenticated (prevents infinite redirects)!
    if (accountStore.authenticated) {
      window.location.replace("/auth/login");
    }

    accountStore.clearAccount();
    return;
  }

  accountStore.setAccount(mapRecordToAccount(model));
}, true);
</script>

<template>
  <VApp>
    <TheAppBar :loading="loadingAuth" />
    <VMain class="app-main">
      <NuxtLayout v-if="!loadingAuth">
        <NuxtPage />
      </NuxtLayout>
      <VProgressCircular v-else class="ma-auto" color="primary" indeterminate size="64" />
    </VMain>
    <TheAppSnackbar />
  </VApp>
</template>

<style lang="scss" scoped>
.app-main {
  display: flex;
  flex-direction: column;
}
</style>
