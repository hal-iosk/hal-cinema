<template>
  <section>

    <div class="flex">
      <div v-if="title.length > 8">
        <b-tooltip type="is-dark" :label="title" position="is-bottom" multilined>
          <h1 style="font-size: 1.3rem;">「{{title.slice(0, 7) + "..."}}」スケジュール編集</h1>
        </b-tooltip>
      </div>
      <div v-else>
        <h1 style="font-size: 1.3rem;">「{{title}}」スケジュール編集</h1>
      </div>
      <button class="button" style="margin-left: auto;" @click="back">戻る</button>
      <button class="button is-primary" style="margin-left: 10px;" @click="add">スケジュール追加</button>
    </div>

    <b-table
      :data="schedules.length === 0 ? [] : schedules"
      :bordered="false"
      :striped="true"
      :narrowed="false"
      :hoverable="true"
      :mobile-cards="true"
    >

    <template slot-scope="props">

      <b-table-column label="ID">
        {{props.row.ID}}
      </b-table-column>

      <b-table-column label="シアター番号" v-if="props.row.release">
        {{props.row.theater_number}}
      </b-table-column>
      <b-table-column label="シアター番号" v-else>
        <b-input v-model="props.row.theater_number"></b-input>
      </b-table-column>

      <b-table-column label="開始時間">
        {{moment(props.row.start_time).format("YYYY/MM/DD HH:mm")}}
      </b-table-column>

      <b-table-column numeric v-if="props.row.release">
        <button class="button is-danger" disabled>公開済</button>
      </b-table-column>
      <b-table-column numeric v-else>
        <div class="tool-box">
          <button class="button" @click="scheduleDelete(props.row.ID)" style="color: red;">削除</button>
        </div>
        <button class="button is-danger" @click="open(props.row.ID)">公開する</button>
      </b-table-column>

      </template>

      <template slot="empty">
        <section class="section">
          <div class="content has-text-grey has-text-centered">
            <p>
              <b-icon
                icon="sentiment_very_dissatisfied"
                size="is-large">
              </b-icon>
            </p>
            <p>Nothing here.</p>
          </div>
        </section>
      </template>
    </b-table>

  </section>
</template>

<script>
import httpUtils from '../../lib/httpUtils'
import validationUtils from '../../lib/validationUtils'
import vueStore from '../../vuex'
import moment from 'moment'

export default {
  name: "schedule",
  data() {
    return {
      schedules: [],
      title: "",
      moment
    }
  },
  methods: {
    save(id) {
      let _schedule = null;

      this.schedules.map((schedule) => {
        if(schedule.ID === id) {
          _schedule = schedule;
        }
      });

      // validation
      if(!validationUtils.isNumber(String(_schedule.theater_number))) { validationUtils.pushMessage("シアター番号を数字で入力してください。", this); return }

      httpUtils.PutSchedule(_schedule)
      .then((res) => {
        if(res.status === 200) {
          this.$toast.open({
            message: 'スケジュールを編集しました。',
            type: 'is-success'
          })
        }
      })
      .catch((err) => console.error(err))
    },
    scheduleDelete(id) {
      httpUtils.DeleteSchedule(id)
      .then((res) => {
        if(res.status === 204) {
          this.schedules.some((schedule, index) => {
            if(schedule.ID === id) this.schedules.splice(index, 1)
          })

          this.$toast.open({
            message: "スケジュールを削除しました。",
            type: "is-success"
          })

          this.$forceUpdate()
        }
      })
      .catch((err) => console.error(err))
    },
    open(id) {
      this.$dialog.confirm({
          message: '映画を公開しますか？この動作は取り消せません。',
          onConfirm: () => {
            let _schedule = null;

            this.schedules.map((schedule) => {
              if(schedule.ID === id) {
                schedule.release = true
                schedule.start_time = moment(schedule.start_time).format("YYYY/MM/DD HH:mm:ss")
                _schedule = schedule;
              }
            });
            httpUtils.PutSchedule(_schedule)
            .then(() => {
              this.$toast.open({
                message: "映画を公開しました。",
                type: "is-success"
              })
            })
          }
      })
    },
    back() {
      this.$router.push({ path: `/admin` });
    },
    add() {
      this.$router.push({ path: `/admin/scheduleadd/${this.$route.params.id}` });
    }
  },
  mounted() {
    const id = this.$route.params.id;
    httpUtils.GetMovieDetail(id)
    .then((res) => {
      const title = res.data.movie_name;
      const watch_time = res.data.watch_time;
      vueStore.commit('set_movie_title', title)
      vueStore.commit('set_watch_time', watch_time)
      this.title = title;
    })
    .catch((err) => console.error(err))

    httpUtils.GetSchedule(id)
    .then((res) => {
      const schedules = res.data.schedules;
      schedules.map((schedule) => {
        schedule.start_time = schedule.start_time
      })
      this.schedules = schedules;
    })
    .catch((err) => console.error(err))
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
.flex {
  margin-bottom: 50px;
  display: flex;
  align-items: center;
}
.tool-box {
  margin-bottom: 10px;
}
</style>
