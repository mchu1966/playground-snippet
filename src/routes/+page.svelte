<script lang="ts">
	import SnippetArea from 'components/snippetArea.svelte';
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
}
`
		}
	];

	function handleAddNew(s: CustomEvent<snippet>) {
		mockSnippets = [...mockSnippets, s.detail];
	}
</script>

<Header on:addNew={handleAddNew} />
<div class="container h-full self-center dark:bg-black">
	<!-- https://eternaldev.com/blog/5-ways-to-perform-for-loop-in-svelte-each-block/ -->
	{#each mockSnippets as snippet, index}
		<div class="scroll-mt-28 p-4 sm:scroll-mt-20" id={snippet.name + index}>
			<div class="flex w-full  flex-row">
				<a href="#{snippet.name + index}">
					<input
						bind:value={snippet.name}
						placeholder={snippet.name}
						class="w-full dark:bg-gray-900"
					/>
				</a>
			</div>
			<SnippetArea code={snippet.code} />
		</div>
	{/each}
</div>
