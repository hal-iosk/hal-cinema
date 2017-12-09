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
import vueStore from '../../vuex'
import HttpUtils from '../../lib/httpUtils'

export default {
  name: "buy-modal",
  data() {
    return {
      changed: false
    }
  },
  methods: {
    change() {
      if(vueStore.state.coin < 3) {
        this.$dialog.alert({
          title: 'ポップコーン交換',
          message: 'コインが足りません！<br/>3コイン以上で交換できます。',
          type: 'is-danger',
          confirmText: 'OK'
        })
        return
      }
      HttpUtils.PostPopcorn()
      .then((res) => {
        this.changed = true;
        vueStore.commit("coinUpdate", res.data.point)
      })
      .catch((err) => console.error(err))
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
