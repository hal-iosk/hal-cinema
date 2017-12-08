<template>
  <section>
    <!-- <reserve-nav></reserve-nav> -->

    <h1>チケット券種選択</h1>

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

    <ul>
      <li>
        <div class="img">
          <img src="/image/seat.png" alt="シートイラスト">
        </div>
        <p>N1</p>
        <b-select
            placeholder="券種を選んでください。"
            size="is-medium"
            expanded
            required
        >
          <option value="flint">Flint</option>
          <option value="silver">Silver</option>
        </b-select>
      </li>
    </ul>

    <div>
      <button class="button" @click="back">戻る</button>
      <button class="button is-primary" @click="next">次へ</button>
    </div>
  </section>
</template>

<script>
export default {
  name: "ticketType",
  methods: {
    back() {
      this.$router.push({ path: "/reserve" })
    },
    next() {
      this.$router.push({ path: "/reserve/complete" })
    }
  }
}
</script>
