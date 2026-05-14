<script setup lang="ts">
import { CButton, CModal, CModalFooter, CModalHeader, CModalTitle } from '@coreui/vue'

const props = defineProps<{
  visible: boolean
  headerText: string
  id: string
}>()
const emit = defineEmits<{
  close: []
  delete: [id: string]
}>()

function handleDelete() {
  emit('delete', props.id)
  emit('close')
}
</script>

<template>
  <CModal backdrop="static" :visible="visible" @close="emit('close')">
    <CModalHeader>
      <CModalTitle>{{ props.headerText }}</CModalTitle>
    </CModalHeader>
    <CModalFooter>
      <CButton color="secondary" @click="emit('close')"> Close </CButton>
      <CButton color="danger" @click="handleDelete">Delete</CButton>
    </CModalFooter>
  </CModal>
</template>

<style scoped></style>
