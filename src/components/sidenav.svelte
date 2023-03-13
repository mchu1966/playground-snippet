<script lang="ts">
	let open: boolean = true;
	function toggle() {
		open = !open;
	}

	let svgClass: string =
		'w-6 h-6 absolute rounded-full -right-3 hover:fill-blue-100 stroke-black fill-white cursor-pointer stroke-2 transition duration-300 ease-in-out';

	import Svg from 'components/svg.svelte';
</script>

<!-- https://www.youtube.com/watch?v=MszSqhEw__8 -->
<div class="flex">
	<div class=" bg-primary-900 h-screen p-5 pt-8 {open ? 'w-72' : 'w-20'} relative">
		{#if open}
			<Svg name="arrow-left-circle" class={svgClass} {open} {toggle} />
		{:else}
			<Svg name="arrow-right-circle" class={svgClass} {open} {toggle} />
		{/if}
	</div>
	<div class="p-7">
		<h1 class="text-2xl font-semibold">Title</h1>
	</div>
</div>
