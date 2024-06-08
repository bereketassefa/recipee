<script setup>
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Menu, Search, User, Utensils } from "lucide-vue-next";
import { useRouter } from "vue-router";
import Maxwidth from "./MaxWidth.vue";
import ThemeToggle from "./ThemeToggle.vue";
import SheetClose from "./ui/sheet/SheetClose.vue";
import SheetFooter from "./ui/sheet/SheetFooter.vue";
import { setContext } from 'apollo-link-context'

const router = useRouter();

const isLogged = () => {
  let result;
  if (localStorage.getItem("token")) {
    result = true;
  } else {
    result = false;
  }
  return result;
};

const logout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("id");
  setContext(async (_, { headers }) => {
    const token = await localStorage.getItem(AUTH_TOKEN)
    return {
      ...headers,
      "x-hasura-role": "anonymous",
    }
  })
  router.push("/");
  router.go();

};
const getid = () => {
  return localStorage.getItem("id");
};
</script>

<template>
  <header className="" id="header">
    <Maxwidth>
      <div className=" ">
        <div className="flex h-16 items-center justify-between">
          <div className="md:flex md:items-center gap-4 md:gap-12">
            <RouterLink to="/" className="flex items-center gap-2 text-lg font-semibold md:text-base">
              <!-- <HandCoins className="h-6 w-6" /> -->
              <Utensils />
              <span className="sr-only">Acme Inc</span>
            </RouterLink>
          </div>

          <div className="hidden md:block">
            <nav aria-label="Global"></nav>
          </div>

          <div className="flex items-center gap-4">
            <div className="flex  gap-1 sm:gap-2  items-center justify-center">
              <Button variant="ghost">
                <RouterLink to="/search">
                  <Search />
                </RouterLink>
              </Button>
              <ThemeToggle />

              <Button variant="outline" asChild v-if="isLogged()">
                <RouterLink to="/new">Create a receipe</RouterLink>
              </Button>

              <div className="hidden sm:flex">
                <DropdownMenu v-if="isLogged()">
                  <DropdownMenuTrigger as-child>
                    <Avatar>
                      <AvatarImage src="https://github.com/radix-vue.png" alt="@radix-vue" />
                      <AvatarFallback>CN</AvatarFallback>
                    </Avatar>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent class="w-56">
                    <DropdownMenuLabel>My Account</DropdownMenuLabel>
                    <DropdownMenuSeparator />
                    <DropdownMenuGroup>
                      <DropdownMenuItem asChild>
                        <RouterLink :to="'/profile/' + getid()">
                          <User class="mr-2 h-4 w-4" />
                          Profile
                          <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
                        </RouterLink>
                      </DropdownMenuItem>
                    </DropdownMenuGroup>

                    <DropdownMenuSeparator />
                    <DropdownMenuItem @click="logout()">
                      <span>Log out</span>
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
                <Button asChild v-if="!isLogged()">
                  <RouterLink to="/login">Login</RouterLink>
                </Button>
              </div>
            </div>

            <div className="block md:hidden">
              <Sheet>
                <SheetTrigger asChild>
                  <Button variant="ghost">
                    <Menu />
                  </Button>
                </SheetTrigger>
                <SheetContent>
                  <SheetHeader></SheetHeader>
                  <nav aria-label="Global">
                    <ul className="flex flex-col items-center gap-6 text-sm">
                      <li>list</li>
                    </ul>
                  </nav>
                  <SheetFooter>
                    <SheetClose asChild></SheetClose>
                  </SheetFooter>
                </SheetContent>
              </Sheet>
            </div>
          </div>
        </div>
      </div>
    </Maxwidth>
  </header>
</template>
