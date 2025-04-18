<script lang="ts" setup>
import {PanelApi, type SetupRequestSetup, type SetupResponseSetup} from "@/api";
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


function finalize(data: SetupRequestSetup) {
  const userStore = useUserStore()
  const configStore = useConfigStore()
  const api = new PanelApi()

  api.panelSetupPost({request: data}).then((res: SetupResponseSetup) => {
    configStore.setSetupComplete()
    localStorage.setItem("token", res.token)
    if (res.user) {
      userStore.setUser(...res.user)
    }
    router.push({name: "HomePage"})
  }).catch((err: AxiosError) => {
    console.log(err)
  })
}


</script>

<template>
  <component :is="isMobile ? MobileView : DesktopView" @finalize="finalize"/>
</template>