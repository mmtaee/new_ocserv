<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {ModelsOcservUserOrGroupConfigs} from "@/api";
import {reactive, ref, toRaw} from "vue";

const emit = defineEmits(['sendResult'])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps({
  data: {
    type: Object as ModelsOcservUserOrGroupConfigs,
    required: true,
  },
})

const formValues: ModelsOcservUserOrGroupConfigs = reactive<ModelsOcservUserOrGroupConfigs>({})

const sendResult = () => {
  emit('sendResult', {
    valid: valid.value,
    result: toRaw(formValues)
  })
}

// const rules = {
// required: (v: string) => requiredRule(v, t),
// }

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
            {{ t('DEFAULT_OCSERV_GROUP_CONFIGURATION') }}
          </span>
        </div>
      </v-col>

      <v-divider class="mt-5 mb-2"/>

      <v-col class="ma-0 px-5" cols="12" md="12" sm="12">
        <div class="text-subtitle-2  mx-2">{{ t('BANDWIDTH_&_SESSION_LIMITS') }}</div>
      </v-col>

      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('HELPS_CONTROL_RESOURCE_USAGE_AND_PREVENT_ACCOUNT_SHARING_OR_ABUSE_OF_BANDWIDTH') }}.
        </p>
      </v-col>

      <v-col class="ma-0 pa-1 px-10" cols="12" md="8" sm="12">
        <v-number-input
            v-model="formValues['rx-data-per-sec']"
            :hint="t('BANDWIDTH_RESTRICTIONS_(IN_BYTES/SECOND)')"
            :label="t('RX_DATA')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
      <v-col class="ma-0 pa-1 px-10" cols="12" md="8" sm="12">
        <v-number-input
            v-model="formValues['tx-data-per-sec']"
            :hint="t('BANDWIDTH_RESTRICTIONS_(IN_BYTES/SECOND)')"
            :label="t('TX_DATA')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
      <v-col class="ma-0 pa-1 px-10" cols="12" md="8" sm="12">
        <v-number-input
            v-model="formValues.maxSameClient"
            :hint="t('LIMIT_THE_NUMBER_OF_IDENTICAL_CLIENTS')"
            :label="t('MAX_SAME_CLIENT')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
      <v-col class="ma-0 pa-1 pb-3 px-10" cols="12" md="8" sm="12">
        <v-number-input
            v-model="formValues.sessionTimeout"
            :hint="t('MAX_SESSION_TIME_IN_SECONDS_BEFORE_FORCED_DISCONNECT')"
            :label="t('SESSION_TIMEOUT')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            variant="underlined"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
    </v-row>
  </v-form>

</template>

<!--https://chatgpt.com/c/6802383a-1398-8005-99c1-02e5fdd3c858-->