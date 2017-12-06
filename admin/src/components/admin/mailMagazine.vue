<template>
  <section>
    <button class="button" @click="isCardModalActive = true">メルマガ作成</button>

    <b-modal :active.sync="isCardModalActive" :width="640" scroll="keep">
      <div class="card">
        <b-field label="本文">
          <b-input maxlength="200" type="textarea"></b-input>
        </b-field>
        <button class="button is-primary" @click="toast">送信</button>
      </div>
    </b-modal>
  </section>
</template>

<script>
export default {
  name: "mailMagazine",
  data() {
    return {
      isCardModalActive: false
    }
  },
  methods: {
    toast() {
      this.isCardModalActive = false;
      this.$toast.open('メールマガジンを送信しました。')
    }
  }
}
</script>
