<script setup>
import { cn } from "@/lib/utils";
import MaxWidth from "./MaxWidth.vue";
import { Button, buttonVariants } from "@/components/ui/button";
const categories = [
  {
    name: "fasting",
    link: "/home/fasting",
  },
  {
    name: "traditional",
    link: "/home/traditional",
  },
  {
    name: "meat",
    link: "/home/meat",
  },
  {
    name: "fruit",
    link: "/home/fruit",
  },
  {
    name: "desserts",
    link: "/home/desserts",
  },
  {
    name: "appetizers",
    link: "/home/appetizers",
  },
];
</script>
<template>
  <MaxWidth>
    <div class="flex items-center justify-center">
      <div class="flex gap-2 flex-row flex-wrap">
        <!-- <RouterLink :class="cn(buttonVariants({ variant: 'outline' }))" to="/">All</RouterLink> -->
        <div v-for="c in categories" :key="c">
          <!-- <Button variant="outline" aschild> -->
          <RouterLink :class="cn(buttonVariants({ variant: 'outline' }))" :to="c.link">{{ c.name }}</RouterLink>
          <!-- </Button> -->
        </div>

      </div>
    </div>
  </MaxWidth>
</template>
<style scoped>
a.router-link-exact-active {
  background-color: hsl(var(--primary));
  color: hsl(var(--primary-foreground));
}
</style>
