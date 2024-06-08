import {
  createMemoryHistory,
  createRouter,
  createWebHistory,
} from "vue-router";
import First from "../components/First.vue";
import Home from "../page/Home.vue";
import Login from "../page/Login.vue";

import Signup from "../page/Signup.vue";
import Receipe from "../page/Receipe.vue";
import NotFound from "../page/NotFound.vue";
import CategorieDetail from "@/page/CategorieDetail.vue";
import ReceipeCategory from "@/page/ReceipeCategory.vue";
import NewRecipe from "@/page/NewRecipe.vue";
import Profile from "@/page/Profile.vue";
import Search from "@/page/Search.vue";

const routes = [
  { path: "/", redirect: "/home/fasting" },
  { path: "/home/:id?", component: Home, name: "Home" },
  { path: "/categories/:name", component: CategorieDetail, name: "category" },
  { path: "/login", component: Login, name: "Login" },
  { path: "/signup", component: Signup, name: "Signup" },
  { path: "/receipe/:id", component: Receipe, name: "Recipe", props: true },
  { path: "/new", component: NewRecipe, name: "newrecipe", props: true },
  { path: "/profile/:id", component: Profile, name: "profile", props: true },
  { path: "/search", component: Search, name: "search", props: true },
  {
    path: "/recipe/tag/:id",
    component: ReceipeCategory,
    name: "Recipetag",
    props: true,
  },
  {
    path: "/:pathMatch(.*)*",
    component: NotFound,

    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
