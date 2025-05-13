<script lang="ts" setup>
import {computed, defineAsyncComponent, reactive, ref, toRaw} from 'vue'
import type {OcOcservDefaultConfigs, PanelSetupData} from "@/api";
import {useLocale} from "vuetify/framework";
import type {AdminConfigurations} from "@/components/setup/types.ts";

const Step1 = defineAsyncComponent(() => import("./Step1.vue"));
const Step2 = defineAsyncComponent(() => import("./Step2.vue"));
const Step3 = defineAsyncComponent(() => import("./Step3.vue"));
const Step4 = defineAsyncComponent(() => import("./Step4.vue"));
const Step5 = defineAsyncComponent(() => import("./Step5.vue"));
const Step6 = defineAsyncComponent(() => import("./Step6.vue"));
const Step7 = defineAsyncComponent(() => import("./Step7.vue"));

const props = defineProps<{
  loading: boolean,
}>();


const step = ref(1)
const {t} = useLocale()
const emit = defineEmits(["finalize"])
const skipStep = ref(0)
const formIsValid = ref(false)
const finalizeData = reactive(<PanelSetupData>{admin: {}, default_ocserv_group: {}})


const validate = (v: boolean) => {
  formIsValid.value = v
}

const configHandler = (data: AdminConfigurations) => {
  finalizeData.admin.username = data.username;
  finalizeData.admin.password = data.password;
  if (!finalizeData.config) {
    finalizeData.config = {
      google_captcha_site_key: "",
      google_captcha_secret_key: ""
    }
  }
  finalizeData.config.google_captcha_site_key = data.google_captcha_site_key || ""
  finalizeData.config.google_captcha_secret_key = data.google_captcha_secret_key || ""
}

const groupHandler = (data: OcOcservDefaultConfigs) => {
  const combined = {
    ...structuredClone(toRaw(data))
  }
  Object.assign(finalizeData.default_ocserv_group, combined)
}


const steps = [
  {
    value: 1,
    component: Step1,
    handler: null,
  },
  {
    value: 2,
    component: Step2,
    handler: configHandler,
  },
  {
    value: 3,
    component: Step3,
    handler: groupHandler,
  },
  {
    value: 4,
    component: Step4,
    handler: groupHandler,
  },
  {
    value: 5,
    component: Step5,
    handler: groupHandler,
  },
  {
    value: 6,
    component: Step6,
    handler: groupHandler,
  },
  {
    value: 7,
    loading: props.loading,
    component: Step7,
    handler: null,
  },
]

function nextStep() {
  if (step.value < steps.length + 1) {
    if (step.value === steps.length) {
      emit("finalize", finalizeData)
      return
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
      return t("FINALIZATION_SETUP_BTN")
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
              :data="finalizeData"
              @result="st.handler"
              @validate="validate"
          />
        </v-window-item>
      </v-window>
    </v-card-text>

    <v-card-actions class="my-2 mx-2">
      <v-btn
          v-if="![1,2].includes(step)"
          :disabled="step === 1 || loading"
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