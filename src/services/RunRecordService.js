import axios from 'axios';

export default class RunRecordService {
  // 取得跑步紀錄
  static GetRunRecord(athleteID) {
    return axios.get(`/api/activities/${athleteID}`).then((o) => o.data);
  }
}
