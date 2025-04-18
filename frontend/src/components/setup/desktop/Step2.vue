<script lang="ts" setup>
import {useLocale} from 'vuetify/framework'
import {reactive, ref, toRaw} from 'vue'
import {requiredRule} from '@/utils/rules'
import type {SetupRequestSetupConfig} from "@/api";

const emit = defineEmits(['sendResult'])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps({
  data: {
    type: Object as SetupRequestSetupConfig,
    required: true,
  },
})

const rules = {
  required: (v: string) => requiredRule(v, t),
}

const formValues: SetupRequestSetupConfig = reactive<SetupRequestSetupConfig>({})

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
  <v-row align="center" justify="center">

    <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
      <center>
      <span class="text-h5">
 Admin User & Login Configuration
      </span>
      </center>
    </v-col>

    <blockquote>
      <p class="mx-8 my-4">
        <v-icon>mdi-balloon</v-icon>
        Create your admin account and provide your CAPTCHA site key and secret key.
        These are required to protect the login page from automated bots.
      </p>
    </blockquote>


    <v-col class="ma-0 pa-0" cols="" md="8" sm="12">
      <v-text-field
          v-model="formValues.admin_username"
          :label="t('ADMIN_USERNAME')"
          :rules="[rules.required]"
          class="my-3"
          clearable
          dense
          prepend-inner-icon="mdi-account"
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />
    </v-col>

    <v-col class="ma-0 pa-0" cols="" md="8" sm="12">
      <v-text-field
          v-model="formValues.admin_password"
          :label="t('ADMIN_PASSWORD')"
          :rules="[rules.required]"
          class="my-3 mb-5"
          clearable
          dense
          prepend-inner-icon="mdi-key"
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />
    </v-col>

    <blockquote>
      <p class="mx-8 mt-3">
        <v-icon>mdi-bullhorn-outline</v-icon>
        Only <strong>Google reCAPTCHA v2</strong> is supported.
        You can generate your site and secret keys from
        <a href="https://www.google.com/recaptcha/" rel="noopener noreferrer" target="_blank">Google reCAPTCHA</a>.
      </p>
    </blockquote>

    <v-col class="ma-0 pa-0" cols="" md="8" sm="12">
      <v-text-field
          v-model="formValues.googleCaptchaSiteKey"
          :label="t('GOOGLE_CAPTCHA_SITE_KEY')"
          class="mt-5 mb-1"
          clearable
          dense
          prepend-inner-icon="mdi-key"
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />
    </v-col>

    <v-col class="ma-0 pa-0" cols="" md="8" sm="12">

      <v-text-field
          v-model="formValues.googleCaptchaSecretKey"
          :label="t('GOOGLE_CAPTCHA_SECRET_KEY')"
          clearable
          dense
          prepend-inner-icon="mdi-key"
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />
    </v-col>

  </v-row>
</template>

<style scoped>

</style>