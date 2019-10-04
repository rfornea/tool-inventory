import "@babel/polyfill";
import Vue from "vue";
import vuetify from "./plugins/vuetify";
import App from "./App.vue";
import store from "./store";
import router from "./routes";
import VueToasted from "vue-toasted";
import "./root.css";

Vue.config.productionTip = process.env.NODE_ENV == "production";

new Vue({
  store,
  router,
  vuetify,
  VueToasted,
  render: h => h(App)
}).$mount("#app");
