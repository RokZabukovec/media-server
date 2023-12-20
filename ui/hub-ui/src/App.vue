<script setup>
import {RouterLink, RouterView} from 'vue-router'
import Navigation from "@/components/Navigation.vue";
</script>

<template>
  <header>
    <div class="container mx-auto">
      <Navigation></Navigation>
    </div>
  </header>
  <RouterView />
</template>

<style scoped>

</style>
