import axios from "axios";

const BASE_URI = process.env.VUE_APP_HOST;
const client = axios.create({
  baseURL: BASE_URI,
  json: true
});

const APIClient = {
  createBook(book) {
    return this.perform("post", "/books", book);
  },

  deleteBook(book) {
    return this.perform("delete", `/books/${book.id}`);
  },
  updateBook(book) {
    return this.perform("put", `/books/${book.id}`, book);
  },

  getBooks() {
    return this.perform("get", "/books");
  },

  async perform(method, resource, data) {
    return client({
      method,
      url: resource,
      data,
      headers: {}
    }).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
