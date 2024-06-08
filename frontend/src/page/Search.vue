<script setup lang="ts">
import Header from "@/components/Header.vue";
import MaxWidth from "@/components/MaxWidth.vue";
import Button from "@/components/ui/button/Button.vue";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import Input from "@/components/ui/input/Input.vue";
import {
  TagsInput,
  TagsInputInput,
  TagsInputItem,
  TagsInputItemDelete,
  TagsInputItemText,
} from "@/components/ui/tags-input";
import Loading from "@/components/Loading.vue";
import { reactive, ref } from "vue";
import { useMutation, useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { useForm } from "vee-validate";
import ReceipeCard from "@/components/ReceipeCard.vue";
import Notright from "@/components/Notright.vue";
import Slider from "@/components/ui/slider/Slider.vue";
import FormDescription from "@/components/ui/form/FormDescription.vue";
const categories = [
  {
    name: "fasting",
  },
  {
    name: "traditional",
  },
  {
    name: "meat",
  },
  {
    name: "fruit",
  },
  {
    name: "desserts",
  },
  {
    name: "appetizers",
  },
];
const Search = ref("");
let val = reactive({
  firstName: "%%",
  lastName: "%%",
  name: "",
  duration: 100,
  tag: "%%",
});
const { result, loading, error, variables } = useQuery(
  gql`
    query MyQuery(
      $firstName: String
      $lastName: String
      $name: String
      $duration: Int
      $tag: String
    ) {
      Recipe(
        where: {
          _or: {
            user: {
              firstName: { _ilike: $firstName }
              lastName: { _ilike: $lastName }
            }
            name: { _ilike: $name }
            duration: { _lt: $duration }
            RecipeTags: { Tag: { tagName: { _ilike: $tag } } }
          }
        }
      ) {
        id
        name
        duration
        recipeImages {
          imageLink
        }
        user {
          firstName
          lastName
          profileImg
        }
        RecipeTags {
          Tag {
            tagName
          }
        }
      }
    }
  `,
  val
);

const modelValue = ref([]);

const { handleSubmit } = useForm();
const onSubmit = handleSubmit((data) => {
  for (const [key, value] of Object.entries(data)) {
    if (value == undefined || value == "") {
      continue;
    }

    val[key] = `%${value}%`;

    // console.log(key, value);
  }
  if (data.duration) {
    val.duration = data.duration[0];
  } else {
    val.duration = 100;
  }
  console.log(val);
  variables.value = val;
});
const handleSearch = (newval) => {
  console.log(newval);
  if (newval == "") {
    val.name = "";
  } else {
    val.name = `%${newval}%`;
  }
  variables.value = val;
};
</script>

<template>
  <Header />
  <MaxWidth>

    <div class="relative max-w-[400px] mx-auto mb-7">
      <label for="Search" class="sr-only"> Search </label>

      <Input type="text" id="Search" v-model="Search" @input="(event) => handleSearch(event.target.value)"
        placeholder="Search for..." />

      <span class="absolute inset-y-0 end-0 grid w-10 place-content-center">
        <button type="button" class="text-gray-600 hover:text-gray-700">
          <span class="sr-only">Search</span>

          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="h-4 w-4">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
          </svg>
        </button>
      </span>
    </div>
    <div class="grid grid-cols-9">
      <div class="w-full col-span-3 px-5">
        <form class="space-y-6" @submit="onSubmit">
          <FormField v-slot="{ componentField }" name="firstName">
            <FormItem>
              <FormLabel>First Name</FormLabel>
              <FormControl>
                <Input type="text" placeholder="First Name" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="lastName">
            <FormItem>
              <FormLabel>Last Name</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Last Name" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField, value }" name="duration">
            <FormItem>
              <FormLabel>Duration</FormLabel>
              <FormControl>
                <Slider v-bind="componentField" :default-value="[100]" :max="100" :min="0" :step="5" />
                <FormDescription class="flex justify-between">
                  <span>{{ value?.[0] }} min</span>
                </FormDescription>
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="ingredient">
            <FormItem>
              <FormLabel>Ingredients</FormLabel>
              <FormControl>
                <TagsInput v-bind="componentField" v-model="modelValue">
                  <TagsInputItem v-for="item in modelValue" :key="item" :value="item">
                    <TagsInputItemText />
                    <TagsInputItemDelete />
                  </TagsInputItem>

                  <TagsInputInput placeholder="Fruits..." />
                </TagsInput>
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="tag">
            <FormItem>
              <FormLabel>Category</FormLabel>
              <FormControl>
                <Select v-bind="componentField">
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="Select a valid category" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectGroup>
                      <SelectItem v-for="val in categories" :value="val.name">
                        {{ val.name }}
                      </SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
          <Button type="submit"> Search </Button>
        </form>
      </div>
      <div class="w-full h-[100px] col-span-6">
        <div v-if="error">Error loading recipe</div>
        <Loading v-else-if="loading"></Loading>
        <div v-else-if="result && result.Recipe"
          class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4">
          <!-- {{ result }} -->
          <ReceipeCard v-for="x in result.Recipe" :id="x.id" :name="x.name" :author="x.user" :duration="x.duration"
            :tags="x.RecipeTags" :image="x.recipeImages" />
        </div>
        <div v-if="Search == ''"></div>
        <Notright v-else-if="result && result.Recipe.length == 0">
          There is no recipe available
        </Notright>
      </div>
    </div>
  </MaxWidth>
</template>
