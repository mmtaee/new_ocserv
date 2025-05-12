<script lang="ts" setup>
import {useTheme} from 'vuetify'
import {defineAsyncComponent, onBeforeUnmount, onMounted} from "vue";
import GlobalSnackbar from "@/components/common/GlobalSnackbar.vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";

const CenterLayout = defineAsyncComponent(() => import("@/components/common/CenterLayout.vue"))
const AppBar = defineAsyncComponent(() => import("@/components/app/AppBar.vue"))
const SideBar = defineAsyncComponent(() => import("@/components/app/SideBar.vue"))

const theme = useTheme()
const useIsMobile = useIsMobileStore()

onMounted(() => {
  checkIsMobile()
  window.addEventListener('resize', checkIsMobile)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', checkIsMobile)
})

const checkIsMobile = () => {
  useIsMobile.setIsMobile(
      /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)
      ||
      window.innerWidth <= 870
  )
}

</script>

<template>
  <v-app :theme="theme.global.name.value">

    <SideBar/>
    <AppBar/>

    <v-main>
      <CenterLayout>
        <RouterView/>
      </CenterLayout>
    </v-main>

    <GlobalSnackbar/>
    
  </v-app>
</template>
