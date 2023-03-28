<script lang="ts" setup>
import { VTable } from "vuetify/components";

import { type PaginationResponse } from "~/api/types";

type TableProps = InstanceType<typeof VTable>["$props"];

interface SimpleTableProps {
  density?: TableProps["density"];
  // NOTE: Do not specify a default value for 'error' prop (causes issues)
  // eslint-disable-next-line vue/require-default-prop
  error?: unknown;
  empty?: boolean;
  emptyMessage?: string;
  fetching?: boolean;
  loading?: boolean;
  pagination?: PaginationResponse;
}

const props = withDefaults(defineProps<SimpleTableProps>(), {
  density: "default",
  empty: false,
  emptyMessage: "No items found",
  fetching: false,
  loading: false,
  pagination: undefined,
});

const emit = defineEmits<{
  (event: "update:pagination", page: number): void;
}>();

const paginationString = computed(() => `${props.pagination?.totalItems ?? 0} total`);
</script>

<template>
  <VSheet :elevation="1" rounded>
    <ProgressLoader v-if="loading" />
    <VAlert v-else-if="empty" type="warning">{{ emptyMessage }}</VAlert>
    <template v-else-if="!empty">
      <VTable :density="density">
        <thead>
          <tr>
            <slot name="head" />
          </tr>
        </thead>
        <tbody>
          <slot name="body" />
        </tbody>
      </VTable>
      <ActionBar v-if="pagination" class="pa-1" right>
        <p class="text-caption">{{ paginationString }}</p>
        <VPagination
          density="compact"
          :disabled="fetching"
          :length="pagination.totalPages"
          :model-value="pagination.page"
          :total-visible="5"
          @update:model-value="(value) => emit('update:pagination', value)"
        />
      </ActionBar>
    </template>
    <VAlert v-else-if="error" type="error">
      {{ getError(error) }}
    </VAlert>
  </VSheet>
</template>
