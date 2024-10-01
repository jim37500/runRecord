<template>
  <div
    v-if="props.Visible"
    class="absolute z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
    :class="props.PositionClass"
    role="menu"
    aria-orientation="vertical"
    aria-labelledby="user-menu-button"
    tabindex="-1"
  >
    <!-- Active: "bg-gray-100", Not Active: "" -->
    <div v-for="item in props.Items" :key="item" class="block px-4 py-2 text-base text-gray-700 hover:bg-gray-100" role="menuitem" tabindex="-1">{{ item.Name }}</div>
  </div>
</template>

<script setup>
const props = defineProps({
  Items: { Array, default: [] },
  Visible: { Boolean, default: false },
  PositionClass: { String, default: 'left-0' },
});
</script>
