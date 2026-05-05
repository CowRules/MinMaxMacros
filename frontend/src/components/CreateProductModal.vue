<script setup lang="ts">
import {
  CButton,
  CChipInput,
  CCol,
  CForm,
  CFormCheck,
  CFormInput,
  CModal,
  CModalBody,
  CModalFooter,
  CModalHeader,
  CModalTitle,
  CRow,
} from '@coreui/vue'
import { type ProductItem } from '@/types.ts'
import { ref } from 'vue'
const props = defineProps<{
  visible: boolean
  shops: Set<string>
  categories: Set<string>
}>()
const emit = defineEmits<{
  close: []
  addProduct: [newProduct: ProductItem]
}>()
const productToAdd = ref<ProductItem>({
  id: crypto.randomUUID(),
  name: '',
  calories: 0,
  protein: 0,
  fiber: 0,
  price: 0,
  weight: 0,
  shops: [],
  categories: [],
})

function handleSubmit() {
  emit('addProduct', { ...productToAdd.value })
  handleClose()
}
function handleClose() {
  productToAdd.value.id = crypto.randomUUID()
  productToAdd.value.name = ''
  productToAdd.value.calories = 0
  productToAdd.value.protein = 0
  productToAdd.value.fiber = 0
  productToAdd.value.price = 0
  productToAdd.value.weight = 0
  productToAdd.value.shops = []
  productToAdd.value.categories = []

  emit('close')
}

function fixNumberInput(event: { target: { value: string } }) {
  const value = event.target.value
  if (value === '') return
  const fixed = Number(value)
  event.target.value = fixed.toString()
}
</script>
<template>
  <CModal backdrop="static" :visible="visible" @close="handleClose">
    <CModalHeader>
      <CModalTitle>Add new product</CModalTitle>
    </CModalHeader>
    <CForm class="g-3" @submit.prevent="handleSubmit">
      <CModalBody>
        <h6>Product information</h6>
        <CRow>
          <CCol>
            <CFormInput
              class="mb-2"
              type="text"
              floatingLabel="Product Name"
              placeholder="Product Name"
              v-model="productToAdd.name"
              required
            />
          </CCol>
        </CRow>
        <CRow>
          <CCol>
            <CFormInput
              @change="fixNumberInput"
              class="mb-2"
              type="number"
              floatingLabel="Weight"
              placeholder="Weight"
              step="0.01"
              min="0"
              v-model="productToAdd.weight"
              required
            />
          </CCol>
          <CCol>
            <CFormInput
              @change="fixNumberInput"
              class="mb-2"
              type="number"
              floatingLabel="Price (€)"
              placeholder="Price (€)"
              step="0.01"
              min="0"
              v-model="productToAdd.price"
              required
            />
          </CCol>
        </CRow>
        <h6>Macros per 100g</h6>
        <CRow>
          <CCol>
            <CFormInput
              @change="fixNumberInput"
              class="mb-2"
              type="number"
              floatingLabel="Calories (kcal)"
              placeholder="Calories (kcal)"
              step="0.01"
              min="0"
              v-model="productToAdd.calories"
              required
            />
          </CCol>
          <CCol>
            <CFormInput
              @change="fixNumberInput"
              class="mb-2"
              type="number"
              floatingLabel="Protein (g)"
              placeholder="Protein (g)"
              step="0.01"
              min="0"
              v-model="productToAdd.protein"
              required
            />
          </CCol>
          <CCol>
            <CFormInput
              @change="fixNumberInput"
              class="mb-2"
              type="number"
              floatingLabel="Fiber (g)"
              placeholder="Fiber (g)"
              step="0.01"
              min="0"
              v-model="productToAdd.fiber"
              required
            />
          </CCol>
        </CRow>
        <h6>Shops</h6>
        <CFormCheck
          inline
          v-for="shop in shops"
          :id="shop"
          :value="shop"
          :label="shop"
          v-model="productToAdd.shops"
        />
        <CChipInput
          v-model="productToAdd.shops"
          separator=""
          placeholder="Enter new shop"
          class="overflow-x-auto"
        />
        <h6>Categories</h6>
        <CFormCheck
          inline
          v-for="category in categories"
          :id="category"
          :value="category"
          :label="category"
          v-model="productToAdd.categories"
        />
        <CChipInput
          v-model="productToAdd.categories"
          separator=""
          placeholder="Enter new category"
          class="overflow-x-auto"
        />
      </CModalBody>
      <CModalFooter>
        <CButton color="secondary" @click="handleClose"> Close </CButton>
        <CButton type="submit" color="primary">Add</CButton>
      </CModalFooter>
    </CForm>
  </CModal>
</template>
