<template>
  <div class="bg">

    <section>

      <h1>HAL CINEMA Admin</h1>

        <b-field
          label="Email"
        >
          <b-input
            type="email"
            maxlength="30"
          />
        </b-field>

        <b-field
          label="Password"
        >
          <b-input
            type="password"
            maxlength="30"
          />
        </b-field>

        <button class="button is-primary login-button" @click="openLoading">ログイン</button>

    </section>

    <b-loading :active.sync="isLoading" />

  </div>

</template>

<script>
import httpUtils from '../../lib/httpUtils'

export default {
  name: "login",
  data() {
    return {
      isLoading: false
    }
  },
  methods: {
    openLoading() {

      ( async () => {
        const vm = this;
        vm.isLoading = true
        const res = await httpUtils.Login();
        setTimeout(() => {
          vm.isLoading = false;
        }, 5 * 1000)
      })()

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
