<script setup>
import { useRoute } from 'vue-router'
import {VideoPlayer} from "@videojs-player/vue";

const route = useRoute()
const name = route.params.name
const url = route.params.url
</script>

<template>
  <div class="mx-auto max-w-7xl sm:px-6 lg:px-8">
    <div class="flex justify-center mt-5">
      <video-player
          :src="url"
          controls
          :loop="true"
          :volume="0.6"
      />
    </div>
    <div class="flex">
      <h3>{{ name }}</h3>
    </div>
  </div>
</template>

<style scoped>

</style>