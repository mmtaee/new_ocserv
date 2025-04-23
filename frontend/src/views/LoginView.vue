<script lang="ts" setup>

import {onBeforeUnmount, onMounted, ref} from "vue";
import DesktopView from "@/components/login/desktop/index.vue";
import MobileView from "@/components/login/mobile/index.vue";


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

const login = () => {

}

</script>

<template>
  <component :is="isMobile ? MobileView : DesktopView" @login="login"/>
</template>

<style scoped>

</style>