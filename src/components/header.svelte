<script lang="ts">
	import Svg from 'components/svg.svelte';

	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher<{ addNew: { name: string; code: string } }>();

	function handleAddNew() {
		dispatch('addNew', { name: 'Type a name', code: `` });
	}

	let darkMode = false;
	function handleSwitchDarkMode(): void {
		darkMode = !darkMode;

		darkMode
			? document.documentElement.classList.add('dark')
			: document.documentElement.classList.remove('dark');
	}
</script>

<div
	class="sticky top-0 flex flex-row content-center  bg-transparent p-2 backdrop-blur transition-colors dark:border-slate-50/[0.06]"
>
	<div class="mx-2 content-center p-4 font-bold"><h1>Playground Snippet</h1></div>
	<button
		class="duration-600 mx-2 w-20 rounded-3xl p-1 shadow-lg shadow-gray-500/50 transition ease-in-out hover:bg-slate-400"
		on:click={handleAddNew}>Add</button
	>

	<button
		class="duration-600 mx-2 w-20 rounded-3xl p-1 shadow-lg shadow-gray-500/50 transition ease-in-out hover:bg-slate-400"
		>Save</button
	>
	<div class="grow" />
	<button
		class="duration-600 mx-2 rounded-3xl p-1 shadow-lg shadow-gray-500/50 transition ease-in-out hover:bg-slate-400 dark:bg-gray-800 dark:hover:shadow-xl  dark:hover:shadow-white/70"
		on:click={handleSwitchDarkMode}
		>{#if darkMode}
			<Svg name="sun" class="m-3 h-5 w-5 fill-none stroke-white stroke-2" />
		{:else}
			<Svg name="moon" class="fill-solid m-3 h-5 w-5 stroke-current stroke-2" />
		{/if}</button
	>
</div>
