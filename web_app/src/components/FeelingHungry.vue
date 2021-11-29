
<template>
  <!-- TODO: Fix this... -->
  <v-card>
    <v-card-title class="headline">Feelin' Hungry</v-card-title>
    <v-card-text>Choose an option below to get started! </v-card-text>
    <v-card-text class="pb-0">
      <v-radio-group
        v-model="search.input"
        class="ma-2"
        v-for="(item, index) in feelingHungryRecipes2"
        :key="index"
      >
        <v-radio v-bind:label="item" v-bind:value="item"></v-radio>
      </v-radio-group>
    </v-card-text>
    <v-card-actions class="pt-0">
      <v-btn text color="blue" class="ml-5" @click="randomRecipes()">
        Shuffle
        <v-icon large dark color="blue">mdi-refresh</v-icon>
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn color="red darken-1" text @click="showFeelingHungry = false"
        >Cancel</v-btn
      >
      <v-btn color="green darken-1" text @click="submit()">Search!</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
  name: "FeelingHungry",
  props: {
    feelingHungryRecipes: Array
  },
  data() {
    return {
      feelingHungryRecipes2: ["test"]
    };
  },
  mounted() {
    this.randomRecipes();
  },
  methods: {
    randomRecipes() {
      this.feelingHungryRecipes2 = this.$props.feelingHungryRecipes
        .sort(() => 0.5 - Math.random())
        .slice(0, 5);
    }
  }
});
</script>