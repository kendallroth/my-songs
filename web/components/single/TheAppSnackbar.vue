<script lang="ts" setup>
const { snackbar } = useSnackbar();
</script>

<template>
  <VSnackbar
    :color="(snackbar.type as string)"
    location="bottom left"
    :model-value="snackbar.visible"
  >
    {{ snackbar.text }}
  </VSnackbar>
</template>
