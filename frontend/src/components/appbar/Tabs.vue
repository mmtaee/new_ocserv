<script lang="ts" setup>
import router from "@/plugins/router.ts";
import {onBeforeMount, type PropType, reactive, ref, watch} from "vue";
import type {ModelsUser} from "@/api";
import {useLocale} from "vuetify/framework";
import {useRoute} from "vue-router";

interface Tab {
  key: number
  title: string
  icon: string
  path: string
  permission?: string
}

const props = defineProps({
  user: {
    type: Object as PropType<ModelsUser>,
    required: true,
  },
})

const route = useRoute();
const {t} = useLocale()
const tab = ref<number | null>(null)
const possibleTabs: Tab[] = [
  {key: 2, title: t("OCSERV_USERS"), icon: "mdi-account-network-outline", path: "/oc_user", permission: "oc_user"},
  {key: 3, title: t("OCSERV_GROUPS"), icon: "mdi-account-group-outline", path: "/oc_group", permission: "oc_group"},
  {key: 4, title: t("OCCTL_COMMANDS"), icon: "mdi-bash", path: "/occtl", permission: "occtl"},
  {key: 5, title: t("STATISTICS"), icon: "mdi-chart-bar-stacked", path: "/statistic", permission: "statistic"},
  {key: 6, title: t("SERVER_LOG"), icon: "mdi-math-log", path: "/log", permission: "see_server_log"},
  {key: 8, title: t("OCSERV_SYSTEM"), icon: "mdi-cogs", path: "/oc_system", permission: "system"},
]

const tabs = reactive<Tab[]>([
  {key: 1, title: t("HOME"), icon: "mdi-home", path: "/"},
])


onBeforeMount(() => {
  if (props.user) {
    if (props.user.is_super_admin) {
      tabs.push(...possibleTabs);
    } else {
      Object.keys(props.user.permission).forEach((permissionKey: string) => {
        possibleTabs.forEach((tab: Tab) => {
          if (tab.permission === permissionKey) {
            tabs.push(tab);
          }
        });
      });
    }
    if (props.user.is_admin) {
      tabs.push({key: 9, title: t("STAFF_MANAGEMENT"), icon: 'mdi-account-tie-hat', path: "/staffs",});
    }
  }
});

watch(
    () => route.path,
    (newPath) => {
      if (props.user) {
        const exist = tabs.find((tab) => tab.path === newPath);
        if (exist) {
          tab.value = exist.key;
        } else {
          tab.value = null;
        }
      }
    },
    {immediate: true}
);

</script>

<template>
  <v-row v-if="user" align="start" justify="start">
    <v-tabs
        v-model="tab"
        align-tabs="start"
        color="primary"
        hide-slider
        stacked
    >
      <v-tab v-for="(t, index) in tabs" :key="`tab-${index}`" :value="t.key" @click="router.push(t.path)">
        <v-icon :icon="t.icon" size="x-large"/>
        <span style="font-size: 12px">{{ t.title }}</span>
      </v-tab>
    </v-tabs>
  </v-row>
</template>
