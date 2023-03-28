<script lang="ts" setup>
import {
  mdiHeart as mdiFavourite,
  mdiHeartOutline as mdiFavouriteEmpty,
  mdiMusic,
  mdiRefresh,
} from "@mdi/js";

import { type BreadcrumbItem } from "~/components/layout/Breadcrumbs";
import { useGetListByIdQuery, useStarListMutation } from "~/slices/list/queries";
import { type List } from "~/slices/list/types";

definePageMeta({
  middleware: "auth-route",
});

const route = useRoute();
const { notifyError } = useSnackbar();

const listId = computed(() => {
  const listIdParam = route.params["id"];
  return Array.isArray(listIdParam) ? listIdParam[0] : listIdParam;
});
const { data, error, isFetching, isLoading, ...listQuery } = useGetListByIdQuery({
  listId: listId.value,
});

const songs = computed(() => data.value?.expand?.songs ?? []);

const { isLoading: isStarring, ...starListMutation } = useStarListMutation();

const handleListStar = async (list: List) => {
  const alreadyStarred = !!list.starredAt;

  try {
    await starListMutation.mutateAsync({
      listId: listId.value,
      starred: !alreadyStarred,
    });
  } catch (e: unknown) {
    notifyError(getError(e));
    return;
  }
};

const breadcrumbs: BreadcrumbItem[] = [
  { title: "Dashboard", to: "/" },
  { title: "Lists", to: "/lists" },
  { title: "Details" },
];
</script>

<template>
  <LayoutStack align-items="stretch" class="pa-4" :spacing="4">
    <TitleBar :title="data?.title ?? 'List'">
      <template #actions>
        <VBtn
          v-if="data"
          color="error"
          density="comfortable"
          :disabled="isFetching || isStarring"
          :icon="data?.starredAt ? mdiFavourite : mdiFavouriteEmpty"
          variant="text"
          @click="handleListStar(data!)"
        />
        <VBtn
          density="comfortable"
          :disabled="isFetching"
          :icon="mdiRefresh"
          variant="text"
          @click="listQuery.refetch"
        />
      </template>
      <template #breadcrumbs>
        <Breadcrumbs :breadcrumbs="breadcrumbs" />
      </template>
    </TitleBar>
    <template v-if="data">
      <VContainer class="pa-0" fluid>
        <VRow>
          <VCol cols="4">
            <VCard class="fill-height pa-3" color="info" variant="tonal">
              <LayoutStack align-items="stretch">
                <LayoutStack direction="row" justify-content="space-between">
                  <strong>Group</strong>
                  <div>{{ data.expand?.group?.name ?? "N/A" }}</div>
                </LayoutStack>
                <LayoutStack direction="row" justify-content="space-between">
                  <strong>Director</strong>
                  <div>{{ data.director }}</div>
                </LayoutStack>
                <LayoutStack direction="row" justify-content="space-between">
                  <strong>Date</strong>
                  <div>{{ data.date ? formatDate(data.date, "MMM YYYY") : "N/A" }}</div>
                </LayoutStack>
                <LayoutStack direction="row" justify-content="space-between">
                  <strong>Rating</strong>
                  <VRating
                    color="primary"
                    density="compact"
                    half-increments
                    :length="5"
                    :model-value="data.rating / 2"
                    readonly
                    size="small"
                    :style="{ opacity: data.rating ? 1 : 0.5 }"
                  />
                </LayoutStack>
              </LayoutStack>
            </VCard>
          </VCol>
          <VCol cols="8">
            <VCard class="fill-height pa-3" color="info" variant="tonal">
              <template v-if="data.description">{{ data.description }}</template>
              <em v-else>No description</em>
              <LayoutStack v-if="data.tags?.length" class="mt-2" direction="row">
                <VChip v-for="tag in data.tags.slice(0, 3)" :key="tag" size="small" :text="tag" />
              </LayoutStack>
            </VCard>
          </VCol>
        </VRow>
      </VContainer>
    </template>
    <h3 class="text-h3">
      Songs
      <VChip
        class="ml-1"
        color="primary"
        :prepend-icon="mdiMusic"
        size="small"
        :text="songs.length.toString()"
      />
    </h3>
    <SimpleTable
      density="compact"
      :empty="!songs.length"
      :error="error"
      :fetching="isFetching"
      :loading="isLoading"
    >
      <template #head>
        <th>Title</th>
        <th>Composer</th>
        <th>Voicing</th>
        <th>Rating</th>
      </template>
      <template #body>
        <tr v-for="song in songs ?? []" :key="song.id">
          <td>
            {{ song.title }}
            <VIcon v-if="song.starredAt" color="error" :icon="mdiFavourite" :size="16" />
          </td>
          <td><em v-if="song.arranged">arr. </em>{{ song.composer }}</td>
          <td>
            <VChip size="small" :text="song.voicing" />
          </td>
          <td>
            <VRating
              color="primary"
              density="compact"
              half-increments
              :length="5"
              :model-value="song.rating / 2"
              readonly
              size="small"
              :style="{ opacity: song.rating ? 1 : 0.5 }"
            />
          </td>
        </tr>
      </template>
    </SimpleTable>
  </LayoutStack>
</template>
