<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {OcOcservDefaultConfigs, PanelSetupData} from "@/api";
import {onMounted, reactive, ref, toRaw} from "vue";
import {domainRule, ipOrRangeRule, ipRule, ipWithRangeRule} from "@/utils/rules.ts";

const emit = defineEmits(['result', "validate"])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps<{
  data: PanelSetupData;
}>();

const formValues: OcOcservDefaultConfigs = reactive<OcOcservDefaultConfigs>({})
const dnsInput = ref("")


const sendResult = () => {
  emit("validate", valid.value)
  emit('result', toRaw(formValues))
}


const rules = {
  ip: (v: string) => ipRule(v, t),
  ipOrRange: (v: string) => ipOrRangeRule(v, t),
  domain: (v: string) => domainRule(v, t),
  ipWithRange: (v: string) => ipWithRangeRule(v, t)
}

function addDNS() {
  let r = dnsInput.value.trim()
  if (r && !formValues.dns?.includes(r)) {
    if (!formValues.dns) {
      formValues.dns = []
    }
    formValues?.dns.push(r)
    dnsInput.value = ''
    sendResult()
  }
}

function removeDNS(r: string) {
  if (formValues.dns) {
    let index = formValues.dns.findIndex((route: string) => route === r)
    formValues.dns.splice(index, 1)
    sendResult()
  }
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
        <div class="text-subtitle-2  mx-2">{{ t('IP_&_DNS_CONFIGURATION') }}</div>
      </v-col>
      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('THESE_SETTINGS_CONTROL_HOW_CLIENTS_RECEIVE_IP_ADDRESSES_AND_RESOLVE_DOMAIN_NAMES_WHILE_CONNECTED') }}.
        </p>
      </v-col>


      <v-col class="ma-0 pa-0 px-10 mb-8 mt-5" cols="12" md="12" sm="12">
        <v-row align="start" justify="center">
          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-text-field
                v-model="dnsInput"
                :hint="t('THE_ADVERTISED_DNS_SERVER')"
                :label="t('DNS')"
                :rules="[rules.domain]"
                clearable
                density="comfortable"
                placeholder="8.8.8.8"
                variant="underlined"
                @keydown.enter="addDNS"
            />
          </v-col>
          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-text-field
                v-model="formValues['ipv4-network']"
                :hint="t('ADDRESS_POOL_FOR_CLIENT_IP_LEASES')"
                :label="t('IPV4NETWORK')"
                :rules="[rules.ipWithRange]"
                clearable
                density="comfortable"
                placeholder="192.168.1.0/24(or 255.255.255.0)"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>

          <v-col class="ma-0 pa-0 px-3 mt-1" cols="12" md="6" sm="12">
            <v-card class="overflow-auto" height="230" subtitle="DNS List: ">
              <v-chip
                  v-for="route in formValues.dns"
                  :key="route"
                  class="ma-2"
                  closable
                  color="primary"
                  @click:close="removeDNS(route)"
              >
                {{ route }}
              </v-chip>
            </v-card>
          </v-col>

          <v-col class="ma-0 pa-0 px-3 mt-5" cols="12" md="6" sm="12">
            <v-row align="center" justify="center">
              <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="12" sm="12">
                <v-text-field
                    v-model="formValues['explicit-ipv4']"
                    :hint="t('STATIC_IPV4_ADDRESS_TO_ASSIGN_TO_CLIENT')"
                    :label="t('EXPLICIT_IPV4')"
                    :rules="[rules.ip]"
                    clearable
                    density="comfortable"
                    placeholder="192.168.1.0"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>
              <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="12" sm="12">
                <v-text-field
                    v-model="formValues.nbns"
                    :hint="t('NETBIOS_NAME_SERVERS_(WINS)_FOR_WINDOWS_CLIENTS')"
                    :rules="[rules.ip]"
                    clearable
                    density="comfortable"
                    label="NBNS"
                    placeholder="192.168.1.3"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>
              <v-col class="ma-0 pa-0 px-3" cols="12" md="12" sm="12">
                <v-text-field
                    v-model="formValues.cgroup"
                    :hint="t('LINUX_CONTROL_GROUP')"
                    clearable
                    density="comfortable"
                    label='CGroup'
                    placeholder="cpuset,cpu:test"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>
              <v-col class="ma-0 pa-0 px-3" cols="12" md="12" sm="12">
                <v-checkbox
                    v-model="formValues['no-udp']"
                    :false-value="false"
                    :label="t('PREVENT_A_UDP_SESSION_(NO-UDP)')"
                    :true-value="true"
                    base-color="grey"
                    density="comfortable"
                    hide-details
                    @change="sendResult"
                />
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-col>
    </v-row>

  </v-form>
</template>
