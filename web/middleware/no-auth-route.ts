import { pocketbase } from "~/api";

/** Prevent viewing some pages while authenticated */
export default defineNuxtRouteMiddleware(() => {
  if (pocketbase.authStore.isValid) {
    return navigateTo("/");
  }
});
