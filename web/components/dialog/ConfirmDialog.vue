<script lang="ts" setup>
type ConfirmDialogProps = {
  cancelText?: string;
  confirmColor?: string;
  confirmText?: string;
  /** Whether dialog should close immediately upon confirmation */
  closeOnConfirm?: boolean;
  /** Whether confirmation is disabled/pending */
  disabled?: boolean;
  modelValue: boolean;
  loading?: boolean;
  title: string;
  /**
   * Confirmation dialog type
   *
   * - "confirm" - Confirm/cancel buttons
   * - "alert" - Confirm button
   */
  type?: "confirm" | "alert";
};

const props = withDefaults(defineProps<ConfirmDialogProps>(), {
  cancelText: undefined,
  confirmColor: "primary",
  confirmText: undefined,
  closeOnConfirm: true,
  disabled: false,
  loading: false,
  type: "confirm",
});

const emit = defineEmits<{
  (event: "update:modelValue", open: boolean): void;
  (event: "confirm"): void;
  (event: "cancel"): void;
}>();

const handleCancel = () => {
  emit("cancel");
  emit("update:modelValue", false);
};

const handleConfirm = () => {
  emit("confirm");

  // Confirmation dialogs might not always close immediately after confirmation (ie. operation pending, etc)
  if (props.closeOnConfirm) {
    emit("update:modelValue", false);
  }
};
</script>

<template>
  <VDialog
    max-width="400"
    min-width="300"
    :model-value="modelValue"
    v-bind="$attrs"
    @update:model-value="(open) => emit('update:modelValue', open)"
  >
    <VCard class="w-fill">
      <VCardTitle>{{ title }}</VCardTitle>
      <VCardText v-if="$slots.default">
        <slot />
      </VCardText>
      <VCardActions>
        <ActionBar right>
          <VBtn v-if="type === 'confirm'" :disabled="loading" @click="handleCancel">
            {{ cancelText ?? "Cancel" }}
          </VBtn>
          <VBtn
            :color="confirmColor"
            :disabled="disabled"
            :loading="loading"
            @click="handleConfirm"
          >
            {{ confirmText ?? "Confirm" }}
          </VBtn>
        </ActionBar>
      </VCardActions>
    </VCard>
  </VDialog>
</template>

<style lang="scss" scoped>
@import "vuetify/settings";

// NOTE: Fix for dialog overlay not expanding (needed to expand smaller dialog content)
/*.v-dialog :deep(.v-overlay__content) {
  width: 100%;
}*/

.v-dialog .v-card-text {
  padding: calc($spacer * 2) calc($spacer * 4) !important;
}
</style>
