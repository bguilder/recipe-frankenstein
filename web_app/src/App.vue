<template>
  <v-app>
    <!-- Start Main Header Bar -->
    <v-app-bar app color="#00C279" dark>
      <v-row align="start" justify="start">
        <v-col cols="12">
          <v-btn x-large color="white" text @click="reset" class="pa-2">
            Recipe Frankenstein</v-btn
          >
        </v-col>
      </v-row>
    </v-app-bar>

    <!-- Start Result Page -->
    <template v-if="searching">
      <v-content>
        <Recipes
          v-bind:search="search"
          v-bind:showRecipeList="showRecipeList"
          v-bind:showFrequency="showFrequency"
        />
      </v-content>
    </template>

    <!-- Start Search Bar -->
    <template v-else>
      <v-container fluid fill-height>
        <v-row no-gutters>
          <v-col>
            <v-row align="center" justify="center">
              <v-col cols="12" class="pa-2">
                <v-img
                  contain
                  max-height="125"
                  src="./assets/frankenstein.png"
                ></v-img>
              </v-col>
            </v-row>
            <v-row align="center" justify="center">
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
                      v-model="search.input"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row align="center" justify="center">
                  <v-spacer />
                  <v-btn
                    outlined
                    color="#00C279"
                    elevation-20
                    large
                    @click="submit"
                    >Search Recipes</v-btn
                  >
                  <v-spacer />
                  <v-spacer />
                  <v-dialog v-model="showFeelingHungry" max-width="400">
                    <template v-slot:activator="{ on }">
                      <v-btn
                        outlined
                        color="#00C279"
                        elevation-20
                        large
                        v-on="on"
                        @click="randomRecipes()"
                        >Feelin' Hungry</v-btn
                      >
                    </template>
                    <v-card>
                      <v-card-title class="headline"
                        >Feelin' Hungry</v-card-title
                      >

                      <v-card-text
                        >Choose an option below to get started!
                      </v-card-text>
                      <v-card-text class="pb-0">
                        <v-radio-group
                          v-model="search.input"
                          class="ma-2"
                          v-for="(item, index) in feelingHungryRecipes2"
                          :key="index"
                        >
                          <v-radio
                            v-bind:label="item"
                            v-bind:value="item"
                          ></v-radio>
                        </v-radio-group>
                      </v-card-text>
                      <v-card-actions class="pt-0">
                        <v-btn
                          text
                          color="blue"
                          class="ml-5"
                          @click="randomRecipes()"
                        >
                          Shuffle
                          <v-icon large dark color="blue">mdi-refresh</v-icon>
                        </v-btn>
                        <v-spacer></v-spacer>
                        <v-btn
                          color="red darken-1"
                          text
                          @click="showFeelingHungry = false"
                          >Cancel</v-btn
                        >
                        <v-btn color="green darken-1" text @click="submit()"
                          >Search!</v-btn
                        >
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                  <v-spacer />
                </v-row>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-container>
    </template>
    <footer class="caption font-weight-thin text-center">
      Icon made by Freepik from www.flaticon.com <br />v0.2.0-alpha<br />
    </footer>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Recipes from "./components/Recipes.vue";
import axios from "axios";

export default Vue.extend({
  name: "App",

  components: {
    Recipes
  },

  data() {
    return {
      activeColor: "green",
      showFeelingHungry: false,
      searching: false,
      search: {
        input: null,
        recipeCount: 7
      },
      feelingHungryRecipes2: ["test"],
      feelingHungryRecipes: [],
      drawer: false,
      showRecipeList: true,
      showFrequency: true
    };
  },
  mounted() {
    axios.get(this.getFeelingHungryUrl()).then(response => {
      this.feelingHungryRecipes = response.data;
    });
  },
  methods: {
    getFeelingHungryUrl() {
      return process.env.VUE_APP_FEELING_HUNGRY_URL;
    },
    submit() {
      this.searching = true;
    },
    feelingHungry() {
      this.showFeelingHungry = true;
    },
    reset() {
      this.searching = false;
      this.search.input = null;
      this.showFeelingHungry = false;
    },
    toggleRecipeList() {
      this.showRecipeList = !this.showRecipeList;
    },
    toggleFrequency() {
      this.showFrequency = !this.showFrequency;
    },
    randomRecipes() {
      this.feelingHungryRecipes2 = this.feelingHungryRecipes
        .sort(() => 0.5 - Math.random())
        .slice(0, 5);
    },
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
