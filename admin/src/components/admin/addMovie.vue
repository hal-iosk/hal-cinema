<template>
  <section>

    <div class="flex">
      <h1>映画追加</h1>
      <div style="margin-left: auto;">
        <button class="button"@click="back">戻る</button>
        <button class="button is-primary" @click="add">追加する</button>
      </div>
    </div>

    <b-field label="タイトル">
      <b-input v-model="movie.movie_name"></b-input>
    </b-field>

    <b-field label="詳細">
      <b-input maxlength="200" type="textarea" v-model="movie.details"></b-input>
    </b-field>

    <div class="thumbnail">
      <p class="thumbnail-title">サムネイル</p>
      <div>
        <b-input v-model="movie.image_path" class="thumbnail-input"></b-input>
        <p class="control">
          <button class="button"　@click="isImageModalActive = true">確認する</button>
        </p>
      </div>
    </div>

    <div class="date">
      <b-field label="公開開始日">
        <b-datepicker v-model="start_time" inline></b-datepicker>
      </b-field>

      <p style="font-size: 2rem;">〜</p>

      <b-field label="公開終了日">
        <b-datepicker v-model="end_time" inline></b-datepicker>
      </b-field>
    </div>

    <b-field label="上映時間" class="time">
      <b-input type="number" v-model="movie.watch_time"></b-input>
    </b-field>

     <b-modal :active.sync="isImageModalActive">
        <p class="image is-4by3">
          <img :src="movie.image_path">
        </p>
      </b-modal>

  </section>
</template>

<script>
import httpUtils from '../../lib/httpUtils'

export default {
  name: "addMovie",
  data() {
    return {
      movie: {},
      start_time: "",
      end_time: "",
      isImageModalActive: false,
    }
  },
  methods: {
    add() {
      httpUtils.CreateMovie(
        this.movie.movie_name,
        this.movie.details,
        this.movie.image_path,
        this.start_time,
        this.end_time,
        this.movie.watch_time
      )
      .then((res) => {
        if(res.status === 201) {
          this.$toast.open({
            message: '映画を追加しました。',
            type: 'is-success'
          })
          setTimeout(() => {
            this.$router.push({ path: `/admin` });
          }, 500)
        }
      })
      .catch((err) => {
        console.error(err)
      })
    },
    back() {
      this.$router.push({ path: `/admin` });
    }
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
.thumbnail {
  margin-bottom: 50px;
  div {
    display: flex;
  }
  .thumbnail-title {
    font-weight: bold;
    margin-bottom: 7px;
  }
  .thumbnail-input {
    width: 50%;
  }
}
</style>

