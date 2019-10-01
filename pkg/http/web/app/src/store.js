import Vue from "vue";
import Vuex from "vuex";
import _ from "lodash";

import APIClient from "./apiClient";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    books: []
  },
  mutations: {
    resetBooks(state, books) {
      state.books = books;
    }
  },
  getters: {
    books(state) {
      return state.books;
    }
  },
  actions: {
    getAllBooksStore({ commit }) {
      APIClient.getBooks().then(data => {
        commit("resetBooks", data.books);
      });
    },
    updateBookInStore({ commit, state }, book) {
      const books = [...state.books];
      books[
        _.findIndex(books, function(bk) {
          return bk.id === book.id;
        })
      ] = book;
      return APIClient.updateBook(book).then(() => {
        commit("resetBooks", books);
      });
    },
    removeBook({ commit, state }, book) {
      const books = [...state.books];
      _.remove(books, function(bk) {
        return bk.id === book.id;
      });
      return APIClient.deleteBook(book).then(() => {
        commit("resetBooks", books);
      });
    },
    addBook({ commit, state }, book) {
      const books = [...state.books];
      return APIClient.createBook(book).then(book => {
        books[books.length] = book;
        commit("resetBooks", books);
      });
    }
  }
});

export default store;
