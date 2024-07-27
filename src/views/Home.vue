<template>
  <TopBar />
  <div class="flex justify-content-center w-12">
    <RouterView />
  </div>
</template>

<script setup>
import TopBar from '../components/TopBar.vue';
</script>
