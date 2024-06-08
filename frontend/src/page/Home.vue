<script setup>
import Header from "../components/Header.vue";
import Categories from "../components/Categories.vue";
import ReceipeCard from "@/components/ReceipeCard.vue";
import MaxWidth from "@/components/MaxWidth.vue";
import Loading from "@/components/Loading.vue";
import { useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { watch, computed, onMounted, ref } from "vue";
import Hero from "@/components/Hero.vue";
import { useRoute, useRouter } from "vue-router";
import Notright from "@/components/Notright.vue";
const route = useRoute();
let vals = ref(route.params.id);


const { result, loading, error, variables } = useQuery(
  gql`
    query MyQuery($id: String!) {
      Recipe(
        where: { RecipeTags: { Tag: { tagName: { _eq: $id } } } }
        limit: 8
      ) {
        description
        duration
        id
        name
        user {
        id
          lastName
          firstName
          email
        }
        RecipeTags {
          Tag {
            tagName
          }
        }
        recipeImages {
          imageLink
        }
        Comments_aggregate {
          aggregate {
            avg {
              rating
            }
          }
        }
      }
      Second: Recipe(
        where: { Comments: { rating: { _gt: 0 } } }
        limit: 8
        order_by: { userLikes_aggregate: { count: asc } }
      ) {
        description
        duration
        id

        name
        user {
        id
          lastName
          firstName
          email
        }
        RecipeTags {
          Tag {
            tagName
          }
        }
        recipeImages {
          imageLink
        }
        Comments_aggregate {
          aggregate {
            avg {
              rating
            }
          }
        }
      }
    }
  `,
  { id: route.params.id },
  { fetchPolicy: "cache-and-network" }
);
watch(
  () => route.params.id,
  (newId, oldId) => {
    console.log("changed", newId);
    variables.value = {
      id: newId,
    };
  }
);
const recipes = computed(() => result.value?.Recipe ?? []);
</script>

<template>
  <Header></Header>
  <MaxWidth>
    <Hero />
    <section class="">
      <div class="mx-auto max-w-screen-xl px-4 py-8 sm:px-6 sm:py-12 lg:px-8 lg:py-16">
        <div class="mx-auto max-w-lg text-center">
          <h2 class="text-3xl font-bold sm:text-4xl">
            Start with the most liked
          </h2>

          <p class="mt-4 text-muted-foreground">
            Rediscover your favorites with our curated list of recipes you've
            liked. Dive back into deliciousness and find your next culinary
            adventure!
          </p>
        </div>
        <Loading v-if="loading"></Loading>

        <Notright v-else-if="error">There is error loading this content</Notright>

        <div v-else-if="result && result.Second.length"
          class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4">
          <!-- {{ JSON.stringify(result.Recipe) }} -->
          <ReceipeCard v-for="x in result.Second" :id="x.id" :name="x.name" :author="x.user" :duration="x.duration"
            :tags="x.RecipeTags" :image="x.recipeImages" :all="x" />
        </div>

        <Notright v-else-if="!result.Recipe.length">We cannot find the top rated recipes.</Notright>
      </div>
    </section>

    <div class="mx-auto max-w-lg text-center mt-5 mb-8">
      <h2 class="text-3xl font-bold sm:text-4xl">
        Browse recipe by Categories
      </h2>

      <p class="mt-4 text-muted-foreground">
        Discover delicious recipes categorized by cuisine, course, dietary
        needs, and more. Find your perfect dish for any occasion!
      </p>
    </div>
    <div class="my-9">
      <Categories />
    </div>
    <Loading v-if="loading"></Loading>

    <Notright v-else-if="error">There is error loading this page</Notright>

    <div v-else-if="result && result.Recipe.length"
      class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4">
      <!-- {{ JSON.stringify(result.Recipe) }} -->
      <ReceipeCard v-for="x in result.Recipe" :id="x.id" :name="x.name" :author="x.user" :duration="x.duration"
        :tags="x.RecipeTags" :image="x.recipeImages" :all="x" />
    </div>

    <Notright v-else-if="!result.Recipe.length">We cannot find any recipe in this category.</Notright>
  </MaxWidth>
</template>
