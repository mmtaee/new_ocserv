<script lang="ts" setup>
import logoUrl from "@/assets/ocserv.png"
import {useTheme} from "vuetify/framework";
import {defineAsyncComponent, onBeforeMount, ref} from "vue";


const theme = useTheme()
const Logout = defineAsyncComponent(() => import("@/components/app/Logout.vue"))
const logoutDialog = ref(false)

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
      <span class="text-subtitle-1">Ocserv Service With User Management </span>
    </template>

    <template v-slot:append>
      <v-btn density="comfortable" icon @click="toggleTheme">
        <v-icon>mdi-theme-light-dark</v-icon>
      </v-btn>
      <v-btn density="comfortable" icon @click="logoutDialog = true">
        <v-icon>mdi-logout</v-icon>
      </v-btn>
    </template>

  </v-app-bar>

  <Logout v-model:logoutDialog="logoutDialog"/>

</template>
