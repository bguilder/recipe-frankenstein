<template>
	<div class="hello">
		<h1>Empenada Recipes</h1>
		<ul>
			<li	v-for="(item, index) in recipes" :key="index">
			<strong>{{ item.Title }}</strong><br>
			<strong>Ingredients</strong>
			<ul>
				<li	v-for="(ing, index) in item.Ingredients" :key="index">
				{{ ing }}<br>
				</li>
			</ul>
			<strong>Directions</strong>
			<ul>
				<li	v-for="(dir, index) in item.Directions" :key="index">
				{{index+1}}. {{ dir }}<br>
				</li>
			</ul>
			<hr>
			</li>
		</ul>
		<h1>Empenada Ingredients</h1>
		<ul>
			<li	v-for="(item, index) in ingredients" :key="index">
			<strong>{{ item.Key }}</strong> - <strong>{{ item.Value }}</strong><br>
			</li>
		</ul>
	</div>
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
	margin: 40px 0 0;
}
ul {
	list-style-type: none;
	padding: 0;
}
li {
	margin: 0 10px;
}
a {
	color: #42b983;
}
</style>
