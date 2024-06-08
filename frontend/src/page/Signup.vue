<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { FormKit } from "@formkit/vue";
import { useRouter, useRoute } from "vue-router";
import { useMutation } from "@vue/apollo-composable";
import gql from "graphql-tag";
const router = useRouter();
import { toast } from "vue-sonner";

import { h } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const formSchema = toTypedSchema(
  z.object({
    email: z.string().min(2).max(50).email(),
    password: z.string().min(2).max(50),
    firstname: z.string().min(2).max(50),
    lastname: z.string().min(2).max(50),
  })
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((data) => {
  SignUser({
    email: data.email,
    password: data.password,
    firstName: data.firstname,
    lastName: data.lastname,
  });
  console.log({
    email: data.email,
    password: data.password,
    firstName: data.firstname,
    lastName: data.lastname,
  });
});

const {
  mutate: SignUser,
  onDone,
  error,
  onError,
} = useMutation(gql`
  mutation Signup(
    $firstName: String!
    $lastName: String!
    $email: String!
    $password: String!
  ) {
    Signup(
      firstName: $firstName
      lastName: $lastName
      email: $email
      password: $password
    ) {
      message
    }
  }
`);

onDone((handle) => {
  toast("Signed up successfully Successful");
  router.push("/login");
});
onError((message) => {
  console.log(message.message);
  toast(message.message);
});

const handleLogin = (data) => {
  SignUser({
    email: data.email,
    password: data.password,
    firstName: data.firstName,
    lastName: data.lastName,
  });
  console.log({
    email: data.email,
    password: data.password,
    firstName: data.firstName,
    lastName: data.lastName,
  });
};
</script>

<template>
  <div class="w-full lg:grid  lg:grid-cols-2 ">
    <div class="flex items-center justify-center py-12 h-screen">
      <div class="mx-auto grid w-[350px] gap-6">
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">Signup</h1>
          <p class="text-balance text-muted-foreground">
            Enter your email below to create new account
          </p>
        </div>

        <form class="space-y-6" @submit="onSubmit">
          <FormField v-slot="{ componentField }" name="firstname">
            <FormItem>
              <FormLabel>First Name</FormLabel>
              <FormControl>
                <Input type="text" placeholder="First Name" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="lastname">
            <FormItem>
              <FormLabel>Last Name</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Last Name" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Email" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Password" v-bind="componentField" />
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
          <Button type="submit"> Signup </Button>
        </form>
        <div class="mt-4 text-center text-sm">
          Don't have an account?
          <a href="#" class="underline"> Sign up </a>
        </div>
      </div>
    </div>
    <div class="hidden bg-muted lg:block">
      <img
        src="https://images.unsplash.com/photo-1555939594-58d7cb561ad1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Nnx8Zm9vZHxlbnwwfHwwfHx8MA%3D%3D"
        alt="Image" width="1920" class="h-screen w-full object-cover dark:brightness-[0.2] dark:grayscale " />
    </div>
  </div>
</template>
