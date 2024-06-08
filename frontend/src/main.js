import { defaultConfig, plugin } from "@formkit/vue";
import { createApp, h, provide } from "vue";
import App from "./App.vue";
import "./style.css";
// import config from "../formkit.config.js";
import {
  ApolloClient,
  createHttpLink,
  InMemoryCache,
} from "@apollo/client/core";
import { DefaultApolloClient } from "@vue/apollo-composable";
import config from "../formkit.config.js";
import router from "./route";

let token = localStorage.getItem("token");
let headers;
if (token) {
  headers = {
    Authorization: `Bearer ${token}`,
  };
} else {
  headers = {
    "x-hasura-role": "anonymous",
  };
}

const httpLink = createHttpLink({
  uri: "http://localhost:8080/v1/graphql",
  headers,
});

// Cache implementation
const cache = new InMemoryCache();

// Create the apollo client
const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});

createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient);
  },

  render: () => h(App),
})
  .use(router, plugin, defaultConfig(config))
  .mount("#app");
// createApp(App).use(plugin, defaultConfig(config)).mount("#app");
