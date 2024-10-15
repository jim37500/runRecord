<!-- eslint-disable vue/max-len -->
<template>
  <div class="min-h-full">
    <nav class="bg-gray-800 w-full fixed z-10">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 items-center justify-between">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <img class="size-10 rounded-full" src="/images/JimAndy.png" alt="Your Company" />
            </div>
            <div class="hidden md:block">
              <div class="ml-2 flex items-center space-x-4">
                <div class="px-3 py-2 font-medium text-2xl cursor-pointer text-yellow-500" @click="DirectToPage('/')" @keydown="() => router.push('/')">Growth</div>
                <div class="narbar-inactive" id="navbar-home" @click="DirectToPage('/')" @keydown="DirectToPage('/')">首頁</div>
                <div class="narbar-inactive" id="navbar-activity" @click="DirectToPage('/activity')" @keydown="DirectToPage('/activity')">我的運動</div>
                <div
                  class="narbar-inactive relative"
                  @mouseover="MenuVisible.ArticleShare = true"
                  @focus="MenuVisible.ArticleShare = true"
                  @mouseleave="MenuVisible.ArticleShare = false"
                  @focusout="MenuVisible.ArticleShare = false"
                >
                  <div class="font-medium">文章分享<i class="pi pi-angle-down ml-1" /></div>
                  <Menu :Items="ArticleShareItems" :Visible="MenuVisible.ArticleShare" :key="MenuVisible.ArticleShare" />
                </div>
              </div>
            </div>
          </div>
          <div class="hidden md:block">
            <div class="ml-4 flex items-center md:ml-6">
              <!-- Profile dropdown -->
              <div class="relative ml-3 cursor-pointer">
                <div>
                  <button
                    type="button"
                    class="relative flex items-center rounded-full p-1 bg-gray-800 text-sm focus:outline-none ring-2 ring-white"
                    id="user-menu-button"
                    aria-expanded="false"
                    aria-haspopup="true"
                    @click="ClickProfile"
                  >
                    <!-- <span class="absolute -inset-1.5" /> -->
                    <span class="sr-only">Open user menu</span>
                    <i class="pi pi-user text-white" style="font-size: 2rem" />
                  </button>
                </div>

                <!--
                Dropdown menu, show/hide based on menu state.

                Entering: "transition ease-out duration-100"
                  From: "transform opacity-0 scale-95"
                  To: "transform opacity-100 scale-100"
                Leaving: "transition ease-in duration-75"
                  From: "transform opacity-100 scale-100"
                  To: "transform opacity-0 scale-95"
              -->
                <Menu :Items="ProfileItems" :Visible="MenuVisible.Profile" :PositionClass="'right-0'" :key="MenuVisible.Profile" />
              </div>
            </div>
          </div>
          <div class="-mr-2 flex md:hidden">
            <!-- Mobile menu button -->
            <button
              type="button"
              class="relative inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
              aria-controls="mobile-menu"
              aria-expanded="false"
              @click="IsMenuOpen = !IsMenuOpen"
            >
              <span class="absolute -inset-0.5" />
              <span class="sr-only">Open main menu</span>
              <!-- Menu open: "hidden", Menu closed: "block" -->
              <svg class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
              </svg>
              <!-- Menu open: "block", Menu closed: "hidden" -->
              <svg class="hidden h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile menu, show/hide based on menu state. -->
      <div id="mobile-menu" :class="IsMenuOpen && IsMobile ? 'max-h-screen' : 'max-h-0'" class="overflow-hidden transition-all duration-500 ease-out">
        <div class="space-y-1 px-2 pb-3 pt-2 sm:px-3">
          <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
          <div
            id="navbar-home-mobile"
            class="cursor-pointer block rounded-md px-3 py-2 text-base font-medium text-gray-300 hover:bg-gray-700 hover:text-white"
            @click="DirectToPage('/')"
            @keydown="DirectToPage('/')"
          >
            首頁
          </div>
          <div
            id="navbar-activity-mobile"
            class="cursor-pointer block rounded-md px-3 py-2 text-base font-medium text-gray-300 hover:bg-gray-700 hover:text-white"
            @click="DirectToPage('/activity')"
            @keydown="DirectToPage('/activity')"
          >
            我的運動
          </div>
          <div class="cursor-pointer block rounded-md px-3 py-2 text-base font-medium text-gray-300 hover:bg-gray-700 hover:text-white">文章分享</div>
          <div v-for="item in ArticleShareItems" :key="item" class="block rounded-md px-3 py-1 ml-4 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white">
            {{ item.Name }}
          </div>
        </div>
        <div class="border-t border-gray-700 pb-3 pt-4">
          <div class="flex items-center px-5">
            <div class="rounded-full h-10 w-10 ring-2 ring-white flex justify-center items-center">
              <i class="pi pi-user text-white" style="font-size: 2rem" />
            </div>
            <!-- <div class="flex-shrink-0">
              <img
                class="h-10 w-10 rounded-full"
                src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                alt=""
              />
            </div> -->
            <div class="ml-3">
              <div class="text-base font-medium leading-none text-white">Tom Cook</div>
              <div class="text-sm font-medium leading-none text-gray-400">tom@example.com</div>
            </div>
          </div>
          <div class="mt-3 space-y-1 px-2">
            <div v-for="item in ProfileItems" :key="item" class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white">
              {{ item.Name }}
            </div>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script setup>
import { onBeforeUnmount, onMounted, ref } from 'vue';
import Menu from './nav-menu.vue';

const { router } = window;

const ArticleShareItems = ref([{ Name: '運動筆記' }, { Name: '程式筆記' }, { Name: '閱讀筆記' }]); // 文章分享選單內容
const ProfileItems = ref([{ Name: '帳號管理' }, { Name: '管理員' }, { Name: '登出' }]); // 帳號選單內容
const MenuVisible = ref({ ArticleShare: false, Profile: false }); // 選單是否顯示(預設皆為否)
const IsMobile = ref(window.innerWidth < 768); // 是否為手機版
const IsMenuOpen = ref(false); // 手機版是否打開選單
const NavbarMap = { '/': 'home', '/activity': 'activity' }; // 導覽列路徑與id名稱對應
const IsClickProfile = ref(false); // 是否點擊帳號頭像

// 為了要點擊帳號頭像時能顯示選單，點擊其他地方會不顯示
// 點擊帳號頭像
const ClickProfile = () => {
  if (MenuVisible.value.Profile) {
    MenuVisible.value.Profile = false;
    return;
  }

  IsClickProfile.value = true;
  MenuVisible.value.Profile = true;
};

// 處理點選帳號頭像之外的事件
const HandleClickOutside = () => {
  if (IsClickProfile.value) {
    setTimeout(() => {
      IsClickProfile.value = false;
    }, 0);
    return;
  }

  MenuVisible.value.Profile = false;
};

const Load = () => {
  let navbarID;
  Object.keys(NavbarMap).forEach((o) => {
    if (router.currentRoute.value.path.startsWith(o)) {
      navbarID = NavbarMap[o];
    }
  })
  
  const element = document.querySelector(`#navbar-${navbarID}`);
  element.classList.add('narbar-active');
  element.classList.remove('narbar-inactive');

  const mobileElement = document.querySelector(`#navbar-${navbarID}-mobile`);
  mobileElement.classList.add('bg-gray-900', 'text-white');
  mobileElement.classList.remove('text-gray-300', 'hover:bg-gray-700', 'hover:text-white');

  document.addEventListener('click', HandleClickOutside);
};

// 導致其他分頁
const DirectToPage = (path) => {
  Object.keys(NavbarMap).forEach((o) => {
    const element = document.querySelector(`#navbar-${NavbarMap[o]}`);
    const mobileElement = document.querySelector(`#navbar-${NavbarMap[o]}-mobile`);

    if (o === path) {
      element.classList.remove('narbar-inactive');
      element.classList.add('narbar-active');
      mobileElement.classList.add('bg-gray-900', 'text-white');
      mobileElement.classList.remove('text-gray-300', 'hover:bg-gray-700', 'hover:text-white');
      return;
    }

    element.classList.remove('narbar-active');
    element.classList.add('narbar-inactive');
    mobileElement.classList.remove('bg-gray-900', 'text-white');
    mobileElement.classList.add('text-gray-300', 'hover:bg-gray-700', 'hover:text-white');
  });

  IsMenuOpen.value = false;
  router.push(`${path}`);
};

onMounted(() => Load());

onBeforeUnmount(() => {
  document.removeEventListener('click', HandleClickOutside);
});
</script>
