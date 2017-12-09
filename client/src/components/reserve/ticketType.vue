<template>
  <section>

    <ul class="navigation">
      <li class="navigation-item">座席選択</li>
      <li>></li>
      <li class="navigation-item isActive">チケット選択</li>
      <li>></li>
      <li class="navigation-item">購入内容の確認</li>
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
        <p>{{seat.id}}</p>
        <b-select
            placeholder="券種を選んでください。"
            size="is-medium"
            expanded
            required
            @input="selected"
        >
          <option :value="`${seat.id}-${price.ID}-${price.price}`" v-for="price in prices">{{price.customer_type}}</option>
        </b-select>
      </li>
    </ul>

    <div class="controller">
      <button class="button" @click="back">戻る</button>
      <button class="button is-primary" @click="next">次へ</button>
    </div>

  </section>
</template>

<script>
import reservePayload from '../../lib/reserve.class'
import ReserveHttp from '../../services/reserve'
import MovieContent from './movieContent.vue'

const selectObj = {}

export default {
  name: "ticketType",
  components: {
    MovieContent
  },
  mounted() {
    if(reservePayload.getRequestData().length <= 0) this.$router.push({ path: "/reserve" });

    if(this.prices.length <= 0) {
      ReserveHttp.GetPrices()
      .then((res) => {
        this.prices = res.data.prices
      })
    }
  },
  data() {
    return {
      seats: reservePayload.getSeats().map((seat) => {
        return {
          id: seat,
          ticketType: null
        }
      }),
      movie: JSON.parse(sessionStorage.getItem("halCinemaReserve")),
      prices: []
    }
  },
  methods: {
    selected(e) {
      const [seat, priceId, price] = e.split("-")
      selectObj[seat] = {priceId, price}
    },
    back() {
      reservePayload.clear()
      location.href = "/reserve"
    },
    next() {
      console.log()
      console.log(selectObj)
      // this.$router.push({ path: "/reserve/complete" })
    }
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
    justify-content: space-between;
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
      margin-right: auto;
      margin-left: 50px;
    }
  }
}
</style>