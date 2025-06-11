<script lang="ts" setup>
import {formatDay} from "@/utils/dates.ts";
import {
  type OcOcservUser,
  type OcservUserCreateOcservUserData,
  OcservUserCreateOcservUserDataTrafficTypeEnum
} from "@/api";
import {defineAsyncComponent, reactive, ref, watch} from "vue";
import {requiredRule} from "@/utils/rules.ts";
import {useI18n} from "vue-i18n";

const Modal = defineAsyncComponent(() => import("@/components/common/ModalLayout.vue"))

const props = withDefaults(
    defineProps<{
      btnLoading: boolean
      groups: string[] | undefined
      modelValue: boolean
      data?: OcOcservUser,
      update?: boolean
    }>(),
    {
      btnLoading: false,
      update: false,
    }
)


const emit = defineEmits(["update:modelValue", "doAction"])

const {t} = useI18n()

const menu = ref(false)
const createUserValid = ref(true)
const today = new Date()
const futureDate = new Date(today)
futureDate.setDate(futureDate.getDate() + 30)
const formattedFutureDate = futureDate.toISOString().split('T')[0]
const createUser = reactive<OcservUserCreateOcservUserData>({
  group: "defaults",
  username: "",
  password: "",
  traffic_size: 10,
  traffic_type: OcservUserCreateOcservUserDataTrafficTypeEnum.TOTALLY_TRANSMIT,
  expire_at: formattedFutureDate,
})

const rules = {
  required: (v: string) => requiredRule(v, t),
}

const addUser = () => {
  emit("doAction", props.update ? "update" : "create", {
    data: createUser,
    uid: props?.data?.uid || null
  })
}

const resetCreateUser = () => {
  Object.assign(createUser, {
    group: "defaults",
    username: "",
    password: "",
    traffic_size: 10,
    traffic_type: OcservUserCreateOcservUserDataTrafficTypeEnum.TOTALLY_TRANSMIT,
    expire_at: formattedFutureDate
  })
}

const close = () => {
  emit('update:modelValue', false)
  resetCreateUser()
}


if (props.update) {
  watch(
      () => props.data,
      (data) => {
        if (data) {
          Object.assign(createUser, {
            group: data.group,
            password: data.password,
            username: "",
            traffic_size: data.traffic_size,
            traffic_type: data.traffic_type,
            expire_at: formatDay(data.expire_at)
          })
        }
      },
      {immediate: true, deep: true}
  )
}


</script>

<template>
  <Modal :persistent="false" :show="modelValue" width="750" @close="close">
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


          <v-col v-if="!update" cols="12" md="4" sm="6">
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
                    clearable
                    density="comfortable"
                    readonly
                    v-bind="props"
                    variant="underlined"
                />
              </template>
              <v-date-picker
                  v-model="createUser.expire_at"
                  :header="t('EXPIRE_AT')"
                  elevation="24"
                  title=""
                  @update:model-value="createUser.expire_at=formatDay(createUser.expire_at); menu = false"
              />
            </v-menu>
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

          <v-col cols="12" md="4" sm="6">
            <v-number-input
                v-model="createUser.traffic_size"
                :hint="t('MAX_TRAFFIC_SIZE_HINT')"
                :label="t('MAX_TRAFFIC_SIZE')"
                :min="0"
                controlVariant="hidden"
                density="comfortable"
                suffix="GB"
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
          @click="close"
      />
      <v-btn
          :disabled="!createUserValid"
          :loading="btnLoading"
          :text="update ? t('UPDATE_OCSERV_USER'): t('ADD_OCSERV_USER')"
          class="me-1"
          color="success"
          variant="outlined"
          @click="addUser"
      />
    </template>

  </Modal>
</template>

