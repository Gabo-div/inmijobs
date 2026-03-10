import { URL, fileURLToPath } from 'node:url'
import { defineConfig } from 'vite'
import { devtools } from '@tanstack/devtools-vite'
import viteReact from '@vitejs/plugin-react'

import tailwindcss from '@tailwindcss/vite'

import { tanstackRouter } from '@tanstack/router-plugin/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    devtools(),
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true,
    }),
    viteReact(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    port: 3001,
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // Apunta a tu backend en Go
        changeOrigin: true,
        secure: false,
        // Esta función inyecta el token para que Go te deje pasar
        configure: (proxy, _options) => {
          proxy.on('proxyReq', (proxyReq, req, _res) => {
            const cookies = req.headers.cookie;
            if (cookies) {
              const match = cookies.match(/(?:^|; )better-auth\.session_token=([^;]*)/);
              if (match && match[1]) {
                proxyReq.setHeader('Authorization', `Bearer ${match[1]}`);
              }
            }
          });
        },
      },
    },
  }
})
