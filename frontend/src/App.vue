<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <nav class="navbar navbar-expand-sm navbar-dark bg-dark">
      <div class="container">
        <div class="navbar-brand">MinMax Macros</div>
        <button
          class="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#MinMaxMacrosNavbar"
          aria-controls="MinMaxMacrosNavbar"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="MinMaxMacrosNavbar">
          <ul class="navbar-nav">
            <li class="nav-item">
              <RouterLink class="nav-link" to="/" active-class="active">Home</RouterLink>
            </li>
            <li class="nav-item">
              <RouterLink class="nav-link" to="/about" active-class="active">About</RouterLink>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>

  <div class="container text-center mt-5">
    <RouterView />
  </div>
</template>
