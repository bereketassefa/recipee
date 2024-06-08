import path from "path";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

import tailwind from "tailwindcss";
import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";
import formkit from "unplugin-formkit/vite";

export default defineConfig({
  server: {
    port: 3001,
  },
  css: {
    postcss: {
      plugins: [tailwind(), autoprefixer()],
    },
  },
  plugins: [formkit({}), vue(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  optimizeDeps: {
    exclude: ["js-big-decimal"],
  },
});
