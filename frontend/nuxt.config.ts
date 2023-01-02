// https://nuxt.com/docs/api/configuration/nuxt-config
const nuxtConfig = defineNuxtConfig({
  modules: ["@pinia/nuxt"],
  css: ["@/assets/scss/main.scss"],
  vite: {
    css: {
      preprocessorOptions: {
        sass: {
          additionalData: '@import "@/assets/scss/main.scss"',
        },
      },
    },
  },
  app: {
    head: {
      title: "Clean your twitter",
      link: [{ rel: "icon", type: "image/x-icon", href: "/icon.ico" }],
    },
  },
});

export default nuxtConfig;
