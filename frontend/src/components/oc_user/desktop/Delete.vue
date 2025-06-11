<script lang="ts" setup>
import {defineAsyncComponent, type PropType} from "vue";
import type {OcOcservUser} from "@/api";
import {useI18n} from "vue-i18n";
import type {OcservDataInterface} from "@/utils/interfaces.ts";

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const props = defineProps({
  deleteDialog: {
    type: Boolean,
    default: false,
  },
  item: {
    type: Object as PropType<OcOcservUser>,
    required: true,
  },
})

const emit = defineEmits(["update:modelValue", "doAction"])

const {t} = useI18n()

const deleteUser = () => {
  let data: OcservDataInterface = {
    uid: props.item.uid,
    data: null
  }
  emit("doAction", "delete", data)
}

</script>

<template>
  <Modal
      :divider="false"
      :persistent="false"
      :show="deleteDialog"
      color="error"
      width="450"
      @close="emit('update:modelValue', false)"
  >
    <template #dialogTitle>
      <span class="text-capitalize">{{ t("OCSERV_USER_DELETE_TITLE") }}</span>
    </template>

    <template #dialogText>
      <div class="text-grey-darken-2 text-h6 mb-3 mt-2"> {{ t("USERNAME") }}: {{ item.username }}</div>
      <div class="text-error"> {{ t("DELETE_USER_QUESTION") }}</div>
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="emit('update:modelValue', false)"
      />
      <v-btn
          :text="t('DELETE')"
          class="me-1"
          color="error"
          variant="outlined"
          @click="deleteUser"
      />
    </template>
  </Modal>
</template>

<style scoped>

</style>