<script lang="ts">
	let groupID = '';
</script>

<div class="hero min-h-screen">
	<div class="hero-content flex-col">
		<div class="card max-w-3xl w-full bg-neutral">
			<div class="card-body gap-4">
				<div class="card-title">Group ID</div>
				<input type="text" class="input" bind:value={groupID} />
				<div class="card-actions justify-center">
					<a href="/{groupID}" class="btn btn-primary">Join</a>
				</div>
			</div>
		</div>
		<div class="card max-w-sm w-full bg-neutral">
			<form class="card-body">
				<div class="form-control">
					<button class="btn btn-primary">Create Group</button>
				</div>
			</form>
		</div>
	</div>
</div>
