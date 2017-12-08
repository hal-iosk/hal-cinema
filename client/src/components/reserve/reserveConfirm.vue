<template>
  <section>
    <h1>予約完了</h1>
  </section>
</template>

<script>
import reservePayload from '../../lib/reserve.class'

export default {
  name: "reserve-confirm",
  mounted() {
    console.log(reservePayload.getRequestData())
  }
}
</script>

