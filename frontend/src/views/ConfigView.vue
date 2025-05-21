<script lang="ts" setup>
import {defineAsyncComponent, onMounted, reactive, ref} from "vue";
import {PanelApi, type PanelConfigData} from "@/api";
import {useSnackbarStore} from "@/stores/snackbar.ts";
import {useLocale} from "vuetify/framework";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import {getAuthorization} from "@/utils/request.ts";


const DesktopView = defineAsyncComponent(() => import("@/components/config/desktop/index.vue"))
const MobileView = null

const loading = ref(false)
const useIsMobile = useIsMobileStore()
const {t} = useLocale()
const snackbar = useSnackbarStore()

const settings = reactive<PanelConfigData>({
  google_captcha_site_key: "",
  google_captcha_secret_key: "",
})


onMounted(async () => {
  const api = new PanelApi()
  const response = await api.panelConfigGet(getAuthorization())

  Object.assign(settings, response.data)
})

const save = (settings: PanelConfigData) => {
  loading.value = true
  const api = new PanelApi()
  api.panelConfigPatch({
    ...getAuthorization(),
    request: settings,
  }).then((response) => {
    Object.assign(settings, response.data)
    snackbar.show({
      id: 1,
      message: t("CONFIG_UPDATED_SUCCESSFULLY"),
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
      :settings="settings"
      @save="save"
  />
</template>