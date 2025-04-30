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

onMounted(()=>{
if (props.data) {
  const combined = {
    ...toRaw(props.data.default_ocserv_group),
  }
  Object.assign(formValues, combined)
}})

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
        <div class="text-subtitle-2  mx-2">{{ t('Connection Behavior & Timeouts') }}</div>
      </v-col>

      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('These settings improve stability and automatically handle idle or disconnected clients') }}.
        </p>
      </v-col>

      <v-col class="ma-0 pa-0 px-10 mb-2 mt-5" cols="12" md="12" sm="12">
        <v-row align="center" justify="center">
          <v-col class="ma-0 pa-1  px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues.keepalive"
                :hint="t('Interval in seconds to send keep-alive pings')"
                :label="t('keepalive')"
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
                :hint="t('Dead Peer Detection timeout in seconds')"
                :label="t('DPD')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>
          <v-col class="ma-0 pa-1 px-3" cols="12" md="6" sm="12">
            <v-number-input
                v-model="formValues['mobile-dpd']"
                :hint="t('DPD timeout specifically for mobile clients')"
                :label="t('mobile DPD')"
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
                :hint="t('Time in seconds before disconnecting idle clients')"
                :label="t('IDLE timeout')"
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
                :hint="t('Idle timeout for mobile clients')"
                :label="t('mobile IDLE timeout')"
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
                :hint="t('Interval in seconds for stats reporting')"
                :label="t('stats report time')"
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
                :hint="t('Tunnel interface MTU to avoid fragmentation')"
                :label="t('MTU')"
                :min="0"
                clearable
                controlVariant="hidden"
                density="comfortable"
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
            :label="t('Force all DNS traffic through the VPN tunnel (tunnel-all-dns)')"
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
            :label="t('Disconnect client if its IP changes (deny-roaming)')"
            :true-value="true"
            density="comfortable"
            hide-details
            @change="sendResult"
        />
      </v-col>

    </v-row>
  </v-form>

</template>