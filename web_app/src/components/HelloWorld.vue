<template>
	<v-container fluid text-wrap>
		<v-row             
            justify="space-around">
			<v-col cols="6">
				<h2 class="display-1 font-weight-bold mb-2">Recipes</h2>
				<div
				v-for="(item, index) in recipes"
				:key="index"
			>
			<v-card>
				<v-card-title v-text="item.Title"></v-card-title>
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
			</v-card>
			<br>
			</div>
			</v-col>
			<v-col cols="6">
			<h1 class="display-1 font-weight-bold mb-2">Ingredients</h1>
			<div
				v-for="(item, indexs) in ingredients"
				:key="indexs"
			>
				{{item.Key}} 
				{{item.Value}}

			</div>
			</v-col>
		</v-row>
	</v-container>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
	name: "HelloWorld",
	props: {
		msg: String
	},
	data() {
		return {
			ingredients: null,
			recipes: null,
		};
	},
	mounted() {
		// axios.get("http://127.0.0.1:8088/example").then(response => {
		// 	this.messageEx = response.data;
		// });
		axios.get("http://127.0.0.1:8088/search").then(response => {
			this.recipes = response.data.Recipes;
			this.ingredients = response.data.Ingredients;
		});
	}
});
</script>
