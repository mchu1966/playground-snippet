<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let code: string;

	let result: string = '';
	let height: number;
	let resultHeight: number;
	$: height = code.split(/\r\n|\r|\n/).length;

	const dispatch = createEventDispatcher();
	const submit = () => dispatch('submit');

	async function handleRun(): Promise<void> {
		let formData: FormData = new FormData();
		formData.append('version', '2');
		formData.append('body', code);
		formData.append('withVet', 'true');

		// call our own backend endpoint later.
		const res = await fetch('https://go.dev/_/compile?backend=', {
			method: 'POST',
			body: formData,
			mode: 'cors',
			headers: {
				host: 'go.dev',
				accept: '*/*'
			}
		})
			.then((res) => res.json())
			.then((data) => {
				result = data.Events[0].Message;
				resultHeight = result.split(/\r\n|\r|\n/).length;
			});
	}

	async function handleFormat(): Promise<void> {
		let formData: FormData = new FormData();
		formData.append('version', '2');
		formData.append('body', code);
		formData.append('withVet', 'true');

		// call our own backend endpoint later.
		const res = await fetch('https://go.dev/_/compile?backend=', {
			method: 'POST',
			body: formData,
			mode: 'cors',
			headers: {
				host: 'go.dev',
				accept: '*/*'
			}
		})
			.then((res) => res.json())
			.then((data) => {
				result = data.Events[0].Message;
				resultHeight = result.split(/\r\n|\r|\n/).length;
			});
	}
</script>

<div class="flex flex-row w-full ">
	<div class="block p-2.5 w-1/2 ">
		<label for="large-input" class="block text-sm font-medium text-gray-900 dark:text-white"
			>Code</label
		>
		<div class="flex flex-row">
			<div class="w-[3%] py-1 text-sm">
				<div>1</div>
				<div>2</div>
				<div>3</div>
				<div>4</div>
				<div>5</div>
				<div>6</div>
			</div>

			<textarea
				id="message"
				rows={height}
				wrap="off"
				class="resize-none leading-4 h-grow w-full p-2 text-sm text-gray-900 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 "
				placeholder="Type your code here..."
				bind:value={code}
			/>
		</div>
		<button type="submit" class="shadow-black rounded-lg border text-sm p-1" on:click={handleRun}
			>Run</button
		>
		<button type="submit" class="shadow-black rounded-lg border text-sm p-1" on:click={handleFormat}
			>Format</button
		>
	</div>
	<div class="block p-2.5 w-1/2">
		<label for="large-input" class="block text-sm font-medium text-gray-900 dark:text-white"
			>Result</label
		>
		<div class="flex flex-row">
			<div class="w-[3%] py-1 text-sm">
				<div>1</div>
				<div>2</div>
				<div>3</div>
				<div>4</div>
				<div>5</div>
				<div>6</div>
			</div>

			<textarea
				disabled
				id="output"
				rows={resultHeight}
				wrap="off"
				class="resize-none leading-4 h-grow w-full p-2 text-sm text-gray-900 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 "
				>{result}</textarea
			>
		</div>
	</div>
</div>
