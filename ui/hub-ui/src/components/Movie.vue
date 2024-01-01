<script setup>
  defineProps({
    name: {
      type: String,
      required: true
    },
    imgUrl: {
      type: String,
      required: true
    },
    playlist: {
      type: String,
      required: true
    }
  })
</script>

<template>
  <router-link :to="{ name: 'movie', params: { name: name, url: playlist }}" class="p-0">
    <div class="card card-compact bg-base-100 shadow-xl">
      <figure>
        <img :src="imgUrl" alt="movie thumbnail" class="aspect-video" />
      </figure>
      <div class="card-body">
        <h2 class="card-title text-black dark:text-white">{{ name }}</h2>
      </div>
    </div>
  </router-link>
</template>


<style scoped>

</style>