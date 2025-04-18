<script lang="ts" setup>
import logoUrl from "@/assets/ocserv.png"
import {useLocale} from "vuetify/framework";

const {t} = useLocale()
</script>

<template>
  <v-row align="center" justify="center">
    <v-col cols="12" md="6" sm="12">
      <v-img :src="logoUrl" alt="My Image"/>
    </v-col>
    <v-col cols="12" md="6" sm="12">
  <span class="text-h6 text-grey">
    Ocserv User Management
  </span>
    </v-col>
    <v-divider/>
    <v-col class="mt-5 text-justify" cols="12" md="12" sm="12">
      <div class="mx-5 text-h6">
        {{ t('WELCOME_TO_THE_OCSERV_USER_MANAGEMENT_PANEL') }}.
      </div>
      <br/>
      <div class="mx-5 text-subtitle-1 text-justify">
        {{ t('LET’S_BEGIN_BY_COMPLETING_A_FEW_QUICK_CONFIGURATION_STEPS_TO_GET_YOUR_PANEL_UP_AND_RUNNING') }} .
      </div>
    </v-col>
  </v-row>
</template>
