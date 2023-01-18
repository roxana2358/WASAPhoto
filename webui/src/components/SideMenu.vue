<script>
import { RouterLink } from 'vue-router'
export default {}
</script>

<template>
    <div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#eye"/></svg>
								Search user
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/profile" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#smile"/></svg>
								Personal profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/uploadPhoto" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#plus-circle"/></svg>
								Upload photo
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>
		</div>
	</div>
</template>

<style>
</style>