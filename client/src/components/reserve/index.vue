<template>
  <section>

    <!-- <reserve-nav></reserve-nav> -->

    <h1>座席選択</h1>

    <ul>
      <li>
        <p>作品名</p>
        <p>あさひなぐ</p>
      </li>
      <li>
        <p>日時</p>
        <p>2017/09/26 16:20 ~ 18:50</p>
      </li>
      <li>
        <p>シアター番号</p>
        <p>1</p>
      </li>
    </ul>

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

    <div>
      <button class="button" @click="back">戻る</button>
      <button class="button is-primary" @click="next">次へ</button>
    </div>

  </section>
</template>

<script>
import ReserveNav from './nav.vue'
import seatsFormat from './seat.format'

let reservedSeats = []

export default {
  name: "reserve",
  components: {
    ReserveNav
  },
  data() {
    return {
      seats: seatsFormat
    }
  },
  methods: {
    seatSelect(seatId) {
      reservedSeats.push(seatId)
      seatsFormat.map((seat) => {
        if(seatId === seat.seatSymbol) seat.isSelected = !seat.isSelected;
      })
    },
    back() {
      alert("映画一覧に戻る")
    },
    next() {
      console.log("予約席: ", reservedSeats)
      console.log("スケジュールID:", 1)
      this.$router.push({ path: "/reserve/ticket" })
    }
  }
}
</script>

<style lang="scss" scoped>
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
</style>
