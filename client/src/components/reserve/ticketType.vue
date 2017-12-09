<template>
  <section>

    <reserve-nav></reserve-nav>

    <movie-content
      :title="movieContent.title"
      :date="movieContent.date"
      :theater="movieContent.theater"
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
            @select="selected"
        >
          <option value="1">Flint</option>
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
import ReserveNav from './nav.vue'
import MovieContent from './movieContent.vue'

export default {
  name: "ticketType",
  mounted() {
    if(reservePayload.getRequestData().length <= 0) this.$router.push({ path: "/reserve" });
  },
  data() {
    return {
      movieContent: {
        title: "あさひなぐ",
        date: "2017/09/26 10:00 〜 11:00",
        theater: "1"
      },
      seats: reservePayload.getSeats().map((seat) => {
        return {
          id: seat,
          ticketType: null
        }
      })
    }
  },
  methods: {
    selected() {
      console.log("hoge")
    },
    back() {
      reservePayload.clear()
      location.href = "/reserve"
    },
    next() {
      console.log(this.selectTicket)
      // reservePayload.setPriceId(1)
      // this.$router.push({ path: "/reserve/complete" })
    }
  },
  components: {
    ReserveNav,
    MovieContent
  },
}
</script>

<style lang="scss" scoped>
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