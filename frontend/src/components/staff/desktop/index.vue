<script lang="ts" setup>
import {computed, defineAsyncComponent, onBeforeMount, reactive, ref} from "vue";
import type {DataTableHeader} from 'vuetify'
import {type ModelsUser, UserApi, type UserChangeStaffPassword, type UserCreateStaffData} from "@/api";
import {useSnackbarStore} from "@/stores/snackbar.ts";
import {useI18n} from "vue-i18n";
import {requiredRule} from "@/utils/rules.ts";
import {useUserStore} from "@/stores/user.ts";

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const {t} = useI18n()
const loading = ref(false)
const btnLoading = ref(false)
const totalItems = ref(0)
const page = ref(1)
const pageSize = ref(10)
const totalRecord = ref(0)
const pageCount = ref(0)
const permDialog = ref(false)
const deleteDialog = ref(false)
const addDialog = ref(false)
const changePasswordDialog = ref(false)
const staffData = reactive<ModelsUser[]>([]);
const selectedUser = reactive<ModelsUser>({
  id: 0,
  is_admin: false,
  is_super_admin: false,
  last_login: "",
  permission: {
    oc_group: false,
    oc_user: false,
    occtl: false,
    see_server_log: false,
    statistic: false,
    system: false,
  },
  uid: "",
  username: ""
});

const createUser = reactive<UserCreateStaffData>({
  is_admin: false,
  password: "",
  permission: {
    oc_group: false,
    oc_user: false,
    occtl: false,
    see_server_log: false,
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
  {title: t("Username"), key: 'username', align: 'center'},
  {title: t("Last Login"), key: 'last_login', align: 'center'},
  {title: t("Action"), key: 'actions', align: 'center'},
]
const rules = {
  required: (v: string) => requiredRule(v, t),
}
const userStore = useUserStore()

const snackbar = useSnackbarStore()

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
  api.userStaffsGet().then((res) => {
    Object.assign(staffData, res.data.result)
    page.value = res.data.meta.page
    pageSize.value = res.data.meta.page_size
    totalRecord.value = res.data.meta?.total_records || 0
    pageCount.value = Math.ceil(totalRecord.value / pageSize.value);
  }).finally(() => {
    loading.value = false
  })
}

const formatDate = (dateString: string | undefined): string => {
  if (!dateString) {
    return t("Not Login Yet")
  }
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}`;
}

const addStaff = () => {
  btnLoading.value = true
  const api = new UserApi()
  api.userStaffsPost({
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
  api.userStaffsIdPasswordPost({id: selectedUser.id, request: newPassword}).then(_ => {
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

const createUserPermissionManager = () => {
  if (createUser.is_admin) {
    Object.keys(createUser.permission).forEach(key => {
      createUser.permission[key as keyof typeof createUser.permission] = true
    })
  } else {
    Object.keys(createUser.permission).forEach(key => {
      createUser.permission[key as keyof typeof createUser.permission] = false
    })
  }
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
      return t("Ocserv User")
    case "oc_group":
      return t("Ocserv Group")
    case "statistic":
      return t("Statistic")
    case "occtl":
      return t("Occtl Commands")
    case "system":
      return t("Ocserv System")
    case "see_server_log":
      return t("Server Log")
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


const isAllowAction = (item: ModelsUser, action: string) => {
  let user = userStore.getUser
  if (user) {
    if (user.is_super_admin && !["delete", "perm"].includes(action)) return true
    return user.id !== item.id && !item.is_admin
  }
}

</script>

<template>
  <v-card max-height="900" min-width="800">
    <v-card-text>
      <v-data-table
          :headers="headers"
          :items="staffData"
          :items-length="totalItems"
          :loading="loading"
          :no-data-text="t('No Data Found')"
          disable-sort
          mobile-breakpoint="sm"
      >
        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title>
              <v-icon class="mb-1" color="medium-emphasis" icon="mdi-account-tie-hat" size="large" start></v-icon>
              {{ t("Staffs") }}
            </v-toolbar-title>
            <v-btn
                :text="t('Add Staff')"
                border
                class="me-2"
                prepend-icon="mdi-plus"
                rounded="lg"
                @click="addDialog=true"
            />
          </v-toolbar>
        </template>
        <template #item.last_login="{ item }">
          {{ formatDate(item.last_login) }}
        </template>
        <template #item.actions="{ item }">
          <v-icon
              v-if="isAllowAction(item, 'perm')"
              color="warning"
              end
              icon="mdi-door-closed-lock"
              @click="permission(item)"
          />
          <v-icon v-else color="warning" disabled end icon="mdi-door-closed-cancel"/>

          <v-icon
              v-if="isAllowAction(item, 'password')"
              color="primary"
              end
              icon="mdi-key"
              @click="Object.assign(selectedUser, item); changePasswordDialog = true"
          />

          <v-icon v-else color="primary" disabled end icon="mdi-key"/>

          <v-icon
              v-if="isAllowAction(item, 'delete')"
              color="error"
              end
              icon="mdi-delete"
              @click="Object.assign(selectedUser, item); deleteDialog = true"
          />
          <v-icon v-else color="error" disabled end icon="mdi-delete-off"/>
        </template>
        <template v-slot:bottom>
          <div v-if="pageCount>1" class="text-center pt-2">
            <v-pagination
                v-model="page"
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

          <v-col class="ma-0 pa-0 px-2 mb-3" cols="12" md="auto" sm="6">
            <v-checkbox
                v-model="createUser.is_admin"
                :hint="t('CREATE_IS_ADMIN_CHECKBOX_HINT')"
                :label="t('MAKE_USER_AS_ADMIN')"
                color="primary"
                density="compact"
                persistent-hint
                @update:modelValue="createUserPermissionManager"
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
          :text="t('Close')"
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
      {{ t("CHANGE_STAFF_PASSWORD") }} <span class="text-capitalize">({{ selectedUser.username }})</span>
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
          :text="t('Close')"
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
      {{ t("Delete Staff ") }} <span class="text-capitalize">({{ selectedUser.username }})</span>
    </template>
    <template #dialogText>
      {{ t('Do you want to delete the staff?') }}
    </template>
    <template #dialogAction>
      <v-btn
          :text="t('Close')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="deleteDialog=false;Object.assign(selectedUser, {})"
      />
      <v-btn
          :loading="btnLoading"
          :text="t('Delete')"
          class="me-1"
          color="error"
          variant="outlined"
          @click="remove"
      />
    </template>
  </Modal>
  <Modal :persistent="false" :show="permDialog" width="auto" @close="permDialog=!permDialog">
    <template #dialogTitle>
      {{ t("Staff Permissions Routes Access") }}
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
          :text="t('Close')"
          class="me-2"
          color="grey"
          variant="outlined"
          @click="permDialog=false;Object.assign(selectedUser, {})"
      />
      <v-btn
          :loading="btnLoading"
          :text="t('Update')"
          class="me-1"
          color="error"
          variant="outlined"
          @click="updatePerm"
      />
    </template>
  </Modal>
</template>
