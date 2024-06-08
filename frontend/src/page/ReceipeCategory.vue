<script setup>
import { useRoute } from "vue-router";
import Header from "../components/Header.vue";
import Loading from "../components/Loading.vue";
import { useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { watch } from "vue";
import ReceipeCard from "@/components/ReceipeCard.vue";
import MaxWidth from "@/components/MaxWidth.vue";
import Categories from "@/components/Categories.vue";
import Notright from "@/components/Notright.vue";
const route = useRoute();
const { result, loading, error, variables } = useQuery(
  gql`
    query MyQuery($id: String!) {
      Recipe(where: { RecipeTags: { Tag: { tagName: { _eq: $id } } } }) {
        description
        duration
        id
        name
        user {
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
      }
    }
  `,
  { id: route.params.id },
  { fetchPolicy: "cache-and-network" }
);
watch(
  () => route.params.id,
  (newId, oldId) => {
    variables.value = {
      id: newId,
    };
  }
);
</script>
<template>
  <Header></Header>
  <Categories></Categories>
  <MaxWidth>
    <Loading v-if="loading"></Loading>

    <Notright v-else-if="error">There is error loading this page</Notright>

    <div
      v-else-if="result && result.Recipe.length"
      class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4"
    >
      <!-- {{ JSON.stringify(result.Recipe) }} -->
      <ReceipeCard
        v-for="x in result.Recipe"
        :id="x.id"
        :name="x.name"
        :author="x.user"
        :duration="x.duration"
        :tags="x.RecipeTags"
        :image="x.recipeImages"
      />
    </div>

    <Notright v-else-if="!result.Recipe.length"
      >We cannot find any recipe in this category.</Notright
    >
  </MaxWidth>
</template>
<style></style>
