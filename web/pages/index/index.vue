<script lang="ts" setup>
import {
  mdiAccountMultiple as mdiGroups,
  mdiViewList as mdiLists,
  mdiRefresh,
  mdiMusic as mdiSongs,
} from "@mdi/js";

import StatCard from "./StatCard.vue";
import { useStatsQuery } from "./queries";

const { data, isError, isFetching, isLoading, refetch } = useStatsQuery();

definePageMeta({
  middleware: "auth-route",
});

const stats = computed(() => ({
  groups: data.value?.groups ?? 0,
  lists: data.value?.lists ?? 0,
  songs: data.value?.songs ?? 0,
}));
</script>

<template>
  <LayoutStack align-items="stretch" class="pa-4" :spacing="4">
    <VContainer class="pa-0">
      <VRow no-gutters>
        <VCol>
          <TitleBar title="Dashboard">
            <template #actions>
              <VBtn
                density="comfortable"
                :disabled="isFetching"
                :icon="mdiRefresh"
                variant="text"
                @click="refetch"
              />
            </template>
          </TitleBar>
        </VCol>
      </VRow>
      <VRow>
        <VCol :cols="12" :sm="4">
          <StatCard
            :error="isError"
            :icon="mdiSongs"
            :loading="isLoading"
            title="Songs"
            to="/songs"
            :value="stats.songs"
          />
        </VCol>
        <VCol :cols="12" :sm="4">
          <StatCard
            :error="isError"
            :icon="mdiLists"
            :loading="isLoading"
            title="Lists"
            to="/lists"
            :value="stats.lists"
          />
        </VCol>
        <VCol :cols="12" :sm="4">
          <StatCard
            :error="isError"
            :icon="mdiGroups"
            :loading="isLoading"
            title="Groups"
            to="/groups"
            :value="stats.groups"
          />
        </VCol>
      </VRow>
    </VContainer>
  </LayoutStack>
</template>
