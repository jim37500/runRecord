import axios from 'axios';

export default class ActivityService {
  // 取得活動
  static GetActivities(athleteID) {
    return axios.get(`/api/activities/${athleteID}`).then((o) => o.data);
  }

  // 取得活動圈數
  static GetActivityLaps(athleteID, type, activityID) {
    return axios.get(`/api/activities/laps/${athleteID}/${type}/${activityID}`).then((o) => o.data);
  }
}
