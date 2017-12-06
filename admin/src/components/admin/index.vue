<template>
  <!-- <sales-manage /> -->
  <mail-magazine />
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
  }
}
</script>
