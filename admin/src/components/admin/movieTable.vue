<template>
  <section>
    <b-table
      :data="movies.length == 0 ? [] : movies"
      :bordered="false"
      :striped="true"
      :narrowed="false"
      :hoverable="true"
      :mobile-cards="true"
    >
      <template slot-scope="props">
        <b-table-column label="ID" width="40" numeric>
          {{ props.row.ID }}
        </b-table-column>

        <b-table-column label="タイトル" width="200">
          {{ props.row.movie_name }}
        </b-table-column>

        <b-table-column label="詳細">
          {{ props.row.details }}
        </b-table-column>

        <!-- <b-table-column label="公開開始日" centered width="100">
          {{ new Date(props.row.start_date).toLocaleDateString() }}
        </b-table-column>

        <b-table-column label="公開終了日" centered width="100">
          {{ new Date(props.row.end_date).toLocaleDateString() }}
        </b-table-column> -->

        <b-table-column label="上映時間" centered width="100">
          {{ props.row.watch_time }}分
        </b-table-column>

        <b-table-column>
          <button class="button" @click="moveSchedule(props.row.id)" style="margin-bottom: 10px;">スケジュール</button>
          <button class="button" @click="movieEdit(props.row.id)">編集</button>
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
import vueStore from '../../vuex'
import httpUtils from '../../lib/httpUtils'

export default {
  name: "movieTable",
  data() {
    return {
      movies: []
    }
  },
  mounted() {
    httpUtils.GetMovies()
    .then((res) => {
      console.log(res.data.movies[0])
      this.movies = res.data.movies;
    })
    .catch((err) => {
      console.error(err)
    })
  },
  methods: {
    movieEdit(id) {
      this.$router.push({ path: `/admin/movieedit/${id}` })
    },
    moveSchedule(id) {
      this.$router.push({ path: `/admin/movieedit/${id}` })
    }
  }
}
</script>
