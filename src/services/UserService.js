import APIService from './APIService';
import AthleteService from './AthleteService';
import { loginStore } from '../stores/LoginStore';

export default class UserService {
  static User = ''; // 使用者帳號

  // 取得 當前使用者資訊
  static async GetCurrentUser() {
    const o = await APIService.get('currentuser');

    APIService.defaults.headers.common.Authorization = `Bearer ${o.data.Token.Token}`;
    UserService.User = o.data.Account; // 使用者帳號
    document.cookie = `JimAndyToken=${JSON.stringify(APIService.defaults.headers.common.Authorization)}`;

    const store = loginStore();
    store.IsLoggedIn = true;

    // 若 有儲存權杖 則 更新權杖
    if (localStorage.token) localStorage.token = `Bearer ${o.data.Token.Token}`;
    AthleteService.MyAthlete = o.data.Athlete;
  }
}
