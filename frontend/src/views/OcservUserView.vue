<script lang="ts" setup>
import {defineAsyncComponent, onBeforeMount, reactive, ref} from "vue";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";
import {
  type OcservUserCreateOcservUserData,
  type OcservUserOcservUsersResponse,
  OcservUsersApi,
  type RequestMeta
} from "@/api";
import {getAuthorization} from "@/utils/request.ts";

const DesktopView = defineAsyncComponent(() => import("@/components/oc_user/desktop/index.vue"))
const MobileView = defineAsyncComponent(() => import("@/components/oc_user/mobile/index.vue"))

const useIsMobile = useIsMobileStore()
const loading = ref(true)
const btnLoading = ref(false)
const meta = reactive<RequestMeta>({
  page: 1,
  page_size: 10,
  total_records: 0
})
const pageCount = ref(0)
const data = reactive<OcservUserOcservUsersResponse>({
  meta: {
    page: 1,
    page_size: 10,
    total_records: 0
  },
  result: undefined
})

const groups = reactive<string[]>([])


onBeforeMount(() => {
  fetchOcUser()
})

const fetchOcUser = () => {
  loading.value = true
  const api = new OcservUsersApi()
  api.ocUsersGet(
      {
        ...getAuthorization(),
        page: meta.page,
        pageSize: meta.page_size,
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

const addStaff = (data: OcservUserCreateOcservUserData) => {
  btnLoading.value = true
  const api = new OcservUsersApi()
  api.ocUsersPost({
    ...getAuthorization(),
    request: data,
  }).then((_) => {
    Object.assign(meta, {
      page: 1,
      page_size: 10,
      total_records: 0
    })
    fetchOcUser()
  }).finally(() => {
    btnLoading.value = false
  })

}

</script>

<template>
  <component
      :is="useIsMobile.isMobile ? MobileView: DesktopView"
      :btnLoading="btnLoading"
      :data="data"
      :groups="groups"
      :loading="loading"
      :pageCount="pageCount"
      @addStaff="addStaff"
      @fetchOcGroups="fetchOcGroups"
      @fetchOcUser="fetchOcUser"
  />
</template>
