<script lang="ts" setup>
import type {SetupRequestSetup} from "@/api";
import type {PropType} from "vue";

// const {t} = useLocale()

defineProps({
  data: {
    type: Object as PropType<SetupRequestSetup>,
    required: true,
  },
})

</script>

<template>
  <v-card class="ma-3 bg-background text-white" variant="tonal">


    <div>
      <div v-for="(value, key) in data" :key="key">
        <strong>{{ key }}:</strong> {{ value }}
      </div>
    </div>
  </v-card>
</template>

<style scoped>

</style>