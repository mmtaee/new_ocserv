<script setup lang="ts">
import {defineAsyncComponent} from "vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";

const useIsMobile = useIsMobileStore()

const DesktopView = defineAsyncComponent(() => import("@/components/staff/desktop/index.vue"))
const MobileView = null


</script>

<template>
  <component :is="useIsMobile.isMobile ? MobileView: DesktopView" />
</template>
