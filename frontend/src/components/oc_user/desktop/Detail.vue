<script lang="ts" setup>
import {formatDate} from "@/utils/dates.ts";
import {useI18n} from "vue-i18n";
import type {OcOcservUser} from "@/api";
import {defineAsyncComponent} from "vue";

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

defineProps<{
  modelValue: boolean
  item: OcOcservUser,
}>()

const emit = defineEmits(["update:modelValue"])

const {t} = useI18n()

</script>

<template>
  <Modal
      :persistent="false"
      :show="modelValue"
      width="450"
      @close="emit('update:modelValue', false)"
  >
    <template #dialogTitle>
      <span class="text-capitalize">{{ t("OCSERV_USER_DETAIL_TITLE") }} </span>
    </template>

    <template #dialogText>
      <div class="text-subtitle-2">
        <div class="mb-1">
          <span class="text-info">{{ t("USERNAME") }}: </span>
          <span class="text-grey-darken-2">{{ item.username }}</span>
        </div>
        <div class="mb-1">
          <span class="text-info">{{ t("CREATED_AT") }}: </span>
          <span class="text-grey-darken-2">{{ formatDate(item.created_at, "") }}</span>
        </div>
        <div class="mb-1">
          <span class="text-info">{{ t("UPDATED_AT") }}: </span>
          <span class="text-grey-darken-2">{{ formatDate(item.updated_at, "") }}</span>
        </div>
        <div class="mb-1">
          <span class="text-info">{{ t("EXPIRE_AT") }}: </span>
          <span class="text-grey-darken-2">{{ formatDate(item.expire_at, "") }}</span>
        </div>
        <div class="mb-1">
          <span class="text-info">{{ t("DEACTIVATED_AT") }}: </span>
          <span v-if="item.deactivated_at" class="text-grey-darken-2">{{ formatDate(item.deactivated_at, "") }}</span>
          <span v-else class="text-grey-darken-2">{{ t("USER_ACTIVE_SEEN") }}</span>
        </div>
        <div class="mb-1">
          <span class="text-info">{{ t("DESCRIPTION") }}: </span>
          <p class="text-grey-darken-2 ms-5 text-subtitle-2">{{ item.description || t("NO_DESCRIPTION") }}</p>
        </div>
      </div>
    </template>
    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="emit('update:modelValue', false)"
      />
    </template>
  </Modal>
</template>

<style scoped>

</style>