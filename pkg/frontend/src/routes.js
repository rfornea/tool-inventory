import Vue from "vue";
import VueRouter from "vue-router";

import ToolsList from "./components/ToolsList";
import UpdateTool from "./components/UpdateTool";
import CreateTool from "./components/CreateTool";
import ToolDetails from "./components/ToolDetails";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", redirect: "/tools" },
    { path: "/tools", name: "view-tools", component: ToolsList },
    { path: "/tools/:id", name: "update-tool", component: UpdateTool },
    { path: "/tool-details/:id", name: "tool-details", component: ToolDetails },
    { path: "/new-tool", name: "create-tool", component: CreateTool },
    { path: "*", redirect: "/tools" }
  ]
});
