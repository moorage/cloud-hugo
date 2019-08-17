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

    <v-content style="height: 100%">
      <v-row no-gutters style="height: 100%" align="stretch">
        <v-col md="auto">
          <v-card
            class="pa-2"
            outlined
            tile
          >
            Variable width contentjlsdkfjsdlkf
          </v-card>
        </v-col>
        <v-col>
          <MdeEditor/>
        </v-col>
      </v-row>
      
    </v-content>
  </v-app>
</template>

<script>
import MdeEditor from './components/MdeEditor.vue';

export default {
  name: 'App',
  components: {
    MdeEditor,
  },
  data: () => ({
    //
  }),
};
</script>
