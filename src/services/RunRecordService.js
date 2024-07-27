import axios from 'axios';

export default class RunRecordService {
  // 取得跑步紀錄
  static GetRunRecord() {
    return axios.get('/api/runRecord').then((o) => o.data);
  }

  // 新增跑步紀錄
  static AddRunRecord(record) {
    return axios.post('/api/runRecord', record);
  }

  // 更新跑步紀錄
  static UpdateRunRecord(ID, record) {
    return axios.put(`/api/runRecord/${ID}`, record);
  }

  // 刪除跑步紀錄
  static DeleteRunRecord(ID) {
    return axios.delete(`/api/runRecord/${ID}`);
  }
}
