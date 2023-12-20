<script setup>
import { useMovieStore } from '@/stores/movieStore.js'
import {computed, onMounted} from "vue";
import MovieCard from "@/components/Movie.vue";

const movieStore = useMovieStore();
onMounted(() => {
  console.log("MOUNTED")
  movieStore.fetchMovies()
      .catch(error => {
        console.error('Error fetching movies:', error);
      });
});

const movies = computed(() => movieStore.movies);
</script>

<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
    <MovieCard
        v-for="movie in movies.data"
        :key="movie.Name"
        :img-url="movie.Thumbnail"
        :name="movie.Name"
        :playlist="movie.Playlist"
    ></MovieCard>
  </div>

</template>
