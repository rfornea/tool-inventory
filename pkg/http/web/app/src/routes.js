import Vue from "vue";
import VueRouter from "vue-router";

import BooksList from "./components/BooksList";
import UpdateBook from "./components/UpdateBook";
import CreateBook from "./components/CreateBook";
import StatusList from "./components/StatusList";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", redirect: "/books" },
    { path: "/books", name: "view-books", component: BooksList },
    { path: "/books-status", name: "books-status", component: StatusList },
    { path: "/books/:id", name: "update-book", component: UpdateBook },
    { path: "/new-book", name: "create-book", component: CreateBook },
    { path: "*", redirect: "/books" }
  ]
});
