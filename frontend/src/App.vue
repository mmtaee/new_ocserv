<script lang="ts" setup>
import {useTheme} from 'vuetify'
import {defineAsyncComponent} from "vue";
import GlobalSnackbar from "@/components/common/GlobalSnackbar.vue";

const CenterLayout = defineAsyncComponent(() => import("@/components/common/CenterLayout.vue"))
const AppBar = defineAsyncComponent(() => import("@/components/common/AppBar.vue"))

const theme = useTheme()

</script>

<template>
  <v-app :theme="theme.global.name.value">
    <AppBar/>
    <!--    <v-btn @click="toggleTheme">Toggle Theme</v-btn>-->
    <CenterLayout>
      <RouterView/>
    </CenterLayout>
    <GlobalSnackbar/>
  </v-app>
</template>
