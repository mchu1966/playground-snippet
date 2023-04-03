<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	let dispatchCloseModel = createEventDispatcher<{ closeModel: boolean }>();
	function closeModal() {
		dispatchCloseModel('closeModel', false);
	}

	let dispatchSwitchSignup = createEventDispatcher<{ switchSignup: boolean }>();
	function switchSignup() {
		dispatchSwitchSignup('switchSignup', false);
	}

	import type { PageData } from '../../routes/$types';
	import ModalInput from 'components/auth/modalInput.svelte';
	export let d: PageData;
	$: supabase = d.supabase;

	async function signUp(e: SubmitEvent) {
		const formData = new FormData(e.target as HTMLFormElement);
		const { data, error } = await supabase.auth.signUp({
			email: formData.get('email') as string,
			password: formData.get('password') as string
		});
		console.log(data, error);
		dispatchCloseModel('closeModel', false);
	}
</script>

<div>
	<div
		class="fixed top-0 left-0 right-0 bottom-0 min-h-screen w-full backdrop-blur-md"
		on:click={closeModal}
		on:keydown={closeModal}
	/>
	<div class="z-50">
		<div class="pointer-events-none fixed top-0 left-0 right-0 bottom-0 h-full w-full">
			<div class="grid min-h-screen place-items-center">
				<div
					class="pointer-events-auto w-11/12 bg-white p-12 dark:bg-black sm:w-8/12 md:w-1/2 lg:w-5/12 "
				>
					<h1 class="text-xl font-semibold dark:text-white">Sign up and enjoy coding!</h1>
					<form class="mt-6" on:submit|preventDefault={signUp}>
						<ModalInput name="email" title="E-mail" placeholder="hello@playground.com" />
						<ModalInput name="password" title="Password" placeholder="********" />
						<button
							type="submit"
							class="mt-6 mb-4 w-full bg-black py-3 font-medium uppercase tracking-widest text-white shadow-lg hover:bg-gray-900 hover:shadow-none hover:shadow-white/70 focus:outline-none dark:border dark:shadow-inner"
						>
							Sign up
						</button>

						<div class="mb-4 flex flex-row">
							<div class="h-px w-full self-center bg-black dark:bg-white" />
							<div class="px-2">or</div>
							<div class="h-px w-full self-center bg-black dark:bg-white" />
						</div>

						<button
							on:click={switchSignup}
							class="w-full bg-black py-3 font-medium uppercase tracking-widest text-white shadow-lg hover:bg-gray-900 hover:shadow-none hover:shadow-white/70 focus:outline-none dark:border dark:shadow-inner"
						>
							Already have account?<br />Login here
						</button>
					</form>
				</div>
			</div>
		</div>
	</div>
</div>
