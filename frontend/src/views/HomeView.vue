<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { type ProductItem } from '@/types.ts'
import ProductList from '@/components/ProductList.vue'
import {
  CButton,
  CContainer,
  CDropdown,
  CDropdownHeader,
  CDropdownItem,
  CDropdownMenu,
  CDropdownToggle,
  CFormCheck,
  CFormInput,
  CInputGroup,
  CInputGroupText,
  CSpinner,
} from '@coreui/vue'
import CreateProductModal from '@/components/CreateProductModal.vue'
import { api } from '@/api/api.ts'

const loaded = ref(false)

const fullListing = ref<ProductItem[]>([])

async function fetchData() {
  fullListing.value = await api.get('/api/products').then((res) => {
    return res.data
  })
}

onMounted(async () => {
  await fetchData()
  loaded.value = true
})

const filteredListing = computed(() => {
  return fullListing.value.filter((item: ProductItem) => {
    return (
      item.name.toLowerCase().includes(searchInput.value.toLowerCase()) &&
      (activeFilters.value.length === 0 ||
        item.categories.some((category) => activeFilters.value.includes(category)))
    )
  })
})
const visibleModal = ref<boolean>(false)

const searchInput = ref<string>('')

const existingShops = computed(() => {
  return new Set(
    fullListing.value
      .flatMap((listing) => {
        return listing.shops
      })
      .sort((a, b) => (a.toLowerCase() > b.toLowerCase() ? 1 : -1)),
  )
})

const existingCategories = computed(() => {
  return new Set(
    fullListing.value
      .flatMap((listing) => {
        return listing.categories
      })
      .sort((a, b) => (a.toLowerCase() > b.toLowerCase() ? 1 : -1)),
  )
})

const activeFilters = ref<string[]>([])
const sortedBy = ref<string>('')
const sortDirection = ref<string>('')

async function handleSort(criteria: string, direction: string) {
  sortedBy.value = criteria
  sortDirection.value = direction
  switch (criteria) {
    case 'calories':
      fullListing.value.sort((a, b) => (a.calories > b.calories === (direction === 'asc') ? 1 : -1))
      break
    case 'protein':
      fullListing.value.sort((a, b) => (a.protein > b.protein === (direction === 'asc') ? 1 : -1))
      break
    case 'ratio':
      fullListing.value.sort((a, b) =>
        a.protein / a.calories > b.protein / b.calories === (direction === 'asc') ? 1 : -1,
      )
      break
    case 'fiber':
      fullListing.value.sort((a, b) => (a.fiber > b.fiber === (direction === 'asc') ? 1 : -1))
      break
    case 'price':
      fullListing.value.sort((a, b) => (a.price > b.price === (direction === 'asc') ? 1 : -1))
      break
    case 'weight':
      fullListing.value.sort((a, b) => (a.weight > b.weight === (direction === 'asc') ? 1 : -1))
      break
    default:
      fullListing.value.sort((a, b) => (a.name.toLowerCase() > b.name.toLowerCase() ? 1 : -1))
      break
  }
}

function handleClose(): void {
  visibleModal.value = false
}

function handleNewProduct(newProduct: ProductItem): void {
  fullListing.value.push(newProduct)
}
</script>

<template>
  <CContainer v-if="loaded" class="px-lg-5">
    <CreateProductModal
      :visible="visibleModal"
      :shops="existingShops"
      :categories="existingCategories"
      @close="handleClose"
      @add-product="handleNewProduct"
    />
    <h1>Welcome to MinMax Macros</h1>

    <div
      class="d-flex mt-5 mb-2"
      :class="{
        'justify-content-between': filteredListing.length > 0,
        'justify-content-center': filteredListing.length === 0,
      }"
    >
      <CInputGroup class="w-25 me-2">
        <CInputGroupText><span class="pi pi-search"></span></CInputGroupText>
        <CFormInput type="text" placeholder="Search by name" v-model="searchInput" />
      </CInputGroup>
      <CDropdown auto-close="outside" :class="{ 'me-auto': filteredListing.length > 0 }">
        <CDropdownToggle :color="activeFilters.length > 0 ? 'warning' : 'secondary'">
          <span
            v-if="activeFilters.length > 0"
            class="pi pi-filter"
            style="font-size: 0.8rem"
          ></span>
          <span v-else class="pi pi-filter-slash" style="font-size: 0.8rem"></span>
          Filter
        </CDropdownToggle>
        <CDropdownMenu>
          <CDropdownHeader class="text-center">
            <CButton color="link" class="p-0" size="sm" @click="() => (activeFilters = [])"
              >Clear all filters</CButton
            >
          </CDropdownHeader>
          <CDropdownItem
            class="py-0 px-1"
            v-for="category in existingCategories"
            :key="category + '_filter'"
          >
            <CFormCheck
              hit-area="full"
              :value="category"
              :label="category"
              v-model="activeFilters"
            />
          </CDropdownItem>
        </CDropdownMenu>
      </CDropdown>
      <CButton
        color="primary"
        class="mx-3"
        @click="
          () => {
            visibleModal = true
          }
        "
      >
        Add product
      </CButton>
    </div>
    <h2 v-if="filteredListing.length === 0" class="my-5">No products found</h2>
    <ProductList
      v-if="filteredListing.length > 0"
      :products="filteredListing"
      :sorted-by
      :sort-direction
      @handle-sort="handleSort"
    />
  </CContainer>
  <CContainer class="flex-grow-1 align-content-center" v-else>
    <CSpinner />
  </CContainer>
</template>
