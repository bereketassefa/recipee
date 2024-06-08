<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { FormKit } from "@formkit/vue";
import { useRouter, useRoute } from "vue-router";
import { useMutation } from "@vue/apollo-composable";
import gql from "graphql-tag";
const router = useRouter();

import { h } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { setContext } from 'apollo-link-context'

import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { toast } from "vue-sonner";

const formSchema = toTypedSchema(
  z.object({
    email: z.string().min(2).max(50).email(),
    password: z.string().min(2).max(50),
  })
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((data) => {
  console.log({ email: data.email, password: data.password });
  LogUser({ email: data.email, password: data.password });
});

const {
  mutate: LogUser,
  onDone,
  error,
  onError,
} = useMutation(gql`
  mutation Log($email: String!, $password: String!) {
    Login(email: $email, password: $password) {
      message
      token
      id
    }
  }
`);

onDone((handle) => {
  console.log(handle);
  localStorage.setItem("token", handle.data.Login.token);
  localStorage.setItem("id", handle.data.Login.id);
  toast("Login Successful");

  // router.push("/");
  router.push('/', () => router.go(0))
});
onError((message) => {
  console.log(message.message);
  toast(message.message);
});
</script>

<template>
  <div class="w-full lg:grid  lg:grid-cols-2 h-screen">
    <div class="flex items-center justify-center py-12">
      <div class="mx-auto grid w-[350px] gap-6">
        <div class="grid gap-2 text-center">
          <h1 class="text-3xl font-bold">Login</h1>
          <p class="text-balance text-muted-foreground">
            Enter your email below to login to your account
          </p>
        </div>

        <form class="space-y-6" @submit="onSubmit">
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
          <Button type="submit"> Login </Button>
        </form>
        <div class="mt-4 text-center text-sm">
          Don't have an account?
          <RouterLink to="/signup" class="underline"> Sign up </RouterLink>
        </div>
      </div>
    </div>
    <div class="hidden bg-muted lg:block">
      <img
        src="https://images.unsplash.com/photo-1555939594-58d7cb561ad1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Nnx8Zm9vZHxlbnwwfHwwfHx8MA%3D%3D"
        alt="Image" width="1920" class="h-full w-full object-cover dark:brightness-[0.2] dark:grayscale h-screen" />
    </div>
  </div>
</template>
