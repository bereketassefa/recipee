<script setup>
import { ref } from "vue";
import { watchOnce } from "@vueuse/core";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import { Card, CardContent } from "@/components/ui/card";

const props = defineProps(["data"]);

const emblaMainApi = ref();
const emblaThumbnailApi = ref();
const selectedIndex = ref(0);

function onSelect() {
  if (!emblaMainApi.value || !emblaThumbnailApi.value) return;
  selectedIndex.value = emblaMainApi.value.selectedScrollSnap();
  emblaThumbnailApi.value.scrollTo(emblaMainApi.value.selectedScrollSnap());
}

function onThumbClick(index) {
  if (!emblaMainApi.value || !emblaThumbnailApi.value) return;
  emblaMainApi.value.scrollTo(index);
}

watchOnce(emblaMainApi, (emblaMainApi) => {
  if (!emblaMainApi) return;

  onSelect();
  emblaMainApi.on("select", onSelect);
  emblaMainApi.on("reInit", onSelect);
});
</script>

<template>
  <div class="w-full sm:w-auto">
    <Carousel
      class="relative w-full max-w-[100%] lg:max-w-[80%]"
      :opts="{
        align: 'start',
      }"
    >
      <CarouselContent class="-ml-1">
        <CarouselItem
          v-for="(val, index) in props.data"
          :key="index"
          class="pl-1 lg:basis-[80%] xl:basis-[60%]"
        >
          <div class="p-1">
            <img :src="val.imageLink" alt="" srcset="" class="aspect-video" />
          </div>
        </CarouselItem>
      </CarouselContent>
      <CarouselPrevious />
      <CarouselNext />
    </Carousel>
  </div>
</template>
