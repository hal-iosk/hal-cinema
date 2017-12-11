<template>
  <div>
    <section>
      <div id="theater_nav">
        <a href="watchFilm">
          <div class="nav_selected">
            <div class="nav_mean">
              <p>上映中</p>
              <p>now playing</p>
            </div>
          </div>
        </a>
        <a href="comingSoon">
          <div class="nav_noselected">
            <div class="nav_mean">
              <p>公開予定</p>
              <p>coming soon</p>
            </div>
          </div>
        </a>
      </div>
    </section>
    <article>
      <div class="flex-container">
        <div class="container">

          <ul class="dates">
            <li v-for="tab in tabs" @click="selectDate(tab.key, tab.day)">
              <p class="active" v-if="tab.isActive">{{tab.day.format("MM/DD")}}</p>
              <p v-else>{{tab.day.format("MM/DD")}}</p>
            </li>
          </ul>

          <p class="container-title"><span class="date">{{currentDate.format("MM/DD")}}</span>の上映スケジュール</p>
          <ul class="movie-container">
            <li v-for="movie in movies" v-if="movie.schedules && movie.schedules.length != 0">
              <div class="movie">
                <div class="top-container">
                  <p class="movie-title">{{movie.movie_name}}</p>
                  <p class="screening-time">(本編: {{movie.watch_time}}分)</p>
                </div>
                <div class="bottom-container">
                  <ul class="schedule-container">
                    <li v-for="schedule in movie.schedules" @click="select(schedule.ID, movie.movie_name, schedule.start_time, movie.watch_time, schedule.theater_number)">
                      <div class="schedule">
                        <div class="screen-name">スクリーン{{schedule.theater_number}}</div>
                        <div class="time-container">
                          <span class="time">
                            {{moment(schedule.start_time).format("HH:mm")}}
                          </span>
                          ~
                          <span class="time">
                            {{moment(schedule.start_time).add(movie.watch_time, "m").format("HH:mm")}}
                          </span>
                        </div>
                      </div>
                    </li>
                  </ul>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </article>
    <b-loading :active.sync="isLoading" :canCancel="true"></b-loading>
  </div>
</template>

<script>
import MovieHttp from '../../services/movie'
import moment from 'moment'
import vueStore from '../../vuex'

export default {
  name: "movie",
  data() {
    return {
      movies: [],
      moment,
      currentDate: moment(),
      isLoading: false,
      tabs: [
        { key: 0, day: moment(), isActive: true },
        { key: 1, day: moment().add(1, "days"), isActive: false },
        { key: 2, day: moment().add(2, "days"), isActive: false },
        { key: 3, day: moment().add(3, "days"), isActive: false },
        { key: 4, day: moment().add(4, "days"), isActive: false },
        { key: 5, day: moment().add(5, "days"), isActive: false },
        { key: 6, day: moment().add(6, "days"), isActive: false }
      ]
    }
  },
  mounted() {
    this.isLoading = true

    MovieHttp.GetOnAirMovies(moment().format("YYYY/MM/DD"))
    .then((res) => {
      this.isLoading = false
      this.movies = res.data.movies
    })
  },
  methods: {
    selectDate(key, day) {
      this.isLoading = true;
      this.tabs.map((tab) => {
        tab.isActive = false
        if(key === tab.key) { tab.isActive = true }
      })
      this.currentDate = day
      MovieHttp.GetOnAirMovies(day.format("YYYY/MM/DD"))
      .then((res) => {
        this.isLoading = false
        this.movies = res.data.movies
      })
    },
    select(scheduleId, title, start_time, watch_time, theater_number) {
      const m = moment(start_time)
      const reserve = {
        scheduleId,
        title,
        theater_number,
        days: `${m.format("YYYY/MM/DD")} ${m.format("HH:mm")} ~ ${m.add(watch_time, "m").format("HH:mm")}`
      }

      sessionStorage.setItem("halCinemaReserve", JSON.stringify(reserve))

      location.href = "/reserve"
    }
  }
}
</script>


<style lang="scss" scoped>
.flex-container {
  display: flex;
  justify-content: center;
}
.container {
  width: 80%;
  margin: 0 auto;
  padding-top: 2rem;
}
.container-title {
  padding: 0.5rem;
  border-bottom: 1px solid #DDD;
}
.container-title .date {
  font-size: 1.5rem;
  font-weight: bold;
}
.movie-container {
  width: 100%;
}
.movie-container li {
  list-style: none;
}
.movie {
  width: 100%;
  padding: 1rem 0;
}
.movie .movie-title {
  font-weight: bold;
  font-size: 1.2rem;
}
.movie .top-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 0;
}
.movie .bottom-container {
  overflow: scroll;
  display: flex;
  align-items: center;
  background-color: rgba(0, 0, 0, .79);
  border-radius: 5px;
}
.movie-image {
  width: 120px;
  padding: 0.5rem;
}
.movie-image img {
  display: block;
  width: 100%;
}
.schedule-container {
  display: flex;
  padding: 1rem;
}
.schedule-container li {
  cursor: pointer;
  list-style: none;
}
.schedule-container li+li {
  margin-left: 0.7rem;
}
.schedule {
  padding: 1rem;
  background-color: #FFF;
  border-radius: 15px;
}
.schedule .screen-name {
  font-size: 0.7rem;
  text-align: center;
}
.schedule .time-container {
  display: flex;
  align-items: center;
  text-align: center;
  padding: 1rem 0;
}
.schedule .time-container .time {
  display: block;
  font-weight: bold;
  padding: 0 0.5rem;
  font-size: 1.1rem;
}
.schedule.disabled {
  background-color: #777;
}
.dates {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
  border-bottom: 1px solid gray;
  li {
    letter-spacing: 3px;
    cursor: pointer;
    p {
      padding: 10px;
    }
    p.active {
      background-color: #e2e2e2;
    }
  }
}
</style>
