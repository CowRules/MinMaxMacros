<script setup lang="ts">
import { CAlert, CButton, CCol, CForm, CFormInput, CFormLabel, CRow } from '@coreui/vue'
import { RouterLink } from 'vue-router'
import { ref } from 'vue'
import api from '@/api/api.ts'
import axios from 'axios'
import router from '@/router'
import { userStorage } from '@/composables/useUserStorage.ts'
import PasswordInput from '@/components/PasswordInput.vue'
import { email, required } from '@vuelidate/validators'
import useVuelidate from '@vuelidate/core'

const loginDetails = ref({
  email: '',
  password: '',
})
const errorMessage = ref<string>('')

const rules = {
  email: [required, email],
  password: [required],
}
const v$ = useVuelidate(rules, loginDetails)

async function handleLogin() {
  try {
    const isValid = await v$.value.$validate()
    if (!isValid) {
      return
    }
    const res = await api.post('/auth/login', {
      email: loginDetails.value.email,
      password: loginDetails.value.password,
    })
    if (res.status === 200) {
      userStorage.value.id = res.data.id
      userStorage.value.username = res.data.username
      userStorage.value.role = res.data.role
      router.push('/')
    }
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      errorMessage.value = 'Incorrect email or password'
      return
    }
    errorMessage.value = 'Error occurred. Try again later.'
  }
}
</script>

<template>
  <div class="bg-white rounded-5 p-4 mx-auto login">
    <h2>Login</h2>
    <CAlert v-if="errorMessage" color="danger" dismissible @close="errorMessage = ''"
      >Login failed: {{ errorMessage }}</CAlert
    >
    <CForm class="py-2" @submit.prevent="handleLogin" novalidate>
      <CFormInput
        class="mb-2"
        type="email"
        id="email"
        label="Email"
        v-model="v$.email.$model"
        :invalid="v$.email.$error"
        :feedback-invalid="
          v$.email.$errors.length > 0 ? v$.email.$errors[0]!.$message.toString() : undefined
        "
      />
      <PasswordInput
        class="mb-2"
        label="Password"
        v-model="v$.password.$model"
        :invalid="v$.password.$error"
        :feedback-invalid="
          v$.password.$errors.length > 0 ? v$.password.$errors[0]!.$message.toString() : undefined
        "
      />
      <CRow :xs="{ cols: 'auto' }" class="mt-4 justify-content-center">
        <CCol>
          <CButton color="primary" type="submit">Login</CButton>
        </CCol>
        <CCol class="align-content-center">
          <RouterLink class="btn-link" to="/register">Register</RouterLink>
        </CCol>
      </CRow>
    </CForm>
  </div>
</template>

<style scoped>
.login {
  min-width: 60% !important;
  border: 5px solid lightskyblue;
}
</style>
