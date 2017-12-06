<template>
  <section>
    <b-table
      :data="tableDataSimple.length == 0 ? [] : tableDataSimple"
      :bordered="false"
      :striped="true"
      :narrowed="false"
      :hoverable="true"
      :mobile-cards="true"
    >
      <template slot-scope="props">
        <b-table-column label="ID" width="40" numeric>
          {{ props.row.id }}
        </b-table-column>

        <b-table-column label="タイトル" width="200">
          {{ props.row.title }}
        </b-table-column>

        <b-table-column label="詳細">
          {{ props.row.detail }}
        </b-table-column>

        <b-table-column label="公開開始日" centered width="100">
          {{ new Date(props.row.start_date).toLocaleDateString() }}
        </b-table-column>

        <b-table-column label="公開終了日" centered width="100">
          {{ new Date(props.row.end_date).toLocaleDateString() }}
        </b-table-column>

        <b-table-column label="上映時間" centered width="100">
          {{ props.row.watch_time }}分
        </b-table-column>

        <b-table-column label="">
          <button class="button">編集</button>
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

export default {
  name: "movieTable",
  data() {
    const tableDataSimple = [
      {
        'id': 1,
        'title': '君の名は。',
        'detail': '父親と東京の都心に住む男子高校生・立花瀧（たちばなたき）と、岐阜県の山奥の糸守町（いともりちょう）に住む女子高生・宮水三葉（みやみずみつは）。朝起きると、ひょんなことから二人の意識が入れ替わり、ドタバタ生活が唐突に始まるところからストーリーが始まります。意識の入れ替わりがなぜ起こるのかは不明。二人は、当初、単純にリアルな夢を見ているのだと思っていました。しかし、朝起きると、週に２～３回ランダムに二人の体が入れ替わっているので、これは夢の出来事ではなく、これが現実世界の出来事であると認識し、互いの実在を確信するようになっていきます。',
        'start_date': '2016-10-15 13:43:27',
        'end_date': '2016-10-15 13:43:27',
        'watch_time': 120
      },
      {
        'id': 2,
        'title': '君の名は。',
        'detail': '父親と東京の都心に住む男子高校生・立花瀧（たちばなたき）と、岐阜県の山奥の糸守町（いともりちょう）に住む女子高生・宮水三葉（みやみずみつは）。朝起きると、ひょんなことから二人の意識が入れ替わり、ドタバタ生活が唐突に始まるところからストーリーが始まります。意識の入れ替わりがなぜ起こるのかは不明。二人は、当初、単純にリアルな夢を見ているのだと思っていました。しかし、朝起きると、週に２～３回ランダムに二人の体が入れ替わっているので、これは夢の出来事ではなく、これが現実世界の出来事であると認識し、互いの実在を確信するようになっていきます。',
        'start_date': '2016-10-15 13:43:27',
        'end_date': '2016-10-15 13:43:27',
        'watch_time': 120
      },
      {
        'id': 3,
        'title': '君の名は。',
        'detail': '父親と東京の都心に住む男子高校生・立花瀧（たちばなたき）と、岐阜県の山奥の糸守町（いともりちょう）に住む女子高生・宮水三葉（みやみずみつは）。朝起きると、ひょんなことから二人の意識が入れ替わり、ドタバタ生活が唐突に始まるところからストーリーが始まります。意識の入れ替わりがなぜ起こるのかは不明。二人は、当初、単純にリアルな夢を見ているのだと思っていました。しかし、朝起きると、週に２～３回ランダムに二人の体が入れ替わっているので、これは夢の出来事ではなく、これが現実世界の出来事であると認識し、互いの実在を確信するようになっていきます。',
        'start_date': '2016-10-15 13:43:27',
        'end_date': '2016-10-15 13:43:27',
        'watch_time': 120
      },
      {
        'id': 4,
        'title': '君の名は。',
        'detail': '父親と東京の都心に住む男子高校生・立花瀧（たちばなたき）と、岐阜県の山奥の糸守町（いともりちょう）に住む女子高生・宮水三葉（みやみずみつは）。朝起きると、ひょんなことから二人の意識が入れ替わり、ドタバタ生活が唐突に始まるところからストーリーが始まります。意識の入れ替わりがなぜ起こるのかは不明。二人は、当初、単純にリアルな夢を見ているのだと思っていました。しかし、朝起きると、週に２～３回ランダムに二人の体が入れ替わっているので、これは夢の出来事ではなく、これが現実世界の出来事であると認識し、互いの実在を確信するようになっていきます。',
        'start_date': '2016-10-15 13:43:27',
        'end_date': '2016-10-15 13:43:27',
        'watch_time': 120
      },
      {
        'id': 5,
        'title': '君の名は。',
        'detail': '父親と東京の都心に住む男子高校生・立花瀧（たちばなたき）と、岐阜県の山奥の糸守町（いともりちょう）に住む女子高生・宮水三葉（みやみずみつは）。朝起きると、ひょんなことから二人の意識が入れ替わり、ドタバタ生活が唐突に始まるところからストーリーが始まります。意識の入れ替わりがなぜ起こるのかは不明。二人は、当初、単純にリアルな夢を見ているのだと思っていました。しかし、朝起きると、週に２～３回ランダムに二人の体が入れ替わっているので、これは夢の出来事ではなく、これが現実世界の出来事であると認識し、互いの実在を確信するようになっていきます。',
        'start_date': '2016-10-15 13:43:27',
        'end_date': '2016-10-15 13:43:27',
        'watch_time': 120
      },
      {
        'id': 6,
        'title': '君の名は。',
        'detail': '父親と東京の都心に住む男子高校生・立花瀧（たちばなたき）と、岐阜県の山奥の糸守町（いともりちょう）に住む女子高生・宮水三葉（みやみずみつは）。朝起きると、ひょんなことから二人の意識が入れ替わり、ドタバタ生活が唐突に始まるところからストーリーが始まります。意識の入れ替わりがなぜ起こるのかは不明。二人は、当初、単純にリアルな夢を見ているのだと思っていました。しかし、朝起きると、週に２～３回ランダムに二人の体が入れ替わっているので、これは夢の出来事ではなく、これが現実世界の出来事であると認識し、互いの実在を確信するようになっていきます。',
        'start_date': '2016-10-15 13:43:27',
        'end_date': '2016-10-15 13:43:27',
        'watch_time': 120
      }
    ]
    return {
      tableDataSimple
    }
  },
  methods: {
  }
}
</script>
