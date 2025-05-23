<script lang="ts" setup>
import {computed, defineAsyncComponent, onBeforeMount, reactive, ref} from "vue";
import type {DataTableHeader} from 'vuetify'
import {
  type ModelsUser,
  type RequestMeta,
  UserApi,
  type UserChangeStaffPassword,
  type UserCreateStaffData
} from "@/api";
import {useSnackbarStore} from "@/stores/snackbar.ts";
import {useI18n} from "vue-i18n";
import {requiredRule} from "@/utils/rules.ts";
import {getAuthorization} from "@/utils/request.ts";
import {formatDate} from "@/utils/dates.ts"

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const {t} = useI18n()
const snackbar = useSnackbarStore()
const loading = ref(false)
const btnLoading = ref(false)

const meta = reactive<RequestMeta>({
  page: 1,
  page_size: 10,
  total_records: 0
})
const pageCount = ref(0)
const permDialog = ref(false)
const deleteDialog = ref(false)
const addDialog = ref(false)
const changePasswordDialog = ref(false)
const staffData = reactive<ModelsUser[]>([]);
const selectedUser = reactive<ModelsUser>({
  id: 0,
  is_admin: false,
  last_login: "",
  permission: {
    oc_groups: false,
    oc_users: false,
    occtl: false,
    see_server_logs: false,
    statistic: false,
    system: false,
  },
  uid: "",
  username: ""
});

const createUser = reactive<UserCreateStaffData>({
  password: "",
  permission: {
    oc_groups: false,
    oc_users: false,
    occtl: false,
    see_server_logs: false,
    statistic: false,
    system: false,
  },
  username: ""
})
const createUserValid = ref(true)
const newPassword = reactive<UserChangeStaffPassword>({
  password: "",
});
const changePasswordValid = ref(true)
const headers: DataTableHeader[] = [
  {title: t("ID"), key: 'id', align: 'center'},
  {title: t("USERNAME"), key: 'username', align: 'center'},
  {title: t("LAST_LOGIN"), key: 'last_login', align: 'center'},
  {title: t("ACTIONS"), key: 'actions', align: 'center'},
]
const rules = {
  required: (v: string) => requiredRule(v, t),
}

onBeforeMount(() => {
  fetchStaffs()
})

const permission = (user: ModelsUser) => {
  Object.assign(selectedUser, JSON.parse(JSON.stringify(user)))
  permDialog.value = true
}

const fetchStaffs = () => {
  loading.value = true
  const api = new UserApi()
  api.userStaffsGet(
      {
        ...getAuthorization(),
        page: meta.page,
        pageSize: meta.page_size,
      }
  ).then((res) => {
    staffData.splice(0, staffData.length, ...res.data.result || [])
    Object.assign(meta, res.data.meta)
    pageCount.value = Math.ceil(meta.total_records / meta.page_size);
  }).finally(() => {
    loading.value = false
  })
}


const addStaff = () => {
  btnLoading.value = true
  const api = new UserApi()
  api.userStaffsPost({
    ...getAuthorization(),
    request: createUser,
  }).then((res) => {
    resetCreateUser()
    staffData.push(res.data)
    snackbar.show({
      id: 1,
      message: t("STAFF_CREATED_SUCCESSFULLY", {username: res.data.username}),
      color: 'success',
      timeout: 4000,
    })
    addDialog.value = false
  }).finally(() => {
    btnLoading.value = false
  })
}

const updatePerm = () => {
  let user = staffData.find(user => user.id === selectedUser.id)
  if (user && JSON.stringify(user.permission) === JSON.stringify(selectedUser.permission)) {
    permDialog.value = false
    return
  }

  btnLoading.value = true
  const api = new UserApi()
  api.userStaffsPermissionsIdPut({
    ...getAuthorization(),
    id: selectedUser?.id,
    request: selectedUser?.permission,
  }).then(() => {
    Object.assign(selectedUser, {})
    permDialog.value = false
  }).finally(() => {
    btnLoading.value = false
  })
}

const changePassword = () => {
  btnLoading.value = true
  const api = new UserApi()
  api.userStaffsIdPasswordPost({
    ...getAuthorization(),
    id: selectedUser.id, request: newPassword
  }).then(_ => {
    changePasswordDialog.value = false
    snackbar.show({
      id: 1,
      message: t("STAFF_CHANGE_PASSWORD_SUCCESSFULLY", {username: selectedUser.username}),
      color: 'success',
      timeout: 4000,
    })
  }).finally(() => {
    newPassword.password = ''
    btnLoading.value = false
  })
}

const remove = () => {
  btnLoading.value = false
  const api = new UserApi()
  api.userStaffsIdDelete({
    ...getAuthorization(),
    id: selectedUser.id,
  }).then(_ => {
    let index = staffData.findIndex(i => i.id === selectedUser.id)
    staffData.splice(index, 1)
    deleteDialog.value = false
    snackbar.show({
      id: 1,
      message: t("STAFF_DELETED_SUCCESSFULLY", {username: selectedUser.username}),
      color: 'success',
      timeout: 4000,
    })
  }).finally(() => {
    btnLoading.value = false
  })

}

const resetCreateUser = () => {
  Object.assign(createUser, {
    is_admin: false,
    password: "",
    permissions: {
      oc_user: false,
      oc_group: false,
      occtl: false,
      see_server_log: false,
      statistic: false,
      system: false
    },
    username: ""
  })
}

const permKeyWrapper = (key: string) => {
  switch (key) {
    case 'oc_user':
      return t("OCSERV_USER")
    case "oc_group":
      return t("OCSERV_GROUP")
    case "statistic":
      return t("STATISTICS")
    case "occtl":
      return t("OCCTL_COMMANDS")
    case "system":
      return t("OCSERV_SYSTEM")
    case "see_server_log":
      return t("SERVER_LOG")
    default:
      return key
  }
}

const sortedCreatePermissionKeys = computed<(keyof typeof createUser.permission)[]>(() =>
    Object.keys(createUser.permission)
        .sort((a, b) => a.localeCompare(b)) as (keyof typeof createUser.permission)[]
)

const sortedUpdatePermissionKeys = computed<(keyof typeof selectedUser.permission)[]>(() =>
    Object.keys(selectedUser.permission)
        .sort((a, b) => a.localeCompare(b)) as (keyof typeof selectedUser.permission)[]
)

</script>

<template>
  <v-card max-height="900" min-width="800">
    <v-card-text>
      <v-data-table
          :headers="headers"
          :items="staffData"
          :items-length="meta.total_records"
          :loading="loading"
          :no-data-text="t('No Data Found')"
          disable-sort
          mobile-breakpoint="sm"
      >
        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title>
              <v-icon class="mb-1" color="medium-emphasis" icon="mdi-account-tie-hat" size="large" start></v-icon>
              {{ t("STAFFS") }}
            </v-toolbar-title>
            <v-btn
                :text="t('ADD_STAFF')"
                border
                class="me-2"
                prepend-icon="mdi-plus"
                rounded="lg"
                @click="addDialog=true"
            />
          </v-toolbar>
        </template>

        <template #item.last_login="{ item }">
          {{ formatDate(item.last_login, t("NOT_LOGIN_YET")) }}
        </template>

        <template #item.actions="{ item }">
          <v-icon
              color="warning"
              end
              icon="mdi-door-closed-lock"
              @click="permission(item)"
          />
          <v-icon
              color="primary"
              end
              icon="mdi-key"
              @click="Object.assign(selectedUser, item); changePasswordDialog = true"
          />
          <v-icon
              color="error"
              end
              icon="mdi-delete"
              @click="Object.assign(selectedUser, item); deleteDialog = true"
          />
        </template>

        <template v-slot:bottom>
          <div v-if="pageCount>1" class="text-center pt-2">
            <v-pagination
                v-model="meta.page"
                :length="pageCount"
                @update:modelValue="fetchStaffs"
            />
          </div>
        </template>

      </v-data-table>
    </v-card-text>


  </v-card>

  <Modal :persistent="false" :show="addDialog" width="650" @close="addDialog=!addDialog">
    <template #dialogTitle>
      {{ t("CREATE_STAFF") }}
    </template>

    <template #dialogText>
      <v-form v-model="createUserValid">
        <v-row align="center" class="pa-1" justify="start">

          <v-col cols="12" md="6" sm="6">
            <v-text-field
                v-model="createUser.username"
                :label="t('USERNAME')"
                :rules="[rules.required]"
                clearable
                density="comfortable"
                variant="underlined"
            />
          </v-col>

          <v-col cols="12" md="6" sm="6">
            <v-text-field
                v-model="createUser.password"
                :label="t('PASSWORD')"
                :rules="[rules.required]"
                density="comfortable"
                variant="underlined"
            />
          </v-col>

          <v-col cols="12" md="12" sm="12">
            <span class="text-grey-darken-2"> {{ t("PERMISSIONS") }}:</span>
          </v-col>

          <v-divider/>

          <v-col cols="12" md="12" sm="6">
            <v-row align="center" class="ms-auto mt-2" justify="center">
              <v-col
                  v-for="(key, index) in sortedCreatePermissionKeys"
                  :key="`perm-${index}`"
                  class="ma-0 pa-0"
                  cols="12"
                  md="4"
                  sm="4"
              >
                <v-checkbox
                    v-model="createUser.permission[key]"
                    :label="permKeyWrapper(key)"
                    color="success"
                    density="compact"
                    hide-details
                />
              </v-col>
            </v-row>
          </v-col>

        </v-row>
      </v-form>
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="addDialog=false;resetCreateUser()"
      />
      <v-btn
          :disabled="!createUserValid"
          :loading="btnLoading"
          :text="t('ADD_STAFF')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="addStaff"
      />
    </template>

  </Modal>

  <Modal :show="changePasswordDialog" width="350" @close="changePasswordDialog=!changePasswordDialog">
    <template #dialogTitle>
      {{ t("CHANGE_STAFF_PASSWORD_TITLE") }} <span class="text-capitalize">({{ selectedUser.username }})</span>
    </template>

    <template #dialogText>
      <v-form v-model="changePasswordValid">
        <v-row align="center" justify="center">
          <v-col cols="12" md="12" sm="6">
            <v-text-field
                v-model="newPassword.password"
                :label="t('NEW_PASSWORD')"
                :rules="[rules.required]"
                density="comfortable"
                variant="underlined"
            />
          </v-col>
        </v-row>
      </v-form>
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="changePasswordDialog=false;Object.assign(selectedUser, {});newPassword.password=''"
      />
      <v-btn
          :disabled="!changePasswordValid"
          :loading="btnLoading"
          :text="t('CHANGE_PASSWORD')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="changePassword"
      />
    </template>
  </Modal>

  <Modal :show="deleteDialog" width="auto" @close="deleteDialog=!deleteDialog">
    <template #dialogTitle>
      {{ t("DELETE_STAFF_TITLE") }} <span class="text-capitalize">({{ selectedUser.username }})</span>
    </template>
    <template #dialogText>
      {{ t('DO_YOU_WANT_TO_DELETE_THE_STAFF') }}?
    </template>
    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="deleteDialog=false;Object.assign(selectedUser, {})"
      />
      <v-btn
          :loading="btnLoading"
          :text="t('DELETE')"
          class="me-1"
          color="error"
          variant="outlined"
          @click="remove"
      />
    </template>
  </Modal>

  <Modal :persistent="false" :show="permDialog" width="auto" @close="permDialog=!permDialog">
    <template #dialogTitle>
      {{ t("STAFF_PERMISSIONS_ROUTES_ACCESS_TITLE") }}
    </template>
    <template #dialogText>
      <v-row align="center" class="ms-auto mt-2" justify="center">
        <v-col
            v-for="(key, index) in sortedUpdatePermissionKeys"
            :key="`perm-${index}`"
            class="ma-0 pa-0"
            cols="12"
            md="4"
            sm="4"
        >
          <v-checkbox
              v-model="selectedUser.permission[key]"
              :label="permKeyWrapper(key)"
              color="success"
              density="compact"
              hide-details
          />
        </v-col>

      </v-row>
    </template>

    <template #dialogAction>
      <v-btn
          :text="t('CLOSE')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="permDialog=false;Object.assign(selectedUser, {})"
      />
      <v-btn
          :loading="btnLoading"
          :text="t('UPDATE')"
          class="me-1"
          color="error"
          variant="outlined"
          @click="updatePerm"
      />
    </template>
  </Modal>
</template>
