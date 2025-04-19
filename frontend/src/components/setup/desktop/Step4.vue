<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {ModelsOcservUserOrGroupConfigs} from "@/api";
import {reactive, ref, toRaw} from "vue";
import {ipOrRangeRule, ipRule, ipWithNetmaskRule} from "@/utils/rules.ts";

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
const dnsInput = ref("")


const sendResult = () => {
  emit('sendResult', {
    valid: valid.value,
    result: toRaw(formValues)
  })
}


const rules = {
  ip: (v: string) => ipRule(v, t),
  ipOrRange: (v: string) => ipOrRangeRule(v, t),
  ipWithNetmask: (v: string) => ipWithNetmaskRule(v, t),
}

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}


function addDNS() {
  console.log("192.168.10.2: ", dnsInput.value)
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
  let index = formValues.dns.findIndex((route: string) => route === r)
  formValues.dns.splice(index, 1)
  sendResult()
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
        <div class="text-subtitle-2  mx-2">{{ t('IP_&_DNS_CONFIGURATION') }}</div>
      </v-col>
      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('These settings control how clients receive IP addresses and resolve domain names while connected') }}.
        </p>
      </v-col>


      <v-col class="ma-0 pa-0 px-10 mb-8 mt-5" cols="12" md="12" sm="12">
        <v-row align="start" justify="center">
          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-text-field
                v-model="dnsInput"
                :hint="t('THE_ADVERTISED_DNS_SERVER')"
                :label="t('DNS')"
                :rules="[rules.ip]"
                clearable
                density="comfortable"
                placeholder="8.8.8.8"
                variant="underlined"
                @keydown.enter="addDNS"
            />
          </v-col>
          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-text-field
                v-model="formValues.ipv4Network"
                :hint="t('Address pool for client IP leases')"
                :label="t('IPV4NETWORK')"
                :rules="[rules.ipWithNetmask]"
                clearable
                density="comfortable"
                placeholder="192.168.1.0/255.255.255.0"
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
                    :label="t('NBNS')"
                    :rules="[rules.ip]"
                    clearable
                    density="comfortable"
                    placeholder="192.168.1.3"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>
              <v-col class="ma-0 pa-0 px-3" cols="12" md="12" sm="12">
                <v-text-field
                    v-model="formValues.cgroup"
                    :hint="t('Linux control group')"
                    :label="t('CGroup')"
                    :rules="[rules.ip]"
                    clearable
                    density="comfortable"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-form>

</template>

<!--https://chatgpt.com/c/6802383a-1398-8005-99c1-02e5fdd3c858-->