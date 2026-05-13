<script setup lang="ts">
import { RouterView } from 'vue-router'
import { CContainer } from '@coreui/vue'
import Navbar from '@/components/Navbar.vue'
</script>

<template>
  <div class="min-vh-100 d-flex flex-column">
    <header>
      <Navbar />
    </header>
    <CContainer class="flex-grow-1 d-flex flex-column text-center pt-5">
      <RouterView />
    </CContainer>
  </div>
</template>
