<script lang="ts" setup>
import {PanelApi, type PanelSetupData} from "@/api";
import {defineAsyncComponent, ref} from "vue";
import router from "@/plugins/router.ts";
import {useConfigStore} from "@/stores/config.ts";
import {useUserStore} from "@/stores/user.ts";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";


const DesktopView = defineAsyncComponent(() => import("@/components/setup/desktop/index.vue"))
const MobileView = null

const useIsMobile = useIsMobileStore()
const loading = ref(false);

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
  }).finally(() => {
    loading.value = false
  })
}

</script>

<template>
  <component :is="useIsMobile.isMobile ? MobileView : DesktopView" :loading="loading" @finalize="finalize"/>
</template>
