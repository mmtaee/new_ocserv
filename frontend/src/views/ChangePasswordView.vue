<script lang="ts" setup>
import {UserApi, type UserChangePasswordData} from "@/api";
import {useSnackbarStore} from "@/stores/snackbar.ts";
import {useLocale} from "vuetify/framework";
import {defineAsyncComponent, ref} from "vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import {getAuthorization} from "@/utils/request.ts";

const DesktopView = defineAsyncComponent(() => import("@/components/change_password/desktop/index.vue"))
const MobileView = null

const {t} = useLocale()
const snackbar = useSnackbarStore()
const loading = ref(false);
const useIsMobile = useIsMobileStore()

const save = (password: UserChangePasswordData) => {
  loading.value = true
  const api = new UserApi()
  api.userPasswordPost({
    ...getAuthorization(),
    request: password
  }).then((_) => {
    snackbar.show({
      id: 1,
      message: t("PASSWORD_CHANGED_SUCCESSFULLY"),
      color: 'success',
      timeout: 4000,
    })
  }).finally(() => {
    loading.value = false
  })
}
</script>

<template>
  <component
      :is="useIsMobile.isMobile ? MobileView : DesktopView"
      :loading="loading"
      @save="save"
  />
</template>
