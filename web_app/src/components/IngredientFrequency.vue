<template>
  <div>
    <h2 class="display-1 font-weight-bold mb-2 text-center">
      Ingredient Frequency
    </h2>
    <v-card class="sticky pa-5">
      <div v-for="(item, indexs) in ingredients" :key="indexs">
        - {{ item.Name }}
        <span v-bind:style="{ color: 'green', fontSize: `.7em` }"
          >({{ item.Count }})</span
        >
      </div>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
  name: "IngredientFrequency",
  props: {
    ingredients: Array
  }
});
</script>


<style scoped>
/* TODO: move this to a stylesheet */
.sticky {
  position: -webkit-sticky;
  position: sticky;
  top: 4rem;
  max-height: 72vh;
  overflow: auto;
}
</style>