import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          'vendor': ['react', 'react-dom'],
          'mantine': ['@mantine/core', '@mantine/hooks', '@mantine/notifications'],
          'icons': ['react-icons']
        }
      }
    }
  },
  server: {
    port: 5173,
    host: true,
    hmr: {
      protocol: 'ws'
    }
  },
  optimizeDeps: {
    include: ['@mantine/core', '@mantine/hooks', '@mantine/notifications', 'react-icons']
  }
})
