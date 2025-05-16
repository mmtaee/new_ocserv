<script lang="ts" setup>
import {computed, defineAsyncComponent, reactive, ref} from "vue";
import {useLocale} from "vuetify/framework";
import {requiredRule} from "@/utils/rules.ts";
import type {PanelLoginData} from "@/api";
import {useConfigStore} from "@/stores/config.ts";
import {storeToRefs} from "pinia";


const Captcha = defineAsyncComponent(() => import('@/components/login/Captcha.vue'));

defineProps({
  loading: {
    type: Boolean,
    default: false,
  },
});


const emit = defineEmits(['login'])

const {t} = useLocale()
const valid = ref(true)
const loginData: PanelLoginData = reactive<PanelLoginData>({
  username: "",
  password: "",
  remember_me: false,
})
const configStore = useConfigStore()
const {config} = storeToRefs(configStore)
const rules = {
  required: (v: string) => requiredRule(v, t),
}

const login = () => {
  emit("login", loginData)
}

const btnDisable = computed(() => {
  if (Boolean(config.value.googleCaptchaSiteKey)) {
    return !(loginData.token != null && valid.value);
  }
  return !valid.value;
})

</script>
<template>
  <v-card
      max-height="800"
      max-width="700"
      min-width="500"
  >
    <v-card-text class="flex-grow-1 overflow-auto">
      <v-row align="center" justify="center">
        <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
          <div style="text-align: center;">
              <span class="text-h5">
                {{ t('LOGIN') }}
              </span>
          </div>
        </v-col>
      </v-row>

      <v-divider class="mt-5"/>

      <v-col class="ma-0 pa-0 mt-5" cols="12" md="12" sm="12">
        <v-form v-model="valid">
          <v-row align="center" justify="center">
            <v-col class="ma-0 pa-0 mt-5" cols="12" md="7" sm="12">
              <v-text-field
                  v-model="loginData.username"
                  :label="t('ADMIN_USERNAME')"
                  :rules="[rules.required]"
                  density="comfortable"
                  prepend-inner-icon="mdi-key"
                  variant="underlined"
              />
            </v-col>

            <v-col class="ma-0 pa-0" cols="12" md="7" sm="12">
              <v-text-field
                  v-model="loginData.password"
                  :class="config?.googleCaptchaSiteKey != null ? 'my-3': 'my-0'"
                  :label="t('ADMIN_PASSWORD')"
                  :rules="[rules.required]"
                  density="comfortable"
                  prepend-inner-icon="mdi-key"
                  variant="underlined"
              />
            </v-col>

            <v-col
                class="ma-0 pa-0" cols="12" md="7"
                sm="12"
            >
              <Captcha
                  v-if="config.googleCaptchaSiteKey"
                  v-model="loginData.token"
                  :siteKey="config.googleCaptchaSiteKey"
              />
            </v-col>

            <v-col
                :class="config?.googleCaptchaSiteKey != null ? 'mt-5': ''"
                class="ma-0 pa-0"
                cols="12"
                md="7"
                sm="12"
            >
              <v-checkbox
                  v-model="loginData.remember_me"
                  :label="t('REMEMBER_ME')"
                  class="pa-0 ma-0"
                  density="compact"
                  hide-details
                  name="remember_me"
              />
            </v-col>
          </v-row>
        </v-form>
      </v-col>
    </v-card-text>

    <v-card-actions class="my-2 mx-2">
      <v-spacer/>
      <v-btn
          :disabled="btnDisable"
          :loading="loading" color="grey"
          variant="outlined"
          @click="login"
      >
        {{ t('LOGIN') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>
