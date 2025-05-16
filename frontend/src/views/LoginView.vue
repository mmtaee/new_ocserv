<script lang="ts" setup>
import {defineAsyncComponent, ref} from "vue";
import {useUserStore} from "@/stores/user.ts";
import {useConfigStore} from "@/stores/config.ts";
import router from "@/plugins/router.ts";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import type {PanelLogin} from "@/api/models/panel-login.ts";
import {PanelApi} from "@/api";

const DesktopView = defineAsyncComponent(() => import("@/components/login/desktop/index.vue"))
const MobileView = defineAsyncComponent(() => import("@/components/login/mobile/index.vue"))

const loading = ref(false)
const useIsMobile = useIsMobileStore()

const login = async (data: PanelLogin) => {
  loading.value = true
  const api = new PanelApi()
  api.panelLoginPost({
    request: data,
  }).then((res) => {

    const userStore = useUserStore()
    const configStore = useConfigStore()

    userStore.setUser(res.data.user)
    configStore.setSetup(true)

    localStorage.setItem("token", res.data.token)
    router.push("/")
  }).finally(() => {
    loading.value = false
  })
}

</script>

<template>
  <component :is="useIsMobile.isMobile ? MobileView : DesktopView" :loading="loading" @login="login"/>
</template>
