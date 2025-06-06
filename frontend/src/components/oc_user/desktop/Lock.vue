<script lang="ts" setup>
import {defineAsyncComponent} from "vue";
import type {OcOcservUser} from "@/api";
import {useI18n} from "vue-i18n";
import type {OcservDataInterface} from "@/utils/interfaces.ts";

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))
const props = defineProps<{
  lockDialog: boolean,
  item: OcOcservUser,
}>()

const emit = defineEmits(["closeDialogs", "doAction"])

const {t} = useI18n()

const lock = () => {
  let data: OcservDataInterface = {
    data: {
      lock: !props.item.is_locked,
    },
    uid: props.item.uid,
  }
  emit("doAction", "lock", data)
}

</script>

<template>
  <Modal
      :divider="false"
      :persistent="false"
      :show="lockDialog"
      width="450"
      @close="emit('closeDialogs')"
  >
    <template #dialogTitle>
      <span v-if="item.is_locked" class="text-capitalize">{{ t("OCSERV_USER_UNLOCK_TITLE") }}</span>
      <span v-else class="text-capitalize">{{ t("OCSERV_USER_LOCK_TITLE") }}</span>

    </template>

    <template #dialogText>
      <div class="text-grey-darken-2 text-h6 mb-3 mt-2"> {{ t("USERNAME") }}: {{ item.username }}</div>
      <div v-if="item.is_locked"> {{ t("UNLOCK_USER_QUESTION") }}</div>
      <div v-else> {{ t("LOCK_USER_QUESTION") }}</div>
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="emit('closeDialogs')"
      />
      <v-btn
          :text="item.is_locked ? t('UNLOCK') : t('LOCK')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="lock"
      />
    </template>
  </Modal>
</template>

<style scoped>

</style>