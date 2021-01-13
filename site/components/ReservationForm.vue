<template>
  <v-card max-width="800">
    <v-card-title>Reservation Form</v-card-title>
    <v-card-text>
      <v-form ref="form" v-model="valid">
        <v-text-field
          v-model="form.first"
          label="First Name"
          :rules="[rules.required]"
        />
        <v-text-field
          v-model="form.last"
          label="Last Name"
          :rules="[rules.required]"
        />
        <v-text-field v-model="form.phone" label="Phone Number" type="tel" />
        <v-text-field
          v-model="form.email"
          label="Email Address"
          type="email"
          :rules="[rules.required, rules.email]"
        />
        <v-select
          v-model="form.apartment"
          label="Apartment"
          :items="apartments"
          :rules="[rules.required]"
        />
        <v-select
          v-model="form.guests"
          label="Guests"
          :items="guests"
          :rules="[rules.required]"
        />
        <v-date-picker
          v-model="form.dates"
          :allowed-dates="allowedDates"
          multiple
        />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn :disabled="!valid" @click="submitReservation">
        Create Reservation
      </v-btn>
    </v-card-actions>
    <v-overlay v-if="error" z-index="1" absolute>
      <v-alert prominent type="error">There was an error. ):</v-alert>
    </v-overlay>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    error: null,
    valid: false,
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: [1, 2, 3, 4],
    rules: {
      required: (v) => !!v || 'Required.',
      email: (v) => {
        const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(v) || 'Invalid Email.'
      },
    },
    form: {
      first: '',
      last: '',
      phone: '',
      email: '',
      apartment: '',
      guests: 0,
      dates: [],
    },
    reservedDates: [],
  }),

  mounted() {
    this.$axios.get('/reservations').then(
      (response) => {
        const dates = response.data
        if (dates != null) {
          this.reservedDates = dates
        }
      },
      (error) => {
        this.error = error
      }
    )
  },

  methods: {
    submitReservation() {
      const userForm = this.form
      userForm.name = userForm.first.concat(' ', userForm.last)
      userForm.created = new Date().toDateString()
      this.$nuxt.$loading.start()
      this.$axios.post('/reservations', this.form).then(
        (response) => {
          this.$refs.form.reset()
          this.$nuxt.$loading.finish()
          this.$router.push('/')
        },
        (error) => {
          this.$nuxt.$loading.finish()
          this.error = error
        }
      )
    },

    allowedDates(val) {
      return !this.reservedDates.includes(val)
    },

    orderDates() {
      const splitDates = this.form.dates.map((x) => x.split('-'))
      const first = splitDates[0]
      const second = splitDates[1]
      const firstUTC = Date.UTC(first[0], first[1], first[2])
      const secondUTC = Date.UTC(second[0], second[1], second[2])
      if (secondUTC < firstUTC) {
        this.form.dates.reverse()
      }
    },
  },
}
</script>
