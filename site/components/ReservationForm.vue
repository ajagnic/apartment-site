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
          @change="filterDates"
        />
        <v-select
          v-model="form.guests"
          label="Guests"
          :items="guests"
          :rules="[rules.required]"
        />
        <v-date-picker
          v-if="form.apartment != ''"
          v-model="pickerDates"
          :min="new Date().toISOString().substr(0, 10)"
          :max="setMaxDate"
          :allowed-dates="allowedDates"
          range
          @change="orderDates"
        />
        <v-text-field v-model="pickerDates[0]" label="Start" readonly />
        <v-text-field v-model="pickerDates[1]" label="End" readonly />
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
    allReservedDates: {
      'Apartment #1': [],
      'Apartment #2': [],
      'Apartment #3': [],
    },
    pickerDates: [],
  }),

  computed: {
    setMaxDate() {
      const year = new Date().getFullYear()
      return new Date(year, 11, 31).toISOString().substr(0, 10)
    },
  },

  mounted() {
    this.$axios.get('/reservations').then(
      (response) => {
        const dates = response.data
        if (dates != null) {
          for (const key in dates) {
            this.allReservedDates[key] = dates[key]
          }
        }
      },
      (error) => {
        this.error = error
      }
    )
  },

  methods: {
    submitReservation() {
      this.fillDates()
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

    filterDates(val) {
      this.reservedDates = this.allReservedDates[val]
    },

    allowedDates(val) {
      return !this.reservedDates.includes(val)
    },

    orderDates() {
      const [first, second] = this.pickerDates.map((x) => x.split('-'))
      const firstUTC = Date.UTC(first[0], first[1], first[2])
      const secondUTC = Date.UTC(second[0], second[1], second[2])
      if (secondUTC < firstUTC) {
        this.pickerDates.reverse()
      }
    },

    fillDates() {
      const endISO = this.pickerDates[1]
      let dateISO = this.pickerDates[0]
      let start = dateISO.split('-')
      start = start.map((x) => parseInt(x))
      const [y, m, d] = [start[0], start[1] - 1, start[2]]
      let cntr = 0
      while (dateISO !== endISO) {
        dateISO = new Date(y, m, d + cntr).toISOString().substr(0, 10)
        this.form.dates.push(dateISO)
        cntr++
      }
    },
  },
}
</script>
