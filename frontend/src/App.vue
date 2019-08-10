<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <span class="font-weight-light">CLUGO</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn tile outlined>
        <span class="mr-2">Save</span>
      </v-btn>
      <v-btn class="ma-2" tile outlined color="blue">
        <v-icon left>mdi-upload</v-icon> Publish
      </v-btn>
      
    </v-app-bar>

    <v-content>
      <HelloWorld/>
    </v-content>
  </v-app>
</template>

<script>
import HelloWorld from './components/HelloWorld';

export default {
  name: 'App',
  components: {
    HelloWorld,
  },
  data: () => ({
    //
  }),
};
</script>
