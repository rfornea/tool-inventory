<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="teal">
                <v-toolbar-title justify-center
                  >Book Status List</v-toolbar-title
                >
              </v-toolbar>
              <v-card-text>
                <span v-for="book in books" v-bind:key="book.id">
                  <span class="book-title-status">
                    {{ book.title }}
                  </span>
                  <span class="book-status-status" v-if="book.available">
                    Available
                  </span>
                  <span class="book-status-status" v-if="!book.available">
                    Checked Out
                  </span>
                  <br />
                  <div class="book-separator"></div>
                  <br />
                </span>
              </v-card-text>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import { mapActions } from "vuex";
import Toasted from "vue-toasted";
import Vue from "vue";

Vue.use(Toasted, {
  duration: 4000
});

export default {
  name: "BooksList",
  data() {
    return {};
  },
  computed: {
    books() {
      return this.$store.state.books;
    }
  },
  created() {
    if (
      this.books === undefined ||
      this.books === null ||
      this.books.length === 0
    ) {
      this.$store.dispatch("getAllBooksStore");
    }
  },
  methods: {
    ...mapActions(["getAllBooksStore"])
  }
};
</script>

<style>
.book-separator {
  border-bottom: 3px solid teal;
}
.book-title-status {
  clear: left;
  float: left;
  width: 60%;
  display: inline-block;
}
book-status-status {
  float: left;
  width: 35%;
  display: inline-block;
}
</style>
