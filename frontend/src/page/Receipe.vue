<script setup>
import { watch, computed, ref } from "vue";
import { useMutation, useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { useRoute } from "vue-router";
import MaxWidth from "@/components/MaxWidth.vue";
import Loading from "@/components/Loading.vue";
import FoodSlider from "@/components/FoodSlider.vue";
import { Bookmark, Clock9, Dot, Star, ThumbsUp } from "lucide-vue-next";
import Badge from "@/components/ui/badge/Badge.vue";
import { marked } from "marked";
import Header from "@/components/Header.vue";
import Textarea from "@/components/ui/textarea/Textarea.vue";
import Button from "@/components/ui/button/Button.vue";
import { toTypedSchema } from "@vee-validate/zod";
import { z } from "zod";
import { FormField } from "@/components/ui/form";
import FormItem from "@/components/ui/form/FormItem.vue";
import FormControl from "@/components/ui/form/FormControl.vue";
import FormMessage from "@/components/ui/form/FormMessage.vue";
import { useForm } from "vee-validate";
import { toast } from "vue-sonner";
const route = useRoute();

const getMarked = (val) => {
  return marked.parse(val);
  // return Marked.parse(result.value.Recipe_by_pk.description);
};
const { result, loading, error, variables } = useQuery(
  gql`
    query MyQuery($id: uuid!, $userId: uuid) {
      Recipe_by_pk(id: $id) {
        description
        duration
        id
        name
        Comments_aggregate {
          aggregate {
            avg {
              rating
            }
          }
        }
        userLikes(where: { userId: { _eq: $userId } }) {
          userId
        }
        Favorites(where: { userId: { _eq: $userId } }) {
          userId
        }
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
        Comments(limit: 7) {
          comment
          rating
          user {
            firstName
            lastName
            profileImg
            id
          }
        }
        recipeSteps {
          order
          description
        }
      }
    }
  `,
  { id: route.params.id, userId: localStorage.getItem("id") || "a1d73fe6-3322-4771-bb7f-87956ae64047" },
  {
    fetchPolicy: "cache-and-network",
  }
);
watch(
  () => route.params.id,
  (newId, oldId) => {
    variables.value = {
      id: newId,
    };
  }
);

const {
  mutate: AddFavorite,
  onDone: favoriteSuccess,
  onError: favoriteError,
} = useMutation(
  gql`
    mutation newFavorite($receipeId: uuid!) {
      insert_Favorite(objects: { receipeId: $receipeId }) {
        affected_rows
      }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);

favoriteSuccess((handle) => {
  toast("Succesfully add recipe to favorite");
});
favoriteError((message) => {
  toast("Error adding recipe to favorite");
});
const {
  mutate: AddLike,
  onDone: LikeSuccess,
  onError: LikeError,
} = useMutation(
  gql`
    mutation newLike($receipeId: uuid!, $userId: uuid!) {
      insert_userLikes_one(object: { recipeId: $receipeId, userId: $userId }) {
        userId
      }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);

LikeSuccess((handle) => {
  toast("Succesfully add recipe to Like");
});
LikeError((message) => {
  toast("Error adding recipe to Like");
});

const {
  mutate: AddComment,
  onDone,
  onError,
} = useMutation(
  gql`
    mutation Addcomment(
      $comment: String!
      $rating: Int!
      $userId: uuid!
      $recipeId: uuid!
    ) {
      insert_Comment(
        objects: {
          comment: $comment
          rating: $rating
          userId: $userId
          recipeId: $recipeId
        }
      ) {
        affected_rows
      }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);

onDone((handle) => {
  toast("Comment added succesfully");
});
onError((message) => {
  toast("Request failed");
});

let rating = ref(0);
const formSchema = toTypedSchema(
  z.object({
    comment: z.string().min(2).max(200),
  })
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((data) => {
  if (rating.value == 0) {
    toast("please enter a valid rating");
  } else {
    AddComment({
      comment: data.comment,
      rating: rating.value,
      userId: localStorage.getItem("id"),
      recipeId: route.params.id,
    });
  }
  console.log(data);
});
const addToFavorite = () => {
  AddFavorite({
    receipeId: route.params.id,
  });
};
const addToLike = () => {
  AddLike({
    receipeId: route.params.id,
    userId: localStorage.getItem("id"),
  });
};
const RemoveFromLike = () => {
  RemoveLike({
    id: route.params.id,
  });
};
const RemoveFromFavorite = () => {
  RemoveFavorite({
    id: route.params.id,
  });
};
const isLogged = () => {
  let result;
  if (localStorage.getItem("token")) {
    result = true;
  } else {
    result = false;
  }
  return result;
};

const {
  mutate: RemoveLike,
  onDone: RemoveLikeDone,
  onError: RemoveLikeError,
} = useMutation(
  gql`
    mutation DeleteLike($id: uuid!) {
      delete_userLikes(where: { recipeId: { _eq: $id } }) {
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

RemoveLikeDone((handle) => {
  toast("Like removed successfully");
  console.log(handle);
});
RemoveLikeError((message) => {
  toast("Request failed");
});

const {
  mutate: RemoveFavorite,
  onDone: RemoveFavoriteDone,
  onError: RemoveFavoriteError,
} = useMutation(
  gql`
    mutation DeleteFavorite($id: uuid!) {
      delete_Favorite(where: { receipeId: { _eq: $id } }) {
        affected_rows
      }
    }
  `,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  })
);

RemoveFavoriteDone((handle) => {
  toast("Comment removed successfully");
});
RemoveFavoriteError((message) => {
  toast("Request failed");
});
</script>

<template>
  <Header></Header>
  <MaxWidth>
    <Loading v-if="loading"></Loading>
    <h1 v-else-if="error">Error loading ui</h1>
    <ul v-else-if="result">
      <p class="text-xl lg:text-4xl font-semibold my-3">
        {{ result.Recipe_by_pk.name }}
      </p>
      <div class="flex gap-3 my-4 items-center justify-between">
        <div class="flex gap-3 max-h-[30px]">
          <div class="flex items-center">
            <span class="inline-block pr-2">
              <Clock9 :size="15" />
            </span>
            <p>{{ " " + result.Recipe_by_pk.duration }} min</p>
          </div>
          <div class="flex flex-row gap-1">
            <Badge v-for="x in result.Recipe_by_pk.RecipeTags" v-if="result.Recipe_by_pk.RecipeTags.length">#{{
              x.Tag.tagName }}</Badge>
          </div>
        </div>
        <div class="" v-if="isLogged()">
          <Button v-if="result.Recipe_by_pk.userLikes[0]?.userId" class="rounded-r-[0]" variant="outline"
            @click="RemoveFromLike()">
            <ThumbsUp class="pr-2" fill="hsl(var(--foreground))" :stroke-width="2" />
            Like
          </Button>
          <Button v-else class="rounded-r-[0]" variant="outline" @click="addToLike()">
            <ThumbsUp class="pr-2" :stroke-width="2" /> Like
          </Button>
          <Button v-if="result.Recipe_by_pk.Favorites[0]?.userId" class="rounded-l-[0]" @click="RemoveFromFavorite()">
            <Bookmark class="pr-2" fill="hsl(var(--foreground))" :stroke-width="2" />
            Add to favorite
          </Button>
          <Button v-else class="rounded-l-[0]" @click="addToFavorite()">
            <Bookmark class="pr-2" :stroke-width="2" /> Add to favorite
          </Button>
        </div>
      </div>
      <p class="my-3">
        By:
        <span class="font-semibold">{{ result.Recipe_by_pk.user.firstName }}
          {{ result.Recipe_by_pk.user.lastName }}</span>
      </p>
      <FoodSlider v-if="result.Recipe_by_pk.recipeImages.length" :data="result.Recipe_by_pk.recipeImages" />
      <!-- {{ JSON.stringify(result.Recipe_by_pk.recipeImages) }} -->
      <div class="prose max-w-[100%]" v-html="getMarked(result.Recipe_by_pk.description)"></div>
      <!-- <p>{{marked.parse( result.Recipe_by_pk.description )}}</p> -->
      <section class="text-gray-600 body-font">
        <div class="py-4 mx-auto flex flex-wrap">
          <div class="flex flex-wrap w-full" v-for="(x, i) in result.Recipe_by_pk.recipeSteps">
            <div class="flex relative pb-12">
              <div v-if="result.Recipe_by_pk.recipeSteps.length != i + 1"
                class="h-full w-10 absolute inset-0 flex items-center justify-center">
                <div class="h-full w-1 bg-gray-200 pointer-events-none"></div>
              </div>
              <div
                class="flex-shrink-0 w-10 h-10 rounded-full bg-primary inline-flex items-center justify-center text-white relative z-10">
                <Dot />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="font-medium title-font text-sm text-primary mb-1 tracking-wider">
                  STEP {{ i + 1 }}
                </h2>
                <p class="leading-relaxed">
                  {{ x.description }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <div class="flex items-start space-x-4" v-if="isLogged()">
        <div class="flex-shrink-0">
          <img class="inline-block h-10 w-10 rounded-full"
            src="https://images.unsplash.com/photo-1550525811-e5869dd03032?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
            alt="" />
        </div>
        <div class="min-w-0 flex-1">
          <form action="#" class="relative" @submit="onSubmit">
            <div class="border border-gray-300 rounded-lg shadow-sm overflow-hidden">
              <FormField v-slot="{ componentField }" name="comment">
                <FormItem>
                  <FormControl>
                    <Textarea v-bind="componentField" rows="3" name="comment" id="comment"
                      class="block w-full py-3 border-0 resize-none focus:ring-0 sm:text-sm"
                      placeholder="Add your comment..." />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>
              <label for="comment" class="sr-only">Add your comment</label>

              <!-- Spacer element to match the height of the toolbar -->
              <div class="py-2" aria-hidden="true">
                <!-- Matches height of button in toolbar (1px border + 36px content height) -->
                <div class="py-px">
                  <div class="h-9" />
                </div>
              </div>
            </div>

            <div class="absolute bottom-0 inset-x-0 pl-3 pr-2 py-2 flex justify-between">
              <div class="flex flex-row text-sm items-center justify-center">
                <Star :fill="index < rating ? 'gold' : 'white'" v-for="(l, index) in 5" @click="rating = index + 1" />
              </div>
              <div class="flex-shrink-0">
                <Button type="submit"> Post </Button>
              </div>
            </div>
          </form>
        </div>
      </div>

      <h1 class="text-xl my-3 font-bold">Comments</h1>
      <div class="border rounded-md p-3 ml-3 my-3" v-for="x in result.Recipe_by_pk.Comments">
        <div class="flex gap-3 items-center">
          <img src="https://avatars.githubusercontent.com/u/22263436?v=4"
            class="object-cover w-8 h-8 rounded-full border-2 border-emerald-400 shadow-emerald-400" />

          <h3 class="font-bold">
            {{ x.user.firstName + " " + x.user.lastName }}
          </h3>
          <div class="flex flex-row text-sm items-center justify-center">
            <Star fill="gold" v-for="l in x.rating" />
          </div>
        </div>

        <p class="mt-2">{{ x.comment }}</p>
      </div>
    </ul>
  </MaxWidth>
</template>
