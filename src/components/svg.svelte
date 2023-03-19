<!-- https://stackoverflow.com/questions/66069647/best-way-to-import-svg-icons-into-a-svelte-app -->
<script lang="ts">
	export let name: string;
	export let toggle: () => void = () => {};

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
		},
		{
			box: 24,
			name: 'arrow-path',
			svg: `<path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />`
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
