<script lang="ts" setup>
import { mdiHeart as mdiFavourite, mdiRefresh } from "@mdi/js";

import { type PaginationInput } from "~/api/types";
import { useSongsQuery } from "~/slices/song/queries";

definePageMeta({
  middleware: "auth-route",
});

const paginationInput = reactive<PaginationInput>({
  page: 1,
  perPage: 10,
});

const { data, error, isFetching, isLoading, ...songsQuery } = useSongsQuery({
  pagination: paginationInput,
});
</script>

<template>
  <LayoutStack align-items="stretch" class="pa-4" :spacing="4">
    <ActionBar>
      <template #left>
        <div class="text-h3">Songs</div>
      </template>
      <template #right>
        <VBtn
          density="comfortable"
          :disabled="isFetching"
          :icon="mdiRefresh"
          variant="text"
          @click="songsQuery.refetch"
        />
      </template>
    </ActionBar>
    <SimpleTable
      :empty="!data?.items.length"
      :error="error"
      :fetching="isFetching"
      :loading="isLoading"
      :pagination="getPaginationResponse(data)"
      @update:pagination="(page) => (paginationInput.page = page)"
    >
      <template #head>
        <th>Title</th>
        <th>Composer</th>
        <th>Voicing</th>
        <th>Rating</th>
      </template>
      <template #body>
        <tr v-for="item in data?.items ?? []" :key="item.id">
          <td>
            {{ item.title }}
            <VIcon v-if="item.starredAt" color="error" :icon="mdiFavourite" :size="16" />
          </td>
          <td><em v-if="item.arranged">arr. </em>{{ item.composer }}</td>
          <td>
            <VChip :text="item.voicing" />
          </td>
          <td>
            <VRating
              color="primary"
              density="compact"
              half-increments
              :length="5"
              :model-value="item.rating / 2"
              readonly
              size="small"
              :style="{ opacity: item.rating ? 1 : 0.5 }"
            />
          </td>
        </tr>
      </template>
    </SimpleTable>
  </LayoutStack>
</template>
