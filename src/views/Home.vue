<template>
  <TopBar />
  <div class="flex">
    <RouterView />
  </div>
  <Footer />
</template>

<script setup>
import TopBar from '../components/TopBar.vue';
import Footer from '../components/Footer.vue';
</script>
