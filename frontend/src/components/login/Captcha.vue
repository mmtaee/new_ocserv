<script lang="ts">
declare global {
  interface Window {
    grecaptcha?: {
      reset: () => void
      render: (...args: any[]) => any
      getResponse: (widgetId?: any) => string
      [key: string]: any
    }
    callbackSuccess: (token: string) => void
    callbackError: () => void
    callbackExpired: () => void
  }
}
</script>

<script lang="ts" setup>
import {onBeforeUnmount, onMounted, ref, watch} from 'vue';
import {useTheme} from "vuetify/framework";

const emit = defineEmits(['update:modelValue', 'validForm']);
const recaptcha = ref<HTMLElement | null>(null);

defineProps({
  modelValue: String, // v-model value
  compact: {
    type: Boolean,
    default: false,
  },
  siteKey: {
    type: String,
    required: true
  }
});


// Callbacks
function callbackSuccess(token: string) {
  emit('update:modelValue', token); // v-model
}

function callbackError() {
  if (window.grecaptcha) window.grecaptcha.reset();
  emit('validForm', false);
}

function callbackExpired() {
  emit('validForm', false);
}

// Script setup
let recaptchaScript: HTMLScriptElement | null = null

onMounted(() => {
  recaptchaScript = document.createElement("script");
  recaptchaScript.src = "https://www.google.com/recaptcha/api.js";
  recaptchaScript.async = true;
  recaptchaScript.defer = true;
  recaptchaScript.id = "captcha";
  document.head.appendChild(recaptchaScript);

  // Bind to window
  window.callbackSuccess = callbackSuccess;
  window.callbackError = callbackError;
  window.callbackExpired = callbackExpired;
});

onBeforeUnmount(() => {
  const existingScript = document.getElementById("captcha");
  if (existingScript) existingScript.remove();
  if (recaptcha.value) recaptcha.value.remove();
});

const theme = useTheme()
const mode = ref(theme.global.name.value)

watch(
    () => theme.global.name.value,
    async (newVal) => {
      mode.value = newVal
      if (window?.grecaptcha) {
        const captchaElement = document.getElementById("recaptcha");
        if (captchaElement) {
          captchaElement.setAttribute("data-theme", newVal);
        }
      }

    }
)
</script>

<template>
  <div
      id="recaptcha"
      ref="recaptcha"
      :data-sitekey="siteKey"
      :data-size="compact ? 'compact' : 'normal'"
      :data-theme="mode"
      class="g-recaptcha"
      data-callback="callbackSuccess"
      data-error-callback="callbackError"
      data-expired-callback="callbackExpired"
      style="min-width: 100px; min-height: 75px; padding: 0; margin: 0;"
  />
</template>

