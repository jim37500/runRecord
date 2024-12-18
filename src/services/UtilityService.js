export default class UtilityService {
  // 提示訊息
  static Alert(title, icon = 'info', timer = 1000) {
    return window.Swal.fire({ title, icon, timer });
  }

  // 確認對話框
  static Confirm(title, text) {
    return window.Swal.fire({
      icon: 'question',
      title,
      text,
      showCancelButton: true,
      confirmButtonText: '確定',
      cancelButtonText: '取消',
    });
  }

  // 檢查Email格式
  static IsEmailValid(email) {
    return /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(String(email).toLowerCase());
  }
}
