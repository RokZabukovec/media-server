<script setup>
import { useMovieStore } from '@/stores/movieStore.js'
import {computed, onMounted} from "vue";
import MovieCard from "@/components/Movie.vue";
import {useCategoryStore} from "@/stores/categoryStore.js";
import Category from "@/components/Category.vue";

const categoryStore = useCategoryStore();
const movieStore = useMovieStore();
onMounted(() => {
  console.log("MOUNTED")
  categoryStore.fetchCategories(0)
      .catch(error => {
        console.error('Error fetching categories:', error);
      });

  movieStore.fetchMovies()
      .catch(error => {
        console.error('Error fetching movies:', error);
      });
});

const movies = computed(() => movieStore.movies);
const categories = computed(() => categoryStore.categories);
</script>

<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 my-3">
    <Category
        v-for="category in categories.data"
        :key="category.Id"
        :name="category.Name"
    ></Category>
  </div>

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
