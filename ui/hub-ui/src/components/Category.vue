<template>
  <div class="flex bg-base-100 p-3 rounded">
    <IconFolder></IconFolder>
    <span class="pl-2 folder-name">{{ name }}</span>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';
import IconFolder from "@/components/icons/IconFolder.vue";

const props = defineProps({
  name: String
});
</script>

<style scoped>
.folder-name {
  font-size: 16px; /* adjust as needed */
}
</style>
