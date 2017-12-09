<template>
  <div class="bg">
    <section>
      <h1>HAL CINEMA<br/>CHECKIN</h1>
      <b-field>
        <b-input
          type="email"
          placeholder="メールアドレス"
          v-model="email"
        />
      </b-field>
      <b-field>
        <b-input
          type="password"
          placeholder="パスワード"
          v-model="password"
        />
      </b-field>
      <button class="button login-button" @click="login">ログイン</button>
    </section>
  </div>
</template>

<script>
import HttpUtils from '../../lib/httpUtils'
import CookieDoc from '../../lib/CookieDoc'

export default {
  name: "chekcin-login",
  data() {
    return {
      email: "",
      password: ""
    }
  },
  methods: {
    login() {
      const email = this.email.trim()
      const password = this.password.trim()

      if(email === "") {
        this.$toast.open({
          message: 'メールアドレスを入力してください。',
          type: 'is-danger'
        })
        return
      }

      if(password === "") {
        this.$toast.open({
          message: 'パスワードを入力してください。',
          type: 'is-danger'
        })
        return
      }

      HttpUtils.UserLogin(email, password)
      .then((res) => {
        CookieDoc.setItem("halCinemaUser", res.data.token)
        this.$router.push({ path: "/checkin" })
      })
      .catch((err) => {
        this.$toast.open({
          message: "ログインに失敗しました。<br/>メールアドレスとパスワードを確認してください。",
          type: "is-danger"
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
h1 {
  color: rgb(250, 196, 27);
  font-size: 2rem;
  margin-bottom: 30px;
  font-weight: bold;
}
.bg {
  width: 100vw;
  height: 100vh;
  background-color: rgb(236, 96, 19);
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.login-button {
  background-color: #fff;
  color: rgb(236, 96, 19);
}
section {
  width: 90%;
  margin: 0 auto;
}
</style>
