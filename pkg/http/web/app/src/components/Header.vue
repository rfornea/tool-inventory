<template>
  <fixed-header class="pa-3 white--text teal-bg" color="teal">
    <div class="navbar">
      <nav>
        <router-link class="books-nav" to="/books">Manage Books</router-link>
        <router-link class="books-nav" to="/new-book">Add New Book</router-link>
      </nav>
    </div>
  </fixed-header>
</template>

<script>
import FixedHeader from "vue-fixed-header";

export default {
  name: "Header",
  components: {
    FixedHeader
  }
};
</script>

<style>
.books-nav {
  color: white !important;
  margin-right: 25px;
}
.teal-bg {
  background-color: #009688 !important;
}
</style>
