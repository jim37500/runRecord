import { ref } from 'vue';
import { defineStore } from 'pinia';
import axios from 'axios';
import APIService from '../services/APIService';
import AthleteService from '../services/AthleteService';
import UserService from '../services/UserService';

export const loginStore = defineStore('login', () => {
  const GoogleClientID = '446800395525-l8441okv81u7tnpeefaiobr6v27rcqa8.apps.googleusercontent.com';
  const IsLoginDialogOpen = ref(true);
  const IsStravaDialogOpen = ref(false);
  const IsLoggedIn = ref(false);
  const Email = ref('');
  const Password = ref('');
  const ConfirmPassword = ref('');

  function OpenLoginDialog() {
    Email.value = '';
    Password.value = '';
    ConfirmPassword.value = '';
    IsLoginDialogOpen.value = true;
  }

  function CloseLoginDialog() {
    IsLoginDialogOpen.value = false;
  }

  function OpenStravaDialog() {
    AthleteService.MyAthlete = {};
    IsStravaDialogOpen.value = true;
  }

  async function login(Email, password) {
    const o = await axios.post('./api/token', { Email, Password: window.sha256(password) });
    APIService.defaults.headers.common.Authorization = `Bearer ${o.data.Token}`;
    document.cookie = `JimAndyToken=${JSON.stringify(APIService.defaults.headers.common.Authorization)}`;

    await UserService.GetCurrentUser();
    IsLoggedIn.value = true;
    window.router.push('/');
  }

  function logout() {
    APIService.RemoveToken();
    AthleteService.MyAthlete = {};
    UserService.User = {};
    IsLoggedIn.value = false;
    window.router.replace('/');
  }

  function Register(Email, password) {
    return axios.post('/api/register', { Email, Password: window.sha256(password) });
  }

  return {
    GoogleClientID,
    IsLoginDialogOpen,
    IsStravaDialogOpen,
    IsLoggedIn,
    Email,
    Password,
    ConfirmPassword,
    OpenLoginDialog,
    CloseLoginDialog,
    OpenStravaDialog,
    login,
    logout,
    Register,
  };
});
