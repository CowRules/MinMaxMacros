<script setup lang="ts">
import type { ProductItem } from '@/types.ts'
import {
  CTable,
  CTableBody,
  CTableDataCell,
  CTableHead,
  CTableHeaderCell,
  CTableRow,
} from '@coreui/vue'
import DropdownList from '@/components/DropdownList.vue'
import TableHeaderCellSort from '@/components/TableHeaderCellSort.vue'

const props = defineProps<{
  products: ProductItem[]
  sortedBy: string
  sortDirection: string
}>()
const emit = defineEmits<{
  handleSort: [criteria: string, direction: string]
}>()

function handleSort(criteria: string) {
  if (props.sortedBy === criteria) {
    if (props.sortDirection === 'desc') {
      emit('handleSort', criteria, 'asc')
    } else {
      emit('handleSort', '', 'desc')
    }
  } else {
    emit('handleSort', criteria, 'desc')
  }
}
</script>

<template>
  <CTable striped hover bordered color="secondary" small>
    <CTableHead>
      <CTableRow color="info">
        <CTableHeaderCell scope="col">Product</CTableHeaderCell>
        <TableHeaderCellSort
          :sorted-by
          :sort-direction
          text="Calories(per 100g)"
          sort-key="calories"
          @click="handleSort"
        />
        <TableHeaderCellSort
          :sorted-by
          :sort-direction
          text="Protein(per 100g)"
          sort-key="protein"
          @click="handleSort"
        />
        <TableHeaderCellSort
          :sorted-by
          :sort-direction
          text="Protein/Calories"
          sort-key="ratio"
          @click="handleSort"
        />
        <TableHeaderCellSort
          :sorted-by
          :sort-direction
          text="Fiber(per 100g)"
          sort-key="fiber"
          @click="handleSort"
        />
        <TableHeaderCellSort
          :sorted-by
          :sort-direction
          text="Price"
          sort-key="price"
          @click="handleSort"
        />
        <TableHeaderCellSort
          :sorted-by
          :sort-direction
          text="Weight"
          sort-key="weight"
          @click="handleSort"
        />
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
