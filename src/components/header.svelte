<script lang="ts">
	import LoginModal from './auth/loginModal.svelte';
	import SignupModal from './auth/signupModal.svelte';
	import DarkModeButton from './darkModeButton.svelte';

	import { createEventDispatcher } from 'svelte';
	import type { snippet } from '$lib/types/snippet';

	const dispatch = createEventDispatcher<{ addNew: snippet }>();
	function handleAddNew() {
		dispatch('addNew', { name: 'Type a name', code: `` });
	}

	let openModal = false;
	let signupModalOn = false;
	function closeModal(e: CustomEvent<boolean>) {
		openModal = e.detail;
		document.getElementsByTagName('BODY')[0].classList.remove('overflow-hidden'); // hide the scroll bar
		document.getElementsByTagName('BODY')[0].classList.add('overflow-auto'); // show the scroll bar
		signupModalOn = false;
	}
	function switchSignup(e: CustomEvent<boolean>) {
		signupModalOn = e.detail;
	}

	import { page } from '$app/stores';
	import type { PageData } from '../routes/$types';
	export let data: PageData;
	$: supabase = data.supabase;
</script>

<div
	class="{openModal
		? 'hidden'
		: 'fixed'} top-0 flex w-full flex-row content-center bg-transparent px-2 backdrop-blur-xl transition-colors dark:border-slate-50/[0.06]"
>
	<div class="inline-block content-center p-4 ">
		<!-- before:absolute before:-inset-1 before:block before:-skew-y-3 before:bg-pink-500  -->
		<div class="text-2xl font-bold text-black ">
			<a href="/"> Playground Snippet </a>
		</div>
	</div>
	<div class="grow" />
	{#if !$page.data.session}
		<button
			class="duration-800 header-btn"
			on:click={() => {
				openModal = true;
				document.getElementsByTagName('BODY')[0].classList.add('overflow-hidden');
			}}>Login</button
		>
	{:else}
		<button
			class="duration-800 header-btn"
			on:click={async () => {
				const { error } = await supabase.auth.signOut();
				if (error) {
					console.log(error);
				}
			}}>Logout</button
		>
		<button class="duration-800 header-btn" on:click={handleAddNew}>Add</button>
		<button class="duration-800 header-btn">Save</button>
	{/if}
	<DarkModeButton />
</div>
<div class="invisible h-[104px] bg-cyan-300 sm:h-[72px]" />

{#if openModal}
	{#if !signupModalOn}
		<LoginModal on:closeModel={closeModal} on:switchSignup={switchSignup} d={data} />
	{:else}
		<SignupModal on:closeModel={closeModal} on:switchSignup={switchSignup} d={data} />
	{/if}
{/if}
