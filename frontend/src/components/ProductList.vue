<script setup lang="ts">
import type { ProductItem } from '@/types.ts'

const props = defineProps<{
  products: ProductItem[]
}>()
</script>

<template>
  <table class="table table-striped table-hover table-sm table-secondary table-bordered">
    <thead>
      <tr class="table-info">
        <th scope="col">Product</th>
        <th scope="col">Calories(per 100g)</th>
        <th scope="col">Protein(per 100g)</th>
        <th scope="col">Protein/Calories</th>
        <th scope="col">Fiber(per 100g)</th>
        <th scope="col">Price</th>
        <th scope="col">Weight</th>
        <th scope="col">Shops</th>
        <th scope="col">Categories</th>
      </tr>
    </thead>
    <tbody class="table-group-divider">
      <tr v-for="product in products" :key="product.id">
        <th scope="row">{{ product.name }}</th>
        <td>{{ product.calories }}</td>
        <td>{{ product.protein }}</td>
        <td>
          {{
            product.protein === 0 || product.calories === 0
              ? '-'
              : (product.protein / product.calories).toFixed(2)
          }}
        </td>
        <td>{{ product.fiber }}</td>
        <td>{{ product.price }}</td>
        <td>{{ product.weight }}</td>
        <td class="dropdown">
          <button
            class="btn btn-sm dropdown-toggle"
            type="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          ></button>
          <ul class="dropdown-menu">
            <li v-for="(shop, index) in product.shops" :key="product.id + '_' + shop">
              <p class="text-center pb-0 mb-0">{{ shop }}</p>
              <hr class="dropdown-divider" v-if="index !== product.shops.length - 1" />
            </li>
          </ul>
        </td>
        <td class="dropdown">
          <button
            class="btn btn-sm dropdown-toggle"
            type="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          ></button>
          <ul class="dropdown-menu">
            <li v-for="(category, index) in product.categories" :key="product.id + '_' + category">
              <p class="text-center pb-0 mb-0">{{ category }}</p>
              <hr class="dropdown-divider" v-if="index !== product.categories.length - 1" />
            </li>
          </ul>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<style scoped></style>
