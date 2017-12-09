class Reserve {

  constructor() {
    this.price_id = null;
    this.schedule_id = null;
    this.seats = [];
  }

  setScheduleId(id) {
    this.schedule_id = id
  }

  setSeats(ids) {
    this.seats = ids
  }

  getSeats() {
    return this.seats
  }

  setPriceId(id) {
    this.price_id = id
  }

  clear() {
    this.schedule_id = null;
    this.seats = [];
    this.price_id = null;
  }

}

export default new Reserve;