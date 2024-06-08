<template>
    <footer class="border-t border-border mt-7">
        <div class="relative mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8 lg:pt-24">
            <div class="absolute end-4 top-4 sm:end-6 sm:top-6 lg:end-8 lg:top-8">
                <a class="inline-block rounded-full bg-primary p-2 text-white shadow transition hover:bg-teal-500 sm:p-3 lg:p-4"
                    href="#header">
                    <span class="sr-only">Back to top</span>

                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd"
                            d="M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z"
                            clip-rule="evenodd" />
                    </svg>
                </a>
            </div>

            <div class="lg:flex lg:items-end lg:justify-between">
                <div>
                    <div class="flex justify-center text-primary lg:justify-start">
                        <Utensils size="35" />
                    </div>

                    <p class="mx-auto mt-6 max-w-md text-center leading-relaxed  lg:text-left">
                        Craving inspiration? Explore a world of recipes! Easy-to-follow dishes for every skill level and
                        every craving. Cook with confidence, from healthy meals to global flavors. Let's get cooking!
                    </p>
                </div>

                <ul class="mt-12 flex flex-wrap justify-center gap-6 md:gap-8 lg:mt-0 lg:justify-end lg:gap-12">
                    <li>
                        <Button variant="link" asChild>
                            <RouterLink to="/">Home </RouterLink>
                        </Button>
                    </li>

                    <li>
                        <Button variant="link" asChild>
                            <RouterLink to="/search"> Search</RouterLink>
                        </Button>
                    </li>

                    <!-- <li>
                        <Button variant="link" asChild>
                            <RouterLink to="/"> Projects</RouterLink>
                        </Button>
                    </li> -->

                    <!-- <li>
                        <Button variant="link" asChild>
                            <RouterLink to="/"> Blog</RouterLink>
                        </Button>
                    </li> -->
                </ul>
            </div>

            <p class="mt-12 text-center text-sm text-gray-500 lg:text-right">
                Copyright &copy; 2022. All rights reserved.
            </p>
        </div>
    </footer>
</template>

<script setup>
import { Utensils } from "lucide-vue-next";

</script>