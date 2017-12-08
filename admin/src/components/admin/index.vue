<template>
  <section>
    <div class="flex">
      <button class="button is-primary movie-add" @click="addMovie">映画追加</button>
      <button class="button is-danger logout" @click="logout">ログアウト</button>
    </div>
    <b-tabs v-model="activeTab">

      <b-tab-item label="映画情報管理">
        <movie-table />
      </b-tab-item>

      <b-tab-item label="メルマガ">
        <mail-magazine />
      </b-tab-item>

    </b-tabs>
  </section>
</template>

<script>
import MovieTable from './movieTable.vue'
import MailMagazine from './mailMagazine.vue'
import CookieDoc from '../../lib/CookieDoc'

export default {
  name: "admin",
  components: {
    MovieTable,
    MailMagazine
  },
  data() {
    return {
      activeTab: 0
    }
  },
  methods: {
    logout() {
      CookieDoc.removeItem("halCinemaAdmin")
      location.pathname = "/admin/login"
    },
    addMovie() {
      this.$router.push({ path: `/admin/movie` })
    }
  }
}
</script>

<style lang="scss" scoped>
.flex {
  display: flex;
  margin: 20px 0;
  .movie-add {
    margin-left: 20px;
  }
  .logout {
    margin-left: auto;
    margin-right: 20px;
  }
}
</style>
