import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default defineConfig({
  base: '/', // Dosya yollarını doğru ayarlamak için
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')
    }
  },
  plugins: [react()],
  server: {
    // In development the API runs locally; in production nginx serves both
    // from the same origin.
    proxy: {
      '/api': process.env.VITE_API_PROXY ?? 'http://127.0.0.1:9000',
    },
  },
  css: {
    postcss: './postcss.config.js',
  },
});
