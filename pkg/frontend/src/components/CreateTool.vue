<template>
  <v-app id="inspire">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="#3d5496">
                <v-toolbar-title justify-center>Add New Tool</v-toolbar-title>
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
                    id="tool-image-file"
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
                    type="button"
                    class="btn btn-sm btn-outline-info tool-clear-btn"
                    @click="clearTool()"
                  >
                    Clear
                  </button>
                  <button
                    :disabled="!isEnabled()"
                    type="button"
                    class="btn btn-sm btn-outline-primary tool-update-btn"
                    @click="createTool(tool)"
                  >
                    Create
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
  name: "CreateTool",
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
      tool: {
        name: "",
        model: "",
        manufacturer: "",
        description: "",
        file: {}
      }
    };
  },
  created() {
    this.tool.file = {};
  },
  computed: {},
  methods: {
    ...mapActions(["addTool"]),
    createTool: function(tool) {
      let toolForm = new FormData();
      toolForm.append("name", tool.name);
      toolForm.append("description", tool.description);
      toolForm.append("model", tool.model);
      toolForm.append("manufacturer", tool.manufacturer);
      toolForm.append("file", tool.file);
      this.$store
        .dispatch("addTool", toolForm)
        .then(() => {
          this.$toasted.show(tool.name + " has been created.", {
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
    clearTool: function() {
      this.tool = {
        name: "",
        model: "",
        manufacturer: "",
        description: "",
        file: {}
      };
      this.dirty = {
        name: false,
        model: false,
        manufacturer: false,
        description: false,
        file: false
      };
      this.isProcessing = false;
      document.getElementById("tool-image-file").valueOf().value = null;
    },
    isEnabled: function() {
      return (
        this.validNameModelManufacturer(this.tool.name) &&
        this.validNameModelManufacturer(this.tool.model) &&
        this.validNameModelManufacturer(this.tool.manufacturer) &&
        this.validDescription(this.tool.description) &&
        !this.isProcessing
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

<style>
label {
  margin-left: 5px;
  font-size: 1.25em;
}
textarea {
  border: 1px solid #cac6c6;
  padding: 1px;
  margin: 0 0 5px 5px;
  width: 100%;
}
input {
  border: 1px solid #cac6c6;
  padding: 1px;
  margin: 0 0 5px 5px;
  width: 100%;
}
/*this class needed for error styles, do not delete*/
.input-error {
  border: 1px solid red;
}
.tool-clear-btn {
  float: right;
}
.tool-update-btn {
  float: right;
  margin-right: 5px;
}
.tool-description-input {
  height: 100px;
}
.dropbox {
  outline: 2px dashed grey; /* the dash box */
  outline-offset: -10px;
  background: lightcyan;
  color: dimgray;
  padding: 10px 10px;
  min-height: 100px !important; /* minimum height */
  max-height: 100px !important; /* minimum height */
  position: relative;
  cursor: pointer;
}

.input-file {
  opacity: 0; /* invisible but it's there! */
  width: 100%;
  height: 100px !important;
  position: absolute;
  cursor: pointer;
}
.file-error {
  border: 2px dashed red !important;
}

.dropbox:hover {
  background: lightblue; /* when mouse over to the drop zone, change color */
}

.dropbox p {
  font-size: 1.2em;
  text-align: center;
  padding: 25px 0;
}
</style>
