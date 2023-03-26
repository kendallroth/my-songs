<script lang="ts" setup>
// NOTE: Possibly an odd use-case for layouts, but allows passing in props. Should not be used
//         with 'NuxtLayout' though (acts as a replacement of sorts)!

type AuthLayoutProps = {
  title: string;
};

withDefaults(defineProps<AuthLayoutProps>(), {});
</script>

<template>
  <LayoutStack align-items="center" flex-grow justify-content="center">
    <VCard class="ma-4 pa-5 auth-card">
      <h2 class="text-h2 mb-5 text-center">{{ title }}</h2>
      <slot />
    </VCard>
    <template v-if="$slots.append">
      <slot name="append" />
    </template>
  </LayoutStack>
</template>

<style lang="scss" scoped>
.auth-card {
  width: 100%;
  max-width: 500px;
}
</style>
