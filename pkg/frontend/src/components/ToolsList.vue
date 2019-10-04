<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="#3d5496">
                <v-toolbar-title justify-center
                  >Tools Inventory</v-toolbar-title
                >
              </v-toolbar>
              <v-card-text>
                <span v-if="tools.length === 0">
                  You don't have any tools in your inventory yet.
                  <router-link to="/new-tool">Click here</router-link>
                  to add some.
                </span>
                <span v-for="tool in tools" v-bind:key="tool.id">
                  <h6 class="tool-name">
                    {{ tool.name }}
                  </h6>
                  <div class="half-width-left">
                    <div class="tool-model">Model: {{ tool.model }}</div>
                    <div class="tool-manufacturer">
                      Manufacturer: {{ tool.manufacturer }}
                    </div>
                  </div>
                  <div class="container half-width-right">
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-danger tool-delete-btn"
                      @click="deleteTool(tool)"
                    >
                      Delete
                    </button>
                    <button
                      type="button"
                      class="btn btn-sm btn-outline-info tool-details-btn"
                      @click="goToDetailPage(tool)"
                    >
                      Details
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
                  <div class="tool-separator"></div>
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
  name: "ToolsList",
  data() {
    return {};
  },
  computed: {
    tools() {
      return this.$store.state.tools;
    }
  },
  created() {
    if (
      this.tools === undefined ||
      this.tools === null ||
      this.tools.length === 0
    ) {
      this.$store.dispatch("getAllToolsStore");
    }
  },
  methods: {
    ...mapActions(["getAllToolsStore", "updateToolInStore", "removeTool"]),
    goToUpdatePage(tool) {
      this.$router.push({
        name: "update-tool",
        location: "tools/" + tool.id,
        params: { id: tool.id, tool: { ...tool } }
      });
    },
    goToDetailPage(tool) {
      this.$router.push({
        name: "tool-details",
        location: "tool-details/" + tool.id,
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
