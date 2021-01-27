<template>
  <p>Your reservation has been confirmed</p>
</template>

<script>
export default {
  mounted() {
    const id = this.$route.query.id
    this.$axios.post('/reservations/confirm', id).then(
      (response) => {
        console.log(response)
      },
      (error) => {
        console.log(error)
      }
    )
  },
}
</script>
