<script lang="ts" setup>
interface TitleBarProps {
  title?: string;
}

withDefaults(defineProps<TitleBarProps>(), {
  title: undefined,
});
</script>

<template>
  <ActionBar>
    <template #left>
      <LayoutStack align-items="center" class="text-h3" direction="row">
        <slot name="title">{{ title }}</slot>
      </LayoutStack>
      <slot name="title:append" />
    </template>
    <template v-if="$slots.actions" #right>
      <slot name="actions" />
    </template>
  </ActionBar>
  <div v-if="$slots.breadcrumbs" class="align-self-start mt-n3">
    <slot name="breadcrumbs" />
  </div>
</template>
