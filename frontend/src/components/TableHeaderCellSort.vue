<script setup lang="ts">
import { CTableHeaderCell } from '@coreui/vue'
const props = defineProps<{
  sortedBy: string
  sortDirection: string
  text: string
  sortKey: string
}>()
const emit = defineEmits<{
  click: [dataType: string]
}>()
</script>

<template>
  <CTableHeaderCell class="sortable-header" scope="col" @click="emit('click', sortKey)">
    {{ text }}
    <span
      class="pi"
      :class="{
        'pi-sort-alt': sortedBy !== sortKey,
        'pi-sort-amount-down': sortedBy === sortKey && sortDirection === 'desc',
        'pi-sort-amount-up-alt': sortedBy === sortKey && sortDirection === 'asc',
      }"
      style="font-size: 0.8rem"
    />
  </CTableHeaderCell>
</template>

<style scoped>
.sortable-header {
  cursor: pointer;
  user-select: none;
}
</style>
