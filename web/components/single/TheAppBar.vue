<script lang="ts" setup>
import { mdiMusicCircle as mdiLogo } from "@mdi/js";

import { pocketbase } from "~/api";
import { useAccountStore } from "~/slices/account/store";

type TheAppBarProps = {
  loading?: boolean;
};

withDefaults(defineProps<TheAppBarProps>(), {
  loading: false,
});

const logoutConfirmation = useDialog();

const accountStore = useAccountStore();

const handleLogout = () => {
  pocketbase.authStore.clear();
};
</script>

<template>
  <VAppBar color="primary" :elevation="2">
    <VAppBarNavIcon :icon="mdiLogo" @click="navigateTo('/')" />
    <VAppBarTitle class="app-bar__title">My Songs</VAppBarTitle>
    <template v-if="!loading" #append>
      <template v-if="accountStore.account">
        <VBtn rounded variant="outlined" @click="logoutConfirmation.show">Logout</VBtn>
        <VAvatar class="ml-4" color="white" variant="elevated">
          {{ accountStore.account.name.charAt(0).toUpperCase() }}
        </VAvatar>
      </template>
      <template v-else>
        <VBtn :active="false" rounded to="/auth/login" variant="elevated">Login</VBtn>
      </template>
      <ConfirmDialog
        v-model="logoutConfirmation.open.value"
        title="Logout?"
        @confirm="handleLogout"
      >
        Are you sure you want to logout?
      </ConfirmDialog>
    </template>
  </VAppBar>
</template>

<style lang="scss" scoped>
.app-bar__title {
  font-variant: small-caps;
}
</style>
