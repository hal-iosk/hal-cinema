<template>
  <div>
    <section>
      <div id="theater_nav">
        <a href="watchFilm.">
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
          <p class="container-title"><span class="date">12/8</span>(金)の上映スケジュール</p>
          <ul class="movie-container">
            <li>
              <div class="movie">
                <div class="top-container">
                  <p class="movie-title">[字幕] 否定と肯定</p>
                  <p class="screening-time">(本編: 140分)</p>
                </div>
                <div class="bottom-container">
                  <div class="movie-image">
                    <img src="movie1.jpg" alt="">
                  </div>
                  <ul class="schedule-container">
                    <li>
                      <div class="schedule">
                        <div class="screen-name">スクリーン10</div>
                        <div class="time-container">
                          <span class="time">8:50</span>~<span class="time">10:50</span>
                        </div>
                      </div>
                    </li>
                    <li>
                      <div class="schedule">
                        <div class="screen-name">スクリーン10</div>
                        <div class="time-container">
                          <span class="time">8:50</span>~<span class="time">10:50</span>
                        </div>
                      </div>
                    </li>
                    <li>
                      <div class="schedule">
                        <div class="screen-name">スクリーン10</div>
                        <div class="time-container">
                          <span class="time">8:50</span>~<span class="time">10:50</span>
                        </div>
                      </div>
                    </li>
                    <li>
                      <div class="schedule">
                        <div class="screen-name">スクリーン10</div>
                        <div class="time-container">
                          <span class="time">8:50</span>~<span class="time">10:50</span>
                        </div>
                      </div>
                    </li>
                    <li>
                      <div class="schedule">
                        <div class="screen-name">スクリーン10</div>
                        <div class="time-container">
                          <span class="time">8:50</span>~<span class="time">10:50</span>
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
  </div>
</template>

<script>
export default {
  name: "movie"
}
</script>


<style lang="scss" scoped>
.flex-container {
  display: flex;
  justify-content: center;
}
.container {
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
.movie-container li {
  list-style: none;
}
.movie {
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
</style>