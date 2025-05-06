<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import {defineAsyncComponent, onMounted, reactive, ref} from "vue";
import {PanelApi, type PanelConfigData} from "@/api";
import {useSnackbarStore} from "@/stores/snackbar.ts";

const MinimalFormLayout = defineAsyncComponent(() => import("@/components/common/MinimalFormLayout.vue"))
const {t} = useLocale()
const loading = ref(false)
const snackbar = useSnackbarStore()

const settings = reactive<PanelConfigData>({
  google_captcha_site_key: "",
  google_captcha_secret_key: "",
})

const settingsClone = reactive<PanelConfigData>({
  google_captcha_site_key: "",
  google_captcha_secret_key: "",
})

onMounted(async () => {
  const api = new PanelApi()
  const response = await api.panelConfigGet()

  Object.assign(settings, response.data)
  Object.assign(settingsClone, structuredClone(response.data))
})


const save = () => {
  loading.value = true
  const api = new PanelApi()
  api.panelConfigPatch({
    request: settings,
  }).then((response) => {
    Object.assign(settings, response.data)
    Object.assign(settingsClone, structuredClone(response.data))
    snackbar.show({
      id: 1,
      message: t("Config Updated Successfully"),
      color: 'success',
      timeout: 4000,
    })
  }).finally(() => {
    loading.value = false
  })
}


</script>

<template>
  <MinimalFormLayout>
    <template #formTitle>
      {{ t('configuration') }}
    </template>

    <template #formHint>
      <p class="mx-8 mt-3 text-grey-darken-1">
        <v-icon color="primary">mdi-bullhorn-outline</v-icon>
        {{ t('ONLY') }} <strong>Google reCAPTCHA v2</strong> {{ t('IS_SUPPORTED') }}.
        {{ t('YOU_CAN_GENERATE_YOUR_SITE_AND_SECRET_KEYS_FROM') }}
        <a href="https://www.google.com/recaptcha/" rel="noopener noreferrer" target="_blank">Google reCAPTCHA</a>.
      </p>
    </template>

    <template #formFields>
      <v-form>
        <v-col class="ma-0 pa-0 px-10 mt-5" cols="12" md="12" sm="12">
          <v-text-field
              v-model="settings.google_captcha_site_key"
              :label="t('GOOGLE_CAPTCHA_SITE_KEY')"
              class="mt-5 mb-1"
              clearable
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
          />
        </v-col>
        <v-col class="ma-0 pa-0 px-10 mb-5" cols="12" md="12" sm="12">

          <v-text-field
              v-model="settings.google_captcha_secret_key"
              :label="t('GOOGLE_CAPTCHA_SECRET_KEY')"
              clearable
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
          />
        </v-col>
      </v-form>
    </template>

    <template #formAction>
      <v-btn
          :loading="loading"
          :text="t('Reset')"
          class="me-1"
          color="grey"
          variant="outlined"
          @click="Object.assign(settings, settingsClone)"
      />

      <v-spacer/>

      <v-btn
          :loading="loading"
          :text="t('Save')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="save"
      />
    </template>

  </MinimalFormLayout>

</template>