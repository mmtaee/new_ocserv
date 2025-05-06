<script lang="ts" setup>
</script>

<template>
  <v-card class="pa-2" max-width="650" min-width="450">
    <v-card-text>
      <v-row align="center" justify="center">

        <v-col class="mt-5 ma-0 pa-0" cols="12" md="6" sm="12">
          <div class="text-h5 text-capitalize" style="text-align: center;">
            <slot name="formTitle"/>
          </div>
        </v-col>

        <v-divider class="mt-5"/>
        <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
          <slot name="formHint"/>
        </v-col>

        <v-col class="mt-5 ma-0 pa-0" cols="12" md="12" sm="12">
          <slot class="my-5" name="formFields"/>
        </v-col>

      </v-row>
    </v-card-text>

    <v-divider class="mb-2"/>

    <v-card-actions class="justify-end me-5 ms-5">
      <slot name="formAction"/>
    </v-card-actions>

  </v-card>
</template>
