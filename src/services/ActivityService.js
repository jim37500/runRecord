import APIService from './APIService';

export default class ActivityService {
  static ActivityCache = null;
  static ActivityLapsCache = {};
  static SportTypes = {
    Run: { Name: '跑步', Value: 'Run', Icon: '🏃', TimeConstant: 60, Unit: 'min / km', IsDisplay: true },
    Ride: { Name: '騎車', Value: 'Ride', Icon: '🚴', TimeConstant: 1, Unit: 'km/hr', IsDisplay: true },
    Swim: { Name: '游泳', Value: 'Swim', Icon: '🏊', TimeConstant: 6, Unit: 'min / 100m', IsDisplay: true },
    Hike: { Name: '健行', Value: 'Hike', Icon: '🧗', TimeConstant: 1, Unit: 'km/hr', IsDisplay: false },
  };

  // 取得活動
  static GetActivities(athleteID) {
    if (this.ActivityCache) return Promise.resolve(this.ActivityCache);

    return APIService.get(`activities/${athleteID}`).then((o) => {
      this.ActivityCache = o.data;
      return o.data;
    });
  }

  // 取得活動圈數
  static GetActivityLaps(athleteID, activityID) {
    if (this.ActivityLapsCache[activityID]) return Promise.resolve(this.ActivityLapsCache[activityID]);

    return APIService.get(`activities/laps/${athleteID}/${activityID}`).then((o) => {
      this.ActivityLapsCache[activityID] = o.data;
      return o.data;
    });
  }

  // 換算距離(m -> km)
  static CalculateDistance(distance) {
    return Math.round(distance / 100) / 10;
  }

  // 換算速度
  static CalculateSpeed(speed, type) {
    // 若 為騎車 則 直接回傳時速
    if (type === 'Ride') return `${(Math.round(speed * 10) / 10).toFixed(1)}`;
    // 運動類型相關常數
    const pace = Math.round((this.SportTypes[type].TimeConstant / speed) * 100) / 100; // 配速
    const decimal = pace - Math.floor(pace); // 取小數部分(100 -> 60)

    return `${Math.floor(pace)}:${Math.round(decimal * 0.6 * 100)
      .toString()
      .padStart(2, '0')}`;
  }

  // 換算時間(sec -> hr:min:sec)
  static CalculateTime(secs) {
    const seconds = Math.round(secs) % 60;
    let minutes = Math.floor(secs / 60);
    const hours = Math.floor(minutes / 60);

    if (hours > 0) {
      minutes %= 60;
      return `${hours}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    }

    return `${minutes}:${seconds.toString().padStart(2, '0')}`;
  }
}
