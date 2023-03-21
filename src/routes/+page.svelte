<script lang="ts">
	import CodeArea from 'components/codeArea.svelte';
	import Header from 'components/header.svelte';
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

	function handleAddNew(s: any) {
		mockSnippets = [...mockSnippets, s.detail];
	}
</script>

<div class="container dark:bg-gray-400 h-screen self-center">
	<Header on:addNew={handleAddNew} />
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
