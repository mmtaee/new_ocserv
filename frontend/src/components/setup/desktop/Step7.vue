<script lang="ts" setup>
import type {PanelSetupData} from "@/api";
import {useLocale} from "vuetify/framework";

const {t} = useLocale()

defineProps<{
  data: PanelSetupData;
}>();

function keyTransform(key: string): string {
  return key
      .replace(/-/g, ' ')
      .replace(/([a-z])([A-Z])/g, '$1 $2')
      .replace(/\s+/g, ' ')
      .trim()
}

</script>

<template>
  <v-row align="center" justify="center">

    <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
      <div style="text-align: center;">
      <span class="text-h5">
        {{ t('REVIEW_SETUP_RESULT') }}
      </span>
      </div>
    </v-col>

    <v-divider class="mt-5"/>
    <v-col class="mt-5 ma-0" cols="12" md="11" sm="12">
      <v-card
          :subtitle="t('ADMIN_&_SITE_CONFIGURATION')"
          class="py-5 px-1"
          elevation="40"
          variant="outlined"
      >
        <v-divider class="mb-5"/>
        <v-row align="start" justify="start">
          <v-col class="px-5 mx-5" cols="12" md="12" sm="12">
            <span class="text-primary">
              {{ t('ADMIN_USERNAME') }}:
            </span>
            <span class="text-grey mx-3">
              {{ data.admin.username }}
            </span>
          </v-col>
          <v-col class="px-5 mx-5" cols="12" md="12" sm="12">

            <span class="text-primary">
              {{ t('ADMIN_PASSWORD') }}:
            </span>
            <v-tooltip :text="data.admin.password">
              <template v-slot:activator="{ props }">
                <span class="text-grey mx-3" v-bind="props">
                  {{ '*'.repeat(data.admin.password.length * 2) }}
                </span>
              </template>
            </v-tooltip>

          </v-col>

          <v-col
              class="px-5 mx-5"
              cols="12"
              md="12"
              sm="12"
          >
            <p>
              <span class="text-primary">
                {{ t('GOOGLE_CAPTCHA_SITE_KEY') }}:
              </span>
              <span class="text-justify text-grey mx-3">
                {{ data.config?.google_captcha_site_key }}
              </span>
            </p>
          </v-col>
          <v-col class="px-5 mx-5" cols="12" md="12" sm="12">
            <span class="text-primary">
              {{ t('GOOGLE_CAPTCHA_SECRET_KEY') }}:
            </span>
            <span class="text-justify text-grey mx-3">
            {{ data.config?.google_captcha_secret_key }}
          </span>
          </v-col>
        </v-row>
      </v-card>
    </v-col>

    <v-col
        v-if="Boolean(data.default_ocserv_group)"
        class="ma-0"
        cols="12"
        md="11"
        sm="12"
    >
      <v-card
          :subtitle="t('OCSERV_DEFAULT_GROUP_CONFIGURATION')"
          class="py-5 px-2"
          elevation="40"
          variant="outlined"
      >
        <v-divider class="mb-5"/>
        <v-row align="start" justify="start">
          <v-col
              v-for="(val, key) in data.default_ocserv_group"
              :key="key"
              class="mx-5"
              cols="12"
              md="5"
              sm="12"
          >
            <p>
              <span class="text-capitalize text-primary">
              {{ keyTransform(key.toString()) }}:
                </span>
              <span v-if="!Array.isArray(val)" class="text-justify">
                {{ val || t("UNDEFINED") }}
              </span>
              <span v-else>
                {{ val.join(",") }}
              </span>
            </p>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>

</template>

