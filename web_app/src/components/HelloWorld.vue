<template >
	<v-container
		v-if="ingredients"
		fluid
		text-wrap
	>
		<v-row>
			<v-col cols="6">
				<h2 class="display-1 font-weight-bold mb-2">Recipes</h2>
				<v-expansion-panels>
					<v-expansion-panel
						v-for="(item, index) in recipes"
						:key="index"
					>
						<v-expansion-panel-header>{{item.Title}}</v-expansion-panel-header>
						<v-expansion-panel-content><a v-bind:href="item.URL">{{item.URL}}</a></v-expansion-panel-content>
						<v-expansion-panel-content>
							<v-card-title>Ingredients</v-card-title>
							<v-card-subtitle
								v-for="(ing, i) in item.Ingredients"
								:key="i"
							> {{ ing }}
							</v-card-subtitle>
							<v-card-title>Directions</v-card-title>
							<v-card-subtitle
								v-for="(dir, ind) in item.Directions"
								:key="ind + 100"
							> {{ind + 1}}. {{ dir }}
							</v-card-subtitle>
						</v-expansion-panel-content>
					</v-expansion-panel>
					<br>
				</v-expansion-panels>
			</v-col>
			<v-col cols="6">
				<div class="sticky">
					<h1 class="display-1 font-weight-bold mb-2">Ingredients</h1>
					<v-card class="pa-5">
						<div
							v-for="(item, indexs) in ingredients"
							:key="indexs"
						>
							{{item.Key}}
							{{item.Value}}

						</div>
					</v-card>
				</div>

			</v-col>
		</v-row>
	</v-container>
	<v-container
		v-else
		fluid
		fill-height
	>
		<v-row
			align="center"
			justify="center"
		>
			<v-progress-circular
				size="100"
				width="10"
				indeterminate
				color="green"
			></v-progress-circular>
		</v-row>
	</v-container>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
	name: "HelloWorld",
	props: {
		msg: String,
		search: Object
	},
	data() {
		return {
			ingredients: null,
			recipes: null
		};
	},

	mounted() {
		axios
			.get(
				"https://1v1zwuknkf.execute-api.us-east-1.amazonaws.com/v1?recipe=" +
					this.search.searchInput +
					"&count=" +
					this.search.recipeCount
			)
			.then(response => {
				this.recipes = response.data.Recipes;
				this.ingredients = response.data.Ingredients;
			});
	}
});
</script>

<style scoped>
.sticky {
	position: -webkit-sticky;
	position: sticky;
	top: 4rem;
	max-height: 80vh;
	overflow: auto;
}
</style>
