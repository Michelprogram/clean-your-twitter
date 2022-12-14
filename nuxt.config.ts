// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  buildDir: "dist",
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
