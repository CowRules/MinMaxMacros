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
          <CDropdown>
            <CDropdownToggle size="sm"></CDropdownToggle>
            <CDropdownMenu>
              <CDropdownItem
                class="no-hover"
                v-for="(shop, index) in product.shops"
                :key="product.id + '_' + shop"
              >
                <p class="text-center pb-0 mb-0">{{ shop }}</p>
                <CDropdownDivider v-if="index !== product.shops.length - 1" />
              </CDropdownItem>
            </CDropdownMenu>
          </CDropdown>
        </CTableDataCell>
        <CTableDataCell>
          <CDropdown>
            <CDropdownToggle size="sm"></CDropdownToggle>
            <CDropdownMenu>
              <CDropdownItem
                class="no-hover"
                v-for="(category, index) in product.categories"
                :key="product.id + '_' + index"
              >
                <p class="text-center pb-0 mb-0">{{ category }}</p>
                <CDropdownDivider v-if="index !== product.categories.length - 1" />
              </CDropdownItem>
            </CDropdownMenu>
          </CDropdown>
        </CTableDataCell>
      </CTableRow>
    </CTableBody>
  </CTable>
</template>

<style scoped>
.no-hover:hover,
.no-hover:focus,
.no-hover:active{
  background-color: transparent;
  color: inherit;
}
</style>
