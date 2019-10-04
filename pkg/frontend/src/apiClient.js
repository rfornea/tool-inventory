import axios from "axios";
const BASE_URI = process.env.VUE_APP_HOST;

const client = axios.create({
  baseURL: BASE_URI
});

const APIClient = {
  getImageUrl(id) {
    let imageURL = BASE_URI + "/assets/images/" + id + ".png";
    return this.performJsonRequest("get", imageURL)
      .then(() => {
        return imageURL;
      })
      .catch(() => {
        return BASE_URI + "/assets/static/no-image.png";
      });
  },

  createTool(tool) {
    return this.performFormRequest("post", "/tools", tool);
  },

  deleteTool(tool) {
    return this.performJsonRequest("delete", `/tools/${tool.id}`);
  },
  updateTool(tool) {
    return this.performFormRequest("put", `/tools/${tool.get("id")}`, tool);
  },

  getTools() {
    return this.performJsonRequest("get", "/tools");
  },

  async performJsonRequest(method, resource, data) {
    return client({
      json: true,
      method,
      url: resource,
      data,
      headers: {
        "Content-Type": "multipart/form-data"
      }
    }).then(req => {
      return req.data;
    });
  },

  async performFormRequest(method, resource, data) {
    return client({
      method,
      url: resource,
      data,
      headers: {
        "Content-Type": "multipart/form-data"
      }
    }).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
