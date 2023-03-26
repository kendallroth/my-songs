<script lang="ts" setup>
import { toTypedSchema } from "@vee-validate/yup";
import { useForm } from "vee-validate";

import AuthLayout from "~/layouts/AuthLayout.vue";
import { useLoginMutation } from "~/slices/account/queries";

import { loginSchema } from "./schema";

definePageMeta({
  middleware: "no-auth-route",
});

const { notifyError, notifyNotImplemented } = useSnackbar();

const loginMutation = useLoginMutation();
/** Track whether page should be locked once redirection is initiated */
const pendingRedirect = ref(false);

const { handleSubmit, isSubmitting, setFieldError } = useForm({
  validationSchema: toTypedSchema(loginSchema),
  initialValues: {
    email: "",
    password: "",
  },
});

const handleLogin = handleSubmit(async (data) => {
  try {
    await loginMutation.mutateAsync(data);
  } catch (e: unknown) {
    notifyError(getError(e));
    setFieldError("email", "Invalid credentials");
    return;
  }

  pendingRedirect.value = true;
  window.location.replace("/");
});
</script>

<template>
  <AuthLayout title="Login">
    <LayoutStack is="form" align-items="stretch" :spacing="4" @submit.prevent="handleLogin">
      <TextField :disabled="pendingRedirect" label="Email" name="email" />
      <TextField :disabled="pendingRedirect" label="Password" name="password" />
      <VBtn
        block
        color="primary"
        :loading="pendingRedirect || isSubmitting"
        rounded
        size="large"
        type="submit"
      >
        Login
      </VBtn>
    </LayoutStack>
    <template #append>
      <VBtn
        :disabled="pendingRedirect || isSubmitting"
        variant="text"
        @click="notifyNotImplemented"
      >
        Forgot Password?
      </VBtn>
    </template>
  </AuthLayout>
</template>
