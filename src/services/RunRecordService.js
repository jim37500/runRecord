import axios from 'axios';

export default class RunRecordService {
  // 取得跑步紀錄
  static GetRunRecord() {
    return axios.get('/api/runRecord').then((o) => o.data);
  }

  // 新增跑步紀錄
  static AddRunRecord(Date, TotalDistance, TotalTime, AvgPace, Temperature, TrainingType, Detail, Feeling) {
    return axios.post('/api/runRecord', { Date, TotalDistance, TotalTime, AvgPace, Temperature, TrainingType, Detail, Feeling });
  }

  // 更新跑步紀錄
  static UpdateRunRecord(ID, Date, TotalDistance, TotalTime, AvgPace, Temperature, TrainingType, Detail, Feeling) {
    return axios.put(`/api/runRecord/${ID}`, { Date, TotalDistance, TotalTime, AvgPace, Temperature, TrainingType, Detail, Feeling });
  }

  // 刪除跑步紀錄
  static DeleteRunRecord(ID) {
    return axios.delete(`/api/runRecord/${ID}`);
  }
}
