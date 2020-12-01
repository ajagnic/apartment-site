<template>
  <v-form>
    <v-text-field label="First Name" />
    <v-text-field label="Last Name" />
    <v-text-field label="Email Address" />
    <v-text-field label="Phone Number" />
    <v-select label="Apartment" :items="apartments" />
    <v-select label="Guests" :items="guests" />
    <v-slider
      label="Guests"
      min="0"
      max="4"
      ticks="always"
      tick-size="4"
      thumb-label="always"
    />
    <v-date-picker />
  </v-form>
</template>

<script>
export default {
  data: () => ({
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: ['1', '2', '3', '4'],
  }),
}
</script>
