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
import { ref, watch } from 'vue'
import { userStorage } from '@/composables/useUserStorage.ts'
import { maxLength, minValue, required } from '@vuelidate/validators'
import useVuelidate from '@vuelidate/core'
const props = defineProps<{
  visible: boolean
  shops: string[]
  categories: string[]
  product?: ProductItem
  headerText: string
  submitText: string
}>()
const emit = defineEmits<{
  close: []
  submit: [newProduct: ProductItem]
}>()
const productData = ref<ProductItem>({
  id: '',
  name: '',
  calories: 0,
  protein: 0,
  fiber: 0,
  price: 0,
  weight: 0,
  shops: [],
  categories: [],
  userId: userStorage.value.id ?? '',
})

watch(
  () => props.product,
  (product: ProductItem | undefined) => {
    productData.value = {
      id: product?.id ?? '',
      name: product?.name ?? '',
      calories: product?.calories ?? 0,
      protein: product?.protein ?? 0,
      fiber: product?.fiber ?? 0,
      price: product?.price ?? 0,
      weight: product?.weight ?? 0,
      shops: product?.shops ? [...product?.shops] : [],
      categories: product?.categories ? [...product?.categories] : [],
      userId: product?.userId ?? '',
    }
  },
)

const rules = {
  name: [required, maxLength(50)],
  calories: [required, minValue(0)],
  protein: [required, minValue(0)],
  fiber: [required, minValue(0)],
  price: [required, minValue(0)],
  weight: [required, minValue(0)],
}

const v$ = useVuelidate(rules, productData)

async function handleSubmit() {
  const isValid = await v$.value.$validate()
  if (!isValid) {
    return
  }
  emit('submit', { ...productData.value })
  handleClose()
}

function handleClose() {
  productData.value.name = ''
  productData.value.calories = 0
  productData.value.protein = 0
  productData.value.fiber = 0
  productData.value.price = 0
  productData.value.weight = 0
  productData.value.shops = []
  productData.value.categories = []
  productData.value.userId = ''

  emit('close')
}
</script>
<template>
  <CModal backdrop="static" :visible="visible" @close="handleClose">
    <CModalHeader>
      <CModalTitle>{{ props.headerText }}</CModalTitle>
    </CModalHeader>
    <CForm class="g-3" @submit.prevent="handleSubmit" novalidate>
      <CModalBody>
        <h6>Product information</h6>
        <CRow>
          <CCol>
            <CFormInput
              class="mb-2"
              type="text"
              floatingLabel="Product Name"
              placeholder="Product Name"
              v-model="productData.name"
              :invalid="v$.name.$error"
              :feedback-invalid="
                v$.name.$errors.length > 0 ? v$.name.$errors[0]!.$message.toString() : undefined
              "
            />
          </CCol>
        </CRow>
        <CRow>
          <CCol>
            <CFormInput
              class="mb-2"
              type="number"
              floatingLabel="Weight"
              placeholder="Weight"
              step="0.01"
              min="0"
              v-model.number="productData.weight"
              :invalid="v$.weight.$error"
              :feedback-invalid="
                v$.weight.$errors.length > 0 ? v$.weight.$errors[0]!.$message.toString() : undefined
              "
            />
          </CCol>
          <CCol>
            <CFormInput
              class="mb-2"
              type="number"
              floatingLabel="Price (€)"
              placeholder="Price (€)"
              step="0.01"
              min="0"
              v-model.number="productData.price"
              :invalid="v$.price.$error"
              :feedback-invalid="
                v$.price.$errors.length > 0 ? v$.price.$errors[0]!.$message.toString() : undefined
              "
            />
          </CCol>
        </CRow>
        <h6>Macros per 100g</h6>
        <CRow>
          <CCol>
            <CFormInput
              class="mb-2"
              type="number"
              floatingLabel="Calories (kcal)"
              placeholder="Calories (kcal)"
              step="0.01"
              min="0"
              v-model.number="productData.calories"
              :invalid="v$.calories.$error"
              :feedback-invalid="
                v$.calories.$errors.length > 0
                  ? v$.calories.$errors[0]!.$message.toString()
                  : undefined
              "
            />
          </CCol>
          <CCol>
            <CFormInput
              class="mb-2"
              type="number"
              floatingLabel="Protein (g)"
              placeholder="Protein (g)"
              step="0.01"
              min="0"
              v-model.number="productData.protein"
              :invalid="v$.protein.$error"
              :feedback-invalid="
                v$.protein.$errors.length > 0
                  ? v$.protein.$errors[0]!.$message.toString()
                  : undefined
              "
            />
          </CCol>
          <CCol>
            <CFormInput
              class="mb-2"
              type="number"
              floatingLabel="Fiber (g)"
              placeholder="Fiber (g)"
              step="0.01"
              min="0"
              v-model.number="productData.fiber"
              :invalid="v$.fiber.$error"
              :feedback-invalid="
                v$.fiber.$errors.length > 0 ? v$.fiber.$errors[0]!.$message.toString() : undefined
              "
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
          v-model="productData.shops"
        />
        <CChipInput
          v-model="productData.shops"
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
          v-model="productData.categories"
        />
        <CChipInput
          v-model="productData.categories"
          separator=""
          placeholder="Enter new category"
          class="overflow-x-auto"
        />
      </CModalBody>
      <CModalFooter>
        <CButton color="secondary" @click="handleClose"> Close </CButton>
        <CButton type="submit" color="primary">{{ props.submitText }}</CButton>
      </CModalFooter>
    </CForm>
  </CModal>
</template>
