<script lang="ts" setup>
import {PanelApi, type PanelRequestSetup, type PanelResponseSetup} from "@/api";
import {onBeforeUnmount, onMounted, ref} from "vue";
import DesktopView from "../components/setup/desktop/index.vue"
import {useUserStore} from "@/stores/user.ts";
import router from "@/plugins/router.ts";
import {useConfigStore} from "@/stores/config.ts";
import type {AxiosError} from "axios";

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

function finalize(data: PanelRequestSetup) {
  console.log("data in send: ", data.config)
  console.log("data in send: ", data.default_ocserv_group)
  const userStore = useUserStore()
  const configStore = useConfigStore()
  const api = new PanelApi()

  api.panelSetupPost({request: data}).then((res: PanelResponseSetup) => {
    localStorage.setItem("token", res.token)
    if (res.user) {
      userStore.setUser(...res.user)
    }
    configStore.setSetupComplete()
    router.push({name: "HomePage"})
  }).catch((err: AxiosError) => {
    console.log(err.response?.data)
  })
}

</script>

<template>
  <component :is="isMobile ? MobileView : DesktopView" @finalize="finalize"/>
</template>