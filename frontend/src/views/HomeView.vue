<script setup lang="ts">
import { ref } from 'vue'
import { type ProductItem, ShopName } from '@/types.ts'
import ProductList from '@/components/ProductList.vue'

const listings = ref<ProductItem[]>([
  {
    id: '123',
    name: 'Lašiša',
    price: 5,
    categories: ['Žuvis'],
    calories: 150,
    protein: 18,
    fiber: 3,
    shops: [ShopName.Maxima, ShopName.Rimi],
  },
  {
    id: '555',
    name: 'Varškė',
    price: 1,
    categories: ['Varškė'],
    calories: 80,
    protein: 16.8,
    fiber: 1,
    shops: [ShopName.Iki],
  },
])
</script>

<template>
  <main>
    <div class="container px-lg-5">
      <h1 class="my-5">Welcome to MinMax Macros</h1>
      <ProductList v-if="listings.length>0" :products="listings" />
      <h2 v-else>No products found</h2>
    </div>
  </main>
</template>
