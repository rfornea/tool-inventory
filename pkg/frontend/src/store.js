import Vue from "vue";
import Vuex from "vuex";
import _ from "lodash";

import APIClient from "./apiClient";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    tools: []
  },
  mutations: {
    resetTools(state, tools) {
      state.tools = tools;
    }
  },
  getters: {
    tools(state) {
      return state.tools;
    }
  },
  actions: {
    getAllToolsStore({ commit }) {
      APIClient.getTools().then(data => {
        commit("resetTools", data.tools);
      });
    },
    updateToolInStore({ commit, state }, tool) {
      return APIClient.updateTool(tool).then(tool => {
        const tools = [...state.tools];
        tools[
          _.findIndex(tools, function(tl) {
            return tl.id === tool.id;
          })
        ] = tool;
        commit("resetTools", tools);
      });
    },
    removeTool({ commit, state }, tool) {
      const tools = [...state.tools];
      _.remove(tools, function(tl) {
        return tl.id === tool.id;
      });
      return APIClient.deleteTool(tool).then(() => {
        commit("resetTools", tools);
      });
    },
    addTool({ commit, state }, tool) {
      const tools = [...state.tools];
      return APIClient.createTool(tool).then(tool => {
        tools[tools.length] = tool;
        commit("resetTools", tools);
      });
    }
  }
});

export default store;
