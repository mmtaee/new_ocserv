<script lang="ts" setup>
import {
  type OcservUserCreateOcservUserData,
  OcservUserCreateOcservUserDataTrafficTypeEnum,
  type OcservUserOcservUsersResponse
} from "@/api";
import type {DataTableHeader} from "vuetify/framework";
import {useI18n} from "vue-i18n";
import {defineAsyncComponent, reactive, ref} from "vue";
import {requiredRule} from "@/utils/rules.ts";
import {formatDay} from "@/utils/dates.ts"

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

defineProps<{
  data: OcservUserOcservUsersResponse
  loading?: boolean
  pageCount?: number
  groups?: string[],
  btnLoading: boolean
}>()

const emit = defineEmits(["fetchOcUser", "fetchOcGroups", "addStaff"])

const {t} = useI18n()
const addDialog = ref(false)
const menu = ref(false)
const today = new Date()
const futureDate = new Date(today)
futureDate.setDate(futureDate.getDate() + 30)
const formattedFutureDate = futureDate.toISOString().split('T')[0]
const headers: DataTableHeader[] = [
  {title: t("GROUP"), key: 'group', align: 'center'},
  {title: t("USERNAME"), key: 'username', align: 'center'},
  {title: t("PASSWORD"), key: 'password', align: 'center'},
  {title: t("DATES"), key: 'created_at', align: 'center'},
  {title: t("TRAFFICS"), key: 'traffic_type', align: 'center'},
  {title: t("TRANSMISSION"), key: 'rx', align: 'center'},
  {title: t("DESCRIPTION"), key: 'description', align: 'center'},
  {title: t("STATUS"), key: 'is_online', align: 'center'},
  {title: t("ACTIONS"), key: 'actions', align: 'center'},
]
const createUserValid = ref(true)
const createUser = reactive<OcservUserCreateOcservUserData>({
  group: "defaults",
  username: "",
  password: "",
  traffic_size: 0,
  traffic_type: OcservUserCreateOcservUserDataTrafficTypeEnum.FREE,
  expire_at: formattedFutureDate
})
const rules = {
  required: (v: string) => requiredRule(v, t),
}

const fetchOcUser = () => {
  emit("fetchOcUser")
}

const fetchOcGroups = () => {
  emit("fetchOcGroups")
}

const addStaff = () => {
  emit("addStaff", createUser)
}

const resetCreateUser = () => {
  Object.assign(createUser, {
    group: "defaults",
    username: "",
    password: "",
    traffic_size: 0,
    traffic_type: OcservUserCreateOcservUserDataTrafficTypeEnum.FREE,
    expire_at: formattedFutureDate
  })
}

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
              <v-icon class="mb-1" color="medium-emphasis" icon="mdi-account-tie-hat" size="large" start></v-icon>
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

        <!--        <template #item.last_login="{ item }">-->
        <!--          {{ formatDate(item.last_login) }}-->
        <!--        </template>-->

        <!--        <template #item.actions="{ item }">-->
        <!--          <v-icon-->
        <!--              color="warning"-->
        <!--              end-->
        <!--              icon="mdi-door-closed-lock"-->
        <!--              @click="permission(item)"-->
        <!--          />-->
        <!--          <v-icon-->
        <!--              color="primary"-->
        <!--              end-->
        <!--              icon="mdi-key"-->
        <!--              @click="Object.assign(selectedUser, item); changePasswordDialog = true"-->
        <!--          />-->
        <!--          <v-icon-->
        <!--              color="error"-->
        <!--              end-->
        <!--              icon="mdi-delete"-->
        <!--              @click="Object.assign(selectedUser, item); deleteDialog = true"-->
        <!--          />-->
        <!--        </template>-->

        <template v-slot:bottom>
          <div v-if="data.meta.total_records>1" class="text-center pt-2">
            <v-pagination
                v-model="data.meta.page"
                :length="pageCount"
                @update:modelValue="fetchOcUser"
            />
          </div>
        </template>

      </v-data-table>
    </v-card-text>

  </v-card>

  <Modal :persistent="false" :show="addDialog" width="750" @close="addDialog=!addDialog">
    <template #dialogTitle>
      {{ t("CREATE_OCSERV_USER") }}
    </template>

    <template #dialogText>
      <v-form v-model="createUserValid">
        <v-row align="center" class="pa-1" justify="start">

          <v-col cols="12" md="4" sm="6">
            <v-select
                v-model="createUser.group"
                :items="groups"
                :no-data-text="t('NO_GROUP_FOUND')"
                :rules="[rules.required]"
                density="comfortable"
                variant="underlined"
            />
          </v-col>


          <v-col cols="12" md="4" sm="6">
            <v-text-field
                v-model="createUser.username"
                :label="t('USERNAME')"
                :rules="[rules.required]"
                clearable
                density="comfortable"
                variant="underlined"
            />
          </v-col>

          <v-col cols="12" md="4" sm="6">
            <v-text-field
                v-model="createUser.password"
                :label="t('PASSWORD')"
                :rules="[rules.required]"
                density="comfortable"
                variant="underlined"
            />
          </v-col>

          <v-col cols="12" md="4" sm="6">
            <v-menu
                v-model="menu"
                :close-on-content-click="false"
                offset-y
                transition="scale-transition"
            >
              <template #activator="{ props }">
                <v-text-field
                    v-model="createUser.expire_at"
                    :label="t('EXPIRE_AT')"
                    :rules="[rules.required]"
                    density="comfortable"
                    readonly
                    v-bind="props"
                    variant="underlined"
                />
              </template>
              <v-date-picker
                  v-model="createUser.expire_at"
                  elevation="24"
                  header=""
                  title=""
                  @update:model-value="createUser.expire_at=formatDay(createUser.expire_at); menu = false"
              />
            </v-menu>
          </v-col>

          <v-col cols="12" md="4" sm="6">
            <v-number-input
                v-model="createUser.traffic_size"
                :label="t('TRAFFIC_SIZE')"
                :min="0"
                :rules="[rules.required]"
                controlVariant="hidden"
                density="comfortable"
                variant="underlined"
            />
          </v-col>

          <v-col cols="12" md="4" sm="6">
            <v-select
                v-model="createUser.traffic_type"
                :items="Object.values(OcservUserCreateOcservUserDataTrafficTypeEnum)"
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

</template>
