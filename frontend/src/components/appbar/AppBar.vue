<script lang="ts" setup>
import logoUrl from "@/assets/ocserv.png"
import {useLocale, useTheme} from "vuetify/framework";
import {useUserStore} from "@/stores/user.ts";
import {computed, defineAsyncComponent, ref, watch} from "vue";
import type {ModelsUser} from "@/api";
import router from "@/plugins/router.ts";

interface Item {
  text: string
  icon: string
  admin: boolean
  to?: string
  action?: () => void
}


const userStore = useUserStore()
const {t} = useLocale()
const user = ref<ModelsUser | null>(userStore.getUser)
const theme = useTheme()
const Logout = defineAsyncComponent(() => import("@/components/appbar/Logout.vue"))
const logoutDialog = ref(false)

watch(() => userStore.getUser, (newVal) => {
  user.value = newVal
})


const menuItems: Item[] = [
  {text: t("CONFIG"), icon: 'mdi-cog', admin: false, to: "/config"},
  {text: t("CHANGE_PASSWORD"), icon: 'mdi-account-key', admin: false, to: "/change_password"},
  {text: t("STAFF_MANAGEMENT"), icon: 'mdi-account-tie-hat', admin: true, to: "/staffs"},
  {
    text: t("LOGOUT"), icon: 'mdi-logout', admin: false, action: () => {
      logoutDialog.value = true
    }
  },
]

const itemsCompute = computed(() => {
  if (!user.value?.is_admin) {
    return menuItems.filter((item) => !item.admin)
  }
  return menuItems
})

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  localStorage.setItem('theme', theme.global.name.value)
}

const menuAction = (item: Item) => {
  if (item.to) {
    router.push(item.to)
    return
  }
  if (item.action) {
    item.action()
  }
  return
}

theme.global.name.value = localStorage.getItem('theme') === 'dark' ? 'dark' : 'light'

</script>

<template>
  <v-app-bar :elevation="12" density="compact">

    <template v-slot:prepend>
      <v-app-bar-nav-icon>
        <v-img :src="logoUrl" alt="ocserv logo" width="45"/>
      </v-app-bar-nav-icon>
    </template>

    <v-app-bar-title class="text-subtitle-1">Ocserv User Management Service</v-app-bar-title>

    <v-btn density="comfortable" icon @click="toggleTheme">
      <v-icon>mdi-theme-light-dark</v-icon>
    </v-btn>

    <v-menu v-if="user">
      <template v-slot:activator="{ props }">
        <v-btn class="me-5" density="comfortable" icon="mdi-account-settings-outline" v-bind="props"/>
      </template>

      <v-list>
        <v-list-item>
          <v-list-item-title class="text-primary">
            <v-icon class="mb-1" color="primary" start>
              {{ user.is_admin ? "mdi-account-star-outline" : "mdi-account" }}
            </v-icon>
            <span class="text-capitalize">
              {{ user.username }}
            </span>
            <span class="text-subtitle-2 text-capitalize">
              ({{ user.is_admin ? t('ADMIN') : t('STAFF') }})
            </span>
          </v-list-item-title>
        </v-list-item>

        <v-divider></v-divider>

        <v-list
            :lines="false"
            density="compact"
            nav
        >
          <v-list-item
              v-for="(item, i) in itemsCompute"
              :key="i"
              :value="item"
              color="primary"
              @click="menuAction(item)"
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

  <Logout v-model:logoutDialog="logoutDialog"/>

</template>
