<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
	<h2 class="display-1 font-weight-bold">Recipe Frankenstein</h2>
	<v-spacer></v-spacer>
	<v-col cols="12" sm="6" md="3">
 <v-text-field 
	label="Search"
    filled hide-details>
	</v-text-field>
		</v-col>
    </v-app-bar>

    <v-content>
      <HelloWorld/>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import HelloWorld from './components/HelloWorld.vue';

export default Vue.extend({
  name: 'App',

  components: {
    HelloWorld,
  },

  data: () => ({
    //
  }),
});
</script>
