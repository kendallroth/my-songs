import { pocketbase } from "~/api";

/** Prevent viewing secured pages while unauthenticated */
export default defineNuxtRouteMiddleware(() => {
  if (!pocketbase.authStore.isValid) {
    return navigateTo("/auth/login");
  }
});
