<template>
  <section>
    <ul>
      <li>
        <p class="content-title">作品名:</p>
        <p class="content">あさひなぐ</p>
      </li>
      <li>
        <p class="content-title">日時:</p>
        <p class="content">2017/09/26 10:00 〜 11:00</p>
      </li>
      <li>
        <p class="content-title">シアター番号:</p>
        <p class="content">1</p>
      </li>
    </ul>
  </section>
</template>

<script>
export default {
  porps: ["title"],
  data() {
    return {
      title: "",
      date: "",
      theater: ""
    }
  },
  name: "movie-content"
}
</script>

<style lang="scss" scoped>
ul {
  display: flex;
  justify-content: space-around;
  margin: 10px 0;
  li {
    display: flex;
    padding: 10px;
    align-items: center;
  }
  .content-title {
    margin-right: 10px;
  }
  .content {
    font-weight: bold;
    font-size: 1.3rem;
  }
}
</style>

