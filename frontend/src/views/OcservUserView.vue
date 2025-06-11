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
import type {OcservDataInterface} from "@/utils/interfaces.ts";


const DesktopView = defineAsyncComponent(() => import("@/components/oc_user/desktop/index.vue"))
const MobileView = defineAsyncComponent(() => import("@/components/oc_user/mobile/index.vue"))

const useIsMobile = useIsMobileStore()
const loading = ref(true)
const btnLoading = ref(false)

const pageCount = ref(0)
const ocservUsers = reactive<OcservUserOcservUsersResponse>({
  meta: {
    page: 1,
    page_size: 25,
    total_records: 0
  },
  result: undefined
})

const currentStep = reactive(
    {
      page: 1,
      page_size: 25,
      desc: false
    }
)


const groups = reactive<string[]>([])
const childRef = ref()

const api = new OcservUsersApi()

onBeforeMount(() => {
  fetchOcUser(currentStep.page, currentStep.page_size, currentStep.desc)
  fetchOcGroups()
})

const fetchOcUser = (page: number, pageSize: number, desc: boolean) => {
  currentStep.page = page
  currentStep.page_size = pageSize
  currentStep.desc = desc

  loading.value = true
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
  btnLoading.value = true
  api.ocUsersUidPatch({
    uid: uid,
    ...getAuthorization(),
    request: user,
  }).then((res) => {
    console.log(res.data)
    const result = ocservUsers.result
    if (result) {
      const index = result.findIndex((u) => u.uid === uid)
      if (index >= 0) {
        result?.splice(index, 1, res.data)
      }
    }
    childRef.value?.closeDialogs()
  }).finally(() => {
    btnLoading.value = false

  })
}

const lockOrUnlock = (uid: string, data: OcservUserLockOcservUserData) => {
  btnLoading.value = true
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
  btnLoading.value = true
  api.ocUsersUidDelete({
    uid,
    ...getAuthorization()
  }).then(() => {
    const result = ocservUsers.result
    if (!result) return

    if (ocservUsers.meta.total_records > result.length) {
      fetchOcUser(currentStep.page, currentStep.page_size, false)
    } else {
      const index = result.findIndex((u) => u.uid === uid)
      if (index >= 0) {
        result.splice(index, 1)
      }
    }

    childRef.value?.closeDialogs()
  }).finally(() => {
    btnLoading.value = false
  })
}

const disconnect = (uid: string, username: string) => {
  api.ocUsersUsernameDisconnectPost({
    username,
    ...getAuthorization()
  }).then(() => {
    const result = ocservUsers.result
    if (!result) return
    const index = result.findIndex((u) => u.uid === uid)
    if (index >= 0) {
      result[index].is_online = false
    }
  })
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

    case "disconnect":
      if (data.uid) {
        disconnect(data.uid, data.data.username)
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
      @fetchOcUser="fetchOcUser"
  />
</template>
