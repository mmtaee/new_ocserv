<script lang="ts" setup>
import {defineAsyncComponent, reactive, watch} from "vue";
import {useLocale} from "vuetify/framework";
import type {PanelConfigData} from "@/api";

const MinimalFormLayout = defineAsyncComponent(() => import("@/components/common/MinimalFormLayout.vue"))

const props = defineProps<{
  settings: PanelConfigData,
  loading: boolean
}>()

const emit = defineEmits(["save"])


const {t} = useLocale()
const settingsForm = reactive(<PanelConfigData>{})


watch(
    () => props.settings,
    (newVal) => {
      if (newVal) {
        Object.assign(settingsForm, props.settings)
      }
    },
    {immediate: true}
)


const save = (): void => {
  emit("save", settingsForm)
}

</script>

<template>
  <MinimalFormLayout>
    <template #formTitle>
      {{ t('CONFIGURATION') }}
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
              v-model="settingsForm.google_captcha_site_key"
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
              v-model="settingsForm.google_captcha_secret_key"
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
          :text="t('SAVE')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="save"
      />
    </template>

  </MinimalFormLayout>
</template>
