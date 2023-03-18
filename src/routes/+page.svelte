<script lang="ts">
	import CodeArea from 'components/codeArea.svelte';
	// fetch database to get the snippets (array of string of size N)
	let mockSnippets: string[] = [
		`// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}`,

		`package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}
`
	];

	function handleAddNew() {
		mockSnippets = [...mockSnippets, ``];
	}

	async function testApi(): Promise<void> {
		const res = await fetch('/api/test', {
			method: 'GET'
		})
			.then((res) => res.json())
			.then((data) => {
				console.log(data);
			});
	}
</script>

<div class="px-2">
	<button class="p-2 border shadow-md" on:click={handleAddNew}>Add new</button>
	<!-- https://eternaldev.com/blog/5-ways-to-perform-for-loop-in-svelte-each-block/ -->
	{#each mockSnippets as snippet, index}
		<CodeArea code={snippet} />
	{/each}

	<button on:click={testApi}>Test</button>
</div>
