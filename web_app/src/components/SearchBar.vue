<template>
	<v-container>
		<v-row>
			<v-col>

				<v-text-field
					label="Search..."
					hide-details
					single-line
					outlined
					rounded
					color="#00C279"
				></v-text-field>
				<!-- <v-btn icon>
					<v-icon>search</v-icon>
				</v-btn> -->
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
		msg: String,
		searchInput: String
	},
	data() {
		return {
			ingredients: null,
			recipes: null
		};
	}
});
</script>