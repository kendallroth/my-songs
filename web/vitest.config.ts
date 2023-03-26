import { defineVitestConfig } from "nuxt-vitest/config";

// NOTE: Have run into 'CSS.supports is not a function' Vuetify error when attempting to build Nuxt context...

export default defineVitestConfig({
  test: {
    deps: {
      inline: ["vuetify"],
    },
    // Optionally enable Nuxt environment in all tests (likely not ideal); otherwise will require
    //   file '.nuxt.' suffix OR adding '// @vitest-environment nuxt' comment to Nuext test files.
    // environment: "nuxt",
    globals: true,
  },
});
