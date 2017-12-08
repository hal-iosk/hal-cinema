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

  setPriceId(id) {
    this.price_id = id
  }

  getRequestData() {
    const { seats, schedule_id, price_id } = this;

    return seats.map((seat_id) => {
      return {
        seat_id,
        schedule_id,
        price_id
      }
    })
  }

  clear() {
    this.schedule_id = null;
    this.seats = [];
    this.price_id = null;
  }

}

export default new Reserve;