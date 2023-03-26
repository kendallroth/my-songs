<script lang="ts" setup>
import { mdiHeart as mdiFavourite, mdiRefresh } from "@mdi/js";

import { type Pagination } from "~/api/types";
import { useSongsQuery } from "~/slices/song/queries";

definePageMeta({
  middleware: "auth-route",
});

const paginationInput = reactive<Pagination>({
  page: 1,
  perPage: 10,
});

const { data, error, isFetching, isLoading, ...songsQuery } = useSongsQuery({
  pagination: paginationInput,
});

const paginationString = computed(() => `${data.value?.totalItems ?? 0} total`);
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
    <VSheet :elevation="1" rounded>
      <ProgressLoader v-if="isLoading" />
      <template v-else-if="data">
        <VTable>
          <thead>
            <tr>
              <th>Title</th>
              <th>Composer</th>
              <th>Voicing</th>
              <th>Rating</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in data.items" :key="item.id">
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
                />
              </td>
            </tr>
          </tbody>
        </VTable>
        <ActionBar class="pa-1" right>
          <p class="text-caption">{{ paginationString }}</p>
          <VPagination
            density="compact"
            :disabled="isFetching"
            :length="data.totalPages"
            :model-value="data.page"
            :total-visible="5"
            @update:model-value="(value) => (paginationInput.page = value)"
          />
        </ActionBar>
      </template>
      <VAlert v-else-if="error" type="error">
        {{ getError(error) }}
      </VAlert>
    </VSheet>
  </LayoutStack>
</template>
