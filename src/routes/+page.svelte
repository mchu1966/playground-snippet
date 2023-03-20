<script lang="ts">
	import CodeArea from 'components/codeArea.svelte';
	import Svg from 'components/svg.svelte';

	// fetch database to get the snippets (array of string of size N)
	type snippet = {
		name: string;
		code: string;
	};
	let mockSnippets: snippet[] = [
		{
			name: 'Hello World',
			code: `// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}`
		}
	];

	function handleAddNew() {
		mockSnippets = [...mockSnippets, { name: 'Type a name', code: `` }];
	}

	let darkMode = false;
	function handleSwitchDarkMode() {
		darkMode = !darkMode;

		darkMode
			? document.documentElement.classList.add('dark')
			: document.documentElement.classList.remove('dark');
	}
</script>

<div class="px-2 dark:bg-gray-400 h-screen">
	<div class="p-2 content-center flex flex-row">
		<button
			class="transition duration-600 ease-in-out rounded-full p-2 border shadow-md"
			on:click={handleAddNew}>Add new</button
		>
		<button class="transition duration-600 ease-in-out rounded-full p-2 border shadow-md"
			>Save</button
		>
		<div class="grow" />
		<button
			class="transition duration-600 ease-in-out rounded-full border-collapse p-2 "
			on:click={handleSwitchDarkMode}
			>{#if darkMode}
				<Svg name="sun" class="w-5 h-5 stroke-2 stroke-white fill-none" />
			{:else}
				<Svg name="moon" class="w-5 h-5 stroke-2 stroke-current fill-solid" />
			{/if}</button
		>
	</div>
	<!-- https://eternaldev.com/blog/5-ways-to-perform-for-loop-in-svelte-each-block/ -->
	{#each mockSnippets as snippet, index}
		<div class="p-4">
			<div class="w-full flex flex-row">
				<input bind:value={snippet.name} placeholder={snippet.name} class="w-full" />
			</div>
			<CodeArea code={snippet.code} />
		</div>
	{/each}
</div>
