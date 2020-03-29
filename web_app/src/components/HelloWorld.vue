<template >
	<v-container
		v-if="ingredients"
		fluid
		text-wrap
	>
		<v-row>
			<v-col cols="8">
				<h2 class="display-1 font-weight-bold mb-2 text-center">Recipe List</h2>
				<v-expansion-panels>
					<v-expansion-panel
						v-for="(item, index) in recipes"
						:key="index"
					>
						<v-expansion-panel-header>
							<p
								color="blue"
								class=" font-weight-bold"
							>{{item.Title}}</p>
						</v-expansion-panel-header>
						<v-expansion-panel-content class="pa-0 ma-0">
							<v-row no-gutters>
								<v-col cols="5">
									<v-card-title>
										<p class="display-5">Ingredients</p>
									</v-card-title>
									<v-card-subtitle
										class="pa-0 ma-0 ml-4"
										v-for="(ing, i) in item.Ingredients"
										:key="i"
									>
										<p class="text-sm mb-2">* {{ ing }}</p>
									</v-card-subtitle>
								</v-col>
								<v-col cols="7">
									<v-card-title>
										<p class="display-5">Directions</p>

									</v-card-title>
									<v-card-subtitle
										class="pa-1 ma-0"
										v-for="(dir, ind) in item.Directions"
										:key="ind + 100"
									> {{ind + 1}}. {{ dir }}
									</v-card-subtitle>
								</v-col>
							</v-row>
							<v-card-subtitle style="font-size:.80em">Reference:
								<a
									class="pt-2"
									v-bind:href="item.URL"
								>{{item.URL}}</a>
							</v-card-subtitle>
						</v-expansion-panel-content>
					</v-expansion-panel>
					<br>
				</v-expansion-panels>
			</v-col>
			<v-col cols="4">
				<div>
					<h2 class="display-1 font-weight-bold mb-2 text-center">Ingredient Frequency</h2>
					<v-card class="sticky pa-5">
						<div
							v-for="(item, indexs) in ingredients"
							:key="indexs"
						>
							- {{item.Key}}
							<span v-bind:style="{ color: activeColor, fontSize: `.7em`}">({{item.Value}})</span>

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
				color="#00C279"
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
			activeColor: "green",
			ingredients: null,
			recipes: null
		};
	},
	methods: {
		getFeelingHungryUrl() {
			return (
				process.env.VUE_APP_RECIPE_SEARCH_URL +
				"?recipe=" +
				this.search.input +
				"&count=" +
				this.search.recipeCount
			);
		}
	},
	mounted() {
		axios.get(this.getFeelingHungryUrl()).then(response => {
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
	max-height: 72vh;
	overflow: auto;
}
</style>
