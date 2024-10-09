<template>
  <TopBar />
  <div class="flex pt-20">
    <RouterView />
  </div>
  <Footer />
</template>

<script setup>
import TopBar from '../components/top-bar.vue';
import Footer from '../components/home-footer.vue';
</script>
