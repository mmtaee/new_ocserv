<script lang="ts" setup>
import {defineAsyncComponent, reactive, ref} from "vue";
import {useLocale} from "vuetify/framework";
import {requiredRule} from "@/utils/rules.ts";

const {t} = useLocale()
const valid = ref(true)
const loading = ref(false);
const MinimalFormLayout = defineAsyncComponent(() => import("@/components/common/MinimalFormLayout.vue"))

const rules = {
  required: (v: string) => requiredRule(v, t),
}

const passwordData = reactive({
  current_password: "",
  new_password: "",
})

const repeatedPassword = ref("")

const save = () => {
  loading.value = true
}

const checkPassword = () => {
  valid.value = passwordData.new_password == repeatedPassword.value && valid.value
}

</script>

<template>
  <MinimalFormLayout :valid="valid">
    <template #formTitle>
      {{ t('ChangePassword') }}
    </template>

    <template #formFields>
      <v-form v-model="valid">
        <v-col class="ma-0 pa-0 px-10 mt-5" cols="12" md="12" sm="12">
          <v-text-field
              v-model="passwordData.current_password"
              :label="t('Current Password')"
              :rules="[rules.required]"
              class="mt-5 mb-1"
              clearable
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
          />
        </v-col>

        <v-col class="ma-0 pa-0 px-10" cols="12" md="12" sm="12">
          <v-text-field
              v-model="passwordData.new_password"
              :label="t('New Password')"
              :rules="[rules.required]"
              clearable
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
              @blur="checkPassword"
              @focus="checkPassword"
          />
        </v-col>

        <v-col class="ma-0 pa-0 px-10 mb-5" cols="12" md="12" sm="12">
          <v-text-field
              v-model="repeatedPassword"
              :label="t('Repeated Password')"
              :rules="[rules.required]"
              clearable
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
              @blur="checkPassword"
              @focus="checkPassword"
          />
        </v-col>
      </v-form>
    </template>

    <template #formAction>
      <v-btn
          :disabled="!valid"
          :loading="loading"
          :text="t('Save')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="save"
      />
    </template>

  </MinimalFormLayout>
</template>

<style scoped>

</style>