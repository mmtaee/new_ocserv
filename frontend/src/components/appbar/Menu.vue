<script lang="ts" setup>
import router from "@/plugins/router.ts";
import {useLocale} from "vuetify/framework";
import type {PropType} from "vue";
import type {ModelsUser} from "@/api";


interface Item {
  text: string
  icon: string
  admin: boolean
  to?: string
  action?: () => void
}

defineProps({
  user: {
    type: Object as PropType<ModelsUser>,
    required: true,
  },
})

const {t} = useLocale()
const emit = defineEmits(["logout"])


const menuItems: Item[] = [
  {text: t("CONFIG"), icon: 'mdi-cog', admin: false, to: "/config"},
  {text: t("CHANGE_PASSWORD"), icon: 'mdi-account-key', admin: false, to: "/change_password"},
  {
    text: t("LOGOUT"), icon: 'mdi-logout', admin: false, action: () => {
      emit("logout", true)
    }
  },
]


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
</script>

<template>


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
            v-for="(item, i) in menuItems"
            :key="i"
            :value="item"
            color="primary"
            @click="menuAction(item)"
        >
          <template v-slot:prepend>
            <v-icon :icon="item.icon"/>
          </template>

          <v-list-item-title v-text="item.text"></v-list-item-title>
        </v-list-item>
      </v-list>

    </v-list>
  </v-menu>
</template>
