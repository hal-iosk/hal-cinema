<template>
  <section>
    <b-field label="タイトル">
      <b-input></b-input>
    </b-field>

    <b-field label="詳細">
      <b-input maxlength="200" type="textarea"></b-input>
    </b-field>

    <b-field label="公開開始日">
      <b-datepicker v-model="date" inline></b-datepicker>
    </b-field>

    <b-field label="公開終了日">
      <b-datepicker v-model="date" inline></b-datepicker>
    </b-field>

    <b-field label="上映時間">
      <b-input></b-input>
    </b-field>

  </section>
</template>

<script>
export default {
  name: "movieTableEdit"
}
</script>
