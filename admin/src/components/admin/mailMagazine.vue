<template>
  <section>
    <b-field label="本文">
      <b-input maxlength="200" type="textarea"></b-input>
    </b-field>
    <button class="button is-primary" @click="toast">送信</button>
  </section>
</template>

<script>
export default {
  name: "mailMagazine",
  methods: {
    toast() {
      this.$toast.open('メールマガジンを送信しました。')
    }
  }
}
</script>

<style lang="scss" scoped>
section {
  width: 90%;
  margin: 10px auto;
}
</style>

