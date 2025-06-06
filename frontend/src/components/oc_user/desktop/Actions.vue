<script lang="ts" setup>
import {type OcOcservUser} from "@/api";
import {defineAsyncComponent, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useSnackbarStore} from "@/stores/snackbar.ts";

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const props = defineProps<{
  item: OcOcservUser
}>()

const emit = defineEmits(["updateUser"])

const {t} = useI18n()
const snackbar = useSnackbarStore()

const detailDialog = ref(false)
const lockDialog = ref(false)
const editDialog = ref(false)
const activitiesDialog = ref(false)
const statisticsDialog = ref(false)
const btnLoading = ref(false)


const disconnect = () => {
//   TODO: handle disconnect
  if (!props.item.is_online) return
}


</script>

<template>
  <v-menu>
    <template v-slot:activator="{ props }">
      <v-btn icon="mdi-dots-vertical" v-bind="props" variant="text"></v-btn>
    </template>

    <v-list>
      <v-list-item @click="detailDialog=true">
        <v-list-item-title>{{ t('MORE_DETAIL') }}</v-list-item-title>
      </v-list-item>

      <v-list-item v-if="item.is_online">
        <v-list-item-title>{{ t('DISCONNECT') }}</v-list-item-title>
      </v-list-item>

      <v-list-item @click="lockDialog=true">
        <v-list-item-title>{{ item.is_locked ? t('UNLOCK') : t('LOCK') }}</v-list-item-title>
      </v-list-item>

      <v-list-item @click="editDialog=true">
        <v-list-item-title>{{ t('EDIT') }}</v-list-item-title>
      </v-list-item>

      <v-list-item>
        <v-list-item-title>{{ t('ACTIVITIES') }}</v-list-item-title>
      </v-list-item>

      <v-list-item>
        <v-list-item-title>{{ t('STATISTICS') }}</v-list-item-title>
      </v-list-item>

    </v-list>
  </v-menu>


  <!--  detail dialog -->


  <!--  lock or unlock dialog -->


  <!--  edit dialog -->
  <Modal
      :persistent="false"
      :show="editDialog"
      width="450"
      @close="editDialog=!editDialog"
  >
    <template #dialogTitle>
      <span class="text-capitalize">{{ t("OCSERV_USER_EDIT_TITLE") }}</span>
      <span class="text-grey-darken-2"> ({{ item.username }})</span>
    </template>

    <template #dialogText>
      <div v-if="item.is_locked"> {{ t("UNLOCK_USER_QUESTION") }}</div>
      <div v-else> {{ t("LOCK_USER_QUESTION") }}</div>
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="editDialog=false"
      />
      <v-btn
          :loading="btnLoading"
          :text="item.is_locked ? t('UNLOCK') : t('LOCK')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="lockOrUnlock"
      />
    </template>
  </Modal>


</template>
