import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		allowedHosts: ['web.nextreleaseplease.com', "reverse-proxy", "reverse-proxy.localhost"],
	},
	plugins: [sveltekit()],
});
