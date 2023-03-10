<script lang="ts">
	let on: boolean = false;
	async function toggleOnOff() {
		await new Promise((r) => setTimeout(r, 400));
		on = !on;
	}
</script>

<div class="flex h-screen">
	<div class="m-auto">
		<div class="flex justify-center space-x-2">
			{#if on}
				<button
					type="button"
					data-te-ripple-init
					data-te-ripple-color="light"
					class="inline-block rounded bg-primary px-6 pt-2.5 pb-2 text-xs font-medium uppercase leading-normal text-white shadow-[0_4px_9px_-4px_#3b71ca] transition duration-150 ease-in-out hover:bg-primary-600 hover:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] focus:bg-primary-600 focus:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] focus:outline-none focus:ring-0 active:bg-primary-700 active:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)]"
					on:click={toggleOnOff}
				>
					TURN OFF
				</button>
			{:else}
				<button
					type="button"
					data-te-ripple-init
					data-te-ripple-color="light"
					class="inline-block rounded bg-primary px-6 pt-2.5 pb-2 text-xs font-medium uppercase leading-normal text-white shadow-[0_4px_9px_-4px_#3b71ca] transition duration-150 ease-in-out hover:bg-primary-600 hover:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] focus:bg-primary-600 focus:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] focus:outline-none focus:ring-0 active:bg-primary-700 active:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)]"
					on:click={toggleOnOff}
				>
					TURN ON
				</button>
			{/if}
		</div>
	</div>
</div>
