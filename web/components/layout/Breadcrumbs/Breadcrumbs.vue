<script lang="ts" setup>
import { mdiChevronRight as mdiBreadcrumbDivider } from "@mdi/js";

import { type BreadcrumbItem } from "./BreadcrumbItem.vue";

interface BreadcrumbsProps {
  breadcrumbs: BreadcrumbItem[];
  disabled?: boolean;
}

withDefaults(defineProps<BreadcrumbsProps>(), {
  disabled: false,
});
</script>

<template>
  <VSheet border class="breadcrumbs" rounded="lg">
    <LayoutStack align-items="center" direction="row">
      <template v-for="(item, idx) in breadcrumbs" :key="item.to">
        <VIcon v-if="idx > 0" class="breadcrumbs__divider" :icon="mdiBreadcrumbDivider" size="16" />
        <BreadcrumbItem :disabled="item.disabled" :title="item.title" :to="item.to" />
      </template>
    </LayoutStack>
  </VSheet>
</template>

<style lang="scss" scoped>
.breadcrumbs {
  padding: 4px 8px;
}

.breadcrumbs__divider {
  margin: 0 -4px;
}
</style>
