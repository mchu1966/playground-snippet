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
	class="sticky top-0 flex flex-row content-center bg-white p-2 backdrop-blur transition-colors dark:border-slate-50/[0.06]"
>
	<button
		class="duration-600 rounded-full border bg-cyan-500 p-2 shadow-md shadow-cyan-500/50 transition ease-in-out"
		on:click={handleAddNew}>Add new</button
	>
	<button
		class="duration-600 rounded-full border bg-cyan-500 p-2 shadow-md shadow-cyan-500/50 transition ease-in-out"
		>Save</button
	>
	<div class="grow" />
	<button
		class="duration-600 border-collapse rounded-full bg-cyan-500 p-2 shadow-cyan-500/50 transition ease-in-out "
		on:click={handleSwitchDarkMode}
		>{#if darkMode}
			<Svg name="sun" class="h-5 w-5 fill-none stroke-white stroke-2" />
		{:else}
			<Svg name="moon" class="fill-solid h-5 w-5 stroke-current stroke-2" />
		{/if}</button
	>
</div>
