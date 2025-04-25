<script lang="ts">
declare global {
  interface Window {
    googleRecaptcha?: {
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
import {computed, onBeforeUnmount, onMounted, ref} from 'vue';

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

const emit = defineEmits(['update:modelValue', 'validForm']);
const recaptcha = ref<HTMLElement | null>(null)
const isDarkTheme = computed(() => localStorage.getItem('theme') === 'dark');

// Callbacks
function callbackSuccess(token: string) {
  emit('update:modelValue', token); // v-model
  emit('validForm', true);
}

function callbackError() {
  if (window.googleRecaptcha) window.googleRecaptcha.reset();
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

</script>

<template>
  <div
      ref="recaptcha"
      :data-sitekey="siteKey"
      :data-size="compact ? 'compact' : 'normal'"
      :data-theme="isDarkTheme ? 'dark' : 'light'"
      class="g-recaptcha"
      data-callback="callbackSuccess"
      data-error-callback="callbackError"
      data-expired-callback="callbackExpired"
  />
</template>
