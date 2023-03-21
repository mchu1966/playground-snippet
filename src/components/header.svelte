<script lang="ts">
	import Svg from 'components/svg.svelte';

	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function handleAddNew() {
		dispatch('addNew', { name: 'Type a name', code: `` });
	}

	let darkMode = false;
	function handleSwitchDarkMode() {
		darkMode = !darkMode;

		darkMode
			? document.documentElement.classList.add('dark')
			: document.documentElement.classList.remove('dark');
	}
</script>

<div class="p-2 content-center flex flex-row">
	<button
		class="transition duration-600 ease-in-out rounded-full p-2 border shadow-md"
		on:click={handleAddNew}>Add new</button
	>
	<button class="transition duration-600 ease-in-out rounded-full p-2 border shadow-md">Save</button
	>
	<div class="grow" />
	<button
		class="transition duration-600 ease-in-out rounded-full border-collapse p-2 "
		on:click={handleSwitchDarkMode}
		>{#if darkMode}
			<Svg name="sun" class="w-5 h-5 stroke-2 stroke-white fill-none" />
		{:else}
			<Svg name="moon" class="w-5 h-5 stroke-2 stroke-current fill-solid" />
		{/if}</button
	>
</div>
