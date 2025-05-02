<script lang="ts" setup>

import {defineAsyncComponent, onBeforeUnmount, onMounted, ref} from "vue";
import {PanelApi, type PanelLogin} from "@/api";
import {useUserStore} from "@/stores/user.ts";
import {useConfigStore} from "@/stores/config.ts";
import router from "@/plugins/router.ts";

const DesktopView = defineAsyncComponent(() => import("@/components/login/desktop/index.vue"))
const MobileView = defineAsyncComponent(() => import("@/components/login/mobile/index.vue"))

const isMobile = ref(false)
const loading = ref(false)

function checkIsMobile() {
  isMobile.value = window.innerWidth <= 768
}

onMounted(() => {
  checkIsMobile()
  window.addEventListener('resize', checkIsMobile)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', checkIsMobile)
})

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
  <component :is="isMobile ? MobileView : DesktopView" :loading="loading" @login="login"/>
</template>
