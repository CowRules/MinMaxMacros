<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { type ProductItem } from '@/types.ts'
import ProductList from '@/components/ProductList.vue'
import {
  CButton,
  CContainer,
  CFormInput,
  CInputGroup,
  CInputGroupText,
  CSpinner,
} from '@coreui/vue'
import CreateProductModal from '@/components/CreateProductModal.vue'
import api from '@/api/api.ts'
import { mergeUniqueSortedStringArrays } from '@/utils/arrayUtils.ts'
import DropdownFilter from '@/components/DropdownFilter.vue'
import { userStorage } from '@/composables/useUserStorage.ts'

const loaded = ref(false)

const fullListing = ref<ProductItem[]>([])

const existingShops = ref<string[]>([])

const existingCategories = ref<string[]>([])

async function fetchData() {
  fullListing.value = await api.get('/api/products').then((res) => {
    if (res.data != null) {
      return res.data
    }
    return []
  })
  existingShops.value = await api.get('/api/products/shops').then((res) => {
    if (res.data != null) {
      return res.data
    }
    return []
  })
  existingCategories.value = await api.get('/api/products/categories').then((res) => {
    if (res.data != null) {
      return res.data
    }
    return []
  })
}

onMounted(async () => {
  await fetchData()
  loaded.value = true
})

const filteredListing = computed(() => {
  handleSort(sortedBy.value, sortDirection.value)
  return fullListing.value.filter((item: ProductItem) => {
    return (
      item.name.toLowerCase().includes(searchInput.value.toLowerCase()) &&
      (activeCategoryFilters.value.length === 0 ||
        item.categories.some((category) => activeCategoryFilters.value.includes(category))) &&
      (activeShopFilters.value.length === 0 ||
        item.shops.some((shop) => activeShopFilters.value.includes(shop)))
    )
  })
})
const visibleModal = ref<boolean>(false)

const searchInput = ref<string>('')

const activeCategoryFilters = ref<string[]>([])
const activeShopFilters = ref<string[]>([])
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

async function handleNewProduct(newProduct: ProductItem) {
  await api.post('/api/products', newProduct).then((res) => {
    fullListing.value.push(res.data)
    existingCategories.value = mergeUniqueSortedStringArrays(
      existingCategories.value,
      newProduct.categories,
    )
    existingShops.value = mergeUniqueSortedStringArrays(existingShops.value, newProduct.shops)
  })
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
      <div :class="{ 'me-auto': filteredListing.length > 0 }">
        <DropdownFilter
          label="Category filter"
          :filter-options="existingCategories"
          v-model="activeCategoryFilters"
        />
        <DropdownFilter
          label="Shop filter"
          :filter-options="existingShops"
          v-model="activeShopFilters"
        />
      </div>
      <CButton
        v-if="userStorage.id"
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
