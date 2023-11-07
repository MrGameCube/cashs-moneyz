<script lang="ts">
	let groupID = '';
</script>

<div class="hero min-h-screen">
	<div class="hero-content flex-col">
		<div class="card max-w-3xl w-full bg-neutral">
			<div class="card-body gap-4">
				<div class="card-title">Group ID</div>
				<div class="flex gap-4">
					<input type="text" class="input" bind:value={groupID} />
					<a href="/{groupID}" class="btn btn-circle bg-primary">
						<svg
							class="h-6 w-6 fill-current md:h-8 md:w-8"
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							viewBox="0 0 24 24"
						>
							<path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" />
						</svg>
					</a>
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
