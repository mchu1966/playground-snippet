<script lang="ts">
	import SnippetArea from 'components/snippetArea.svelte';
	import Header from 'components/header.svelte';
	// fetch database to get the snippets (array of string of size N)

	import type { snippet } from '$lib/types/snippet';
	function handleAddNew(s: CustomEvent<snippet>) {
		data.mockSnippets = [...data.mockSnippets, s.detail];
	}

	import type { PageData } from './$types';
	export let data: PageData;
</script>

<Header on:addNew={handleAddNew} {data} />
<div class="container h-full self-center dark:bg-black">
	<!-- https://eternaldev.com/blog/5-ways-to-perform-for-loop-in-svelte-each-block/ -->
	{#each data.mockSnippets as snippet, index}
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
