<script lang="ts" setup>

import {useLocale} from "vuetify/framework";
import {defineAsyncComponent, onMounted, reactive} from "vue";
import {PanelApi, type PanelConfigResponse} from "@/api";

const MinimalFormLayout = defineAsyncComponent(() => import("@/components/common/MinimalFormLayout.vue"))
const {t} = useLocale()

const settings = reactive<PanelConfigResponse>({
  google_captcha_site_key: "",
  google_captcha_secret_key: "",
})

onMounted(async () => {
  const api = new PanelApi()
  const response = await api.panelConfigGet()
  Object.assign(settings, response.data)
})

</script>

<template>
  <MinimalFormLayout>
    <template #formTitle>
      {{ t('login configuration') }}
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
      <v-col class="ma-0 pa-0 px-10" cols="" md="12" sm="12">
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
      <v-col class="ma-0 pa-0 px-10" cols="" md="12" sm="12">

        <v-text-field
            v-model="settings.google_captcha_secret_key"
            :label="t('GOOGLE_CAPTCHA_SECRET_KEY')"
            clearable
            density="comfortable"
            prepend-inner-icon="mdi-key"
            variant="underlined"
        />
      </v-col>
    </template>

    <template #formAction>
      <v-btn
          :text="t('Save')"
          class="me-1"
          color="success"
          variant="outlined"
      />
    </template>

  </MinimalFormLayout>

</template>