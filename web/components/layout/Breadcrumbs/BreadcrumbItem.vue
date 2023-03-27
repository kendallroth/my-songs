<script lang="ts" setup>
export interface BreadcrumbItem {
  disabled?: boolean;
  title: string;
  to?: string;
}

withDefaults(defineProps<BreadcrumbItem>(), {
  disabled: false,
  to: undefined,
});
</script>

<template>
  <NuxtLink
    class="breadcrumbs__item text-caption"
    :disabled="disabled || disabled"
    :to="!disabled ? to : undefined"
  >
    {{ title }}
  </NuxtLink>
</template>

<style lang="scss" scoped>
@use "sass:map";
@import "vuetify/settings";

.breadcrumbs__item {
  text-decoration: none;
  color: map.get($colors, "grey", "darken-4");
  user-select: none;

  &[href] {
    color: map.get($colors, "blue", "darken-2");
    cursor: pointer;
  }

  &[disabled="true"] {
    color: map.get($colors, "grey", "base");
    cursor: not-allowed;
  }
}
</style>
