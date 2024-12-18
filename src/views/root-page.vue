<template>
  <TopBar :key="LoginOrLogoutKey" />
  <div class="flex pt-20">
    <RouterView :key="LoginOrLogoutKey" />
  </div>
  <Footer />

  <primevue-dialog v-model:visible="IsLoginDialogOpen" modal :header="IsRegister ? '註冊' : '登入'" class="w-10/12 sm:w-4/12" :draggable="false">
    <div class="flex flex-col items-center gap-4 p-4">
      <!-- Input fields -->
      <div class="w-full">
        <primevue-input-text v-model="Email" placeholder="電子郵件" class="w-full" />
      </div>
      <div class="w-full">
        <primevue-input-text v-model="Password" type="password" placeholder="密碼" class="w-full" />
      </div>
      <div v-if="IsRegister" class="w-full">
        <primevue-input-text v-model="ConfirmPassword" type="password" placeholder="請再次輸入密碼" class="w-full" />
      </div>

      <button v-if="IsRegister" type="button" class="px-4 py-2 mx-2 rounded font-semibold text-white bg-sky-700 hover:bg-sky-600 transition-colors" @click="Register">註冊</button>
      <button v-else type="button" class="px-4 py-2 mx-2 rounded font-semibold text-white bg-sky-700 hover:bg-sky-600 transition-colors" @click="Login">登入</button>

      <!-- Divider -->
      <div class="relative w-full py-3">
        <div class="absolute inset-0 flex items-center">
          <div class="w-full border-t border-gray-300"></div>
        </div>
        <div class="relative flex justify-center">
          <span class="bg-white px-2 text-sm text-gray-500">或</span>
        </div>
      </div>

      <!-- Google Sign In Button -->
      <div class="w-full">
        <button type="button" class="w-full flex items-center justify-center px-4 py-2 border border-gray-300 rounded-md shadow-sm bg-white hover:bg-gray-50 transition-colors">
          <img class="w-7 mr-2 rounded-full" src="/images/google.png" alt="Google logo" />
          <span class="text-gray-700 font-semibold">使用Google登入</span>
        </button>
      </div>
    </div>

    <template #footer>
      <div v-if="IsRegister" class="flex justify-center w-full">
        已經有帳號？ 點此<span class="ml-0.5 text-blue-700 hover:text-blue-500 cursor-pointer underline underline-offset-4" @click="IsRegister = false">登入</span>
      </div>
      <div v-else class="flex justify-center w-full">
        還沒有帳號？ 點此<span class="ml-0.5 text-blue-700 hover:text-blue-500 cursor-pointer underline underline-offset-4" @click="IsRegister = true">註冊</span>
      </div>
    </template>
  </primevue-dialog>
  <primevue-dialog v-model:visible="IsStravaDialogOpen" modal header="授權Strava" class="w-10/12 sm:w-6/12" :draggable="false">
    <template #header>
      <div class="flex items-center">
        <div class="font-semibold text-xl">授權Strava</div>
        <i v-tooltip.top="'授權Strava教學'" class="ml-2 pi pi-question-circle text-xl mx-1" @click="IsStravaTutorialOpen = true" />
      </div>
    </template>
    <div class="flex flex-col items-center gap-4 p-4">
      <div class="w-full">
        <primevue-input-text v-model="MyAthlete.ClientID" placeholder="用戶端 ID" class="w-full" />
      </div>
      <div class="w-full">
        <primevue-input-text v-model="MyAthlete.ClientSecret" placeholder="用戶端密碼" class="w-full" />
      </div>
      <div class="w-full">
        <primevue-input-text v-model="MyAthlete.AuthorizationCode" placeholder="授權代碼" class="w-full" />
      </div>

      <div class="flex items-center w-full gap-1">
        <primevue-checkbox v-model="IsStravaAgree" class="ml-2" inputId="strava_agree" binary />
        <label for="strava_agree">我已閱讀授權條款，並同意授權</label>
        <primevue-button size="small">閱讀授權條款</primevue-button>
      </div>
    </div>
    <template #footer>
      <div class="w-full flex justify-center">
        <button type="button" class="px-4 py-2 mx-2 rounded font-semibold text-white bg-sky-700 hover:bg-sky-600 transition-colors" @click="AuthorizeStrava">進行Strava授權</button>
      </div>
    </template>
  </primevue-dialog>

  <primevue-dialog v-model:visible="IsStravaTutorialOpen" modal header="授權Strava教學" class="w-10/12" :draggable="false">
    <div class="flex flex-col items-center gap-4 p-4">
      <div class="w-full text-xl font-semibold">
        1. 登入Strava，若還未有Strava帳號 <a class="underline text-blue-700 hover:text-blue-500 underline-offset-4" href="https://www.strava.com/register" target="_blank">點此註冊</a>
      </div>
      <div class="w-full">
        <div class="text-xl font-semibold">
          2. 登入後，前往
          <a class="underline text-blue-700 hover:text-blue-500 underline-offset-4" href="https://www.strava.com/settings/api" target="_blank">https://www.strava.com/settings/api</a> 取得用戶端
          ID及用戶端密碼
        </div>
        <div class="w-full flex justify-center">
          <img src="/images/strava-getting-started-1.png" />
        </div>
      </div>
      <div class="w-full">
        <div class="text-xl font-semibold">3. 輸入用戶端 ID並複製以下網址於瀏覽器開啟，勾選並點選授權</div>
        <div class="flex items-center">
          <primevue-input-text v-model="MyAthlete.ClientID" placeholder="用戶端 ID" class="mr-2" />
          <primevue-button size="small" @click="CopyStravaAuthorizationUrl">複製授權網址</primevue-button>
        </div>
        <div class="w-full flex justify-center">
          <img src="/images/strava-getting-started-2.png" />
        </div>
      </div>
      <div class="w-full">
        <div class="text-xl font-semibold">4. 複製以下授權碼並填寫授權碼</div>
        <div class="w-full flex justify-center">
          <img src="/images/strava-getting-started-3.png" />
        </div>
      </div>
    </div>
  </primevue-dialog>
</template>

<script setup>
import { ref, watch, watchEffect } from 'vue';
import { loginStore } from '../stores/LoginStore';
import { storeToRefs } from 'pinia';
import UtilityService from '../services/UtilityService';
import AthleteService from '../services/AthleteService';
import TopBar from '../components/top-bar.vue';
import Footer from '../components/home-footer.vue';

const { Alert, IsEmailValid } = UtilityService;

const store = loginStore();
const MyAthlete = ref(AthleteService.MyAthlete);
const { IsLoginDialogOpen, IsStravaDialogOpen, Email, Password, ConfirmPassword } = storeToRefs(store);
const IsRegister = ref(false); // 是否為註冊
const LoginOrLogoutKey = ref(0);
const IsStravaTutorialOpen = ref(false);
const IsStravaAgree = ref(false);
const StravaAuthorizationUrl = ref('');

function Login() {
  if (!IsEmailValid(Email.value)) return Alert('請輸入正確的電子郵件', 'error');

  store
    .login(Email.value, Password.value)
    .then(() => Alert('登入成功！', 'success'))
    .then(() => store.CloseLoginDialog())
    .catch(() => Alert('登入失敗', 'error'));
}

function Register() {
  if (!IsEmailValid(Email.value)) return Alert('請輸入正確的電子郵件', 'error');
  if (Password.value !== ConfirmPassword.value) return Alert('密碼與確認密碼不一致', 'error');

  store
    .Register(Email.value, Password.value)
    .then(() => Alert('註冊成功！', 'success'))
    .then(() => store.CloseLoginDialog())
    .catch((o) => {
      if (o.response.status === 400) {
        return Alert('此帳號已經註冊過', 'error');
      }
      return Alert('註冊失敗', 'error');
    });
}

function AuthorizeStrava() {
  if (!MyAthlete.value.ClientID) return Alert('請輸入用戶端 ID', 'error');
  if (!MyAthlete.value.ClientSecret) return Alert('請輸入用戶端密碼', 'error');
  if (!MyAthlete.value.AuthorizationCode) return Alert('請輸入授權代碼', 'error');
  if (!IsStravaAgree.value) return Alert('請閱讀並同意授權條款', 'error');

  AthleteService.MyAthlete = { ClientID: MyAthlete.value.ClientID, ClientSecret: MyAthlete.value.ClientSecret, AuthorizationCode: MyAthlete.value.AuthorizationCode };

  AthleteService.AuthorizeStrava()
    .then(() => Alert('授權Strava成功', 'success'))
    .then(() => (IsStravaDialogOpen.value = false))
    .catch(() => Alert('授權Strava失敗', 'error'));
}

function CopyStravaAuthorizationUrl() {
  navigator.clipboard.writeText(StravaAuthorizationUrl.value).then(() => Alert('已複製授權網址！'));
}

watch(
  () => store.IsLoggedIn,
  () => {
    LoginOrLogoutKey.value += 1;
  },
);

watchEffect(() => {
  StravaAuthorizationUrl.value = `https://www.strava.com/oauth/authorize?client_id=${MyAthlete.value.ClientID}&response_type=code&redirect_uri=http://localhost/exchange_token&approval_prompt=force&scope=read,activity:read_all`;
});
</script>
