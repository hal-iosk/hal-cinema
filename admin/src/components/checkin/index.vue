<template>
  <section class="bg">

    <div class="circle" @click="check">
      <img src="/image/checkin.svg" alt="checkin-icon">
    </div>

    <div class="footer">
      <button class="button point is-warning" @click="showModal">{{coin}}コイン</button>
      <div class="user-info" @click="reset">
        <p class="user-info-name">konojunya</p>
      </div>
    </div>

  </section>
</template>

<script>
import moment from 'moment'
import BuyModal from './buyModal.vue'
import vueStore from '../../vuex'
import HttpUtils from '../../lib/httpUtils'

export default {
  name: "chekcin",
  date() {
    return {
      canGeoGet: false
    }
  },
  computed: {
    coin() {
      return vueStore.state.coin
    }
  },
  methods: {
    showModal() {
      this.$modal.open({
        parent: this,
        component: BuyModal,
        hasModalCard: false
      })
    },
    reset() {
      localStorage.removeItem("hal-cinema-checkin-day")
    },
    chekin() {

      HttpUtils.PostCheckin()
      .then((res) => {
        let coin = res.data.point

        this.$dialog.alert({
            title: "チェックイン",
            message: `<p style="font-size: 1rem; font-weight: bold;">おめでとうございます！</p><br/><p style="font-size: 0.9rem;">${coin - 1}コイン → ${coin}コインになりました！</p>`,
            confirmText: 'OK',
            type: 'is-warning'
        })
        vueStore.commit("coinUpdate", coin)
        localStorage.setItem("hal-cinema-checkin-day", moment(new Date()))
      })
      .catch((err) => console.error(err))
    },
    check() {
      const lastDay = localStorage.getItem("hal-cinema-checkin-day")
      const diffDay = moment(new Date()).diff(lastDay, "days")

      if(diffDay < 1) {
        this.$dialog.alert({
          title: 'チェックイン',
          message: '1日に2回<br/>チェックインできません。',
          type: 'is-danger',
          confirmText: 'OK'
        })
        return
      }

      if(!this.canGeoGet) {
        navigator.geolocation.getCurrentPosition(() => {
          this.canGeoGet = true
          this.chekin()
        })
        return
      }
      this.chekin()
    }
  }
}
</script>

<style lang="scss" scoped>
.bg {
  width: 100vw;
  height: 100vh;
  background-color: rgb(236, 96, 19);
  display: flex;
  justify-content: center;
  align-items: center;
}
.circle {
  width: 70vw;
  height: 70vw;
  background-color: white;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 5px solid rgb(250, 196, 27);
  box-shadow: 0 2px 2px 0 rgba(0,0,0,0.14), 0 3px 1px -2px rgba(0,0,0,0.2), 0 1px 5px 0 rgba(0,0,0,0.12);
  img {
    width: 30%;
    display: block;
  }
}
.footer {
  width: 100vw;
  position: fixed;
  bottom: 0;
  left: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: rgb(236, 96, 19);
  padding: 15px 10px;
  background-color: #fff;
}
</style>
