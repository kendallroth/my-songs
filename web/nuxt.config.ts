import vuetify from "vite-plugin-vuetify";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      title: "My Songs",
      link: [{ rel: "icon", type: "image/png", href: "/logo.svg" }],
    },
  },
  build: {
    transpile: ["vuetify"],
  },
  components: [
    {
      path: "~/components",
      pathPrefix: false,
    },
  ],
  // NOTE: Do not include app scss file or it will always be overridden by Vuetify
  css: ["vuetify/lib/styles/main.sass"],
  hooks: {
    // Remove unwanted TS autocompletion paths/aliases
    // Source: https://github.com/nuxt/nuxt/issues/14816#issuecomment-1397366520
    "prepare:types": ({ tsConfig }) => {
      const aliasesToRemove = ["~~", "~~/*", "@@", "@@/*"];
      aliasesToRemove.forEach((alias) => {
        if (tsConfig.compilerOptions?.paths[alias]) {
          delete tsConfig.compilerOptions?.paths[alias];
        }
      });

      tsConfig.compilerOptions?.types.push("vitest/globals");
    },
    // Source: https://codybontecou.com/how-to-use-vuetify-with-nuxt-3.html
    "vite:extendConfig": (config) => {
      config.plugins?.push(vuetify());
    },
  },
  modules: ["nuxt-vitest", "@pinia/nuxt"],
  runtimeConfig: {
    app: {
      // Overridden by 'NUXT_APP_API_URL'
      apiUrl: "http://localhost:8090",
    },
  },
  // TODO: Enable this sometime when there is time/patience for figuring out how to properly
  //         check for authentication state and return properly rendered state to client...
  ssr: false,
  typescript: {
    shim: false,
  },
});
