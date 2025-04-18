<script lang="ts" setup>
import {reactive, ref, toRaw} from 'vue'
import {useLocale} from 'vuetify/framework'
import {ipOrRangeRule, ipRule} from "@/utils/rules.js";
import type {SetupDefaultOcservGroup} from "@/api";

const emit = defineEmits(['sendResult'])
const valid = ref(true)
const {t} = useLocale()

const rules = {
  ip: (v: string) => ipRule(v, t),
  ipOrRange: (v: string) => ipOrRangeRule(v, t),
}

const props = defineProps({
  data: {
    type: Object as SetupDefaultOcservGroup,
    required: true,
  },
})

const formValues: SetupDefaultOcservGroup = reactive<SetupDefaultOcservGroup>({})

const sendResult = () => {
  emit('sendResult', {
    valid: valid.value,
    ...toRaw(formValues),
  })
}

const selectBoolItems = [
  {text: t('TRUE'), value: 'true'},
  {text: t('FALSE'), value: 'false'},
]

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}

</script>

<template>
  <v-form v-model="valid">
    <v-row align="center" class="mx-5" justify="start">
      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues['rx-data-per-sec']"
            :hint="t('BANDWIDTH_RESTRICTIONS_(IN_BYTES/SECOND)')"
            :label="t('RX-DATA-PER-SEC')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues['tx-data-per-sec']"
            :hint="t('BANDWIDTH_RESTRICTIONS_(IN_BYTES/SECOND)')"
            :label="t('TX-DATA-PER-SEC')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.maxSameClient"
            :hint="t('LIMIT_THE_NUMBER_OF_IDENTICAL_CLIENTS')"
            :label="t('MAX_SAME_CLIENT')"
            :min="0"
            clearable
            controlVariant="hidden"
            density="comfortable"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3">
        <v-text-field
            v-model="formValues.dns1"
            :hint="t('THE_ADVERTISED_DNS_SERVER')"
            :label="t('DNS')"
            :rules="[rules.ip]"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
      <v-col cols="12" md="3">
        <v-text-field
            v-model="formValues.dns2"
            :hint="t('THE_ADVERTISED_DNS_SERVER')"
            :label="t('DNS')"
            :rules="[rules.ip]"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3">
        <v-text-field
            v-model="formValues.ipv4Network"
            :hint="t('THE_POOL_OF_ADDRESSES_THAT_LEASES_WILL_BE_GIVEN_FROM')"
            :label="t('IPV4NETWORK')"
            :rules="[rules.ipOrRange]"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3">
        <v-select
            v-model="formValues.noUDP"
            :hint="t('PREVENT_A_UDP_SESSION')"
            :items="selectBoolItems"
            :label="t('NO_UDP')"
            clearable
            item-title="text"
            item-value="value"
            @update:modelValue="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.keepAlive"
            :hideInput="false"
            :hint="t('KEEPALIVE_IN_SECONDS')"
            :inset="false"
            :label="t('KEEPALIVE')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.dpd"
            :hideInput="false"
            :hint="t('DEAD_PEER_DETECTION_IN_SECONDS')"
            :inset="false"
            :label="t('DPD')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.mobileDpd"
            :hideInput="false"
            :hint="t('DEAD_PEER_DETECTION_FOR_MOBILE_CLIENTS_IN_SECONDS')"
            :inset="false"
            :label="t('MOBILE_DPD')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3">
        <v-select
            v-model="formValues.tunnelAllDns"
            :hint="t('TUNNEL_ALL_DNS_QUERIES_VIA_THE_VPN')"
            :items="selectBoolItems"
            :label="t('TUNNEL_ALL_DNS')"
            clearable
            item-title="text"
            item-value="value"
            @update:modelValue="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3">
        <v-select
            v-model="formValues.restrictUserToRoute"
            :items="selectBoolItems"
            :label="t('RESTRICT_USER_TO_ROUTES')"
            clearable
            item-title="text"
            item-value="value"
            @update:modelValue="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.statsReportTime"
            :hideInput="false"
            :inset="false"
            :label="t('STATS_REPORT_TIME(SECONDS)')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.mtu"
            :hideInput="false"
            :inset="false"
            :label="t('MTU')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.idleTimeout"
            :hideInput="false"
            :inset="false"
            :label="t('IDLE_TIMEOUT(SECONDS)')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.mobileIdleTimeout"
            :hideInput="false"
            :inset="false"
            :label="t('MOBILE_IDLE_TIMEOUT(SECONDS)')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="3" sm="12">
        <v-number-input
            v-model="formValues.sessionTimeout"
            :hideInput="false"
            :inset="false"
            :label="t('SESSION_TIMEOUT(SECONDS)')"
            :min="0"
            :reverse="false"
            clearable
            controlVariant="stacked"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
    </v-row>
  </v-form>
</template>

<style scoped>

</style>