<template>
  <section>

    <div class="flex">
      <h1>映画スケジュール編集</h1>
      <button class="button" style="margin-left: auto;" @click="back">戻る</button>
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
          <button class="button is-danger" @click="scheduleDelete(props.row.ID)">削除</button>
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

export default {
  name: "schedule",
  data() {
    return {
      schedules: []
    }
  },
  methods: {
    scheduleDelete(id) {
      httpUtils.DeleteSchedule(id)
      .then((res) => {
        if(res.status === 204) {
          this.$toast.open({
            message: '削除完了しました。',
            type: 'is-success'
          })
          setTimeout(() => {
            this.$router.push({ path: `/admin` });
          }, 500)
        }
      })
    },
    back() {
      this.$router.push({ path: `/admin` });
    }
  },
  mounted() {
    const id = this.$route.params.id;
    httpUtils.GetMovieDetail(id)
    .then((res) => {
      console.log(res.data)
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
