import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default defineConfig({
  base: './', // Dosya yollarını doğru ayarlamak için
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')
    }
  },
  plugins: [react()],
  css: {
    postcss: './postcss.config.js',
  },
  build: {
    sourcemap: false, // Disable sourcemaps
  }
});
