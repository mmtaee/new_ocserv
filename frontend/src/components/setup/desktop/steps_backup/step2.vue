<script lang="ts" setup>
import {useLocale} from 'vuetify/framework'
import {reactive, toRaw} from 'vue'
import type {SetupConfig} from "@/api";

const {t} = useLocale()
const emit = defineEmits(['sendResult'])

const props = defineProps({
  data: {
    type: Object as SetupConfig,
    required: true,
  },
})

const formValues: SetupConfig = reactive<SetupConfig>({
  googleCaptchaSiteKey: null,
  googleCaptchaSecretKey: null,
})

const sendResult = () => {
  emit('sendResult', {
    valid: true,
    ...toRaw(formValues),
  })
}

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}
</script>

<template>
  <v-card-text>
    <v-form>
      <v-row align="center" justify="center">
        <v-col cols="12" md="6" sm="12">
          <v-textarea
              v-model="formValues.googleCaptchaSiteKey"
              :label="t('GOOGLE_CAPTCHA_SITE_KEY')"
              hide-details="auto"
              rows="4"
              variant="solo-filled"
              @keyup="sendResult"
              @click:clear="sendResult"
          />
        </v-col>
        <v-col cols="12" md="6" sm="12">
          <v-textarea
              v-model="formValues.googleCaptchaSecretKey"
              :label="t('GOOGLE_CAPTCHA_SECRET_KEY')"
              hide-details="auto"
              rows="4"
              variant="solo-filled"
              @keyup="sendResult"
              @click:clear="sendResult"
          />
        </v-col>
      </v-row>
    </v-form>
  </v-card-text>
</template>

<style scoped>

</style>