<script setup>
import { RouterView } from "vue-router";
import Login from "./page/Login.vue";
import { Toaster } from "@/components/ui/sonner";
import Footer from '@/components/Footer.vue'
</script>

<template>
  <!-- <div class="flex gap-2">
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/signup">Signup</RouterLink>
    <RouterLink to="/login">login</RouterLink>
    <RouterLink to="/receipe/3337a965-4ca7-4e7e-9ee2-7b907f7e3351"
      >Bereket</RouterLink
    >
    <RouterLink to="/profile/3337a965-4ca7-4e7e-9ee2-7b907f7e3351"
      >second</RouterLink
    >
  </div>
  <br /> -->
  <RouterView />
  <Footer />
  <Toaster />

  <!-- <Login></Login> -->
</template>

<style scoped>
a.router-link-exact-active {
  background-color: red;
}
</style>
