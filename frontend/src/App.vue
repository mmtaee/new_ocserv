<script lang="ts" setup>
import {useTheme} from 'vuetify'
import {defineAsyncComponent} from "vue";

const CenterLayout = defineAsyncComponent(() => import("@/components/layouts/CenterLayout.vue"))

const theme = useTheme()
theme.global.name.value = localStorage.getItem('theme') === 'dark' ? 'dark' : 'light'

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  localStorage.setItem('theme', theme.global.name.value)
}
</script>

<template>
  <v-app :theme="theme.global.name.value">
        <v-btn @click="toggleTheme">Toggle Theme</v-btn>
    <CenterLayout>
      <RouterView/>
    </CenterLayout>
  </v-app>
</template>
