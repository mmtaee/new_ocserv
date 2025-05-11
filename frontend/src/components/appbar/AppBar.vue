<script lang="ts" setup>
import logoUrl from "@/assets/ocserv.png"
import {useTheme} from "vuetify/framework";
import {useUserStore} from "@/stores/user.ts";
import {defineAsyncComponent, onBeforeMount, ref, watch} from "vue";
import type {ModelsUser} from "@/api";
import Tabs from "@/components/appbar/Tabs.vue";
import Menu from "@/components/appbar/Menu.vue";

const userStore = useUserStore()
const user = ref<ModelsUser | null>(userStore.getUser)
const theme = useTheme()
const Logout = defineAsyncComponent(() => import("@/components/appbar/Logout.vue"))
const logoutDialog = ref(false)

watch(() => userStore.getUser, (newVal) => {
  user.value = newVal
})

onBeforeMount(() => {
  theme.global.name.value = localStorage.getItem('theme') === 'dark' ? 'dark' : 'light';
});

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  localStorage.setItem('theme', theme.global.name.value)
}

</script>

<template>
  <v-app-bar :elevation="12" density="comfortable">

    <template v-slot:prepend>
      <v-app-bar-nav-icon>
        <v-img :src="logoUrl" alt="ocserv logo" width="45"/>
      </v-app-bar-nav-icon>
    </template>

    <template v-slot:title>
      <span class="text-subtitle-1">Ocserv User Management Service</span>
    </template>
    
    <Tabs v-if="user" :user="user"/>

    <template v-slot:append>
      <v-btn density="comfortable" icon @click="toggleTheme">
        <v-icon>mdi-theme-light-dark</v-icon>
      </v-btn>
      <Menu v-if="user" :user="user" @logout="(val) => logoutDialog = val"/>
    </template>

  </v-app-bar>

  <Logout v-model:logoutDialog="logoutDialog"/>

</template>
