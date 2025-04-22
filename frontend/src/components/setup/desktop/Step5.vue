<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {OcOcservDefaultConfigs} from "@/api";
import {reactive, ref, toRaw} from "vue";
import {domainRule, ipOrRangeRule, ipRule} from "@/utils/rules.ts";

const emit = defineEmits(['sendResult'])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps({
  data: {
    type: Object as OcOcservDefaultConfigs,
    required: true,
  },
})

const formValues: OcOcservDefaultConfigs = reactive<OcOcservDefaultConfigs>({})

const routeInput = ref("")
const noRouteInput = ref("")
const splitDNSInput = ref("")


const sendResult = () => {
  emit('sendResult', {
    valid: valid.value,
    result: toRaw(formValues)
  })
}


const rules = {
  ip: (v: string) => ipRule(v, t),
  ipOrRange: (v: string) => ipOrRangeRule(v, t),
  domain: (v: string) => domainRule(v, t)
}

if (props.data) {
  Object.assign(formValues, structuredClone(props.data))
}

function addRoute() {
  let r = routeInput.value.trim()
  if (r && !formValues.route?.includes(r) && !formValues["no-route"]?.includes(r)) {
    if (!formValues.route) {
      formValues.route = []
    }
    formValues.route.push(r)
    routeInput.value = ''
    sendResult()
  }
}

function removeRoute(r: string) {
  let index = formValues.route.findIndex((route: string) => route === r)
  formValues.route.splice(index, 1)
  sendResult()
}

function addNoRoute() {
  let r = noRouteInput.value.trim()
  if (r && !formValues.route?.includes(r) && !formValues["no-route"]?.includes(r)) {
    if (!formValues["no-route"]) {
      formValues["no-route"] = []
    }
    formValues["no-route"].push(r)
    noRouteInput.value = ''
    sendResult()
  }
}

function removeNoRoute(r: string) {
  let index = formValues["no-route"].findIndex((route: string) => route === r)
  formValues["no-route"].splice(index, 1)
  sendResult()
}


function addDomain() {
  let r = splitDNSInput.value.trim()
  if (r && !formValues["split-dns"]?.includes(r)) {
    if (!formValues["split-dns"]) {
      formValues["split-dns"] = []
    }
    formValues["split-dns"].push(r)
    splitDNSInput.value = ''
    sendResult()
  }
}

function removeDomain(d: string) {
  let index = formValues["split-dns"].findIndex((domain: string) => domain === d)
  formValues["split-dns"].splice(index, 1)
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
        <div class="text-subtitle-2  mx-2">{{ t('ROUTING_&_ACCESS_CONTROL') }}</div>
      </v-col>
      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('THESE_SETTINGS_CONTROL_HOW_CLIENTS_RECEIVE_IP_ADDRESSES_AND_RESOLVE_DOMAIN_NAMES_WHILE_CONNECTED') }}.
        </p>
      </v-col>

      <v-col class="ma-0 pa-0 px-10 mb-2 mt-5" cols="12" md="12" sm="12">
        <v-row align="center" justify="center">
          <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="6" sm="12">
            <v-text-field
                v-model="routeInput"
                :hint="t('Routes pushed to the client for routing traffic')"
                :label="t('routes')"
                :rules="[rules.ipOrRange]"
                clearable
                density="comfortable"
                placeholder="192.168.1.0/24"
                variant="underlined"
                @keydown.enter="addRoute"
            />
          </v-col>

          <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="6" sm="12">
            <v-text-field
                v-model="noRouteInput"
                :hint="t('List of networks to exclude from VPN routing')"
                :label="t('no-routes')"
                :rules="[rules.ipOrRange]"
                clearable
                density="comfortable"
                placeholder="192.168.1.0/24"
                variant="underlined"
                @keydown.enter="addNoRoute"

            />
          </v-col>

          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-card :subtitle="t('routes:')" class="overflow-auto" height="150">
              <v-chip
                  v-for="route in formValues.route"
                  :key="route"
                  class="ma-2"
                  closable
                  color="primary"
                  @click:close="removeRoute(route)"
              >
                {{ route }}
              </v-chip>
            </v-card>
          </v-col>

          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-card :subtitle="t('no routes:')" class="overflow-auto" height="150">
              <v-chip
                  v-for="route in formValues['no-route']"
                  :key="route"
                  class="ma-2"
                  closable
                  color="primary"
                  @click:close="removeNoRoute(route)"
              >
                {{ route }}
              </v-chip>
            </v-card>
          </v-col>
        </v-row>
      </v-col>

      <v-col class="ma-0 pa-0 px-10 mb-8 mt-5" cols="12" md="12" sm="12">
        <v-row align="start" justify="center">
          <v-col class="ma-0 pa-0 px-3 mb-1 mt-2" cols="12" md="6" sm="12">
            <v-text-field
                v-model="splitDNSInput"
                :hint="t('Domains to resolve via VPN DNS servers')"
                :label="t('Split DNS')"
                :rules="[rules.domain]"
                clearable
                density="comfortable"
                placeholder="example.com"
                variant="underlined"
                @keydown.enter="addDomain"
            />
          </v-col>
          <v-col class="ma-0 pa-0 px-3 mb-1 mt-2" cols="12" md="6" sm="12">
            <v-text-field
                v-model="formValues.iroute"
                :hint="t('Internal route available')"
                :label="t('iroute')"
                :rules="[rules.ipOrRange]"
                clearable
                density="comfortable"
                placeholder="192.168.2.0/24"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>

          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-card :subtitle="t('Split DNS: ')" class="overflow-auto" height="200">
              <v-chip
                  v-for="domain in formValues['split-dns']"
                  :key="domain"
                  class="ma-2"
                  closable
                  color="primary"
                  @click:close="removeDomain(domain)"
              >
                {{ domain }}
              </v-chip>
            </v-card>
          </v-col>


          <v-col cols="12" md="6" sm="12">
            <v-row align="center" justify="center">
              <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="12" sm="12">
                <v-number-input
                    v-model="formValues['net-priority']"
                    :hint="t('Priority for routes')"
                    :label="t('net priority')"
                    :min="0"
                    clearable
                    controlVariant="hidden"
                    density="comfortable"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>


              <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="12" sm="12">
                <v-text-field
                    v-model="formValues['restrict-user-to-ports']"
                    :hint="t('list of allowed or blocked protocols and ports')"
                    :label="t('restrict user to ports')"
                    clearable
                    density="comfortable"
                    placeholder="tcp(443),tcp(80) or !(tcp(443),tcp(80)"
                    variant="underlined"
                    @keyup="sendResult"
                    @click:clear="sendResult"
                />
              </v-col>

              <v-col class="ma-0 pa-0 px-3 mb-1" cols="12" md="12" sm="12">
                <v-checkbox
                    v-model="formValues['restrict-user-to-routes']"
                    :label="t('Allow client access only to defined routes (restrict-user-to-routes)')"
                    base-color="grey"
                    density="comfortable"
                    hide-details
                    @click="sendResult"
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