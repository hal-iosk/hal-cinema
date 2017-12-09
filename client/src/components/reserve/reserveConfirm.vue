<template>
  <section>

    <ul class="navigation">
      <li class="navigation-item">座席選択</li>
      <li>></li>
      <li class="navigation-item">チケット選択</li>
      <li>></li>
      <li class="navigation-item isActive">購入内容の確認</li>
    </ul>

    <movie-content
      :title="movie.title"
      :date="movie.days"
      :theater="movie.theater_number"
    ></movie-content>

    <ul class="container">
      <li v-for="seat in seats">
        <div class="img">
          <img src="/image/seat.png" alt="シートイラスト">
        </div>
        <p class="seat-id">{{seat.seatId}}</p>
        <p class="price">{{seat.customer_type}} / {{seat.price}}円</p>
      </li>
    </ul>

    <hr>

    <div class="sum-box">
      <p class="sum">合計: {{sum}}円</p>
    </div>

    <div class="controller">
      <button class="button" @click="back">最初に戻る</button>
      <button class="button is-primary" @click="reserveComplete">予約を確定する</button>
    </div>

  </section>
</template>

<script>
import reservePayload from '../../lib/reserve.class'
import ReserveNav from './nav.vue'
import MovieContent from './movieContent.vue'

export default {
  name: "reserve-confirm",
  mounted() {
    if(reservePayload.getSeats().length <= 0) this.$router.push({ path: "/reserve" })

    const seatObj = JSON.parse(sessionStorage.getItem("halCinemaReserveSeats"))
    this.seats = Object.keys(seatObj).map((seat) => { return seatObj[seat] })
    this.seats.map((seat) => { this.sum += parseInt(seat.price, 10) })
  },
  data() {
    return {
      seats: [],
      movie: JSON.parse(sessionStorage.getItem("halCinemaReserve")),
      sum: 0
    }
  },
  methods: {
    back() {
      this.$dialog.confirm({
        message: '全ての操作を取り消して、映画一覧へ戻ります。<br/>よろしいですか？',
        onConfirm: () => { this.clear(); location.href = "/watchFilm" }
      })
    },
    reserveComplete() {
      this.$toast.open({
        message: "予約しました。",
        type: "is-success"
      })
      this.clear()
      setTimeout(() => {
        location.href = "/"
      }, 500)
    },
    clear() {
      sessionStorage.removeItem("halCinemaReserveSeats")
      sessionStorage.removeItem("halCinemaReserve")
    }
  },
  components: {
    ReserveNav,
    MovieContent
  }
}
</script>

<style lang="scss" scoped>
.navigation {
  display: flex;
  justify-content: space-around;
  align-items: center;
  padding: 30px 20%;
  box-shadow: 0px 0px 1px black;
  li.navigation-item {
    letter-spacing: 1.5px;
  }
  li.isActive { font-weight: bold; color: #B71C1C; box-shadow: 0px 0px 1px #B71C1C; }
  li {
    padding: 10px;
  }
}
.controller {
  width: 50%;
  margin: 0 auto 50px;
  display: flex;
  justify-content: space-around;
}
.container {
  width: 80%;
  margin: 20px auto;
  padding: 30px 0;
  li {
    margin: 30px 0;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    .img {
      width: 50px;
      img {
        width: 100%;
        display: block;
      }
    }
    p {
      font-size: 1.5rem;
      margin-left: 50px;
    }
    .seat-id {
      margin-right: auto;
    }
  }
}
.sum-box {
  width: 80%;
  margin: 30px auto;
  display: flex;
  justify-content: flex-end;
  .sum {
    font-size: 1.5rem;
  }
}
</style>
