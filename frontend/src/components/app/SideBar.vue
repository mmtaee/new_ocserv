<script lang="ts" setup>
import {onMounted, reactive, ref, watch} from "vue";
import avatarUrl from "@/assets/torvalds.jpg"
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import {useLocale} from "vuetify/framework";
import {useUserStore} from "@/stores/user.ts";
import type {ModelsUser} from "@/api";


interface sideBar {
  title: string
  icon: string
  to: string
  value: string
  permission?: string
  divider?: boolean
}

const {t} = useLocale()
const drawer = ref(true)
const rail = ref(false)
const useIsMobile = useIsMobileStore()
const userStore = useUserStore()
const user = ref<ModelsUser | null>(userStore.getUser)

const possibleSideBarItems: sideBar[] = [
  {
    title: t("OCSERV_USER"),
    icon: "mdi-account-network-outline",
    to: "/oc_user",
    value: "oc_user",
    permission: "oc_user"
  },
  {
    title: t("OCSERV_GROUP"),
    icon: "mdi-account-group-outline",
    to: "/oc_group",
    value: "oc_group",
    permission: "oc_group"
  },
  {title: t("OCCTL_COMMANDS"), icon: "mdi-bash", to: "/occtl", value: "occtl", permission: "occtl"},
  {
    title: t("STATISTICS"),
    icon: "mdi-chart-bar-stacked",
    to: "/statistic",
    value: "statistic",
    permission: "statistic"
  },
  {title: t("SERVER_LOG"), icon: "mdi-math-log", to: "/log", value: "see_server_log", permission: "see_server_log"},
  {title: t("OCSERV_SYSTEM"), icon: "mdi-cogs", to: "/oc_system", value: "system", permission: "system"},
]

const sidebarItems = reactive<sideBar[]>([
  {title: t("HOME"), icon: "mdi-home", to: "/", value: "home"},
])

const userSidebarItems = reactive<sideBar[]>([
  {title: t("CONFIG"), icon: "mdi-cog", to: "/config", value: "config", divider: true},
  {title: t("CHANGE_PASSWORD"), icon: "mdi-account-key", to: "/change_password", value: "change_password"}
])


onMounted(() => {
  let user = userStore.getUser
  if (user) {
    if (user.is_admin) {
      sidebarItems.push(...possibleSideBarItems);
    } else {
      Object.keys(user.permission).forEach((permissionKey: string) => {
        possibleSideBarItems.forEach((item: sideBar) => {
          if (item.permission === permissionKey) {
            sidebarItems.push(item);
          }
        });
      });
    }
    if (user.is_admin) {
      userSidebarItems.unshift({
        title: t("STAFF_MANAGEMENT"),
        icon: 'mdi-account-tie-hat',
        to: "/staffs",
        value: "staffs"
      });
    }
  }
});

watch(() => userStore.getUser, (newVal) => {
  user.value = newVal
})

watch(
    () => useIsMobile.isMobile,
    (newVal) => {
      rail.value = newVal
    },
    {immediate: true}
)

</script>

<template>
  <v-navigation-drawer
      v-if="user"
      v-model="drawer"
      :rail="rail"
      permanent
      @click="useIsMobile.isMobile? false : rail = false"
  >
    <v-list>
      <v-list-item :prepend-avatar="avatarUrl">
        <v-list-item-title>
          <span class="text-capitalize">
            {{ user.username }}
          </span>
        </v-list-item-title>

        <v-list-item-subtitle>
          <span class="text-capitalize">
            {{ user.is_admin ? t('ADMIN') : t('STAFF') }}
          </span>
        </v-list-item-subtitle>

        <template v-slot:append>
          <v-btn
              v-if="!useIsMobile.isMobile"
              icon="mdi-chevron-left"
              variant="text"
              @click.stop="rail = !rail"
          />
        </template>

      </v-list-item>
    </v-list>

    <v-divider></v-divider>

    <v-list density="compact" nav>

      <v-list-item
          v-for="(i, index) in sidebarItems"
          :key="`sidebar-items-${index}`"
          :prepend-icon="i.icon"
          :title="i.title"
          :to="i.to"
          :value="i.value"
          color="success"
      />

      <v-divider class="mt-3 mb-5"/>

      <v-list-item
          v-for="(i, index) in userSidebarItems"
          :key="`user-sidebar-items-${index}`"
          :prepend-icon="i.icon"
          :title="i.title"
          :to="i.to"
          :value="i.value"
          color="primary"
      />

    </v-list>

    <template v-if="!rail" #append>
      <v-divider class="mb-2"/>
      <div style="text-align: center; font-size: 0.9rem; color: #555; margin-bottom: 10px">
        <div>Built with ❤️ in 2025</div>
        <div>
          Need help? Contact
          <a
              href="https://github.com/mmtaee/ocserv-users-management/issues"
              style="color: #007BFF; text-decoration: none;"
              target="_blank"
          >
            Github
          </a>
        </div>
      </div>
    </template>

  </v-navigation-drawer>

</template>
