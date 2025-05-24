<script lang="ts" setup>
import {
  type OcOcservUser,
  type OcservUserCreateOcservUserData,
  OcservUserCreateOcservUserDataTrafficTypeEnum,
  type OcservUserOcservUsersResponse
} from "@/api";
import type {DataTableHeader} from "vuetify/framework";
import {useI18n} from "vue-i18n";
import {defineAsyncComponent, ref} from "vue";
import {bytesToGB} from "@/utils/convertors.ts";


const Pagination = defineAsyncComponent(() => import("@/components/common/Pagination.vue"))
const Actions = defineAsyncComponent(() => import("@/components/oc_user/desktop/Actions.vue"))
const Create = defineAsyncComponent(() => import("@/components/oc_user/desktop/Create.vue"))

defineProps<{
  data: OcservUserOcservUsersResponse
  loading?: boolean
  pageCount?: number
  groups?: string[],
  btnLoading: boolean
}>()

const emit = defineEmits(["fetchOcUser", "fetchOcGroups", "addUser", "updateUser"])

const addDialog = ref(false)
const {t} = useI18n()

const headers: DataTableHeader[] = [
  {title: t("USER"), key: 'user', align: 'start'},
  {title: t("TRAFFICS"), key: 'traffics', align: 'start'},
  {title: t("TRANSMISSION"), key: 'transmission', align: 'start'},
  {title: t("STATUS"), key: 'status', align: 'center'},
  {title: t("ACTIONS"), key: 'actions', align: 'center'},
]


const fetchOcUser = (page: number, itemPerPage: number, desc: boolean) => {
  emit("fetchOcUser", page, itemPerPage, desc)
}

const fetchOcGroups = () => {
  emit("fetchOcGroups")
}


const addUser = (user: OcservUserCreateOcservUserData) => {
  emit("addUser", user)
}

const updateUser = (user: OcOcservUser) => {
  emit("updateUser", user)
}


const closeDialogs = () => {
  addDialog.value = false
}

defineExpose([closeDialogs])

</script>

<template>
  <v-card max-height="900" min-width="800">
    <v-card-text>
      <v-data-table
          :headers="headers"
          :items="data.result || []"
          :items-length="data.meta.total_records"
          :loading="loading"
          :no-data-text="t('No Data Found')"
          disable-sort
          mobile-breakpoint="sm"
      >
        <template v-slot:top>
          <v-toolbar flat>
            <v-toolbar-title class="text-capitalize">
              <v-icon class="mb-1" color="medium-emphasis" icon="mdi-account-network" size="large" start></v-icon>
              {{ t("OCSERV_USERS") }}
            </v-toolbar-title>
            <v-btn
                :text="t('ADD_OCSERV_USER')"
                border
                class="me-2"
                prepend-icon="mdi-plus"
                rounded="lg"
                @click="addDialog=!addDialog; fetchOcGroups()"
            />
          </v-toolbar>
        </template>

        <template #item.user="{ item }">
          <div>
            <span class="text-primary text-capitalize">{{ t("GROUP") }}: </span>
            <span class="text-grey-darken-2">{{ item.group }}</span>
          </div>
          <div>
            <span class="text-primary text-capitalize">{{ t("USERNAME") }}: </span>
            <span class="text-grey-darken-2">{{ item.username }}</span>
          </div>
          <div>
            <span class="text-primary text-capitalize">{{ t("PASSWORD") }}: </span>
            <span class="text-grey-darken-2">{{ item.password }}</span>
          </div>
        </template>

        <template #item.traffics="{ item }">
          <div>
            <span class="text-primary text-capitalize">{{ t("TYPE") }}: </span>
            <span class="text-grey-darken-2">{{ item.traffic_type }}</span>
          </div>

          <div>
            <span class="text-primary text-capitalize">{{ t("SIZE_USAGE") }}: </span>
            <span class="text-grey-darken-2">
              {{
                item.traffic_type == OcservUserCreateOcservUserDataTrafficTypeEnum.FREE ? '-' : item.traffic_size + ' (GB)'
              }}
            </span>
          </div>
        </template>


        <template #item.transmission="{ item }">
          <div>
            <span class="text-primary text-capitalize">{{ t("RX") }}: </span>
            <span class="text-grey-darken-2">{{ bytesToGB(item.rx) }} (GB)</span>
          </div>

          <div>
            <span class="text-primary text-capitalize">{{ t("TX") }}: </span>
            <span class="text-grey-darken-2">{{ bytesToGB(item.tx) }} (GB)</span>
          </div>
        </template>

        <template #item.status="{ item }">
          <v-tooltip v-if="item.is_locked" :text="t('IS_LOCKED')">
            <template #activator="{ props }">
              <v-icon color="red" v-bind="props">mdi-lock-alert</v-icon>
            </template>
          </v-tooltip>

          <v-tooltip v-else :text="t('IS_ACTIVE')">
            <template #activator="{ props }">
              <v-icon color="green" v-bind="props">mdi-lock-open-alert</v-icon>
            </template>
          </v-tooltip>

          <v-tooltip v-if="item.is_online" :text="t('ONLINE')">
            <template #activator="{ props }">
              <v-icon class="ms-2" color="green" v-bind="props">mdi-lan-connect</v-icon>
            </template>
          </v-tooltip>

          <v-tooltip v-else :text="t('OFFLINE')">
            <template #activator="{ props }">
              <v-icon class="ms-2" color="red" v-bind="props">mdi-lan-disconnect</v-icon>
            </template>
          </v-tooltip>
        </template>


        <template #item.actions="{ item }">
          <Actions :item="item" @updateUser="updateUser"/>
        </template>

        <template v-slot:bottom>
          <Pagination
              :page="data.meta.page"
              :pageCount="pageCount"
              :totalRecords="data.meta.total_records"
              @reFetch="fetchOcUser"
          />
        </template>

      </v-data-table>
    </v-card-text>

  </v-card>

  <Create
      v-model="addDialog"
      :btnLoading="btnLoading"
      :data="data"
      :groups="groups"
      :loading="loading"
      :pageCount="pageCount"
      @addUser="addUser"
  />
</template>
