<script lang="ts" setup>
import {defineAsyncComponent, reactive, ref} from "vue";
import {useLocale} from "vuetify/framework";
import {requiredRule} from "@/utils/rules.ts";
import {UserApi, type UserChangePasswordData} from "@/api";
import {useSnackbarStore} from "@/stores/snackbar.ts";

const {t} = useLocale()
const valid = ref(true)
const loading = ref(false);
const MinimalFormLayout = defineAsyncComponent(() => import("@/components/common/MinimalFormLayout.vue"))
const snackbar = useSnackbarStore()
const rules = {
  required: (v: string) => requiredRule(v, t),
  commonPassword: () => {
    return passwordData.new_password === repeatedPassword.value ? true : t("invalid repeated password")
  },
}

const passwordData = reactive<UserChangePasswordData>({
  current_password: "",
  new_password: "",
})

const repeatedPassword = ref("")

const save = () => {
  loading.value = true
  const api = new UserApi()
  api.userPasswordPost({
    request: passwordData
  }).then((_) => {
    snackbar.show({
      id: 1,
      message: t("Password Changed Successfully"),
      color: 'success',
      timeout: 4000,
    })
  }).finally(() => {
    loading.value = false
  })
}
</script>

<template>
  <MinimalFormLayout :valid="valid" :width="500">
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
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
          />
        </v-col>

        <v-col class="ma-0 pa-0 px-10 mb-5" cols="12" md="12" sm="12">
          <v-text-field
              v-model="repeatedPassword"
              :label="t('Repeated Password')"
              :rules="[rules.required, rules.commonPassword]"
              density="comfortable"
              prepend-inner-icon="mdi-key"
              variant="underlined"
          />
        </v-col>
      </v-form>
    </template>

    <template #formAction>
      <v-btn
          :disabled="!valid || passwordData.new_password !== repeatedPassword"
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
