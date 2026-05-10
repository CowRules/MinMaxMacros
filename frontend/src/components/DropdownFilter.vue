<script setup lang="ts">
import {
  CButton,
  CDropdown,
  CDropdownHeader,
  CDropdownItem,
  CDropdownMenu,
  CDropdownToggle,
  CFormCheck,
} from '@coreui/vue'

const model = defineModel<string[]>({ required: true })

const props = defineProps<{
  label: string
  filterOptions: string[]
}>()
</script>

<template>
  <CDropdown auto-close="outside" class="me-2">
    <CDropdownToggle :color="model.length > 0 ? 'warning' : 'secondary'">
      <span v-if="model.length > 0" class="pi pi-filter" style="font-size: 0.8rem"></span>
      <span v-else class="pi pi-filter-slash" style="font-size: 0.8rem"></span>
      {{ label }}
    </CDropdownToggle>
    <CDropdownMenu>
      <CDropdownHeader class="text-center">
        <CButton color="link" class="p-0" size="sm" @click="model = []">Clear filters</CButton>
      </CDropdownHeader>
      <CDropdownItem class="py-0 px-1" v-for="option in filterOptions" :key="option + '_filter'">
        <CFormCheck
          hit-area="full"
          :id="option"
          :value="option"
          :label="option"
          v-model="model"
        />
      </CDropdownItem>
    </CDropdownMenu>
  </CDropdown>
</template>

<style scoped></style>
