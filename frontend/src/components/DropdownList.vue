<script setup lang="ts">
import {
  CDropdown,
  CDropdownDivider,
  CDropdownItem,
  CDropdownMenu,
  CDropdownToggle,
} from '@coreui/vue'
const props = defineProps<{
  items: any[]
}>()
</script>

<template>
  <CDropdown auto-close="outside">
    <CDropdownToggle size="sm"></CDropdownToggle>
    <CDropdownMenu>
      <CDropdownItem class="no-hover" v-for="(item, index) in items" :key="index + '_' + item">
        <p class="text-center pb-0 mb-0">{{ item }}</p>
        <CDropdownDivider v-if="index !== items.length - 1" />
      </CDropdownItem>
    </CDropdownMenu>
  </CDropdown>
</template>

<style scoped>
.no-hover:hover,
.no-hover:focus,
.no-hover:active {
  background-color: transparent;
  color: inherit;
}
</style>
