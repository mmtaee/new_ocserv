<script lang="ts" setup>
import logoUrl from "@/assets/ocserv.png"
import {useTheme} from "vuetify/framework";
import {useUserStore} from "@/stores/user.ts";
import {ref} from "vue";
import type {ModelsUser} from "@/api";

const userStore = useUserStore()

const user = ref<ModelsUser | null>(userStore.getUser)

const items = [
  {text: 'My Files', icon: 'mdi-folder'},
  {text: 'Shared with me', icon: 'mdi-account-multiple'},
  {text: 'Starred', icon: 'mdi-star'},
  {text: 'Recent', icon: 'mdi-history'},
  {text: 'Offline', icon: 'mdi-check-circle'},
  {text: 'Uploads', icon: 'mdi-upload'},
  {text: 'Backups', icon: 'mdi-cloud-upload'},
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

    <v-menu v-if="user" persistent>
      <template v-slot:activator="{ props }">
        <v-btn icon="$menu" v-bind="props" variant="text"></v-btn>
      </template>

      <v-list>
        <v-list-item
            :subtitle="user.is_admin ? 'Admin': 'Staff'"
            :title="user.username?.toUpperCase()"
        >
        </v-list-item>
        <v-divider></v-divider>

        <v-list
            :lines="false"
            density="compact"
            nav
        >
          <v-list-item
              v-for="(item, i) in items"
              :key="i"
              :value="item"
              color="primary"
          >
            <template v-slot:prepend>
              <v-icon :icon="item.icon"></v-icon>
            </template>

            <v-list-item-title v-text="item.text"></v-list-item-title>
          </v-list-item>
        </v-list>

      </v-list>
    </v-menu>

  </v-app-bar>
</template>

<style scoped>

</style>