<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {OcOcservDefaultConfigs, PanelSetupData} from "@/api";
import {onMounted, reactive, ref, toRaw} from "vue";
import {domainRule, ipOrRangeRule, ipRule} from "@/utils/rules.ts";

const emit = defineEmits(['result', "validate"])
const valid = ref(true)
const {t} = useLocale()

const props = defineProps<{
  data: PanelSetupData;
}>();


const formValues: OcOcservDefaultConfigs = reactive<OcOcservDefaultConfigs>({
  "restrict-user-to-routes": false
})

const routeInput = ref("")
const noRouteInput = ref("")
const splitDNSInput = ref("")


const sendResult = () => {
  emit("validate", valid.value)
  emit('result', toRaw(formValues))
}


const rules = {
  ip: (v: string) => ipRule(v, t),
  ipOrRange: (v: string) => ipOrRangeRule(v, t),
  domain: (v: string) => domainRule(v, t)
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
  if (formValues.route) {
    let index = formValues.route.findIndex((route: string) => route === r)
    formValues.route.splice(index, 1)
    sendResult()
  }
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
  if (formValues["no-route"]) {
    let index = formValues["no-route"].findIndex((route: string) => route === r)
    formValues["no-route"].splice(index, 1)
    sendResult()
  }
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
  if (formValues["split-dns"]) {
    let index = formValues["split-dns"].findIndex((domain: string) => domain === d)
    formValues["split-dns"].splice(index, 1)
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
    }
)

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
                :hint="t('ROUTES_PUSHED_TO_THE_CLIENT_FOR_ROUTING_TRAFFIC')"
                :label="t('ROUTES')"
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
                :hint="t('LIST_OF_NETWORKS_TO_EXCLUDE_FROM_VPN_ROUTING')"
                :label="t('NO-ROUTES')"
                :rules="[rules.ipOrRange]"
                clearable
                density="comfortable"
                placeholder="192.168.1.0/24"
                variant="underlined"
                @keydown.enter="addNoRoute"

            />
          </v-col>

          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-card :subtitle="t('ROUTES:')" class="overflow-auto" height="150">
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
            <v-card :subtitle="t('NO-ROUTES:')" class="overflow-auto" height="150">
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
                :hint="t('DOMAINS_TO_RESOLVE_VIA_VPN_DNS_SERVERS')"
                :label="t('SPLIT_DNS')"
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
                :hint="t('INTERNAL_ROUTE_AVAILABLE')"
                :rules="[rules.ipOrRange]"
                clearable
                density="comfortable"
                label='iroute'
                placeholder="192.168.2.0/24"
                variant="underlined"
                @keyup="sendResult"
                @click:clear="sendResult"
            />
          </v-col>

          <v-col class="ma-0 pa-0 px-3" cols="12" md="6" sm="12">
            <v-card :subtitle="t('SPLIT-DNS:')" class="overflow-auto" height="200">
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
                    :hint="t('PRIORITY_FOR_ROUTES')"
                    :label="t('NET_PRIORITY')"
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
                    :hint="t('LIST_OF_ALLOWED_OR_BLOCKED_PROTOCOLS_AND_PORTS')"
                    :label="t('RESTRICT_USER_TO_PORTS')"
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
                    :false-value="false"
                    :label="t('ALLOW_CLIENT_ACCESS_ONLY_TO_DEFINED_ROUTES_(RESTRICT-USER-TO-ROUTES)')"
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
