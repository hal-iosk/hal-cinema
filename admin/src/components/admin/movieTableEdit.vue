<template>
  <section>

    <div class="flex">
      <h1>映画編集</h1>
      <button class="comp-button button is-primary">完了</button>
    </div>

    <b-field label="タイトル">
      <b-input></b-input>
    </b-field>

    <b-field label="詳細">
      <b-input maxlength="200" type="textarea"></b-input>
    </b-field>

    <div class="date">
      <b-field label="公開開始日">
        <b-datepicker v-model="date" inline></b-datepicker>
      </b-field>

      <p style="font-size: 2rem;">〜</p>

      <b-field label="公開終了日">
        <b-datepicker v-model="date" inline></b-datepicker>
      </b-field>
    </div>

    <b-field label="上映時間" class="time">
      <b-input type="number"></b-input>
    </b-field>

  </section>
</template>

<script>
import httpUtils from '../../lib/httpUtils'

export default {
  name: "movieTableEdit",
  data() {
    return {
      movie: {}
    }
  },
  mounted() {
    const id = this.$route.params.id;
    httpUtils.GetMovieDetail(id)
    .then((res) => {
      console.log(res.data)
      // this.movie = res.data.movies;
    })
    .catch((err) => {
      console.error(err)
    })
  }
}
</script>

<style lang="scss" scoped>
section {
  width: 90%;
  margin: 30px auto;
}
h1 {
  font-size: 2rem;
}
.date {
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.time {
  width: 30%;
}
.flex {
  margin-bottom: 50px;
  display: flex;
  align-items: center;
}
.comp-button {
  margin-left: auto;
}
</style>

