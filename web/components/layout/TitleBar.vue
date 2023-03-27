<script lang="ts" setup>
interface TitleBarProps {
  title: string;
}

withDefaults(defineProps<TitleBarProps>(), {});
</script>

<template>
  <ActionBar>
    <template #left>
      <div class="text-h3">
        <slot name="title">{{ title }}</slot>
      </div>
      <slot name="title:append" />
    </template>
    <template v-if="$slots.actions" #right>
      <slot name="actions" />
    </template>
  </ActionBar>
</template>
