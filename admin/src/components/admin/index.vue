<template>
  <section>
    <div class="logout">
      <button class="button is-danger">ログアウト</button>
    </div>
    <b-tabs v-model="activeTab">

      <b-tab-item label="映画情報管理">
        <movie-table />
      </b-tab-item>

      <b-tab-item label="売上管理">
        <sales-manage />
      </b-tab-item>

      <b-tab-item label="メルマガ">
        <mail-magazine />
      </b-tab-item>

    </b-tabs>
  </section>
</template>

<script>
import MovieTable from './movieTable.vue'
import SalesManage from './salesManage.vue'
import MailMagazine from './mailMagazine.vue'

export default {
  name: "admin",
  components: {
    MovieTable,
    SalesManage,
    MailMagazine
  },
  data() {
    return {
      activeTab: 0
    }
  }
}
</script>

<style lang="scss" scoped>
.logout {
  display: flex;
  margin: 20px 0;
  .button {
    margin-left: auto;
    margin-right: 20px;
  }
}
</style>
