<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	let dispatch = createEventDispatcher<{ closeModel: boolean }>();
	function closeModal() {
		dispatch('closeModel', false);
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
				<div class="pointer-events-auto w-11/12 bg-white p-12 sm:w-8/12 md:w-1/2 lg:w-5/12">
					<h1 class="text-xl font-semibold">
						Welcome back, <span class="font-normal">sign in to continue</span>
					</h1>
					<form class="mt-6">
						<label for="email" class="block text-xs font-semibold uppercase text-gray-600"
							>E-mail</label
						>
						<input
							id="email"
							type="email"
							name="email"
							placeholder="john.doe@company.com"
							autocomplete="email"
							class="mt-2 block w-full appearance-none bg-gray-200 p-3 text-gray-700 focus:bg-gray-300 focus:shadow-inner focus:outline-none"
							required
						/>
						<label for="password" class="mt-2 block text-xs font-semibold uppercase text-gray-600"
							>Password</label
						>
						<input
							id="password"
							type="password"
							name="password"
							placeholder="********"
							autocomplete="current-password"
							class="mt-2 block w-full appearance-none bg-gray-200 p-3 text-gray-700 focus:bg-gray-300 focus:shadow-inner focus:outline-none"
							required
						/>
						<div class="grid gap-4">
							<button
								type="submit"
								class="mt-6 w-full bg-black py-3 font-medium uppercase tracking-widest text-white shadow-lg hover:bg-gray-900 hover:shadow-none focus:outline-none "
							>
								Sign in
							</button>
							<div class="h-px w-full bg-black" />
							<!-- oauth buttons here -->
							<button
								class=" w-full bg-black py-3 font-medium uppercase tracking-widest text-white shadow-lg hover:bg-gray-900 hover:shadow-none focus:outline-none"
							>
								Sign up
							</button>
						</div>
						<p
							class="mt-4 inline-block  cursor-pointer justify-between text-xs text-gray-500 hover:text-black"
						>
							Forgot password?
						</p>
					</form>
				</div>
			</div>
		</div>
	</div>
</div>
