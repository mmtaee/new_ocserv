<script lang="ts" setup>
import {PanelApi, type PanelSetupData} from "@/api";
import {defineAsyncComponent, onBeforeUnmount, onMounted, ref} from "vue";

import router from "@/plugins/router.ts";
import {useConfigStore} from "@/stores/config.ts";
import {useUserStore} from "@/stores/user.ts";


const loading = ref(false);
const DesktopView = defineAsyncComponent(() => import("@/components/setup/desktop/index.vue"))

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
  loading.value = true
  const api = new PanelApi()
  const userStore = useUserStore()
  const configStore = useConfigStore()

  api.panelSetupPost({request: data}).then((res) => {
    userStore.setUser(res.data.user)
    configStore.setSetup(true)
    localStorage.setItem("token", res.data.token)
    router.push("/")
  }).finally(()=>{
    loading.value = false
  })

}

</script>

<template>
  <component :is="isMobile ? MobileView : DesktopView" @finalize="finalize" :loading="loading" />
</template>

<style scoped>

</style>