<template>
  <mwc-drawer hasHeader type="modal">
    <span slot="title">Drawer Title</span>
    <span slot="subtitle">subtitle</span>
    <div class="drawer-content">
      <p>Drawer content</p>
      <p><a v-on:click="selectPage('/')">Home</a></p>
      <p><a v-on:click="selectPage('/prime')">Prime</a></p>
      <p><a v-on:click="selectPage('/about')">About</a></p>
    </div>
    <router-view />
  </mwc-drawer>
</template>

<style lang="scss">
.drawer-content {
  padding: 0px 16px 0 16px;
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import "@material/mwc-icon-button";
import "@material/mwc-drawer";
import { Drawer } from "@material/mwc-drawer";

@Component
export default class AppComponent extends Vue {
  selectPage(path: string) {
    if (path !== this.$router.currentRoute.path) {
      const drawer = this.$el as Drawer;
      drawer.open = false;
      this.$router.push(path);
    }
    return false;
  }
}
</script>
