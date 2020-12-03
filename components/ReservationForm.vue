<template>
  <v-form>
    <v-row>
      <v-col cols="3">
        <v-text-field label="First Name" :rules="[rules.required]" />
      </v-col>
      <v-col cols="3">
        <v-text-field label="Last Name" :rules="[rules.required]" />
      </v-col>
      <v-col>
        <v-text-field label="Phone Number" type="tel" />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="5">
        <v-text-field
          label="Email Address"
          type="email"
          :rules="[rules.required, rules.email]"
        />
      </v-col>
      <v-col>
        <v-select
          label="Apartment"
          :items="apartments"
          :rules="[rules.required]"
        />
      </v-col>
      <v-col cols="2">
        <v-select label="Guests" :items="guests" :rules="[rules.required]" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-date-picker v-model="dates" range />
      </v-col>
      <v-col>
        <v-text-field
          v-model="dates[0]"
          label="Start Date"
          type="date"
          readonly
        />
      </v-col>
      <v-col>
        <v-text-field
          v-model="dates[1]"
          label="End Date"
          type="date"
          readonly
        />
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
export default {
  data: () => ({
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: ['1', '2', '3', '4'],
    dates: [],
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
