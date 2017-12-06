<template>
  <div class="bg">

    <section>

      <h1>HAL CINEMA Admin</h1>

        <form action="/api/admin/signin" method="post">
          <b-field
            label="Email"
          >
            <b-input
              type="email"
              maxlength="30"
              v-model="email"
            />
          </b-field>

          <b-field
            label="Password"
          >
            <b-input
              type="password"
              maxlength="30"
              v-model="password"
            />
          </b-field>

          <button class="button is-primary login-button" type="submit" @click="openLoading">ログイン</button>
        </form>

    </section>

  </div>

</template>

<script>
import httpUtils from '../../lib/httpUtils'
import cookie from 'cookie'

export default {
  name: "login",
  data() {
    return {
      email: "",
      password: ""
    }
  },
  methods: {
    openLoading() {
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

      httpUtils.Login(email, password)
      .then((res) => {
        const token = res.data.token
        document.cookie = cookie.serialize("halCinemaAdmin", token)
        location.href = "/admin";
      })
      .catch((err) => {
        console.error(err);
      });
    }
  }
}
</script>

<style lang="scss" scoped>
h1 {
  color: white;
  font-size: 2rem;
  margin-bottom: 30px;
}
.bg {
  width: 100vw;
  height: 100vh;
  background-color: rgb(112, 70, 209);
  display: flex;
  justify-content: center;
  align-items: center;
}
section {
  width: 30%;
}
.login-button {
  background-color: #fff;
  color: rgb(112, 70, 209);
}
</style>
