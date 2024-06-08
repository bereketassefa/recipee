<script setup>
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
import { FieldArray } from "vee-validate";
import { ref } from "vue";

import FormDescription from "@/components/ui/form/FormDescription.vue";
import Slider from "@/components/ui/slider/Slider.vue";
import { Textarea } from "@/components/ui/textarea";

import { cn } from "@/lib/utils";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { z } from "zod";
import { useMutation } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { toast } from "vue-sonner";
import Label from "@/components/ui/label/Label.vue";
const formSchema = toTypedSchema(
  z.object({
    name: z.string().min(2).max(50),
    description: z.string().min(2),
    duration: z.array(z.number().min(0).max(100), {
      message: "List of ingredients are required",
    }),
    ingredient: z.array(z.string()),
    tag: z.string(),
    urls: z.array(
      z.object({
        value: z.string(),
      })
    ),
  })
);

const {
  mutate: CreateRecipe,
  onDone,
  error,
  onError,
} = useMutation(gql`
  mutation CreateRecipe(
    $description: String!
    $name: String!
    $duration: Int!
    $userid: uuid!
    $tag: uuid!
    $ingredients: [Ingredients_insert_input!]!
    $steps: [recipeSteps_insert_input!]!
    $img: String!
  ) {
    insert_Recipe_one(
      object: {
        description: $description
        duration: $duration
        name: $name
        userId: $userid
        RecipeTags: { data: { tagId: $tag } }
        Ingredients: { data: $ingredients }
        recipeSteps: { data: $steps }
        recipeImages: { data: { imageLink: $img } }
      }
    ) {
      name
    }
  }
`);
onDone((handle) => {
  toast("Recipe created succesfully");
});
onError((message) => {
  toast("Form failed");
});
const { handleSubmit } = useForm({
  validationSchema: formSchema,
  initialValues: {
    urls: [{ value: "" }],
  },
});
const onSubmit = handleSubmit((data) => {
  const ing = data.ingredient.map((each) => ({
    name: each,
  }));
  const st = data.urls.map((each, index) => ({
    order: index,
    description: each.value,
  }));
  console.log(ing);
  console.log(st);
  CreateRecipe({
    name: data.name,
    description: data.description,
    duration: data.duration[0],
    userid: localStorage.getItem("id"),
    tag: data.tag,
    ingredients: ing,
    steps: st,
    img: imageLink.value,
  });
});
const categories = [
  {
    id: "ecd8a10d-62ec-439f-86e7-1eaf23d86f6e",
    tagName: "chechebsa",
  },
  {
    id: "ef1e6fb6-6699-4f30-a4a6-772906feb9df",
    tagName: "gerami",
  },
  {
    id: "df107139-54f1-4436-97f6-d282f332e02d",
    tagName: "fasting",
  },
  {
    id: "76c530d3-4ae1-4096-9f42-6d844b8ed304",
    tagName: "traditional",
  },
  {
    id: "c3b6e62b-1f35-498e-b866-074328bba2bf",
    tagName: "meat",
  },
  {
    id: "75d0804a-449b-44ee-a9c2-ffa46fc37c7d",
    tagName: "fruit",
  },
  {
    id: "ec906b54-4ad2-43db-988f-7048b63b4374",
    tagName: "desserts",
  },
  {
    id: "b83be18c-e731-41c3-b250-a2a970c013ec",
    tagName: "appetizers",
  },
];
const imageLink = ref("");
const modelValue = ref(["banana"]);
const previewFiles = (event) => {
  console.log(event.target.files);
};
const handleFileChange = (event) => {
  const file = event.target.files[0];
  error.value = null; // Clear previous errors

  if (!file) {
    return;
  }

  const formData = new FormData();
  formData.append("image", file);

  const uploadUrl = "http://localhost:3000/upload"; // Replace with your server endpoint

  fetch(uploadUrl, {
    method: "POST",
    body: formData,
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Upload failed");
      }
      return response.json();
    })
    .then((data) => {
      console.log("Upload successful:", data);
      imageLink.value = data.url;
      console.log(imageLink.value);
      // Handle successful upload response (data)
      emit("fileUploaded", file); // Emit an event with the uploaded file
    })
    .catch((error) => {
      console.error("Upload error:", error);
      error.value = "Upload failed. Please try again.";
    });
};
</script>
<template>
  <Header></Header>
  <MaxWidth>
    <div>
      <h1 class="font-semibold text-3xl">Create a new recipe</h1>
      <form class="space-y-6" @submit="onSubmit">
        <FormField v-slot="{ componentField }" name="name">
          <FormItem>
            <FormLabel>Recipe Name</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Name" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="description">
          <FormItem>
            <FormLabel>Description</FormLabel>
            <FormControl>
              <Textarea placeholder="Tell use about your recipe" class="resize-none" v-bind="componentField" />
            </FormControl>

            <FormMessage />
          </FormItem>
        </FormField>
        {{ imageLink }}
        <div class="grid w-full max-w-sm items-center gap-1.5">
          <Label for="picture">Picture</Label>
          <Input @change="handleFileChange" id="picture" type="file" />
        </div>

        <FieldArray v-slot="{ fields, push, remove }" name="urls">
          <div v-for="(field, index) in fields" :key="`urls-${field.key}`">
            <FormField v-slot="{ componentField }" :name="`urls[${index}].value`">
              <FormItem>
                <FormLabel :class="cn(index !== 0 && 'sr-only')">
                  Steps
                </FormLabel>
                <FormDescription :class="cn(index !== 0 && 'sr-only')">
                  Add steps for the recipe
                </FormDescription>
                <div class="relative flex items-center">
                  <FormControl>
                    <Input type="text" v-bind="componentField" />
                  </FormControl>
                  <button type="button" class="absolute py-2 pe-3 end-0 text-muted-foreground" @click="remove(index)">
                    <Cross1Icon class="w-3" />
                  </button>
                </div>
                <FormMessage />
              </FormItem>
            </FormField>
          </div>

          <Button type="button" variant="outline" size="sm" class="text-xs w-20 mt-2" @click="push({ value: '' })">
            Add Steps
          </Button>
          <Button type="button" variant="outline" size="sm" class="text-xs w-20 mt-2" @click="remove()">
            Remove
          </Button>
        </FieldArray>

        <FormField v-slot="{ componentField, value }" name="duration">
          <FormItem>
            <FormLabel>Duration</FormLabel>
            <FormControl>
              <Slider v-bind="componentField" :default-value="[30]" :max="100" :min="0" :step="5" />
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
                    <SelectItem v-for="val in categories" :value="val.id">
                      {{ val.tagName }}
                    </SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </FormControl>

            <FormMessage />
          </FormItem>
        </FormField>
        <Button type="submit"> Submit </Button>
      </form>
    </div>
  </MaxWidth>
</template>
<style></style>
