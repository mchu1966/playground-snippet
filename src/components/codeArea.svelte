<script lang="ts">
	import Svg from 'components/svg.svelte';
	export let code: string;

	let result: string = '';
	let height: number;
	let resultHeight: number;
	$: height = code.split(/\r\n|\r|\n/).length;

	let loading: boolean = false;
	async function handleRun(): Promise<void> {
		let isDevEnv: boolean = import.meta.env.DEV;
		if (isDevEnv) {
			await handleFormat().then((valid) => {
				console.log(valid);
			});
			let formData: FormData = new FormData();
			formData.append('version', '2');
			formData.append('body', code);
			formData.append('withVet', 'true');

			// call our own backend endpoint later.
			loading = true;
			await fetch('https://go.dev/_/compile?backend=', {
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
			loading = false;
		} else {
			await handleFormat().then((valid) => {
				console.log(valid);
			});
			loading = true;
			await fetch('/api/run', {
				method: 'POST',
				headers: {
					Accept: 'application/json',
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					code: code
				})
			})
				.then((res) => res.json())
				.then((data) => {
					result = data.Events[0].Message;
					resultHeight = result.split(/\r\n|\r|\n/).length;
				});
			loading = false;
		}
	}

	async function handleFormat(): Promise<boolean> {
		let isDevEnv: boolean = import.meta.env.DEV;
		if (isDevEnv) {
			let valid: boolean = false;
			loading = true;
			let formData: FormData = new FormData();
			formData.append('body', code);
			formData.append('imports', 'true');

			let form: string =
				encodeURIComponent('body') +
				'=' +
				encodeURIComponent(code) +
				'&' +
				encodeURIComponent('imports') +
				'=' +
				encodeURIComponent('true');

			// call our own backend endpoint later.
			await fetch('https://go.dev/_/fmt?backend=', {
				method: 'POST',
				body: form,
				headers: {
					'content-type': 'application/x-www-form-urlencoded'
				}
			})
				.then((res) => res.json())
				.then((data) => {
					valid = data.Error == '';
					if (valid) {
						code = data.Body;
						height = code.split(/\r\n|\r|\n/).length;
					} else {
						result = data.Error;
						resultHeight = result.split(/\r\n|\r|\n/).length;
					}
				});
			loading = false;
			return valid;
		} else {
			let valid: boolean = false;
			loading = true;
			await fetch('/api/format', {
				method: 'POST',
				headers: {
					Accept: 'application/json',
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					code: code
				})
			})
				.then((res) => res.json())
				.then((data) => {
					valid = data.Error == '';
					if (valid) {
						code = data.Body;
						height = code.split(/\r\n|\r|\n/).length;
					} else {
						result = data.Error;
						resultHeight = result.split(/\r\n|\r|\n/).length;
					}
				});
			loading = false;
			return valid;
		}
	}
</script>

<div class="flex flex-wrap">
	<div class="block p-2.5 w-full md:w-1/2">
		<label for="large-input" class="block text-sm font-medium text-gray-900">Code</label>
		<div class="flex flex-row">
			<div class="w-[5%] py-1 pr-1 text-sm">
				{#each Array(height) as _, index}
					<div class="text-right">{index + 1}</div>
				{/each}
			</div>

			<textarea
				id="message"
				rows={height}
				wrap="off"
				class="resize-none leading-5 h-grow w-full py-1 px-2 text-sm text-gray-900 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 "
				placeholder="Type your code here..."
				bind:value={code}
			/>
		</div>
		<div class="p-2">
			{#if !loading}
				<button
					disabled={loading}
					type="submit"
					class="shadow-black rounded-lg border text-sm p-1 {loading ? 'cursor-not-allowed' : ''}"
					on:click={handleRun}
				>
					Run
				</button>

				<button
					disabled={loading}
					type="submit"
					class="shadow-black rounded-lg border text-sm p-1 {loading ? 'cursor-not-allowed' : ''}"
					on:click={handleFormat}
				>
					Format
				</button>
			{:else}
				<div class="p-3">
					<Svg name="arrow-path" class="animate-spin w-4 h-4 stroke-1 stroke-current fill-none" />
				</div>
			{/if}
		</div>
	</div>
	<div class="block p-2.5 w-full md:w-1/2">
		<label for="large-input" class="block text-sm font-medium text-gray-900">Result</label>
		<div class="flex flex-row">
			<div class="w-[3%] py-1 text-sm">
				{#each Array(resultHeight) as _, index}
					<div>{index + 1}</div>
				{/each}
			</div>

			<textarea
				disabled
				id="output"
				rows={resultHeight}
				wrap="off"
				class="resize-none leading-5 h-grow w-full py-1 px-2 text-sm text-gray-900 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 "
				>{result}</textarea
			>
		</div>
	</div>
</div>
