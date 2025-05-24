<script lang="ts" setup>
import {type OcOcservUser, OcservUsersApi} from "@/api";
import {defineAsyncComponent, ref} from "vue";
import {useI18n} from "vue-i18n";
import {formatDate} from "@/utils/dates.ts"
import {getAuthorization} from "@/utils/request.ts";
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

const lockOrUnlock = () => {
  btnLoading.value = true
  let data = {
    lock: !props.item.is_locked
  }
  const api = new OcservUsersApi()
  api.ocUsersUidLockPost({
    uid: props.item.uid,
    ...getAuthorization(),
    request: data,
  }).then(() => {
    const user = {...props.item}
    user.is_locked = !user.is_locked
    emit("updateUser", user)
    snackbar.show({
      id: 1,
      message: t("OCSERV_USER_SUCCESSFULLY", {lock: user.is_locked ? t('LOCKED') : t('UNLOCKED')}),
      color: 'success',
      timeout: 4000,
    })
  }).finally(() => {
    btnLoading.value = false
    lockDialog.value = false
  })
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
  <Modal
      :persistent="false"
      :show="detailDialog"
      width="450"
      @close="detailDialog=!detailDialog"
  >
    <template #dialogTitle>
      <span class="text-capitalize">{{ t("OCSERV_USER_DETAIL_TITLE") }} </span>
      <span class="text-grey-darken-2"> ({{ item.username }})</span>
    </template>

    <template #dialogText>
      <div class="text-subtitle-2">
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
          @click="detailDialog=false"
      />
    </template>
  </Modal>

  <!--  lock or unlock dialog -->
  <Modal
      :divider="false"
      :persistent="false"
      :show="lockDialog"
      width="450"
      @close="lockDialog=!lockDialog"
  >
    <template #dialogTitle>
      <span v-if="item.is_locked" class="text-capitalize">{{ t("OCSERV_USER_UNLOCK_TITLE") }}</span>
      <span v-else class="text-capitalize">{{ t("OCSERV_USER_LOCK_TITLE") }}</span>
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
          @click="lockDialog=false"
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

  <!--  edit dialog -->
  <Modal
      :persistent="false"
      :show="editDialog"
      width="450"
      @close="editDialog=!editDialog"
  >
    <template #dialogTitle>
      <span v-if="item.is_locked" class="text-capitalize">{{ t("OCSERV_USER_UNLOCK_TITLE") }}</span>
      <span v-else class="text-capitalize">{{ t("OCSERV_USER_LOCK_TITLE") }}</span>
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
          @click="lockDialog=false"
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
