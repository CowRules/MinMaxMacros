<script setup lang="ts">
import { CAlert, CButton, CCol, CForm, CFormInput, CRow } from '@coreui/vue'
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import { alphaNum, email, maxLength, minLength, required } from '@vuelidate/validators'
import useVuelidate from '@vuelidate/core'
import api from '@/api/api.ts'
import router from '@/router'
import axios from 'axios'
import PasswordInput from '@/components/PasswordInput.vue'

const user = ref({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
})

const registerError = ref<string>('')

const passwordsMatch = (value: string) => value === user.value.password

const rules = {
  username: [required, minLength(5), maxLength(20), alphaNum],
  email: [required, email],
  password: [required, minLength(8), maxLength(100)],
  confirmPassword: [passwordsMatch],
}
const v$ = useVuelidate(rules, user)
async function handleRegister() {
  try {
    const isValid = await v$.value.$validate()
    if (!isValid) {
      return
    }
    const res = await api.post('/auth/register', user.value)
    if (res.status === 201) {
      router.push('/login')
      return
    }
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 409)
      registerError.value = 'User with that email already exists'
    else registerError.value = 'Error occurred. Try again later.'
  }
}
</script>

<template>
  <div class="bg-white rounded-5 p-4 mx-auto register">
    <h2>Register</h2>
    <CAlert v-if="registerError" color="danger" dismissible @close="registerError = ''"
      >Login failed: {{ registerError }}</CAlert
    >
    <CForm class="py-2" @submit.prevent="handleRegister()" novalidate>
      <CFormInput
        class="mb-2"
        type="text"
        id="username"
        label="Username"
        v-model="v$.username.$model"
        :invalid="v$.username.$error"
        :feedback-invalid="
          v$.username.$errors.length > 0 ? v$.username.$errors[0]!.$message.toString() : ''
        "
      />
      <CFormInput
        class="mb-2"
        type="email"
        id="email"
        label="Email"
        v-model="v$.email.$model"
        :invalid="v$.email.$error"
        :feedback-invalid="
          v$.email.$errors.length > 0 ? v$.email.$errors[0]!.$message.toString() : ''
        "
      />
      <PasswordInput
        class="mb-2"
        label="Password"
        v-model="v$.password.$model"
        :invalid="v$.password.$error"
        :feedback-invalid="
          v$.password.$errors.length > 0 ? v$.password.$errors[0]!.$message.toString() : ''
        "
      />
      <PasswordInput
        class="mb-2"
        label="Confirm Password"
        v-model="v$.confirmPassword.$model"
        :invalid="v$.confirmPassword.$error"
        :feedback-invalid="v$.confirmPassword.$error ? 'Passwords do not match' : ''"
      />
      <CRow :xs="{ cols: 'auto' }" class="mt-4 justify-content-center">
        <CCol>
          <CButton color="primary" type="submit">Register</CButton>
        </CCol>
        <CCol class="align-content-center">
          <RouterLink class="btn-secondary" to="/login">Login</RouterLink>
        </CCol>
      </CRow>
    </CForm>
  </div>
</template>

<style scoped>
.register {
  min-width: 60% !important;
  border: 5px solid lightskyblue;
}
</style>
