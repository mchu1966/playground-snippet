<script lang="ts">
	export let name: string;
	export let title: string;
	export let placeholder: string;
</script>

<label for={name} class="block text-xs font-semibold uppercase text-gray-600">{title}</label>
<input
	id={name}
	type={name}
	{name}
	{placeholder}
	autocomplete={name}
	class="my-2 block w-full appearance-none bg-gray-200 p-3 text-gray-700 focus:bg-gray-300 focus:shadow-inner focus:outline-none"
	required
/>
