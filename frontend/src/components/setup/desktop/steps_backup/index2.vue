<script lang="ts" setup>
import {defineAsyncComponent, ref} from "vue";
import {useLocale} from "vuetify/framework";
import type {
  ModelsOcservUserOrGroupConfigs,
  SetupRequestSetup,
  SetupRequestSetupAdmin,
  SetupRequestSetupConfig
} from "@/api";

const step = ref(1)
const {t} = useLocale()
const loading = ref(false)
const formIsValid = ref(false)
const emit = defineEmits(["finalize"])
const finalizeData: SetupRequestSetup = {
  admin: null,
  config: null,
  default_ocserv_group: null,
}


const finalize = () => {
  emit("finalize", finalizeData)
}

const adminHandler = (result: SetupRequestSetupAdmin) => {
  formIsValid.value = result?.valid
  delete result.valid
  finalizeData.admin = {...result}
}

const configHandler = (result: SetupRequestSetupConfig) => {
  formIsValid.value = result?.valid
  delete result.valid
  finalizeData.config = {...result}
}

const ocservDefaultGroupHandler = (result: ModelsOcservUserOrGroupConfigs) => {
  formIsValid.value = result?.valid
  delete result.valid
  finalizeData.default_ocserv_group = {...result}
}


function nextStep() {
  if (step.value < steps.length + 1) {
    steps[step.value - 1].complete = true
    if (step.value === steps.length) {
      steps[steps.length - 1].complete = true
      loading.value = true
      emit("finalize", finalize)
    }
    step.value++
  }
}

function prevStep() {
  if (step.value > 1) {
    steps[step.value - 2].complete = false
    step.value--
  }
}

const steps = [
  {
    complete: false,
    title: t('ADMIN_REGISTRY_SETUP'),
    component: defineAsyncComponent(() => import("./step1.vue")),
    data: "admin",
    handler: adminHandler,
  },
  {
    complete: false,
    title: t('SITE_CONFIG_SETUP'),
    component: defineAsyncComponent(() => import("./step2.vue")),
    data: "config",
    handler: configHandler,
  },
  {
    complete: false,
    title: t('DEFAULT_OCSERV_GROUP_SETUP'),
    component: defineAsyncComponent(() => import("./step3.vue")),
    handler: ocservDefaultGroupHandler,
    data: "default_ocserv_group",
  },
  {
    complete: false,
    title: t('FINALIZATION_SETUP'),
    component: defineAsyncComponent(() => import("./step4.vue")),
    loading: false,
    data: null,
  },
]

</script>

<template>

  <v-card color="secondary" width="70%">
    <v-card-title class="pa-3 ma-3 text-white">
      {{ t('SETUP_SERVICE') }}
    </v-card-title>

    <v-card-text class="my-3">
      <v-stepper alt-labels non-linear>
        <v-stepper-header dark vertical>
          <v-stepper-item
              v-for="(st, index) in steps"
              :key="index + 1"
              :complete="st.complete"
              :disabled="step !== index + 1"
              :value="index + 1"
          >
            <template v-slot:title>
              {{ st.title }}
            </template>
          </v-stepper-item>
        </v-stepper-header>

        <v-stepper-window v-for="(st, index) in steps" :key="index + 1" class="ma-0 ml-3 mt-2">
          <component
              :is="st.component"
              v-if="Boolean(st.component) && step === index + 1"
              :data="st.data ? finalizeData[st.data] : finalizeData "
              @sendResult="st.handler"
          />
        </v-stepper-window>

        <v-stepper-actions>
          <template v-slot:next>
            <v-btn
                :disabled="step > steps.length + 1 || !formIsValid"
                :loading="loading"
                color="primary"
                variant="tonal"
                @click="nextStep"
            >
              {{ step < steps.length ? t('Next_BTN') : t('FINALIZATION_CONFIG') }}
            </v-btn>
          </template>

          <template v-slot:prev>
            <v-btn
                :disabled="step === 1 || loading"
                :loading="false"
                color="primary"
                variant="tonal"
                @click="prevStep"
            >
              {{ t('PREV_BTN') }}
            </v-btn>
          </template>
        </v-stepper-actions>

      </v-stepper>
    </v-card-text>
  </v-card>
</template>

<style scoped>

</style>