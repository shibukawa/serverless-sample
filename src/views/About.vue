<template>
  <div slot="appContent">
    <mwc-top-app-bar>
      <mwc-icon-button
        slot="navigationIcon"
        icon="menu"
        v-on:click="toggleDrawer"
      ></mwc-icon-button>
      <div slot="title">About</div>
    </mwc-top-app-bar>
    <div class="main-content">
      <p>
        Sample WebSite for serverless admin screen.
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import "@material/mwc-top-app-bar";
import "@material/mwc-icon-button";
import { Drawer } from "@material/mwc-drawer";

@Component
export default class AboutPage extends Vue {
  toggleDrawer(e: Event) {
    const drawer = this.$parent.$el as Drawer;
    drawer.open = !drawer.open;
  }
}
</script>
