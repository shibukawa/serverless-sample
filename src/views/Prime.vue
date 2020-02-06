<template>
  <div slot="appContent">
    <mwc-top-app-bar>
      <mwc-icon-button
        slot="navigationIcon"
        icon="menu"
        v-on:click="toggleDrawer"
      ></mwc-icon-button>
      <div slot="title">Prime</div>
    </mwc-top-app-bar>
    <div class="main-content">
      <div>
        <mwc-textfield
          :value="num"
          @change="onchange($event.target.value)"
          label="Input Number"
        />
      </div>
      <div>
        <mwc-button @click="check()">Is it Prime Number?</mwc-button>
      </div>
    </div>
    <mwc-dialog :open="dialogOpen" @closing="closingDialog">
      <div>{{ this.result }}</div>
      <mwc-button slot="primaryAction" dialogAction="discard">
        OK
      </mwc-button>
    </mwc-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import "@material/mwc-top-app-bar";
import "@material/mwc-icon-button";
import "@material/mwc-button";
import "@material/mwc-dialog";
import "@material/mwc-textfield";
import AppComponent from '../App.vue';

@Component
export default class AboutPage extends Vue {
  result: string = "";
  num: string = "0";
  dialogOpen: boolean = false;

  closingDialog() {
    this.dialogOpen = false;
  }

  onchange(value: string) {
    this.num = value;
  }

  async check() {
    const res = await fetch("/api", {
      method: "post",
      body: JSON.stringify({
        jsonrpc: "2.0",
        method: "CheckPrimeNumber",
        params: [parseInt(this.num)],
        id: 1
      }),
      headers: {
        "content-type": "application/json",
        accept: "application/json"
      }
    });
    if (res.ok) {
      try {
        const result = await res.json() as any;
        if (result.result) {
          this.result = `${this.num} is prime number`;
        } else {
          this.result = `${this.num} is not prime number`;
        }
      } catch (e) {
        this.result = "parse error";
      }
    } else {
      this.result = "server access error";
    }
    this.dialogOpen = true;
  }

  toggleDrawer(e: Event) {
    const parent = this.$parent as AppComponent;
    parent.openDrawer = !parent.openDrawer;
  }
}
</script>
