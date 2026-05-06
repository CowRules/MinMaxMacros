<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import {
  CCollapse,
  CContainer,
  CNavbar,
  CNavbarBrand,
  CNavbarNav,
  CNavbarToggler,
  CNavItem,
} from '@coreui/vue'
</script>

<template>
  <header>
    <CNavbar expand="sm" color-scheme="dark" class="bg-dark">
      <CContainer>
        <CNavbarBrand>MinMax Macros</CNavbarBrand>
        <CNavbarToggler />
        <CCollapse class="navbar-collapse">
          <CNavbarNav>
            <CNavItem>
              <RouterLink class="nav-link" to="/" active-class="active">Home</RouterLink>
            </CNavItem>
            <CNavItem>
              <RouterLink class="nav-link" to="/about" active-class="active">About</RouterLink>
            </CNavItem>
          </CNavbarNav>
        </CCollapse>
      </CContainer>
    </CNavbar>
  </header>

  <CContainer class="text-center mt-5">
    <RouterView />
  </CContainer>
</template>
