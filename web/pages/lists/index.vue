<script lang="ts" setup>
import { mdiHeart as mdiFavourite, mdiMusic, mdiRefresh, mdiMagnify as mdiView } from "@mdi/js";

import { type PaginationInput } from "~/api/types";
import { type BreadcrumbItem } from "~/components/layout/Breadcrumbs";
import { useGetListsQuery } from "~/slices/list/queries";

definePageMeta({
  middleware: "auth-route",
});

const paginationInput = reactive<PaginationInput>({
  page: 1,
  perPage: 10,
});

const { data, error, isFetching, isLoading, ...listsQuery } = useGetListsQuery({
  pagination: paginationInput,
});

const breadcrumbs: BreadcrumbItem[] = [{ title: "Dashboard", to: "/" }, { title: "Lists" }];
</script>

<template>
  <LayoutStack align-items="stretch" class="pa-4" :spacing="4">
    <TitleBar title="Lists">
      <template #actions>
        <VBtn
          density="comfortable"
          :disabled="isFetching"
          :icon="mdiRefresh"
          variant="text"
          @click="listsQuery.refetch"
        />
      </template>
      <template #breadcrumbs>
        <Breadcrumbs :breadcrumbs="breadcrumbs" />
      </template>
    </TitleBar>
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
        <th>Group</th>
        <th>Director</th>
        <th>Date</th>
        <th>Songs</th>
        <th>Rating</th>
        <th :style="{ width: '40px' }" />
      </template>
      <template #body>
        <v-hover v-for="item in data?.items ?? []" :key="item.id">
          <template #default="{ isHovering, props }">
            <tr v-bind="props" :class="{ 'bg-grey-lighten-5': isHovering }">
              <td>
                {{ item.title }}
                <VIcon v-if="item.starredAt" color="error" :icon="mdiFavourite" :size="16" />
              </td>
              <td>{{ item.expand?.group?.name ?? "N/A" }}</td>
              <td>{{ item.director }}</td>
              <td>{{ formatDate(item.date, "MMM YYYY") }}</td>
              <td>
                <VChip
                  color="primary"
                  :prepend-icon="mdiMusic"
                  :text="item.expand?.songs?.length.toString() ?? 'N/A'"
                />
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
              <td class="px-2">
                <VBtn
                  color="primary"
                  :icon="mdiView"
                  size="small"
                  :to="`/lists/${item.id}`"
                  variant="tonal"
                />
              </td>
            </tr>
          </template>
        </v-hover>
      </template>
    </SimpleTable>
  </LayoutStack>
</template>
