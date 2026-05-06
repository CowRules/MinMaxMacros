<script setup lang="ts">
import type { ProductItem } from '@/types.ts'
import {
  CDropdown,
  CDropdownDivider,
  CDropdownItem,
  CDropdownMenu,
  CDropdownToggle,
  CTable,
  CTableBody,
  CTableDataCell,
  CTableHead,
  CTableHeaderCell,
  CTableRow,
} from '@coreui/vue'
import DropdownList from '@/components/DropdownList.vue'

const props = defineProps<{
  products: ProductItem[]
}>()
</script>

<template>
  <CTable striped hover bordered color="secondary" small>
    <CTableHead>
      <CTableRow color="info">
        <CTableHeaderCell scope="col">Product</CTableHeaderCell>
        <CTableHeaderCell scope="col">Calories(per 100g)</CTableHeaderCell>
        <CTableHeaderCell scope="col">Protein(per 100g)</CTableHeaderCell>
        <CTableHeaderCell scope="col">Protein/Calories</CTableHeaderCell>
        <CTableHeaderCell scope="col">Fiber(per 100g)</CTableHeaderCell>
        <CTableHeaderCell scope="col">Price</CTableHeaderCell>
        <CTableHeaderCell scope="col">Weight</CTableHeaderCell>
        <CTableHeaderCell scope="col">Shops</CTableHeaderCell>
        <CTableHeaderCell scope="col">Categories</CTableHeaderCell>
      </CTableRow>
    </CTableHead>
    <CTableBody class="table-group-divider">
      <CTableRow v-for="product in products" :key="product.id">
        <CTableHeaderCell scope="row">{{ product.name }}</CTableHeaderCell>
        <CTableDataCell>{{ product.calories }}</CTableDataCell>
        <CTableDataCell>{{ product.protein }}</CTableDataCell>
        <CTableDataCell>
          {{
            product.protein === 0 || product.calories === 0
              ? '-'
              : (product.protein / product.calories).toFixed(2)
          }}
        </CTableDataCell>
        <CTableDataCell>{{ product.fiber }}</CTableDataCell>
        <CTableDataCell>{{ product.price }}</CTableDataCell>
        <CTableDataCell>{{ product.weight }}</CTableDataCell>
        <CTableDataCell>
          <DropdownList :items="product.shops" />
        </CTableDataCell>
        <CTableDataCell>
          <DropdownList :items="product.categories" />
        </CTableDataCell>
      </CTableRow>
    </CTableBody>
  </CTable>
</template>

<style scoped></style>
