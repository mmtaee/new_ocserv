<script lang="ts" setup>
import {defineAsyncComponent} from "vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";

const DesktopView = defineAsyncComponent(() => import("@/components/staff/desktop/index.vue"))
const MobileView = null

const useIsMobile = useIsMobileStore()

</script>

<template>
  <component :is="useIsMobile.isMobile ? MobileView: DesktopView"/>
</template>
