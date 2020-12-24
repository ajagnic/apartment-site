<template>
  <v-card max-width="800">
    <v-card-title>Reservation Form</v-card-title>
    <v-card-text>
      <v-form>
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
        <v-row>
          <v-col>
            <v-date-picker v-model="form.dates" range />
          </v-col>
          <v-col>
            <v-text-field v-model="form.dates[0]" label="Start Date" readonly />
            <v-text-field v-model="form.dates[1]" label="End Date" readonly />
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn @click="submitReservation">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: ['1', '2', '3', '4'],
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
      guests: '',
      dates: [],
    },
  }),

  methods: {
    submitReservation() {
      this.$axios.post('/', this.form)
    },
  },
}
</script>
