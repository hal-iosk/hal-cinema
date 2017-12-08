<template>
  <section>
    <p>ポップコーンにコインを交換しよう！</p>
    <div class="img">
      <img src="/image/popcorn.jpg" alt="ポップコーンのイラスト" v-if="!changed">
      <img src="/image/qr.gif" alt="QRコード" v-else>
    </div>

    <button class="button is-warning change" v-if="!changed" @click="change">交換する！</button>
  </section>
</template>

<script>
export default {
  name: "buy-modal",
  data() {
    return {
      changed: false
    }
  },
  methods: {
    change() {
      this.changed = true;
    }
  }
}
</script>

<style lang="scss" scoped>
section {
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
}
p {
  font-size: 0.8rem;
  padding: 10px 0px;
}
.img {
  width: 60%;
  margin: 50px auto;
  img {
    display: block;
    width: 100%;
  }
}
.change {
  width: 100%;
}
</style>
