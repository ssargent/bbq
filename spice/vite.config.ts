import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
import tsconfigPaths from 'vite-tsconfig-paths';

export default defineConfig(() => {
  //const env = loadEnv(mode, process.cwd(), '');
  return {
    plugins: [tsconfigPaths(), react()],
    server: {
      port: 4100,
    },
  };
});
