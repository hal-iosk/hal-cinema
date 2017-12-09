<template>
  <section>

    <ul class="navigation">
      <li class="navigation-item">座席選択</li>
      <li>></li>
      <li class="navigation-item">チケット選択</li>
      <li>></li>
      <li class="navigation-item isActive">購入内容の確認</li>
    </ul>

    <h1>予約完了</h1>
  </section>
</template>

<script>
import reservePayload from '../../lib/reserve.class'
import ReserveNav from './nav.vue'

export default {
  name: "reserve-confirm",
  mounted() {
    if(reservePayload.getRequestData().length <= 0) this.$router.push({ path: "/reserve" });
    console.log(reservePayload.getRequestData())
  },
  components: {
    ReserveNav
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
</style>
