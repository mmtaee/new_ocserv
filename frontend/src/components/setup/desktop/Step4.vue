<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {ModelsOcservUserOrGroupConfigs} from "@/api";
import {reactive, ref, toRaw} from "vue";
import {ipOrRangeRule, ipRule} from "@/utils/rules.ts";

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


const rules = {
  ip: (v: string) => ipRule(v, t),
  ipOrRange: (v: string) => ipOrRangeRule(v, t),
}

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
        <div class="text-subtitle-2  mx-2">{{ t('IP_&_DNS_CONFIGURATION') }}</div>
      </v-col>

      <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
        <p class="mx-8 mb-4 text-grey-darken-1">
          <v-icon color="primary">mdi-bullhorn-outline</v-icon>
          {{ t('These settings control how clients receive IP addresses and resolve domain names while connected') }}.
        </p>
      </v-col>

      <v-col cols="12" md="8" sm="12">
        <v-text-field
            v-model="formValues.ipv4Network"
            :hint="t('THE_POOL_OF_ADDRESSES_THAT_LEASES_WILL_BE_GIVEN_FROM')"
            :label="t('IPV4NETWORK')"
            :rules="[rules.ip]"
            class="my-3"
            clearable
            dense
            variant="solo-filled"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="8" sm="12">
        <v-text-field
            v-model="formValues.dns"
            :hint="t('THE_ADVERTISED_DNS_SERVER')"
            :label="t('DNS')"
            :rules="[rules.ip]"
            class="my-3"
            clearable
            dense
            variant="solo-filled"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
      <v-col cols="12" md="8" sm="12">
        <v-text-field
            v-model="formValues.nbns"
            :hint="t('NETBIOS_NAME_SERVERS_(WINS)_FOR_WINDOWS_CLIENTS')"
            :label="t('NBNS')"
            :rules="[rules.ip]"
            class="my-3"
            clearable
            dense
            variant="solo-filled"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>

      <v-col cols="12" md="8" sm="12">
        <v-text-field
            v-model="formValues.dns"
            :hint="t('STATIC_IPV4_ADDRESS_TO_ASSIGN_TO_CLIENT')"
            :label="t('EXPLICIT_IPV4')"
            :rules="[rules.ip]"
            class="my-3"
            clearable
            dense
            variant="solo-filled"
            @keyup="sendResult"
            @click:clear="sendResult"
        />
      </v-col>
    </v-row>
  </v-form>

</template>

<!--https://chatgpt.com/c/6802383a-1398-8005-99c1-02e5fdd3c858-->