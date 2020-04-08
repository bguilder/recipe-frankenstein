<template >
  <v-container v-if="ingredients" fluid text-wrap>
    <v-row>
      <!-- Show Recipes List -->
      <v-col v-if="!isMobile()" cols="8">
        <RecipeList v-bind:recipes="recipes" />
      </v-col>
      <v-col v-else-if="!showFrequency" cols="12">
        <RecipeList v-bind:recipes="recipes" />
      </v-col>

      <!-- Show Ingredient Frequency -->
      <v-col v-if="!isMobile()" cols="4">
        <IngredientFrequency v-bind:ingredients="ingredients" />
      </v-col>
      <v-col v-else-if="showFrequency" cols="12" text-center>
        <IngredientFrequency v-bind:ingredients="ingredients" />
      </v-col>
    </v-row>
  </v-container>

  <!--  searching progress bar -->
  <v-container v-else fluid fill-height>
    <v-row align="center" justify="center">
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
import RecipeList from "./RecipeList.vue";
import IngredientFrequency from "./IngredientFrequency.vue";

export default Vue.extend({
  name: "Recipes",
  components: {
    RecipeList,
    IngredientFrequency
  },
  props: {
    search: Object,
    showFrequency: Boolean
  },
  data() {
    return {
      ingredients: null,
      recipes: null,
      toggleColor: "green"
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
    },
    isMobile() {
      if (
        /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
          navigator.userAgent
        )
      ) {
        return true;
      } else {
        return false;
      }
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
