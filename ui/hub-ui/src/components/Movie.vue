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
  <router-link :to="{ name: 'movie', params: { name: name, url: playlist }}">
    <div class="card card-compact max-w-md mx-auto bg-base-100 shadow-xl md:w-96 md:mx-auto">
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