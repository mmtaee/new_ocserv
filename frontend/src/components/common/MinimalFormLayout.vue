<script lang="ts" setup>
import {ref} from "vue";

const valid = ref(true)
</script>

<template>
  <v-card class="pa-2" max-width="650">
    <v-card-text>
      <v-form v-model="valid">
        <v-row align="center" justify="center">
          <v-col class="mt-5 ma-0 pa-0" cols="12" md="6" sm="12">
            <div style="text-align: center;">
            <span class="text-h5 text-capitalize">
              <slot name="formTitle"/>
            </span>
            </div>
          </v-col>
          <v-divider class="mt-5"/>
          <v-col class="ma-0 pa-0" cols="12" md="12" sm="12">
            <slot name="formHint"/>
          </v-col>
          <slot name="formFields"/>
        </v-row>
      </v-form>
    </v-card-text>
    <v-card-actions class="justify-end">
      <slot name="formAction"/>
    </v-card-actions>
  </v-card>
</template>
