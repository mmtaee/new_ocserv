<script lang="ts" setup>

import {ref, watch} from "vue";

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
})

const showDialog = ref(false)

watch(() => props.show, (newVal) => {
  showDialog.value = newVal
})

</script>

<template>
  <v-dialog
      v-model="showDialog"
      persistent
      transition="dialog-top-transition"
      width="auto"
  >
    <v-card>
      <v-card-title class="bg-primary">
        <slot name="dialogTitle"/>
      </v-card-title>

      <v-card-text class="text-subtitle-1 text-capitalize">
        <slot name="dialogText"/>
      </v-card-text>

      <v-card-actions class="justify-end">
        <slot name="dialogAction"/>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

