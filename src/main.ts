import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";

Vue.config.productionTip = false;

Vue.config.ignoredElements = [
  "mwc-drawer",
  "mwc-top-app-bar",
  "mwc-button",
  "mwc-icon-button",
  "mwc-dialog",
  "mwc-textfield"
];

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
