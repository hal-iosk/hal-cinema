<template>
  <section>

    <div class="flex">
      <h1>映画編集</h1>
      <div style="margin-left: auto;">
        <button class="button" @click="back">戻る</button>
        <button class="button is-danger" @click="deleteMovie">削除</button>
        <button class="button is-primary" @click="complate">完了</button>
      </div>
    </div>

    <b-field label="タイトル">
      <b-input v-model="movie_name"></b-input>
    </b-field>

    <b-field label="詳細">
      <b-input maxlength="200" type="textarea" v-model="details"></b-input>
    </b-field>

    <div class="thumbnail">
      <p class="thumbnail-title">サムネイル</p>
      <div>
        <b-input v-model="image_path" class="thumbnail-input"></b-input>
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
      <b-input type="number" v-model="watch_time"></b-input>
    </b-field>

     <b-modal :active.sync="isImageModalActive">
        <p class="image is-4by3">
          <img :src="image_path">
        </p>
      </b-modal>

      <b-loading :active.sync="isLoading" :canCancel="true"></b-loading>

  </section>
</template>

<script>
import httpUtils from '../../lib/httpUtils'
import validationUtils from '../../lib/validationUtils'

export default {
  name: "movieTableEdit",
  data() {
    return {
      isLoading: false,
      movie_name: "",
      details: "",
      image_path: "",
      start_time: new Date(),
      end_time: new Date(),
      watch_time: 0,
      isImageModalActive: false,
    }
  },
  mounted() {
    this.isLoading = true;
    const id = this.$route.params.id;
    httpUtils.GetMovieDetail(id)
    .then((res) => {
      this.isLoading = false;

      this.movie = res.data;
      this.movie_name = res.data.movie_name
      this.details = res.data.details
      this.image_path = res.data.image_path
      console.log(res.data.start_date)
      this.start_time = new Date(res.data.start_date)
      this.end_time = new Date(res.data.end_date)
      this.watch_time = res.data.watch_time
    })
    .catch((err) => console.error(err))
  },
  methods: {
    complate() {

      // validation
      if(validationUtils.isBlank(this.movie_name)) { validationUtils.pushMessage("映画名を入力してください。", this); return }
      if(validationUtils.isBlank(this.details)) { validationUtils.pushMessage("映画詳細を入力してください。", this); return }
      if(validationUtils.isBlank(this.image_path)) { validationUtils.pushMessage("サムネイルを入力してください。", this); return }
      if(!validationUtils.isURL(this.image_path)) { validationUtils.pushMessage("サムネイルがURLの形式ではありません。", this); return }
      if(!validationUtils.isNumber(String(this.watch_time))) { validationUtils.pushMessage("上映時間を数字で入力してください。", this); return }

      httpUtils.PutMovieDetail(
        this.$route.params.id,
        this.movie_name,
        this.details,
        this.image_path,
        this.start_time,
        this.end_time,
        this.watch_time
      )
      .then((res) => {
        if(res.status === 200) {
          this.$toast.open({
            message: '編集完了しました。',
            type: 'is-success'
          })
          setTimeout(() => this.$router.push({ path: `/admin` }), 500)
        }
      })
      .catch((err) => {
        console.error(err)
      })
    },
    deleteMovie() {
      this.$dialog.confirm({
        title: '映画削除',
        message: '映画を削除します。よろしいですか？',
        confirmText: '削除',
        type: 'is-danger',
        onConfirm: () => {
          httpUtils.DeleteMovieDetail(this.$route.params.id)
          .then((res) => {
            if(res.status === 204) {
              this.$toast.open({
                message: '削除完了しました。',
                type: 'is-success'
              })
              setTimeout(() => this.$router.push({ path: `/admin` }), 500)
            }
          })
          .catch((err) => console.error(err))
        }
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

