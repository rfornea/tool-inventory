<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="#3d5496">
                <v-toolbar-title justify-center>Tool Detail</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <span>
                  <h3 class="tool-name">
                    {{ tool.name }}
                  </h3>
                  <img
                    class="tool-image"
                    :src="imageUrl"
                    alt="picture of tool"
                    onerror="this.onerror=null;this.src='../assets/no-image.png';"
                  />
                  <div class="tool-description">{{ tool.description }}</div>
                  <div class="tool-model">Model: {{ tool.model }}</div>
                  <div class="tool-manufacturer">
                    Manufacturer: {{ tool.manufacturer }}
                  </div>
                  <div class="container">
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-danger tool-delete-btn"
                      @click="deleteTool(tool)"
                    >
                      Delete
                    </button>
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-primary tool-update-btn"
                      @click="goToUpdatePage(tool)"
                    >
                      Update
                    </button>
                  </div>
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
import Vue from "vue";
import Toasted from "vue-toasted";
import APIClient from "../apiClient";

Vue.use(Toasted, {
  duration: 4000
});

export default {
  name: "ToolDetails",
  data() {
    return {
      tool: { ...this.$route.params.tool },
      imageUrl: ""
    };
  },
  created() {
    if (this.tool === undefined || this.tool === null) {
      this.$router.push({ name: "view-tools", location: "tools" });
    }
    APIClient.getImageUrl(this.tool.id)
      .then(res => {
        this.imageUrl = res;
      })
      .catch(res => {
        this.imageUrl = res;
      });
  },
  methods: {
    ...mapActions(["removeTool"]),
    goToUpdatePage(tool) {
      this.$router.push({
        name: "update-tool",
        location: "tools/" + tool.id,
        params: { id: tool.id, tool: { ...tool } }
      });
    },
    deleteTool: function(tool) {
      let proceedWithDelete = confirm(
        "Sure you want to delete " + tool.name + "?"
      );
      if (proceedWithDelete) {
        this.$store
          .dispatch("removeTool", tool)
          .then(() => {
            this.$toasted.show(tool.name + " has been deleted.", {
              type: "success",
              theme: "bubble"
            });
            this.$router.push({ name: "view-tools", location: "tools" });
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

<style></style>
