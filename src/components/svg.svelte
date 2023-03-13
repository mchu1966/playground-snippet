<!-- https://stackoverflow.com/questions/66069647/best-way-to-import-svg-icons-into-a-svelte-app -->
<script lang="ts">
	export let name: string;
	export let toggle: () => void;

	type icon = {
		box: number;
		name: string;
		svg: string;
	};

	let icons: icon[] = [
		{
			box: 24,
			name: 'arrow-left-circle',
			svg: `<path stroke-linecap="round" stroke-linejoin="round" d="M11.25 9l-3 3m0 0l3 3m-3-3h7.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>`
		},
		{
			box: 24,
			name: 'arrow-right-circle',
			svg: `<path stroke-linecap="round" stroke-linejoin="round" d="M12.75 15l3-3m0 0l-3-3m3 3h-7.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>`
		}
	];
	let displayIcon: icon = icons.find((e) => e.name === name)!;
</script>

<svg
	xmlns="http://www.w3.org/2000/svg"
	class={$$props.class}
	viewBox="0 0 {displayIcon.box} {displayIcon.box}"
	on:click={toggle}
	on:keypress={toggle}
>
	{@html displayIcon.svg}
</svg>
