<template>
  <mwc-drawer hasHeader type="modal" :open="openDrawer">
    <span slot="title">Drawer Title</span>
    <span slot="subtitle">subtitle</span>
    <div class="drawer-content">
      <p>Drawer content</p>
      <p><a @click="selectPage('/')">Home</a></p>
      <p><a @click="selectPage('/prime')">Prime</a></p>
      <p><a @click="selectPage('/about')">About</a></p>
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

@Component
export default class AppComponent extends Vue {
  openDrawer: boolean = false;

  mounted() {
    this.$router.afterEach(() => {
      this.openDrawer = false;
    });
  }

  selectPage(path: string) {
    if (path !== this.$router.currentRoute.path) {
      this.$router.push(path);
    }
    return false;
  }
}
</script>
