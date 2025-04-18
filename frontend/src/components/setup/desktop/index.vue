<script lang="ts" setup>
import {computed, defineAsyncComponent, ref} from 'vue'
import type {SetupRequestSetup, SetupRequestSetupConfig} from "@/api";
import {useLocale} from "vuetify/framework";

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

const configHandler = (data: SetupRequestSetupConfig) => {
  console.log("data", data)
  formIsValid.value = data?.valid
  finalizeData.config = {...data.result}
}

const steps = [
  {
    value: 1,
    complete: false,
    title: t('SITE_CONFIG_SETUP'),
    component: defineAsyncComponent(() => import("./Step1.vue")),
    data: "admin",
    handler: null,
  },
  {
    value: 2,
    complete: false,
    title: t('SITE_CONFIG_SETUP'),
    component: defineAsyncComponent(() => import("./Step2.vue")),
    data: "admin",
    handler: configHandler,
  },
  {
    value: 3,
    complete: false,
    title: t('SITE_CONFIG_SETUP'),
    component: defineAsyncComponent(() => import("./Step3.vue")),
    data: "admin",
    handler: configHandler,
  },
]


const nextBtnDisable = computed(() => {
  if (step.value === 1) return false
  console.log(step.value > steps.length + 1 )
  console.log(!formIsValid)
  return step.value > steps.length + 1 || !formIsValid.value
})

</script>

<template>
  <v-card class="mx-auto d-flex flex-column" max-height="800" min-height="650" width="800">

    <v-card-text class="flex-grow-1 overflow-auto">

      <v-window
          v-for="(st, index) in steps"
          :key="`step-${index}`"
      >
        <v-window-item :value="st.value">
          <component
              :is="st.component"
              v-if="Boolean(st.component) && step === index + 1"
              :data="st.data ? finalizeData[st.data] : finalizeData "
              @sendResult="st.handler"
          />
        </v-window-item>
      </v-window>
    </v-card-text>

    <v-card-actions class="my-2 mx-2">

      <v-btn
          :disabled="step === 1 || loading"
          :hidden="[1,2].includes(step)"
          :loading="false"
          color="primary"
          variant="tonal"
          @click="prevStep"
      >
        {{ t('PREV_BTN') }}
      </v-btn>

      <v-spacer></v-spacer>

      <v-btn
          :disabled="nextBtnDisable"
          :loading="loading"
          color="primary"
          variant="tonal"
          @click="nextStep"
      >
        {{ step < steps.length ? t('Next_BTN') : t('FINALIZATION_CONFIG') }}
      </v-btn>

    </v-card-actions>


  </v-card>
</template>