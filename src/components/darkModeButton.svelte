<script lang="ts">
	import Svg from 'components/svg.svelte';
	let darkMode = false;
	function handleSwitchDarkMode(): void {
		darkMode = !darkMode;

		darkMode
			? document.documentElement.classList.add('dark')
			: document.documentElement.classList.remove('dark');
	}
</script>

<button class="duration-800 header-btn w-12 " on:click={handleSwitchDarkMode}
	>{#if darkMode}
		<Svg name="sun" class="m-3 h-5 w-5 fill-none stroke-white stroke-2" />
	{:else}
		<Svg name="moon" class="fill-solid m-3 h-5 w-5 stroke-current stroke-2" />
	{/if}</button
>
