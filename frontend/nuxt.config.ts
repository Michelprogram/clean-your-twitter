// https://nuxt.com/docs/api/configuration/nuxt-config
const nuxtConfig = defineNuxtConfig({
  modules: ["@pinia/nuxt", "@nuxt/devtools"],
  css: ["@/assets/scss/main.scss"],
  runtimeConfig: {
    public: {
      API_ENDPOINT: process.env.NUXT_API_ENDPOINT || "",
      API_REDIRECT_URL: process.env.NUXT_API_REDIRECT_URL || "",
      API_WEBSOCKET: process.env.NUXT_API_WEBSOCKET || "",
      TWITTER_CLIENT_ID: process.env.NUXT_TWITTER_CLIENT_ID || "",
      TWITTER_CODE_CHALLENGE: process.env.NUXT_TWITTER_CODE_CHALLENGE || "",
      version: process.env.NUXT_VERSION || "",
    },
  },
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
  devtools: {
    // Enable devtools (default: true)
    enabled: true,
    // VS Code Server options
    vscode: {},
    // ...other options
  },
});

export default nuxtConfig;
