import APIService from './APIService';

export default class AthleteService {
  static MyAthlete = {};

  static AuthorizeStrava() {
    this.MyAthlete.ClientID = Number(this.MyAthlete.ClientID);
    return APIService.post('athlete', this.MyAthlete);
  }
}
