<template>
  <section>

    <div class="flex">
      <div v-if="title.length > 8">
        <b-tooltip type="is-dark" :label="title" position="is-bottom" multilined>
          <h1 style="font-size: 1.3rem;">「{{title.slice(0, 7) + "..."}}」スケジュール追加</h1>
        </b-tooltip>
      </div>
      <div v-else>
        <h1 style="font-size: 1.3rem;">「{{title}}」スケジュール追加</h1>
      </div>
      <div style="margin-left: auto;">
        <button class="button"@click="back">戻る</button>
        <button class="button is-primary" @click="add">追加する</button>
      </div>
    </div>

    <b-field label="シアターナンバー" class="theater-number">
      <b-input v-model="theater_number" type="number" min="1" max="10"></b-input>
    </b-field>

    <div class="flex-box">
      <b-datepicker v-model="start_time" inline></b-datepicker>

      <b-field label="映画開始時間">
        <b-timepicker v-model="start_time" inline></b-timepicker>
      </b-field>

      <p style="font-size: 2rem;">〜</p>

      <b-field label="映画終了時間">
        <b-timepicker v-model="end_time" inline></b-timepicker>
      </b-field>
    </div>

  </section>
</template>

<script>
import httpUtils from '../../lib/httpUtils'
import validationUtils from '../../lib/validationUtils'
import vueStore from '../../vuex'
import moment from 'moment'

export default {
  name: "addSchedule",
  data() {
    return {
      title: "",
      theater_number: 1,
      start_time: new Date(),
      end_time: new Date(),
      watch_time: 0
    }
  },
  watch: {
    start_time(new_start_time) {
      const d = moment(new_start_time)
      this.end_time = d.add(120, "minutes").toDate()
    }
  },
  mounted() {
    const title = vueStore.state.movie_title;
    const watch_time = vueStore.state.watch_time;
    if(title === "") {
      const id = this.$route.params.id;
      httpUtils.GetMovieDetail(id)
      .then((res) => {
        const title = res.data.movie_name
        this.title = title;
        this.watch_time = res.data.watch_time
      })
      .catch((err) => console.error(err))
    } else {
      this.watch_time = vueStore.state.watch_time;
      this.title = vueStore.state.movie_title;
    }
  },
  methods: {
    add() {

      // validation
      if(!validationUtils.isNumber(String(this.theater_number))) { validationUtils.pushMessage("シアター番号を数字で入力してください。", this); return }

      httpUtils.CreateSchedule(this.$route.params.id, this.theater_number, this.start_time)
      .then((res) => {
        if(res.status === 201) {
          this.$toast.open({
            message: 'スケジュールを追加しました。',
            type: 'is-success'
          })
          setTimeout(() => this.$router.push({ path: `/admin/schedule/${this.$route.params.id}` }), 500)
        }
      })
      .catch((err) => console.error(err))
    },
    back() {
      this.$router.push({ path: `/admin/schedule/${this.$route.params.id}` });
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
.flex {
  margin-bottom: 50px;
  display: flex;
  align-items: center;
}
.flex-box {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.theater-number {
  width: 30%;
  margin-bottom: 50px;
}
</style>

