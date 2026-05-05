<script setup lang="ts">
import { computed, ref } from 'vue'
import { type ProductItem } from '@/types.ts'
import ProductList from '@/components/ProductList.vue'
import { CButton } from '@coreui/vue'
import CreateProductModal from '@/components/CreateProductModal.vue'

const listings = ref<ProductItem[]>([])
const visible = ref<boolean>(false)

const existingShops = computed(() => {
  return new Set(
    listings.value.flatMap((listing) => {
      return listing.shops
    }),
  )
})

const existingCategories = computed(() => {
  return new Set(
    listings.value.flatMap((listing) => {
      return listing.categories
    }),
  )
})

function handleClose(): void {
  visible.value = false
}

function handleNewProduct(newProduct: ProductItem): void {
  listings.value.push(newProduct)
}
</script>

<template>
  <main>
    <div class="container px-lg-5">
      <CreateProductModal
        :visible="visible"
        :shops="existingShops"
        :categories="existingCategories"
        @close="handleClose"
        @add-product="handleNewProduct"
      />
      <h1>Welcome to MinMax Macros</h1>
      <h2 v-if="listings.length === 0" class="my-5">No products found</h2>
      <div
        class="d-flex mb-2"
        :class="{
          'justify-content-end': listings.length > 0,
          'justify-content-center': listings.length === 0,
        }"
      >
        <CButton
          color="primary"
          class="mx-3"
          @click="
            () => {
              visible = true
            }
          "
        >
          Add product
        </CButton>
      </div>
      <ProductList v-if="listings.length > 0" :products="listings" />
    </div>
  </main>
</template>
