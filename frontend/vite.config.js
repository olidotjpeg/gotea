import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import sveltePreprocess from 'svelte-preprocess'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte({
    preprocess: sveltePreprocess()
  })],
  server: {
    watch: {
      usePolling: true,
    },
    host: true,
    strictPort: true,
    port: 5173,
  }
})
