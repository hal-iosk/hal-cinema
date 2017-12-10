<template>
  <section>

    <reserve-nav></reserve-nav>

    <movie-content
      :title="movie.title"
      :date="movie.days"
      :theater="movie.theater_number"
    ></movie-content>

    <div class="seat-container">
      <div class="screen"></div>
      <div class="screen-name"></div>
      <ul>
        <li v-for="seat in seats" @click="seatSelect(seat.seatSymbol)">
          <div class="seat-icon purchased" v-if="seat.isSelected"></div>
          <div class="seat-icon" v-else></div>
        </li>
      </ul>
    </div>

    <div class="controller">
      <button class="button" @click="back">戻る</button>
      <button class="button is-primary" @click="next">次へ</button>
    </div>

  </section>
</template>

<script>
import ReserveNav from './nav.vue'
import seatsFormat from './seat.format'
import reservePayload from '../../lib/reserve.class'
import MovieContent from './movieContent.vue'

export default {
  name: "reserve",
  components: {
    ReserveNav,
    MovieContent
  },
  mounted() {
    const movie = JSON.parse(sessionStorage.getItem("halCinemaReserve"))
    movie ? this.movie = movie : location.href = "/"
    reservePayload.clear()
  },
  data() {
    return {
      seats: seatsFormat,
      movie: {}
    }
  },
  methods: {
    seatSelect(seatId) {
      seatsFormat.map((seat) => {
        if(seatId === seat.seatSymbol) seat.isSelected = !seat.isSelected;
      })
    },
    back() {
      location.href = "/watchFilm"
    },
    next() {
      const selectSeats = this.seats.filter((seat) => { if(seat.isSelected) { return seat.seatSymbol } })
      const reservedSeats = selectSeats.map((seat) => { return seat.seatSymbol })

      if(Object.keys(reservedSeats).length <= 0) {
        this.$dialog.alert({
          message: "座席を選択してください。",
          type: "is-danger"
        })
        return
      }

      reservePayload.setSeats(reservedSeats)
      this.$router.push({ path: "/reserve/ticket" })
    }
  }
}
</script>

<style lang="scss" scoped>
.movie-reserve-container {
  display: flex;
  justify-content: space-around;
  margin: 10px 0;
  li {
    display: flex;
    padding: 10px;
    align-items: center;
  }
  .content-title {
    margin-right: 10px;
  }
  .content {
    font-weight: bold;
    font-size: 1.3rem;
  }
}
.seat-container {
  background-color: #6B6A6A;
  border: 5px solid #000;
  margin: 2.5rem auto;
  padding: 1.5rem;
  width: 442px;
}
.seat-container .screen {
  width: 80%;
  height: 7px;
  background-color: #FFF;
  margin: 0.5rem auto;
}
.seat-container .screen-name {
  text-align: center;
  font-size: 1.2rem;
  color: #FFF;
  padding-bottom: 3rem;
}
.seat-container ul {
  display: flex;
  flex-wrap: wrap;
  list-style: none;
}
.seat-container ul li div {
  display: block;
  width: 20px;
  height: 30px;
  margin: 5px 2px;
  text-align: center;
  color: #FFF;
}
.seat-container ul .seat-icon {
  background-color: #FFF;
  cursor: pointer;
}
.seat-container ul .seat-icon.selected {
  background-color: #D0011B;
}
.seat-container ul .seat-icon.purchased {
  background-color: #4A4A4A;
}
.controller {
  width: 50%;
  margin: 0 auto 50px;
  display: flex;
  justify-content: space-around;
}
</style>
