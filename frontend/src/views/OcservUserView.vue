<script lang="ts" setup>
import {defineAsyncComponent, onBeforeMount, reactive, ref} from "vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import {
  type OcservUserCreateOcservUserData,
  type OcservUserLockOcservUserData,
  type OcservUserOcservUsersResponse,
  OcservUsersApi,
  type OcservUserUpdateOcservUserData
} from "@/api";
import {getAuthorization} from "@/utils/request.ts";
import {useI18n} from "vue-i18n";
import type {OcservDataInterface} from "@/utils/interfaces.ts";


const DesktopView = defineAsyncComponent(() => import("@/components/oc_user/desktop/index.vue"))
const MobileView = defineAsyncComponent(() => import("@/components/oc_user/mobile/index.vue"))

const {t} = useI18n()
const useIsMobile = useIsMobileStore()
const loading = ref(true)
const btnLoading = ref(false)

const pageCount = ref(0)
const ocservUsers = reactive<OcservUserOcservUsersResponse>({
  meta: {
    page: 1,
    page_size: 5,
    total_records: 0
  },
  result: undefined
})

const groups = reactive<string[]>([])
const childRef = ref()

onBeforeMount(() => {
  fetchOcUser(ocservUsers.meta.page, ocservUsers.meta.page_size, false)
})

const fetchOcUser = (page: number, pageSize: number, desc: boolean) => {
  loading.value = true
  const api = new OcservUsersApi()
  api.ocUsersGet(
      {
        ...getAuthorization(),
        page: page,
        pageSize: pageSize,
        sort: desc ? "DESC" : "ASC"
      }
  ).then((res) => {
    Object.assign(ocservUsers, res.data)
    pageCount.value = Math.ceil(ocservUsers.meta.total_records / ocservUsers.meta.page_size);
  }).finally(() => {
    loading.value = false
  })
}

const fetchOcGroups = () => {
  console.log("get group list name")
//   TODO: get ocserv groups list name
  groups.unshift("defaults")
}

const createUser = (createData: OcservUserCreateOcservUserData) => {
  btnLoading.value = true
  const api = new OcservUsersApi()
  api.ocUsersPost({
    ...getAuthorization(),
    request: createData,
  }).then((res) => {
    ocservUsers.result?.unshift(res.data)
    childRef.value?.closeDialogs()
    // TODO: add snackbar create
  }).finally(() => {
    btnLoading.value = false
  })
}

const updateUser = (uid: string, user: OcservUserUpdateOcservUserData) => {
  // btnLoading.value = true
  //
  // const result = data.result
  // if (!result) return
  //
  // const index = result.findIndex((u) => u.uid === user.uid)
  // if (index >= 0) {
  //   data.result?.splice(index, 1, user)
  // }
}

const lockOrUnlock = (uid: string, data: OcservUserLockOcservUserData) => {
  btnLoading.value = true
  const api = new OcservUsersApi()
  api.ocUsersUidLockPost({
    uid: uid,
    ...getAuthorization(),
    request: data,
  }).then(() => {
    const result = ocservUsers.result
    if (!result) return
    const index = result.findIndex((u) => u.uid === uid)
    if (index >= 0) {
      result[index].is_locked = data.lock
    }
    childRef.value?.closeDialogs()
  }).finally(() => {
    btnLoading.value = false
  })
}


const deleteUser = (uid: string) => {
}

const doAction = (action: string, data: OcservDataInterface) => {
  switch (action) {
    case "create":
      createUser(data.data as OcservUserCreateOcservUserData)
      break
    case "update":
      if (data.uid) {
        updateUser(data.uid, data.data as OcservUserUpdateOcservUserData)
        break
      }
      throw new Error('uid not implemented.')
    case "lock":
      if (data.uid) {
        lockOrUnlock(data?.uid, data.data as OcservUserLockOcservUserData)
        break
      }
      throw new Error('uid not implemented.')
    case "delete":
      if (data.uid) {
        deleteUser(data.uid)
        break
      }
      throw new Error('uid not implemented.')
    default:
      throw new Error('action not implemented.')
  }
}

</script>

<template>
  <component
      :is="useIsMobile.isMobile ? MobileView: DesktopView"
      ref="childRef"
      :btnLoading="btnLoading"
      :data="ocservUsers"
      :groups="groups"
      :loading="loading"
      :pageCount="pageCount"
      @doAction="doAction"
      @fetchOcGroups="fetchOcGroups"
      @fetchOcUser="fetchOcUser"
  />
</template>
