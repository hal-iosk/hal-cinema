<template>
  <div class="bg">
    <section>
      <h1>HAL CINEMA<br/>CHECKIN</h1>
      <b-field>
        <b-input
          type="email"
          placeholder="メールアドレス"
        />
      </b-field>
      <b-field>
        <b-input
          type="password"
          placeholder="パスワード"
        />
      </b-field>
      <button class="button login-button">ログイン</button>
    </section>
  </div>
</template>

<script>
export default {
  name: "chekcin-login"
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
