<script setup lang="ts">
import {useLocale} from "vuetify/framework";
import {defineAsyncComponent, onBeforeMount, reactive, ref} from "vue";
import type { DataTableHeader } from 'vuetify'
import {type ModelsUser, UserApi} from "@/api";

const ModalLayout = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const {t} = useLocale()
const loading = ref(false)
const btnLoading = ref(false)
const totalItems = ref(0)

const headers :DataTableHeader[] = [
  { title: t("ID"), key: 'id', align: 'center' },
  { title: t("Username"), key: 'username', align: 'center' },
  { title: t("Last Login"), key: 'last_login', align: 'center' },
  { title: t("Action"), key: 'actions', align: 'center' },
]

const page =ref(1)
const pageSize = ref(10)
const totalRecord =ref(0)
const pageCount = ref(0)
const  permDialog = ref(false)
const staffData = reactive<ModelsUser[]>([]);
const selectedUser = reactive<ModelsUser>({
  id: 0,
  is_admin: false,
  last_login: "",
  permission: {},
  uid: "",
  username: ""
});


onBeforeMount(()=>{
  fetchStaffs()
})


const add =()=> {
  console.log("add")
}

const edit =(id: string)=> {
  console.log(id)
}

const remove =(id: string)=> {
  console.log(id)
}

const permission =(user: ModelsUser)=> {
  Object.assign(selectedUser, user)
  permDialog.value = true
}

const fetchStaffs = ()=>{
  loading.value = true
  const api = new UserApi()
  api.userStaffsGet().then((res)=>{
    Object.assign(staffData, res.data.result)
    page.value = res.data.meta.page
    pageSize.value = res.data.meta.page_size
    totalRecord.value = res.data.meta?.total_records || 0
    pageCount.value = Math.ceil(totalRecord.value / pageSize.value);
  }).finally(()=>{
    loading.value = false
  })
}

const formatDate = (dateString: string | undefined): string => {
  if (!dateString) {
    return ""
  }
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}`;
}


const updatePerm = ()=> {
  btnLoading.value = true
    const api = new UserApi()
    api.userStaffsPermissionsIdPut({
      id: selectedUser?.id,
      request: selectedUser?.permission,
    }).then(()=>{
      Object.assign(selectedUser, {})
      permDialog.value = false
    }).finally(()=>{
      btnLoading.value = false
    })
}

const permKeyWrapper = (key: string)=>{
  switch(key){
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


</script>

<template>
  <v-card min-width="800" max-height="900">
    <v-card-text >
      <v-data-table
          :headers="headers"
          :items="staffData"
          :items-length="totalItems"
          :loading="loading"
          mobile-breakpoint="sm"
          disable-sort
          :no-data-text="t('No Data Found')"
      >

        <template v-slot:bottom >
          <div class="text-center pt-2" v-if="pageCount>1">
            <v-pagination
                v-model="page"
                :length="pageCount"
                @update:modelValue="fetchStaffs"
            />
          </div>
        </template>

        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title>
              <v-icon color="medium-emphasis" icon="mdi-account-tie-hat" size="large" start class="mb-1"></v-icon>
              {{ t("Staffs") }}
            </v-toolbar-title>

            <v-btn
                class="me-2"
                prepend-icon="mdi-plus"
                rounded="lg"
                :text="t('Add Staff')"
                border
                @click="add"
            />
          </v-toolbar>
        </template>

        <template #item.last_login="{ item }">
          {{ formatDate(item.last_login) }}
        </template>


        <template #item.actions="{ item }">
            <v-icon end color="medium-emphasis" icon="mdi-eye" @click="permission(item)"></v-icon>
            <v-icon end color="medium-emphasis" icon="mdi-pencil"  @click="edit(item?.uid ?? '')"></v-icon>
            <v-icon end color="medium-emphasis" icon="mdi-delete"  @click="remove(item?.uid ?? '')"></v-icon>
        </template>

      </v-data-table>
    </v-card-text>
  </v-card>

  <ModalLayout :show="permDialog" :persistent="false" @close="permDialog=!permDialog" width="auto" >
      <template #dialogTitle>
        {{ t("User Permissions Routes Access") }}
      </template>

<template #dialogText>
  <v-row align="center" justify="center" class="ms-auto mt-2">
    <v-col cols="12" md="4" sm="4" class="ma-0 pa-0" v-for="(_, k, index) in selectedUser.permission" :key="`perm-${index}`">
      <v-checkbox
        v-model="selectedUser.permission![k]"
        hide-details
        color="success"
        :label="permKeyWrapper(k)"
        density="compact"
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
          @click="permDialog=false"
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


  </ModalLayout>

</template>
