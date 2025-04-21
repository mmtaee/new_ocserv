<script lang="ts" setup>
import {computed, defineAsyncComponent, ref} from 'vue'
import type {OcOcservDefaultConfigs, PanelRequestSetup, PanelRequestSetupConfig} from "@/api";
import {useLocale} from "vuetify/framework";

const step = ref(1)
const {t} = useLocale()
const loading = ref(false)
const formIsValid = ref(false)
const emit = defineEmits(["finalize"])

const skipStep = ref(0)

const finalizeData: PanelRequestSetup = {
  config: null,
  default_ocserv_group: null,
}

const finalize = () => {
  emit("finalize", finalizeData)
}


const configHandler = (data: PanelRequestSetupConfig) => {
  formIsValid.value = data?.valid
  finalizeData.config = {...data.result}
}

const ocservDefaultGroupHandler = (data: OcOcservDefaultConfigs) => {
  formIsValid.value = data?.valid
  delete data.valid
  finalizeData.default_ocserv_group = {...data.result}
}

const steps = [
  {
    value: 1,
    component: defineAsyncComponent(() => import("./Step1.vue")),
    data: null,
    handler: null,
  },
  {
    value: 2,
    component: defineAsyncComponent(() => import("./Step2.vue")),
    data: "config",
    handler: configHandler,
  },
  {
    value: 3,
    component: defineAsyncComponent(() => import("./Step3.vue")),
    data: "default_ocserv_group",
    handler: ocservDefaultGroupHandler,
  },
  {
    value: 4,
    component: defineAsyncComponent(() => import("./Step4.vue")),
    data: "default_ocserv_group",
    handler: ocservDefaultGroupHandler,
  },
  {
    value: 5,
    component: defineAsyncComponent(() => import("./Step5.vue")),
    data: "default_ocserv_group",
    handler: ocservDefaultGroupHandler,
  },
  {
    value: 6,
    component: defineAsyncComponent(() => import("./Step6.vue")),
    data: "default_ocserv_group",
    handler: ocservDefaultGroupHandler,
  },
  {
    value: 7,
    loading: false,
    component: defineAsyncComponent(() => import("./Step7.vue")),
    data: null,
    handler: null,
  },
]

function nextStep() {
  if (step.value < steps.length + 1) {
    if (step.value === steps.length) {
      loading.value = true
      emit("finalize", finalize)
    }
    step.value++
  }
}

function prevStep() {
  if (skipStep.value != 0) {
    step.value = skipStep.value
    skipStep.value = 0
  } else {
    if (step.value > 1) {
      step.value--
    }
  }

}

const nextBtnDisable = computed(() => {
  if (step.value === 1) return false
  return step.value > steps.length + 1 || !formIsValid.value
})

const nextBtnText = computed(() => {
  switch (step.value) {
    case 1:
      return t("START_BTN")
    case steps.length :
      return t("FINALIZATION_SETUP")
    default:
      return t("NEXT_BTN")
  }
})

const skipBtn = () => {
  skipStep.value = step.value
  step.value = steps.length
}


</script>

<template>
  <v-card class="mx-auto d-flex flex-column" height="800" width="800">
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
          v-if="![1,2].includes(step)"
          :disabled="step === 1 || loading"
          :loading="false"
          color="grey"
          variant="outlined"
          @click="prevStep"
      >
        {{ t('PREV_BTN') }}
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn
          v-if="[3,4,5].includes(step)"
          :loading="loading"
          class="mx-2"
          color="warning"
          variant="outlined"
          @click="skipBtn"
      >
        {{ t("SKIP_GROUP_CONFIGURATION_BTN") }}
      </v-btn>
      <v-btn
          :disabled="nextBtnDisable"
          :loading="loading"
          color="primary"
          variant="outlined"
          @click="nextStep"
      >
        {{ nextBtnText }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>