<script setup lang="ts">
import {
  CCollapse,
  CContainer,
  CDropdown,
  CDropdownItem,
  CDropdownMenu,
  CDropdownToggle,
  CNavbar,
  CNavbarBrand,
  CNavbarNav,
  CNavbarToggler,
  CNavItem,
} from '@coreui/vue'
import { RouterLink } from 'vue-router'
import router from '@/router'
import api from '@/api/api.ts'
import { userStorage } from '@/composables/useUserStorage.ts'
async function logout() {
  userStorage.value = {
    id: null,
    username: null,
    role: null,
  }
  await api.post('/auth/logout')
  router.push('/login')
}
</script>

<template>
  <CNavbar expand="sm" color-scheme="dark" class="bg-dark">
    <CContainer>
      <CNavbarBrand>MinMax Macros</CNavbarBrand>
      <CNavbarToggler />
      <CCollapse class="navbar-collapse">
        <CNavbarNav class="w-100">
          <CNavItem>
            <RouterLink class="nav-link" to="/" active-class="active">Home</RouterLink>
          </CNavItem>
          <CNavItem>
            <RouterLink class="nav-link" to="/about" active-class="active">About</RouterLink>
          </CNavItem>
          <CNavItem>
            <CDropdown v-if="userStorage.id" variant="nav-item" :popper="true">
              <CDropdownToggle color="secondary"
                >Logged in as: {{ userStorage.username }}</CDropdownToggle
              >
              <CDropdownMenu>
                <CDropdownItem class="text-center user-select-none" @click="logout">
                  Logout
                </CDropdownItem>
              </CDropdownMenu>
            </CDropdown>
            <RouterLink v-else class="nav-link" to="/login" active-class="active">Login</RouterLink>
          </CNavItem>
        </CNavbarNav>
      </CCollapse>
    </CContainer>
  </CNavbar>
</template>

<style scoped></style>
