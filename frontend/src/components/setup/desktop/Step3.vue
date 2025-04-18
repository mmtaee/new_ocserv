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
  console.log(!!formValues.admin_username || !!formValues.admin_password)
  if (!!formValues.admin_username || !!formValues.admin_password) {
    valid.value = false
  }

  emit('sendResult', {
    valid: valid.value,
    ...toRaw(formValues)
  })
}

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}
</script>

<template>
  <v-card-text>
    <v-form v-model="valid">
      <v-card-subtitle class="mb-5">
        admin authentication data
      </v-card-subtitle>
      <v-text-field
          v-model="formValues.admin_username"
          :label="t('ADMIN_USERNAME')"
          :rules="[rules.required]"
          class="my-3"
          clearable
          dense
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />

      <v-text-field
          v-model="formValues.admin_password"
          :label="t('ADMIN_PASSWORD')"
          :rules="[rules.required]"
          class="my-3 mb-5"
          clearable
          dense
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />

      <v-divider/>

      <v-card-subtitle class="mt-5">
        config site login
      </v-card-subtitle>
      <v-text-field
          v-model="formValues.googleCaptchaSiteKey"
          :label="t('GOOGLE_CAPTCHA_SITE_KEY')"
          class="mt-5 mb-1"
          clearable
          dense
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />

      <v-text-field
          v-model="formValues.googleCaptchaSecretKey"
          :label="t('GOOGLE_CAPTCHA_SECRET_KEY')"
          clearable
          dense
          variant="solo-filled"
          @keyup="sendResult"
          @click:clear="sendResult"
      />

    </v-form>
  </v-card-text>

</template>

<style scoped>

</style>