<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<svelte:head>
    <title>Yap</title>
</svelte:head>

{@render children()}
