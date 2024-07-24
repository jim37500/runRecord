export default class UtilityService {
  // 提示訊息
  static Alert(title, icon = 'info', timer = 1000) {
    return window.Swal.fire({ icon, title, timer });
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
}
