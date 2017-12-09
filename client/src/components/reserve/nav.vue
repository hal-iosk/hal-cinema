<template>
  <ul>
    <li class="navigation-item isActive">座席選択</li>
    <li>></li>
    <li class="navigation-item">チケット選択</li>
    <li>></li>
    <li class="navigation-item">購入内容の確認</li>
  </ul>
</template>

<script>
export default {
  name: "nav"
}
</script>

<style lang="scss" scoped>
ul {
  display: flex;
  justify-content: space-around;
  align-items: center;
  padding: 30px 20%;
  box-shadow: 0px 0px 1px black;
  li.navigation-item {
    letter-spacing: 1.5px;
  }
  li.isActive { font-weight: bold; color: #B71C1C; box-shadow: 0px 0px 1px #B71C1C; }
  li {
    padding: 10px;
  }
}
</style>
