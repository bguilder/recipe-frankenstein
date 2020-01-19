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

		<template v-if="searching">
			<v-content>
				<HelloWorld v-bind:searchInput="searchInput" />
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
									<v-text-field
										label="Search..."
										hide-details
										single-line
										outlined
										rounded
										v-model="searchInput"
									></v-text-field>
								</v-row>
								<v-row
									align="center"
									justify="center"
								>
									<v-btn
										color="primary"
										text
										@click="submit"
										class="pa-7"
									>Submit</v-btn>
								</v-row>

							</v-col>
						</v-row>
					</v-col>
				</v-row>
			</v-container>
		</template>
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
			searchInput: null
		};
	},
	methods: {
		submit() {
			this.searching = true;
		}
	}
});
</script>
