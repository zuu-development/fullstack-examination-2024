export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  modules: ["@nuxt/eslint"],
  vite: {
    server: {
      proxy: {
        "/api/": {
          target: process.env.PROXY_API_URL || "http://localhost:8080",
          secure: false
        }
      }
    }
  },
});
