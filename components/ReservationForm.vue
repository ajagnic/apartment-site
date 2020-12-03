<template>
  <v-form>
    <v-row>
      <v-col>
        <v-text-field
          v-model="form.first"
          label="First Name"
          :rules="[rules.required]"
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model="form.last"
          label="Last Name"
          :rules="[rules.required]"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          v-model="form.phone"
          label="Phone Number"
          type="tel"
          :rules="[rules.required]"
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model="form.email"
          label="Email Address"
          type="email"
          :rules="[rules.required, rules.email]"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-select
          v-model="form.apartment"
          label="Apartment"
          :items="apartments"
          :rules="[rules.required]"
        />
      </v-col>
      <v-col>
        <v-select
          v-model="form.guests"
          label="Guests"
          :items="guests"
          :rules="[rules.required]"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-date-picker v-model="form.dates" full-width range />
      </v-col>
      <v-col>
        <v-text-field v-model="form.dates[0]" label="Start Date" type="date" />
        <v-text-field v-model="form.dates[1]" label="End Date" type="date" />
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
export default {
  data: () => ({
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: ['1', '2', '3', '4'],
    form: {
      first: '',
      last: '',
      phone: '',
      email: '',
      apartment: '',
      guests: '',
      dates: [],
    },
    rules: {
      required: (v) => !!v || 'Required.',
      email: (v) => {
        const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(v) || 'Invalid email.'
      },
    },
  }),
}
</script>
