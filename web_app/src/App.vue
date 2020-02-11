<template>
	<v-app>
		<v-app-bar
			app
			color="#00C279"
			dark
		>
			<v-row
				align="start"
				justify="start"
			>
				<v-col cols="12">
					<v-btn
						x-large
						color="white"
						text
						@click="reset"
						class="pa-2"
					> Recipe Frankenstein</v-btn>
				</v-col>
			</v-row>
		</v-app-bar>

		<template v-if="searching">
			<v-content>
				<HelloWorld v-bind:search="search" />
			</v-content>
		</template>
		<template v-else>
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
									max-height="125"
									src="./assets/frankenstein.png"
								></v-img>
							</v-col>
						</v-row>
						<v-row
							align="center"
							justify="center"
						>
							<v-col
								class="pa-2 text-xs-center"
								cols="8"
								lg="5"
								md="5"
								sm="5"
								xs="2"
							>
								<v-row>
									<v-col cols="12">
										<v-text-field
											label="Search Recipes..."
											outlined
											solo
											color="#00C279"
											rounded
											v-model="search.searchInput"
										></v-text-field>
									</v-col>
									<!-- <v-col cols="3">
										<v-text-field
											label="Count"
											outlined
											rounded
											type="number"
											v-model="search.recipeCount"
										></v-text-field>
									</v-col> -->
								</v-row>
								<v-row
									align="center"
									justify="center"
								>
									<v-btn
										outlined
										color="#00C279"
										elevation-20
										large
										@click="submit"
									>Submit</v-btn>
								</v-row>

							</v-col>
						</v-row>
					</v-col>
				</v-row>
			</v-container>
		</template>
		<footer class="caption font-weight-thin text-center">Icon made by Freepik from www.flaticon.com
			<br>v0.1.0-alpha<br>
		</footer>
	</v-app>
</template>

<script lang="ts">
import Vue from "vue";
import HelloWorld from "./components/HelloWorld.vue";
import SearchBar from "./components/SearchBar.vue";

export default Vue.extend({
	name: "App",

	components: {
		HelloWorld
	},

	data() {
		return {
			searching: false,
			search: {
				searchInput: null,
				recipeCount: 5
			}
		};
	},
	methods: {
		submit() {
			this.searching = true;
		},
		reset() {
			this.searching = false;
			this.search.searchInput = null;
		}
	}
});
</script>
