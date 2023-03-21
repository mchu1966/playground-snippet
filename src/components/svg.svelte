<!-- https://stackoverflow.com/questions/66069647/best-way-to-import-svg-icons-into-a-svelte-app -->
<script lang="ts">
	export let name: string;
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function svgClick() {
		dispatch('svgClick');
	}

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
		},
		{
			box: 24,
			name: 'sun',
			svg: `<path stroke-linecap="round" stroke-linejoin="round" d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" />`
		},
		{
			box: 24,
			name: 'moon',
			svg: `<path stroke-linecap="round" stroke-linejoin="round" d="M21.752 15.002A9.718 9.718 0 0118 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 003 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 009.002-5.998z" />`
		}
	];
	let displayIcon: icon = icons.find((e) => e.name === name) as icon;
</script>

<svg
	xmlns="http://www.w3.org/2000/svg"
	class={$$props.class}
	viewBox="0 0 {displayIcon.box} {displayIcon.box}"
	on:click={svgClick}
	on:keypress={svgClick}
>
	{@html displayIcon.svg}
</svg>
