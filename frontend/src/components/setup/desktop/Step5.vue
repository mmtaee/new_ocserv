<script lang="ts" setup>
import {useLocale} from "vuetify/framework";
import type {ModelsOcservUserOrGroupConfigs} from "@/api";
import {reactive} from "vue";
import {requiredRule} from "@/utils/rules.ts";

const {t} = useLocale()

const formValues: ModelsOcservUserOrGroupConfigs = reactive<ModelsOcservUserOrGroupConfigs>({})

// const rules = {
  // required: (v: string) => requiredRule(v, t),
// }

</script>

<template>
  <v-row align="center" justify="center">
    <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
      <div style="text-align: center;">
        <span class="text-h5">
          {{ t('DEFAULT_OCSERV_GROUP_CONFIGURATION') }}
        </span>

      </div>
    </v-col>

    <v-divider class="mt-5 mb-2"/>

    <div class="text-subtitle-2 mt-2">{{ t('IP_&_DNS_CONFIGURATION') }}</div>

    <blockquote>
      <p class="mx-8 my-4 text-grey-darken-1">
        <v-icon color="primary">mdi-bullhorn-outline</v-icon>
        {{ t('These settings control how clients receive IP addresses and resolve domain names while connected') }}.
      </p>
    </blockquote>


    <v-col cols="12" md="6" sm="12"></v-col>
    <v-col cols="12" md="6" sm="12"></v-col>
    <v-col cols="12" md="6" sm="12"></v-col>
    <v-col cols="12" md="6" sm="12"></v-col>
    <v-col cols="12" md="6" sm="12"></v-col>


  </v-row>


</template>

<style scoped>

</style>