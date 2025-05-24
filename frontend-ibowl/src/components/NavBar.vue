<template>
    <nav class="bg-indigo-600 text-white p-4 shadow-md">
      <div class="container mx-auto flex justify-between items-center">
        <!-- Logo/Link a inicio -->
        <router-link 
          to="/" 
          class="text-2xl font-bold hover:text-indigo-200 transition-colors"
        >
          iBowl
        </router-link>
        
        <!-- Menú principal -->
        <div class="hidden md:flex space-x-6">
          <router-link 
            v-for="link in links" 
            :key="link.path"
            :to="link.path"
            class="hover:underline hover:text-indigo-200 transition-colors"
            active-class="font-semibold underline"
          >
            {{ link.name }}
          </router-link>
        </div>
  
        <!-- Menú móvil (opcional) -->
        <button 
          class="md:hidden focus:outline-none"
          @click="isMobileMenuOpen = !isMobileMenuOpen"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
          </svg>
        </button>
      </div>
  
      <!-- Menú móvil desplegable -->
      <div 
        v-if="isMobileMenuOpen"
        class="md:hidden bg-indigo-700 px-4 py-2 space-y-2"
      >
        <router-link 
          v-for="link in links" 
          :key="`mobile-${link.path}`"
          :to="link.path"
          class="block py-2 hover:bg-indigo-600 px-2 rounded transition-colors"
          active-class="font-semibold bg-indigo-800"
          @click="isMobileMenuOpen = false"
        >
          {{ link.name }}
        </router-link>
      </div>
    </nav>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const isMobileMenuOpen = ref(false);
  
  // Links de navegación (puedes pasarlos como props si lo prefieres)
  const links = [
    { path: '/', name: 'Inicio' },
    { path: '/mascotas', name: 'Mascotas' },
    { path: '/dashboard', name: 'Dashboard' },
  ];
  </script>
  
  <style scoped>
  
  </style>