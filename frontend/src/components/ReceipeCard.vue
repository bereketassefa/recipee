<script setup lang="ts">
import { Card } from "@/components/ui/card";
import { cn } from "@/lib/utils";
import { Clock9, Ellipsis, Star } from "lucide-vue-next";
import Badge from "./ui/badge/Badge.vue";
import { buttonVariants } from "./ui/button";
import Button from "./ui/button/Button.vue";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useMutation } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { toast } from "vue-sonner";
const props = defineProps([
  "name",
  "duration",
  "tags",
  "author",
  "id",
  "image",
  "profile",
  "all",
]);
const isLogged = () => {
  let result;
  if (
    localStorage.getItem("id") &&
    localStorage.getItem("id") == props.author.id
  ) {
    result = true;
  } else {
    result = false;
  }

  return result;
};
const lt = () => {
  return localStorage.getItem("id");
};
const {
  mutate: RemoveLike,
  onDone,
  onError,
} = useMutation(
  gql`
    mutation Log($id: uuid!) {
      delete_userLikes(where: { id: { _eq: $id } }) {
        returning {
          id
        }
      }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);
onDone((handle) => {
  toast("Succesfully removed recipe from your liaked list");
});
onError((message) => {
  toast("Error removing liked recipe");
});
const {
  mutate: RemoveFavorite,
  onDone: favdone,
  onError: faverr,
} = useMutation(
  gql`
    mutation Log($id: uuid!) {
      delete_Favorite(where: { id: { _eq: $id } }) {
        returning {
          id
        }
      }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);
favdone((handle) => {
  toast("Succesfully removed recipe from your favorite list");
});
faverr((message) => {
  toast("Error removing favorite recipe");
});
const {
  mutate: RemoveRecipe,
  onDone: recdone,
  onError: recerr,
} = useMutation(
  gql`
    mutation Log($id: uuid!) {
      delete_Recipe(where: {id: {_eq: $id}}) {
    affected_rows
  }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);
recdone((handle) => {
  toast("Succesfully removed recipe from your favorite list");
});
recerr((message) => {
  toast("Error removing favorite recipe");
});
</script>

<template>
  <RouterLink :to="`/receipe/${props.id}`">
    <Card class="group">
      <img alt="" :src="props.image[0]?.imageLink"
        class="h-56 w-full rounded-t-xl rounded-b-[0] object-cover shadow-xl transition group-hover:grayscale-[50%]" />
      <!-- <FoodSlider v-if="props.image.length" :data="props.image" /> -->

      <div class="p-4">
        <div class="flex flex-row justify-between">
          <a href="#">
            <h3 class="text-lg font-medium">
              {{ props.name }}
            </h3>
          </a>
          <DropdownMenu v-if="
            (props.profile == 'favorite' ||
              props.profile == 'profile' ||
              props.profile == 'liked') && isLogged()
          ">
            <DropdownMenuTrigger as-child>
              <Button variant="ghost">
                <Ellipsis />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent class="w-56">
              <!-- <DropdownMenuLabel>Actions</DropdownMenuLabel> -->

              <DropdownMenuItem v-if="props.profile == 'profile'">
                <span>Update</span>
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem v-if="props.profile == 'profile'" @click="RemoveRecipe({ id: props.id })">
                <span>Delete</span>
              </DropdownMenuItem>

              <DropdownMenuItem @click="RemoveFavorite({ id: props.all.id })" v-if="props.profile == 'favorite'">
                <span>remove</span>
              </DropdownMenuItem>

              <DropdownMenuItem @click="RemoveLike({ id: props.all.id })" v-if="props.profile == 'liked'">
                <span>remove</span>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>

        <RouterLink :to="'/profile/' + props.author.id" :class="cn(buttonVariants({ variant: 'link' }), 'px-0')">
          {{ props.author.firstName }} {{ props.author.lastName }}
        </RouterLink>
        <div class="mt-2 line-cla mp-3 text-sm/relaxed text-gray-500 flex justify-between">
          <div class="flex flex-row gap-1">
            <Badge v-for="x in props.tags" v-if="props.tags.length">{{
              x.Tag.tagName
            }}</Badge>
          </div>
          <div class="flex justify-between items-center gap-1">

            <div class="flex justify-between items-center">
              <span class="inline-block pr-1">
                <Star fill="gold" stroke-width="0" :size="15" />
              </span>
              <p>{{ Math.floor(props.all?.Comments_aggregate?.aggregate.avg.rating) }}</p>
              <!-- {{ props.all?.Comments_aggregate }} -->
            </div>
            <div class="flex justify-between items-center">
              <span class="inline-block pr-1">
                <Clock9 :size="15" />
              </span>
              <p>{{ props.duration }}</p>
            </div>
          </div>
        </div>
      </div>
    </Card>
  </RouterLink>
</template>
