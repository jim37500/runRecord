<template>
  <div class="flex flex-col xl:flex-row justify-center mx-auto w-full max-w-7xl mt-10 px-4 sm:px-6 lg:px-10">
    <div class="w-full xl:w-6/12 flex flex-col justify-center">
      <div class="text-6xl text-yellow-500 font-semibold">Growth</div>
      <div class="text-3xl font-semibold mt-5 mb-2">每天進步一點</div>
      <div class="text-3xl font-semibold">五年十年後大不同</div>
      <div class="mt-5">
        <button
          v-if="!IsLoggedIn"
          type="button"
          class="rounded-lg font-semibold text-white text-lg px-6 py-2 bg-sky-700 hover:bg-sky-500 active:bg-sky-700"
          @click="store.OpenLoginDialog()"
        >
          登入
        </button>
      </div>
    </div>
    <div class="w-full xl:w-6/12 mt-10 xl:mt-0 grid grid-cols-1 sm:grid-cols-2 gap-8">
      <div
        v-for="item in HeroItems"
        :key="item"
        class="bg-zinc-50 border-2 border-zinc-100 px-8 py-8 h-72 rounded-2xl hover:bg-blue-50 hover:border-blue-200 shadow-md cursor-pointer"
        @click="router.push(item.Path)"
      >
        <div class="w-12 h-12 flex justify-center items-center rounded-md" :class="item.IconBgColor">
          <i :class="item.IconClass" style="font-size: 1.875rem; color: white" />
        </div>
        <div class="text-2xl mt-3 font-semibold">{{ item.Title }}</div>
        <div class="mt-1 text-lg leading-8 indent-9">{{ item.Content }}</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { loginStore } from '../stores/LoginStore';
import { storeToRefs } from 'pinia';

const store = loginStore();
const { router } = window;
const { IsLoggedIn } = storeToRefs(store);
const HeroItems = ref([
  {
    IconClass: 'pi pi-map-marker',
    IconBgColor: 'bg-sky-700',
    Title: '運動紀錄',
    Content: '連動Strava追蹤，讓你能掌握自己的運動狀況與規劃訓練表。',
    Path: 'activity',
  },
  {
    IconClass: 'pi pi-pencil',
    IconBgColor: 'bg-lime-700',
    Title: '運動筆記',
    Content: '包含游泳、自行車及跑步，馬拉松與鐵人賽事準備。',
    Path: '',
  },
  {
    IconClass: 'pi pi-code',
    IconBgColor: 'bg-yellow-700',
    Title: '程式筆記',
    Content: '程式小白自學筆記，從軟體前後端必讀知識點，到資料結構與演算法。',
    Path: '',
  },
  {
    IconClass: 'pi pi-book',
    IconBgColor: 'bg-pink-700',
    Title: '閱讀筆記',
    Content: '日常閱讀筆記，包含如何提升生產力、提高學習效果。',
    Path: '',
  },
]);
</script>
