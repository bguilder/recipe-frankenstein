<template>
	<v-app>
		<v-app-bar
			app
			color="primary"
			dark
		>
			<v-row
				align="start"
				justify="start"
			>
				<v-col cols="12">
					<h2 class="font-weight-bold">Recipe Frankenstein</h2>
				</v-col>
			</v-row>
		</v-app-bar>

		<v-content v-if="searchInput">
			<HelloWorld v-bind:searchInput="searchInput" />
		</v-content>
		<v-container
			fluid
			fill-height
		>
			<v-row no-gutters>
				<v-col>
					<v-row
						align="center"
						justify="center"
					>
						<v-col
							cols="12"
							class="pa-2"
						>
							<v-img
								contain
								aspect-ratio="3"
								src="./assets/frankenstein.png"
							></v-img>
						</v-col>
					</v-row>
					<v-row
						align="baseline"
						justify="center"
					>
						<v-col
							class="pa-2"
							cols="8"
							lg="5"
							md="5"
							sm="5"
							xs="2"
						>
							<SearchBar />
						</v-col>
					</v-row>
				</v-col>
			</v-row>
		</v-container>
	</v-app>
</template>

<script lang="ts">
import Vue from "vue";
import HelloWorld from "./components/HelloWorld.vue";
import SearchBar from "./components/SearchBar.vue";

export default Vue.extend({
	name: "App",

	components: {
		HelloWorld,
		SearchBar
	},

	data() {
		return {
			searchInput: null
		};
	}
});
</script>
