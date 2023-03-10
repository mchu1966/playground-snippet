<script lang="ts">
	let count: number = 0;
	function handleClick(): void {
		count += 1;
	}

	let name: string = 'Jacky';

	$: if (count >= 10) {
		alert('count is high');
		count = 1;
	}

	let numbers: number[] = [1, 2, 3, 4];
	function addNumber() {
		numbers = [...numbers, numbers.push(numbers.length + 1)];
	}

	$: sum = numbers.reduce((t, n) => t + n, 0);

	let user = { loggedIn: false };

	function toggle() {
		user.loggedIn = !user.loggedIn;
	}

	let passcode: string;
	let pass: boolean = false;
	$: if (passcode == '123') {
		pass = true;
	}

	import Collapse from 'components/collapse.svelte';
</script>

<Collapse />

{#if pass}
	<h1>Welcome to SvelteKit</h1>
	<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
	<p class="text-teal-500">Tailwind is working</p>
	<button on:click={handleClick}>
		Clicked {count}
		{count <= 1 ? 'time' : 'times'}
	</button>

	<h1>Hello {name}!</h1>

	<p>{numbers.join(' + ')} = {sum}</p>
	<button on:click={addNumber}> add a number</button>

	<br />
	{#if user.loggedIn}
		<button on:click={toggle}> Log out </button>
	{/if}

	{#if !user.loggedIn}
		<button on:click={toggle}> Log in </button>
	{/if}

	<br />
	<a data-sveltekit-preload-data="tap" href="/home">go to /home</a>
{:else}
	<input
		type="text"
		placeholder="input pass code"
		class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
		bind:value={passcode}
	/>
{/if}
