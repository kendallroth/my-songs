import "vuetify/styles";
import { createVuetify, type ThemeDefinition } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi-svg";

const lightTheme: ThemeDefinition = {
  dark: false,
  colors: {
    background: "#f9f9f9",
  },
};

export default defineNuxtPlugin((nuxtApp) => {
  const vuetify = createVuetify({
    icons: {
      defaultSet: "mdi",
      aliases,
      sets: {
        mdi,
      },
    },
    ssr: true,
    theme: {
      defaultTheme: "light",
      themes: {
        light: lightTheme,
      },
    },
  });

  nuxtApp.vueApp.use(vuetify);
});
