<script>

	let data = [
		{id: 1, first: 'Jill', last: 'Smith', point: 50},
		{id: 2, first: 'Eve', last: 'Jackson', point: 94},
		{id: 3, first: 'Adam', last: 'Johnson', point: 67},
	]
</script>

<div class="flex flex-col w-full">
	<h2>Deployments</h2>

	<table class="table-auto w-full border-collapse border border-app-dark dark:border-app-light">
		<thead>
		<tr>
			<th>First Name</th>
			<th>Last Name</th>
			<th>Points</th>
		</tr>
		</thead>
		<tbody>
		{#each data as d (d.id)}
			<tr class="even:bg-app-light/50 odd:bg-app-lightest dark:even:bg-app-dark/50 dark:odd:bg-app-darkest">
				<td>{d.first}</td>
				<td>{d.last}</td>
				<td>{d.point}</td>
			</tr>
		{/each}
		</tbody>
	</table>
</div>