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
          <b-input v-model="props.row.ID"></b-input>
        </b-table-column>

        <b-table-column label="映画番号">
          <b-input v-model="props.row.movie_id"></b-input>
        </b-table-column>

        <b-table-column label="シアター番号">
          <b-input v-model="props.row.theater_number"></b-input>
        </b-table-column>

        <b-table-column label="開始時間" centered width="100">
          <b-datepicker
            v-model="props.row.start_time"
            :first-day-of-week="1"
          ></b-datepicker>
        </b-table-column>

        <b-table-column numeric>
          <button class="button is-success" @click="complate(props.row.ID)">完了</button>
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
import vueStore from '../../vuex'

export default {
  name: "schedule",
  data() {
    return {
      schedules: [],
      title: ""
    }
  },
  methods: {
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
      schedules.map((schedule) => { schedule.start_time = new Date(schedule.start_time)})
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
</style>
