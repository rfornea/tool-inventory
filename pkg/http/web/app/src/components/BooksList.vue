<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="teal">
                <v-toolbar-title justify-center>Books</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <span v-for="book in books" v-bind:key="book.id">
                  <h3 class="book-title">
                    {{ book.title }}
                  </h3>
                  <div class="book-description">{{ book.description }}</div>
                  <div class="book-author-name">
                    {{ book.authorFirstName }} {{ book.authorLastName }}
                  </div>
                  <div class="book-isbn">ISBN: {{ book.isbn }}</div>
                  <div class="book-status" v-if="book.available">
                    Status: Available
                  </div>
                  <div class="book-status" v-if="!book.available">
                    Status: Checked Out
                  </div>
                  <div class="container clear-all">
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-danger book-delete-btn"
                      @click="deleteBook(book)"
                    >
                      Delete
                    </button>
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-info book-checkout-btn"
                      @click="toggleBookStatus(book)"
                    >
                      Check {{ book.available ? " Out" : " In" }}
                    </button>
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-primary book-update-btn"
                      @click="goToUpdatePage(book)"
                    >
                      Update
                    </button>
                  </div>
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
    ...mapActions(["getAllBooksStore", "updateBookInStore", "removeBook"]),
    goToUpdatePage(book) {
      this.$router.push({
        name: "update-book",
        location: "books/" + book.id,
        params: { id: book.id, book: { ...book } }
      });
    },
    toggleBookStatus: function(book) {
      book.available = !book.available;
      this.$store
        .dispatch("updateBookInStore", book)
        .then(() => {
          let status = book.available ? "in." : "out.";
          let message = book.title + " has been checked " + status;
          this.$toasted.show(message, {
            type: "success",
            theme: "bubble"
          });
        })
        .catch(() => {
          this.$toasted.show("There was an error.", {
            type: "danger",
            theme: "bubble"
          });
        });
    },
    deleteBook: function(book) {
      let proceedWithDelete = confirm(
        "Sure you want to delete " + book.title + "?"
      );
      if (proceedWithDelete) {
        this.$store
          .dispatch("removeBook", book)
          .then(() => {
            this.$toasted.show(book.title + " has been deleted.", {
              type: "success",
              theme: "bubble"
            });
          })
          .catch(() => {
            this.$toasted.show("There was an error.", {
              type: "danger",
              theme: "bubble"
            });
          });
      }
    }
  }
};
</script>

<style>
.book-checkout-btn {
  float: right;
  margin-right: 5px;
}
.book-delete-btn {
  float: right;
}
.book-update-btn {
  float: right;
  margin-right: 5px;
}
.book-separator {
  border-bottom: 3px solid teal;
}
.book-status {
  clear: left;
  display: block;
}
.book-author-name {
  clear: left;
  display: block;
}
.book-isbn {
  clear: left;
  display: block;
}
.book-title {
  float: left;
  width: 35%;
  display: inline-block;
}
.clear-all {
  clear: both;
}
.book-description {
  width: 60%;
  float: right;
  overflow: auto;
  height: 150px;
}
</style>
