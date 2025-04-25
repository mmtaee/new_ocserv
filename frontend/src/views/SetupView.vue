<script lang="ts" setup>
import {PanelApi, type PanelSetupData} from "@/api";
import {onBeforeUnmount, onMounted, ref} from "vue";
import DesktopView from "../components/setup/desktop/index.vue"
import router from "@/plugins/router.ts";
import {useConfigStore} from "@/stores/config.ts";
import {useUserStore} from "@/stores/user.ts";


const MobileView = null
const isMobile = ref(false)

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

async function finalize(data: PanelSetupData) {
  const api = new PanelApi()
  const res = await api.panelSetupPost({request: data})

  if (res.status !== 201) {
    console.error(res.status)
    return
  }

  const userStore = useUserStore()
  const configStore = useConfigStore()
  
  userStore.setUser(res.data.user)
  configStore.setSetup(true)
  localStorage.setItem("token", res.data.token)
  await router.push("/")
}

</script>

<template>
  <component :is="isMobile ? MobileView : DesktopView" @finalize="finalize"/>
</template>