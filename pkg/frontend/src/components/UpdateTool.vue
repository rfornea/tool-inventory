<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="#3d5496">
                <v-toolbar-title justify-center>Update Tool</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <label>Name:</label>
                <input
                  required
                  :class="validNameModelManufacturerStyle(tool.name, 'name')"
                  v-model="tool.name"
                  placeholder="* Enter Name"
                  @blur="
                    () => {
                      dirty.name = true;
                    }
                  "
                />
                <label>Model:</label>
                <input
                  required
                  :class="validNameModelManufacturerStyle(tool.model, 'model')"
                  v-model="tool.model"
                  placeholder="* Enter Model"
                  @blur="
                    () => {
                      dirty.model = true;
                    }
                  "
                />
                <label>Manufacturer:</label>
                <input
                  required
                  :class="
                    validNameModelManufacturerStyle(
                      tool.manufacturer,
                      'manufacturer'
                    )
                  "
                  v-model="tool.manufacturer"
                  placeholder="* Enter Manufacturer"
                  @blur="
                    () => {
                      dirty.manufacturer = true;
                    }
                  "
                />
                <label>Tool Description:</label>
                <textarea
                  required
                  :class="validDescriptionStyle(tool.description)"
                  class="tool-description-input"
                  v-model="tool.description"
                  placeholder="* Enter Tool Description"
                  @blur="
                    () => {
                      dirty.description = true;
                    }
                  "
                ></textarea>
                <label>Upload image:</label>
                <div :class="this.validImageStyle()">
                  <input
                    type="file"
                    :multiple="false"
                    :name="uploadFieldName"
                    :disabled="isProcessing"
                    @change="onFileChange"
                    accept="image/*"
                    class="input-file"
                  />
                  <p v-if="!isProcessing && tool.file.name === undefined">
                    Drag your file here or click to browse
                  </p>
                  <span v-if="!isProcessing && tool.file.name !== undefined">
                    <p>
                      File: {{ tool.file.name }}
                      <v-icon color="green">mdi-check</v-icon>
                    </p>
                  </span>
                  <p v-if="isProcessing">Processing file...</p>
                </div>
                <div class="container">
                  <button
                    :disabled="!isEnabled()"
                    type="button"
                    class="btn btn-sm btn-outline-primary tool-update-btn"
                    @click="updateTool(tool)"
                  >
                    Update
                  </button>
                </div>
                <br />
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
import _ from "lodash";

Vue.use(Toasted, {
  duration: 4000
});

export default {
  name: "UpdateTool",
  data() {
    return {
      uploadFieldName: "toolImage",
      imageInputClass: "dropbox",
      isProcessing: false,
      dirty: {
        name: false,
        model: false,
        manufacturer: false,
        description: false,
        file: false
      },
      tool: { ...this.$route.params.tool },
      originalTool: { ...this.$route.params.tool }
    };
  },
  created() {
    if (this.tool === undefined || this.tool === null) {
      this.$router.push({ name: "view-tools", location: "tools" });
    }
    this.tool.file = {};
    this.originalTool.file = {};
  },
  methods: {
    ...mapActions(["updateToolInStore"]),
    updateTool: function(tool) {
      let toolForm = new FormData();
      toolForm.append("id", tool.id);
      toolForm.append("name", tool.name);
      toolForm.append("description", tool.description);
      toolForm.append("model", tool.model);
      toolForm.append("manufacturer", tool.manufacturer);
      toolForm.append("file", tool.file);
      this.$store
        .dispatch("updateToolInStore", toolForm)
        .then(() => {
          this.$toasted.show(tool.name + " has been updated.", {
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
    },
    isEnabled: function() {
      return (
        this.validNameModelManufacturer(this.tool.name) &&
        this.validNameModelManufacturer(this.tool.model) &&
        this.validNameModelManufacturer(this.tool.manufacturer) &&
        this.validDescription(this.tool.description) &&
        !_.isEqual(this.originalTool, this.tool)
      );
    },
    validNameModelManufacturerStyle: function(input, prop) {
      return [
        this.validNameModelManufacturer(input) || !this.dirty[prop]
          ? ""
          : "input-error"
      ];
    },
    validNameModelManufacturer(str) {
      let letters = /^[()A-Za-z0-9 '.+-/]+$/;
      return (
        str.length >= 1 &&
        str.match(letters) !== undefined &&
        str.match(letters) !== null &&
        str.length === str.trim().length
      );
    },
    validDescription: function(description) {
      return (
        description.length > 0 &&
        description.length === description.trim().length
      );
    },
    validDescriptionStyle: function(input) {
      return [
        this.validDescription(input) || !this.dirty.description
          ? ""
          : "input-error"
      ];
    },
    validImage: function() {
      return this.tool.file.size !== undefined && this.tool.file.size <= 250000;
    },
    validImageStyle: function() {
      return [
        !this.dirty.file || this.validImage() ? "dropbox" : "file-error dropbox"
      ];
    },
    onFileChange(e) {
      this.dirty.file = true;
      this.isProcessing = true;
      let file = e.target.files[0];

      if (file.size > 250000) {
        this.imageInputClass = "file-error dropbox";
        this.$toasted.show("File size must under .25 MB!", {
          type: "danger",
          theme: "bubble"
        });
        this.isProcessing = false;
        return;
      }
      this.tool.file = file;
      this.isProcessing = false;
    }
  }
};
</script>

<style></style>
