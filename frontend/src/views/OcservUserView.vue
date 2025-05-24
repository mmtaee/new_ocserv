<script lang="ts" setup>
import {defineAsyncComponent, onBeforeMount, reactive, ref} from "vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import {
  type OcOcservUser,
  type OcservUserCreateOcservUserData,
  type OcservUserOcservUsersResponse,
  OcservUsersApi
} from "@/api";
import {getAuthorization} from "@/utils/request.ts";

const DesktopView = defineAsyncComponent(() => import("@/components/oc_user/desktop/index.vue"))
const MobileView = defineAsyncComponent(() => import("@/components/oc_user/mobile/index.vue"))

const useIsMobile = useIsMobileStore()
const loading = ref(true)
const btnLoading = ref(false)

const pageCount = ref(0)
const data = reactive<OcservUserOcservUsersResponse>({
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
  fetchOcUser(data.meta.page, data.meta.page_size, false)
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
    Object.assign(data, res.data)
    pageCount.value = Math.ceil(data.meta.total_records / data.meta.page_size);
  }).finally(() => {
    loading.value = false
  })
}

const fetchOcGroups = () => {
  console.log("get group list name")
//   TODO: get ocserv groups list name
  groups.unshift("defaults")
}

const addUser = (updateData: OcservUserCreateOcservUserData) => {
  btnLoading.value = true
  const api = new OcservUsersApi()
  api.ocUsersPost({
    ...getAuthorization(),
    request: updateData,
  }).then((res) => {
    childRef.value?.closeDialogs()
    data.result?.unshift(res.data)
  }).finally(() => {
    btnLoading.value = false
  })
}

const updateUser = (user: OcOcservUser) => {
  const result = data.result
  if (!result) return

  const index = result.findIndex((u) => u.uid === user.uid)
  if (index >= 0) {
    data.result?.splice(index, 1, user)
  }
}

</script>

<template>
  <component
      :is="useIsMobile.isMobile ? MobileView: DesktopView"
      ref="childRef"
      :btnLoading="btnLoading"
      :data="data"
      :groups="groups"
      :loading="loading"
      :pageCount="pageCount"
      @addUser="addUser"
      @fetchOcGroups="fetchOcGroups"
      @fetchOcUser="fetchOcUser"
      @updateUser="updateUser"
  />
</template>
