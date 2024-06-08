<template>
  <Header />
  <MaxWidth>
    <article>
      <!-- Profile header -->
      <div>
        <div>
          <img class="h-32 w-full object-cover lg:h-48" :src="profile.coverImageUrl" alt="" />
        </div>

        <div class="mx-auto px-4 sm:px-6 lg:px-8" v-if="result && result.users_by_pk">
          <div class="-mt-12 sm:-mt-16 sm:flex sm:items-end sm:space-x-5">
            <div class="flex">
              <img class="h-24 w-24 rounded-full ring-4 ring-white sm:h-32 sm:w-32"
                :src="result.users_by_pk?.profileImg" alt="" />
            </div>
            <div class="mt-6 sm:flex-1 sm:min-w-0 sm:flex sm:items-center sm:justify-end sm:space-x-6 sm:pb-1">
              <div class="sm:hidden 2xl:block mt-6 min-w-0 flex-1">
                <h1 class="text-2xl font-bold  truncate" v-if="result && result.Recipe">
                  {{
                    result.users_by_pk.firstName +
                    result.users_by_pk.lastName
                  }}
                </h1>
              </div>
              <div class="mt-6 flex flex-col justify-stretch space-y-3 sm:flex-row sm:space-y-0 sm:space-x-4">
                <!-- <Sheet>
                  <SheetTrigger asChild>
                    <Button variant="outline" class="">
                      <Set class="-ml-1 mr-2 h-5 w-5 text-gray-400" aria-hidden="true" />
                      <span>Settings</span>
                    </Button>
                  </SheetTrigger>
                  <SheetContent>
                    <SheetHeader>Settings</SheetHeader>
                    <Settings />
                  </SheetContent>
                </Sheet> -->
                <Dialog v-if="isLogged()">
                  <DialogTrigger as-child>
                    <Button variant="outline">
                      Edit Profile
                    </Button>
                  </DialogTrigger>
                  <DialogContent class="sm:max-w-[425px]">
                    <DialogHeader>
                      <DialogTitle>Edit profile</DialogTitle>
                      <DialogDescription>
                        Make changes to your profile here. Click save when you're done.
                      </DialogDescription>
                    </DialogHeader>
                    <div class="grid w-full max-w-sm items-center gap-1.5">
                      <Label for="picture">Picture</Label>
                      <Input @change="handleFileChange" id="picture" type="file" />
                    </div>
                    <DialogFooter>
                      <Button @click="upload">
                        Save changes
                      </Button>
                    </DialogFooter>
                  </DialogContent>
                </Dialog>
                <!-- <button
                  type="button"
                  class="inline-flex justify-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-pink-500"
                >
                  <PhoneIcon
                    class="-ml-1 mr-2 h-5 w-5 text-gray-400"
                    aria-hidden="true"
                  />
                  <span>Call</span>
                </button> -->
              </div>
            </div>
          </div>
          <div class="hidden sm:block 2xl:hidden mt-6 min-w-0 flex-1">
            <h1 class="text-2xl font-bold  truncate" v-if="result && result.Recipe">
              {{
                result.users_by_pk.firstName +
                result.users_by_pk.lastName
              }}
            </h1>
          </div>
        </div>
      </div>

      <Tabs default-value="recipes">
        <TabsList class="justify-start gap-3 border-0">
          <TabsTrigger default value="recipes"> Recipes </TabsTrigger>
          <TabsTrigger value="favorite"> Favorite </TabsTrigger>
          <TabsTrigger value="liked"> Liked </TabsTrigger>
        </TabsList>
        <TabsContent value="recipes">
          <Notright v-if="error">Error loading recipe</Notright>
          <Loading v-else-if="loading"></Loading>
          <div v-else-if="result && result.Recipe"
            class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4">
            <ReceipeCard v-for="x in result.Recipe" :id="x.id" :name="x.name" :author="x.user" :duration="x.duration"
              :tags="x.RecipeTags" :image="x.recipeImages" profile="profile" :all="x" />
          </div>
          <Notright v-else-if="result && (result.Recipe.lenth > 0)">The user does't create any recipe yet</Notright>

        </TabsContent>
        <TabsContent value="favorite">

          <Notright v-if="error">Error loading recipe</Notright>
          <Loading v-else-if="loading"></Loading>
          <div v-else-if="result && result.Favorite"
            class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4">

            <ReceipeCard v-for="x in result.Favorite" :id="x.Recipe.id" :name="x.Recipe.name" :author="x.Recipe.user"
              :duration="x.Recipe.duration" :tags="x.Recipe.RecipeTags" :image="x.Recipe.recipeImages"
              profile="favorite" :all="x" />
          </div>
          <Notright v-else-if="result && (result.Favorite.lenth > 0)">No favorite added yet</Notright>
        </TabsContent>
        <TabsContent value="liked">
          <div v-if="error">Error loading recipe</div>
          <Loading v-else-if="loading"></Loading>
          <div v-else-if="result && result.userLikes"
            class="grid grid-cols-1 sm:grid-cols-3 lg:grid-cols-3 xl:grid-cols-4 gap-4 my-4">
            <ReceipeCard v-for="x in result.userLikes" :id="x.Recipe.id" :name="x.Recipe.name" :author="x.Recipe.user"
              :duration="x.Recipe.duration" :tags="x.Recipe.RecipeTags" :image="x.Recipe.recipeImages" profile="liked"
              :all="x" />
          </div>
          <Notright v-else-if="result && (result.userLikes.lenth > 0)">{{ result.userLikes.lenth }}Did't like any video
            yet</Notright>
        </TabsContent>
      </Tabs>
    </article>
  </MaxWidth>
</template>
<script setup>
import Header from "@/components/Header.vue";
import Loading from "@/components/Loading.vue";
import MaxWidth from "@/components/MaxWidth.vue";
import Notright from "@/components/Notright.vue";
import ReceipeCard from "@/components/ReceipeCard.vue";
import Settings from "@/components/Settings.vue";
import Button from "@/components/ui/button/Button.vue";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import Input from "@/components/ui/input/Input.vue";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { useMutation, useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { Settings as Set } from "lucide-vue-next";
import { ref, watch } from "vue";
import { useRoute } from "vue-router";
import { toast } from "vue-sonner";
const route = useRoute();

const profile = {
  name: "Ricardo Cooper",
  imageUrl:
    "https://images.unsplash.com/photo-1463453091185-61582044d556?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=1024&h=1024&q=80",
  coverImageUrl:
    "https://images.unsplash.com/photo-1444628838545-ac4016a5418a?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1950&q=80",
  about: `
    <p>Tincidunt quam neque in cursus viverra orci, dapibus nec tristique. Nullam ut sit dolor consectetur urna, dui cras nec sed. Cursus risus congue arcu aenean posuere aliquam.</p>
    <p>Et vivamus lorem pulvinar nascetur non. Pulvinar a sed platea rhoncus ac mauris amet. Urna, sem pretium sit pretium urna, senectus vitae. Scelerisque fermentum, cursus felis dui suspendisse velit pharetra. Augue et duis cursus maecenas eget quam lectus. Accumsan vitae nascetur pharetra rhoncus praesent dictum risus suspendisse.</p>
  `,
};



const { result, loading, error, variables, refetch } = useQuery(
  gql`
    query MyQuery($id: uuid! ) {
      
      users_by_pk(id: $id) {
        firstName
        lastName
        isPublic
        profileImg
        email
      }

      Recipe(where: { userId: { _eq: $id } ,userLike: {userId: {_eq: $id}} }) {
        description
        duration
        id
        name
        user {
          lastName
          firstName
          email
          id
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
      Favorite(where: { userId: { _eq: $id } }) {
        id
        Recipe {
          description
          duration
          id
          name
          user {
            lastName
            firstName
            email
            id
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
      userLikes(where: { userId: { _eq: $id } }) {
        id
        userId
        Recipe {
          description
          duration
          id
          name
          user {
            lastName
            firstName
            email
            id
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
    }
  `,
  { id: route.params.id },
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
const imageLink = ref("");


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
      emit("fileUploaded", file); // Emit an event with the uploaded file

    })
    .catch((error) => {
      console.error("Upload error:", error);
      error.value = "Upload failed. Please try again.";
    });
};
const isLogged = () => {
  let result;
  if (localStorage.getItem("token") && route.params.id == localStorage.getItem("id")) {
    result = true;
  } else {
    result = false;
  }
  return result;
};

const upload = () => {
  if (imageLink.value == '') {
    toast("please upload image")
  } else {
    uploadProfile({ id: localStorage.getItem("id"), url: imageLink.value })
  }
}

const {
  mutate: uploadProfile,
  onDone,
  onError,
} = useMutation(gql`
    mutation changeprofile( $id: uuid! , $url: String!) {
      update_users_by_pk(pk_columns: {id: $id}, _set: {profileImg: $url}) {
        id
        profileImg
      }
    }
` ,
  () => ({
    refetchQueries: "active",
    fetchPolicy: "no-cache",
  }));
onDone((handle) => {
  toast("Image uploaded succesfully");
});
onError((message) => {
  toast("There is an error uploading image");
});

</script>
<style></style>
