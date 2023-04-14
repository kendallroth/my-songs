<script lang="ts" setup>
import { mdiHelpCircle as mdiDefault, mdiAlertCircle as mdiError } from "@mdi/js";

import { NuxtLink } from "#components";

interface StatCardProps {
  error?: boolean | string;
  icon?: string;
  loading?: boolean;
  title: string;
  to?: string;
  value: number;
}

withDefaults(defineProps<StatCardProps>(), {
  error: undefined,
  icon: mdiDefault,
  loading: false,
  to: undefined,
});
</script>

<template>
  <VCard :is="NuxtLink" :to="to">
    <VResponsive :aspect-ratio="4 / 3">
      <LayoutStack align-items="center" class="fill-height" justify-content="center" :spacing="4">
        <VIcon color="primary" :icon="icon" size="48" />
        <span v-if="!error" class="stat-card__value">{{ !loading ? value : "-" }}</span>
        <VIcon v-else-if="error" color="error" :icon="mdiError" size="64" />
        <VCardTitle class="stat-card__title">{{ title }}</VCardTitle>
      </LayoutStack>
      <VOverlay class="align-center justify-center" contained :model-value="loading">
        <VProgressCircular color="primary" indeterminate size="64" />
      </VOverlay>
    </VResponsive>
  </VCard>
</template>

<style lang="scss" scoped>
.layout-stack {
  position: relative;
  padding: 16px;
}

.stat-card__title {
  font-size: 2rem;
  line-height: 2rem;
  font-weight: 300;
}

.stat-card__value {
  font-size: 3rem;
  line-height: 3rem;
  font-weight: 400;
}
</style>
