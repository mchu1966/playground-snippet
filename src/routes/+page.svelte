<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	let code: string = '';
	let result: string = 'testestestset';
	let height: number;
	$: height = code.split(/\r\n|\r|\n/).length;

	const dispatch = createEventDispatcher();
	const submit = () => dispatch('submit');

	async function handleSubmit() {
		let formData = new FormData();
		formData.append('version', '2');
		formData.append('body', code);
		formData.append('withVet', 'true');

		const res = await fetch('https://go.dev/_/compile', {
			method: 'POST',
			body: formData
		});

		alert(res);
	}
</script>

<div>
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
			<button
				type="submit"
				class="shadow-black rounded-lg border text-sm p-1"
				on:click={handleSubmit}>Run</button
			>
			<button
				type="submit"
				class="shadow-black rounded-lg border text-sm p-1"
				on:click={handleSubmit}>Format</button
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
					id="result"
					rows="4"
					class="w-[97%] p-2 text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					>{result}</textarea
				>
			</div>
		</div>
	</div>
</div>
