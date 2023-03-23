<script lang="ts">
	import Svg from 'components/svg.svelte';
	import LoginModal from './auth/loginModal.svelte';

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

	let loggedIn = false;
	let open = false;
	function closeModal(e: CustomEvent<boolean>) {
		open = e.detail;
		document.getElementsByTagName('BODY')[0].classList.remove('overflow-hidden');
		document.getElementsByTagName('BODY')[0].classList.add('overflow-auto');
	}

	async function handleLog() {
		loggedIn = !loggedIn;
	}
</script>

<div
	class="{open
		? 'hidden'
		: 'fixed'} top-0 flex w-full flex-row content-center bg-transparent p-2 backdrop-blur-xl transition-colors dark:border-slate-50/[0.06]"
>
	<div class="inline-block content-center p-4 ">
		<!-- before:absolute before:-inset-1 before:block before:-skew-y-3 before:bg-pink-500  -->
		<div class="text-2xl font-bold text-black ">
			<a href="/"> Playground Snippet </a>
		</div>
	</div>
	<div class="grow" />
	{#if !loggedIn}
		<button
			class="duration-800 header-btn"
			on:click={() => {
				open = true;
				document.getElementsByTagName('BODY')[0].classList.add('overflow-hidden');
			}}>Login</button
		>
	{:else}
		<button class="duration-800 header-btn" on:click={handleLog}>Logout</button>
	{/if}
	<button class="duration-800 header-btn" on:click={handleAddNew}>Add</button>
	<button class="duration-800 header-btn">Save</button>
	<button class="duration-800 header-btn w-12 " on:click={handleSwitchDarkMode}
		>{#if darkMode}
			<Svg name="sun" class="m-3 h-5 w-5 fill-none stroke-white stroke-2" />
		{:else}
			<Svg name="moon" class="fill-solid m-3 h-5 w-5 stroke-current stroke-2" />
		{/if}</button
	>
</div>
<div class="invisible h-[104px] bg-cyan-300 sm:h-[72px]" />

{#if open}
	<LoginModal on:closeModel={closeModal} />
{/if}
