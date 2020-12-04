<template>
  <v-card max-width="800">
    <v-card-title>Reservation Form</v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field label="First Name" />
        <v-text-field label="Last Name" />
        <v-text-field label="Phone Number" type="tel" />
        <v-text-field label="Email Address" type="email" />
        <v-select label="Apartment" :items="apartments" />
        <v-select label="Guests" :items="guests" />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn>Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    apartments: ['Apartment #1', 'Apartment #2', 'Apartment #3'],
    guests: ['1', '2', '3', '4'],
  }),
}
</script>
