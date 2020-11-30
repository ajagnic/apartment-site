<template>
  <div>
    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = true" />
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" app>
      <v-list nav>
        <v-list-item-group>
          <v-list-item>
            <v-list-item-title>Home</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  data: () => ({
    drawer: null,
  }),
}
</script>
