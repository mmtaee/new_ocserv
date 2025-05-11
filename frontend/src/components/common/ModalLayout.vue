<script lang="ts" setup>
import {ref, watch} from "vue";

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  persistent: {
    type: Boolean,
    default: true
  },
  width : {
    type: String,
    default: "auto"
  }
})

const emit = defineEmits(["close"])

const showDialog = ref(false)

watch(() => props.show, (newVal) => {
  showDialog.value = newVal
})

const close = ()=>{
  emit("close", true)
}

</script>

<template>
  <v-dialog
      v-model="showDialog"
      :persistent="persistent"
      transition="dialog-top-transition"
      :width="width"
      @update:modelValue="close"
  >
    <v-card>
      <v-card-title class="bg-primary">
        <slot name="dialogTitle"/>
      </v-card-title>

      <v-card-text class="text-subtitle-1 text-capitalize">
        <slot name="dialogText"/>
      </v-card-text>

      <v-divider class="mb-3" />

      <v-card-actions class="justify-end me-2 mb-2">
        <slot name="dialogAction"/>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

