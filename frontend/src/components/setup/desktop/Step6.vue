<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {OcOcservDefaultConfigs, PanelSetupData} from "@/api";
import {onMounted, reactive, ref, toRaw} from "vue";

const emit = defineEmits(['result', "validate"])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps<{
  data: PanelSetupData;
}>();

const formValues: OcOcservDefaultConfigs = reactive<OcOcservDefaultConfigs>({})

const sendResult = () => {
  emit("validate", valid.value)
  emit('result', toRaw(formValues))
}

onMounted(() => {
  if (props.data) {
    const combined = {
      ...toRaw(props.data.default_ocserv_group),
    }
    Object.assign(formValues, combined)
  }
})

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
        <div class="text-subtitle-2  mx-2">{{ t('CONNECTION_BEHAVIOR_&_TIMEOUTS') }}</div>
      </v-col>

      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('THESE_SETTINGS_IMPROVE_STABILITY_AND_AUTOMATICALLY_HANDLE_IDLE_OR_DISCONNECTED_CLIENTS') }}.
        </p>
      </v-col>

      <v-col class="ma-0 pa-0 px-10 mb-2 mt-5" cols="12" md="12" sm="12">
        <v-row align="center" justify="center">
          <v-col class="ma-0 pa-1  px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues.keepalive"
                :hint="t('INTERVAL_IN_SECONDS_TO_SEND_KEEP-ALIVE_PINGS')"
                :label="t('KEEPALIVE')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1  px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues.dpd"
                :hint="t('DEAD_PEER_DETECTION_TIMEOUT_IN_SECONDS')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                label="DPD"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1 px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues['mobile-dpd']"
                :hint="t('DPD_TIMEOUT_SPECIFICALLY_FOR_MOBILE_CLIENTS')"
                :label="t('MOBILE_DPD')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1 pb-3 px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues['idle-timeout']"
                :hint="t('TIME_IN_SECONDS_BEFORE_DISCONNECTING_IDLE_CLIENTS')"
                :label="t('IDLE_TIMEOUT')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1 pb-3 px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues['mobile-idle-timeout']"
                :hint="t('IDLE_TIMEOUT_FOR_MOBILE_CLIENTS')"
                :label="t('MOBILE_IDLE_TIMEOUT')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1 pb-3 px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues['stats-report-time']"
                :hint="t('INTERVAL_IN_SECONDS_FOR_STATS_REPORTING')"
                :label="t('STATS_REPORT_TIME')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1 pb-3 px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues.mtu"
                :hint="t('TUNNEL_INTERFACE_MTU_TO_AVOID_FRAGMENTATION')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                label="MTU"
                placeholder="1400"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
        </v-row>
      </v-col>


      <v-col class="ma-0 pa-0 px-10" cols="12" md="12" sm="12">
        <v-checkbox
            v-model="formValues['tunnel-all-dns']"
            :false-value="false"
            :label="t('FORCE_ALL_DNS_TRAFFIC_THROUGH_THE_VPN_TUNNEL_(TUNNEL-ALL-DNS)')"
            :true-value="true"
            density="comfortable"
            hide-details
            @change="sendResult"
        />
      </v-col>

      <v-col class="ma-0 pa-0 px-10" cols="12" md="12" sm="12">
        <v-checkbox
            v-model="formValues['deny-roaming']"
            :false-value="false"
            :label="t('DISCONNECT_CLIENT_IF_ITS_IP_CHANGES_(DENY-ROAMING)')"
            :true-value="true"
            density="comfortable"
            hide-details
            @change="sendResult"
        />
      </v-col>

    </v-row>
  </v-form>

</template>