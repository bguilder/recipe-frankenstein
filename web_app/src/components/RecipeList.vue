<template>
  <div class="sticky">
    <h2 class="display-1 font-weight-bold mb-2 text-center">Recipe List</h2>
    <v-expansion-panels>
      <v-expansion-panel v-for="(item, index) in recipes" :key="index">
        <v-expansion-panel-header>
          <p color="blue" class=" font-weight-bold">{{ item.Title }}</p>
        </v-expansion-panel-header>
        <v-expansion-panel-content class="pa-0 ma-0">
          <!-- Not Mobile -->
          <v-row v-if="!isMobile()" no-gutters>
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
              >
                {{ ind + 1 }}. {{ dir }}
              </v-card-subtitle>
            </v-col>
          </v-row>

          <!-- Mobile -->
          <v-row v-else no-gutters>
            <v-col cols="12">
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
            <v-col cols="12">
              <v-card-title>
                <p class="display-5">Directions</p>
              </v-card-title>
              <v-card-subtitle
                class="pa-1 ma-0"
                v-for="(dir, ind) in item.Directions"
                :key="ind + 100"
              >
                {{ ind + 1 }}. {{ dir }}
              </v-card-subtitle>
            </v-col>
          </v-row>

          <v-card-subtitle style="font-size:.80em"
            >Reference:
            <a class="pt-2" v-bind:href="item.URL">{{ item.URL }}</a>
          </v-card-subtitle>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <br />
    </v-expansion-panels>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
  name: "RecipeList",
  props: {
    recipes: Array
  },
  methods: {
    isMobile() {
      // TODO: Add responsive small screen here
      // TODO: Don't copy this function everywhere
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