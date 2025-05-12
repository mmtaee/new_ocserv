<script lang="ts" setup>
import {defineAsyncComponent} from "vue";
import {useUserStore} from "@/stores/user.ts";
import {useLocale} from "vuetify/framework";
import router from "@/plugins/router.ts";

defineProps({
  logoutDialog: Boolean
})

const emit = defineEmits(['update:logoutDialog'])
const {t} = useLocale()
const userStore = useUserStore()
const ModalLayout = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const close = () => {
  emit("update:logoutDialog", false)
}


const logout = async () => {
  await userStore.logout()
  close()
  localStorage.removeItem("token")
  await router.push("/login")
}
</script>

<template>
  <ModalLayout :show="logoutDialog">
    <template #dialogTitle>
      {{ t('LOGOUT') }}
    </template>

    <template #dialogText>
      {{ t("Are you sure you want to log out") }}?
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('Close')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="close"
      />
      <v-btn
          :text="t('Logout')"
          class="me-1"
          color="error"
          variant="outlined"
          @click="logout"
      />
    </template>

  </ModalLayout>
</template>
