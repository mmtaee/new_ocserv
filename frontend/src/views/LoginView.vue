<script lang="ts" setup>

import {onBeforeUnmount, onMounted, ref} from "vue";
import DesktopView from "@/components/login/desktop/index.vue";
import MobileView from "@/components/login/mobile/index.vue";
import {PanelApi, type PanelLogin} from "@/api";
import {useUserStore} from "@/stores/user.ts";
import {useConfigStore} from "@/stores/config.ts";
import router from "@/plugins/router.ts";


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
  const res = await api.panelLoginPost({
    request: data,
  })
  if (res.status !== 200) {
    console.error(res.status)
  } else {
    const userStore = useUserStore()
    await userStore.setUser(res.data.user)
    const configStore = useConfigStore()
    await configStore.setSetup(true)
    localStorage.setItem("token", res.data.token)
    await router.push("/")
  }
  loading.value = false
}

</script>

<template>
  <component :is="isMobile ? MobileView : DesktopView" :loading="loading" @login="login"/>
</template>
