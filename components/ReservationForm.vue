<template>
  <v-form>
    <v-text-field label="First Name" :rules="defaultRules" />
    <v-text-field label="Last Name" :rules="defaultRules" />
    <v-text-field label="Email Address" :rules="defaultRules" />
    <v-text-field label="Phone Number" :rules="defaultRules" />
    <v-select label="Apartment" :items="apartments" :rules="defaultRules" />
    <v-select label="Guests" :items="guests" :rules="defaultRules" />
    <v-date-picker v-model="dates" range />
  </v-form>
</template>

<script>
export default {
  data: () => ({
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: ['1', '2', '3', '4'],
    dates: [],
    defaultRules: [(v) => !!v || 'Required'],
  }),
}
</script>
