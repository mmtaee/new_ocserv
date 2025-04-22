<script lang="ts" setup>
import {useLocale} from 'vuetify/framework'
import {reactive, ref, toRaw} from 'vue'
import {requiredRule} from '@/utils/rules'
import type {PanelRequestSetupConfig} from "@/api";

const emit = defineEmits(['sendResult'])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps({
  data: {
    type: Object as PanelRequestSetupConfig,
    required: true,
  },
})

const rules = {
  required: (v: string) => requiredRule(v, t),
}

const formValues: PanelRequestSetupConfig = reactive<PanelRequestSetupConfig>({})

const sendResult = () => {
  valid.value = !(!formValues.admin_username || !formValues.admin_password);
  emit('sendResult', {
    valid: valid.value,
    result: toRaw(formValues)
  })
}

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}
</script>

<template>
  <v-form v-model="valid">
    <v-row align="center" justify="center">

      <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
        <div style="text-align: center;">
      <span class="text-h5">
        {{ t('ADMIN_USER_&_LOGIN_CONFIGURATION') }}
      </span>
        </div>
      </v-col>

      <v-divider class="mt-5"/>

      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 my-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('CREATE_YOUR_ADMIN_ACCOUNT_AND_PROVIDE_YOUR_CAPTCHA_SITE_KEY_AND_SECRET_KEY') }}.
          {{ t('THESE_ARE_REQUIRED_TO_PROTECT_THE_LOGIN_PAGE_FROM_AUTOMATED_BOTS') }}.
        </p>
      </v-col>


      <v-col class="ma-0 pa-0 px-10" cols="" md="8" sm="12">
        <v-text-field
            v-model="formValues.admin_username"
            :label="t('ADMIN_USERNAME')"
            :rules="[rules.required]"
            class="my-3"
            clearable
            density="comfortable"
            prepend-inner-icon="mdi-account"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col class="ma-0 pa-0 px-10" cols="12" md="8" sm="12">
        <v-text-field
            v-model="formValues.admin_password"
            :label="t('ADMIN_PASSWORD')"
            :rules="[rules.required]"
            class="my-3 mb-5"
            clearable
            density="comfortable"
            prepend-inner-icon="mdi-key"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mt-3 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('ONLY') }} <strong>Google reCAPTCHA v2</strong> {{ t('IS_SUPPORTED') }}.
          {{ t('YOU_CAN_GENERATE_YOUR_SITE_AND_SECRET_KEYS_FROM') }}
          <a href="https://www.google.com/recaptcha/" rel="noopener noreferrer" target="_blank">Google reCAPTCHA</a>.
        </p>
      </v-col>

      <v-col class="ma-0 pa-0 px-10" cols="" md="8" sm="12">
        <v-text-field
            v-model="formValues.googleCaptchaSiteKey"
            :label="t('GOOGLE_CAPTCHA_SITE_KEY')"
            class="mt-5 mb-1"
            clearable
            density="comfortable"
            prepend-inner-icon="mdi-key"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col class="ma-0 pa-0 px-10" cols="" md="8" sm="12">

        <v-text-field
            v-model="formValues.googleCaptchaSecretKey"
            :label="t('GOOGLE_CAPTCHA_SECRET_KEY')"
            clearable
            density="comfortable"
            prepend-inner-icon="mdi-key"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

    </v-row>
  </v-form>
</template>
