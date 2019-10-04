<template>
  <fixed-header class="pa-3 white--text header-bg">
    <div class="navbar">
      <nav>
        <router-link class="tools-nav" to="/tools">Manage Tools</router-link>
        <router-link class="tools-nav" to="/new-tool">Add New Tool</router-link>
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

<style></style>
