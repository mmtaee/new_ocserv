<script lang="ts" setup>
import logoUrl from "@/assets/ocserv.png"
import {useTheme} from "vuetify/framework";

const items = [
  {title: 'Click Me'},
  {title: 'Click Me'},
  {title: 'Click Me'},
  {title: 'Click Me 2'},
]

const theme = useTheme()
theme.global.name.value = localStorage.getItem('theme') === 'dark' ? 'dark' : 'light'

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  localStorage.setItem('theme', theme.global.name.value)
}

</script>

<template>

  <v-app-bar :elevation="12" density="compact">


    <template v-slot:prepend>
      <v-app-bar-nav-icon>
        <v-img :src="logoUrl" alt="ocserv logo" width="45"/>
      </v-app-bar-nav-icon>
    </template>


    <v-app-bar-title class="text-subtitle-1">Ocserv User Management Service</v-app-bar-title>

    <v-btn icon @click="toggleTheme">
      <v-icon>mdi-theme-light-dark</v-icon>
    </v-btn>

    <v-menu>
      <template v-slot:activator="{ props }">
        <v-btn icon="$menu" v-bind="props" variant="text"></v-btn>
      </template>

      <v-list>
        <v-list-item
            v-for="(item, i) in items"
            :key="i"
            :value="i"
        >
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>


    <!--    <template v-slot:append>-->
    <!--      <v-menu>-->
    <!--        <template v-slot:activator="{ props }">-->
    <!--          <v-app-bar-nav-icon v-bind:props></v-app-bar-nav-icon>-->
    <!--        </template>-->

    <!--        <v-list>-->
    <!--          <v-list-item-->
    <!--              v-for="(item, i) in items"-->
    <!--              :key="i"-->
    <!--              :value="i"-->
    <!--          >-->
    <!--            <v-list-item-title>{{ item.title }}</v-list-item-title>-->
    <!--          </v-list-item>-->
    <!--        </v-list>-->
    <!--      </v-menu>-->

    <!--    </template>-->
  </v-app-bar>
</template>

<style scoped>

</style>