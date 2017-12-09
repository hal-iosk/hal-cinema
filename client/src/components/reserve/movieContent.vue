<template>
  <section>
    <ul>
      <li>
        <p class="content-title">作品名:</p>
        <p class="content">{{title}}</p>
      </li>
      <li>
        <p class="content-title">日時:</p>
        <p class="content">{{date}}</p>
      </li>
      <li>
        <p class="content-title">シアター番号:</p>
        <p class="content">{{theater}}</p>
      </li>
    </ul>
  </section>
</template>

<script>
export default {
  props: ["title", "date", "theater"],
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

