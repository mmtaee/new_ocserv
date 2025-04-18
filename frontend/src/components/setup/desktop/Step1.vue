<script setup lang="ts">
import logoUrl from "@/assets/ocserv.png"

</script>

<template>
  <v-row align="center" justify="center">
    <v-col cols="12" md="6" sm="12">
      <v-img :src="logoUrl" alt="My Image" />
    </v-col>
    <v-col cols="12" md="6" sm="12">
  <span class="text-h6 text-grey">
    Ocserv User Management
  </span>
    </v-col >
    <v-divider />
    <v-col cols="12" md="12" sm="12" class="mt-5 text-justify" >
      <div class="mx-5 text-h6">
      Welcome to the Ocserv User Management Panel.
      </div>
      <br/>
      <div class="mx-5 text-subtitle-1 text-justify">
      Let’s begin by completing a few quick configuration steps to get your panel up and running.
      </div>


    </v-col>
  </v-row>


</template>

<style scoped>

</style>