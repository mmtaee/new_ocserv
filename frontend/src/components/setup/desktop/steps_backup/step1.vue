<script lang="ts" setup>
import {useLocale} from 'vuetify/framework'
import {reactive, ref, toRaw} from 'vue'
import {requiredRule} from '@/utils/rules.js'
import type {SetupAdmin} from "@/api";

const emit = defineEmits(['sendResult'])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps({
  data: {
    type: Object as SetupAdmin,
    required: true,
  },
})

const rules = {
  required: (v: string) => requiredRule(v, t),
}

const formValues: SetupAdmin = reactive<SetupAdmin>({})

const sendResult = () => {
  if (!Boolean(formValues.username) || !Boolean(formValues.password)) {
    valid.value = false
  } else {
    emit('sendResult', {
      valid: valid.value,
      ...toRaw(formValues)
    })
  }
}

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}
</script>

<template>
  <v-card-text>
    <v-form v-model="valid">
      <v-row align="center" justify="center">
        <v-col class="pa-0 ma-0" cols="12" md="6" sm="12">
          <v-text-field
              v-model="formValues.username"
              :label="t('USERNAME')"
              :rules="[rules.required]"
              variant="solo-filled"
              @keyup="sendResult"
              @click:clear="sendResult"
          />
        </v-col>
      </v-row>
      <v-row align="center" justify="center">
        <v-col class="pa-0 ma-0" cols="12" md="6" sm="12">
          <v-text-field
              v-model="formValues.password"
              :label="t('PASSWORD')"
              :rules="[rules.required]"
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