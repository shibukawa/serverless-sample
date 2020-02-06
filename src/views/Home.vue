<template>
  <div slot="appContent">
    <mwc-top-app-bar>
      <mwc-icon-button
        slot="navigationIcon"
        icon="menu"
        v-on:click="toggleDrawer"
      ></mwc-icon-button>
      <div slot="title">Home</div>
    </mwc-top-app-bar>
    <div class="main-content">
      <img alt="Vue logo" src="../assets/logo.png" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import "@material/mwc-top-app-bar";
import "@material/mwc-icon-button";
import AppComponent from '../App.vue';

@Component
export default class HomeComponent extends Vue {
  toggleDrawer(e: Event) {
    const parent = this.$parent as AppComponent;
    parent.openDrawer = !parent.openDrawer;
  }
}
</script>
