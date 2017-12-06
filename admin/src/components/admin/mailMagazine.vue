<template>
  <section>
    <b-field label="本文">
      <b-input maxlength="200" type="textarea" v-model.trim="text"></b-input>
    </b-field>
    <button class="button is-primary" @click="send">送信</button>
  </section>
</template>

<script>
export default {
  name: "mailMagazine",
  data() {
    return {
      text: ""
    }
  },
  methods: {
    send() {
      const text = this.text.trim()
      if(text === "") {
        this.$toast.open({
          message: '文字を入力してください。',
          type: "is-danger"
        });
        return;
      }
      this.$toast.open({
        message: 'メールマガジンを送信しました。',
        type: "is-success"
      });
      this.text = ""
    }
  }
}
</script>

<style lang="scss" scoped>
section {
  width: 90%;
  margin: 20px auto;
}
</style>

