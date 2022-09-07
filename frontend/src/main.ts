// @ts-ignore
import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		name: 'gnar world'
	}
});

export default app;