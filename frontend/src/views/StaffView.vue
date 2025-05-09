<script setup lang="ts">
import {useLocale} from "vuetify/framework";
import {onBeforeMount, onBeforeUnmount, onMounted, ref} from "vue";
import type { DataTableHeader } from 'vuetify'
import {type ModelsUser, UserApi} from "@/api";
import {useIsMobileStore} from "@/stores/isMobile.js.ts";

const {t} = useLocale()
const loading = ref(false)
const totalItems = ref(0)


const headers :DataTableHeader[] = [
  { title: t("ID"), key: 'uid', align: 'start' },
  { title: t("Username"), key: 'username', align: 'start' },
  { title: t("Last Login"), key: 'last_login', align: 'start' },
  { title: "Action", key: 'actions', align: 'center' },
]

const page =ref(1)
const pageSize = ref(10)
const totalRecord =ref(0)
const pageCount = ref(0)

const staffData = ref<ModelsUser[]>([]);

onBeforeMount(()=>{
  fetchStaffs()
})


const add =()=> {
  console.log(staffData)
}

const edit =(id: string)=> {
  console.log(id)
}

const remove =(id: string)=> {
  console.log(id)
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

const useIsMobile = useIsMobileStore()


</script>

<template>
  <v-card min-width="800" max-height="900" v-if="!useIsMobile.isMobile">
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

      <template #item.actions="{ item }">
        <div class="d-flex ga-2 justify-end">
          <v-icon color="medium-emphasis" icon="mdi-pencil" size="small" @click="edit(item?.uid ?? '')"></v-icon>

          <v-icon color="medium-emphasis" icon="mdi-delete" size="small" @click="remove(item?.uid ?? '')"></v-icon>
        </div>
      </template>

    </v-data-table>
    </v-card-text>
  </v-card>

  <v-card v-else>
    <v-card-text >
      {{t("Not Active in mobile view")}}
    </v-card-text>
  </v-card>
</template>
