<template>
    <!-- Contenedor principal del menú lateral -->
    <aside class="fixed inset-y-0 left-0 w-64 bg-indigo-700 text-white shadow-lg transform transition-all duration-300 ease-in-out"
           :class="{ '-translate-x-full': !isOpen, 'translate-x-0': isOpen }">
      
      <!-- Logo y título -->
      <div class="p-4 border-b border-indigo-600 flex items-center space-x-2">
        <router-link to="/" class="text-2xl font-bold flex items-center">
          <svg class="w-8 h-8 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
          </svg>
          iBowl
        </router-link>
      </div>
  
      <!-- Menú de navegación -->
      <nav class="p-4 space-y-1">
        <router-link 
          v-for="link in links" 
          :key="link.path"
          :to="link.path"
          class="flex items-center px-4 py-3 rounded-lg transition-colors"
          :class="{
            'bg-indigo-800 text-white font-semibold': $route.path === link.path,
            'hover:bg-indigo-600 text-indigo-100': $route.path !== link.path
          }"
        >
          <span class="ml-3">{{ link.name }}</span>
        </router-link>
      </nav>
  
      <!-- Botón para cerrar (opcional) -->
      <button 
        @click="toggleSidebar"
        class="absolute top-15 right-0 translate-x-full bg-indigo-700 text-white p-2 rounded-r-lg focus:outline-none"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
          :d="isOpen ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'"></path>
        </svg>
      </button>
    </aside>
  
    <!-- Overlay para móviles -->
    <div v-if="isOpen" 
         @click="toggleSidebar"
         class="fixed inset-0 bg-black bg-opacity-50 z-10 md:hidden">
    </div>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue';
  import { useRoute } from 'vue-router';
  
  const route = useRoute();
  const isOpen = ref(true);
  
  // Links de navegación
  const links = [
    { path: '/', name: 'Inicio', icon: 'home' },
    { path: '/dashboard', name: 'Dashboard', icon: 'chart-bar' },
    { path: '/mascotas', name: 'Mascotas', icon: 'paw' },
    { path: '/configuracion', name: 'Configuración', icon: 'cog' },
  ];
  
  // Función para alternar la visibilidad
  const toggleSidebar = () => {
    isOpen.value = !isOpen.value;
  };
  
  // Cerrar sidebar al cambiar de ruta en móviles
  watch(route, () => {
    if (window.innerWidth < 768) {
      isOpen.value = false;
    }
  });
  </script>
  
  <style scoped>
  /* Transición suave para el menú */
  aside {
    z-index: 20;
  }

  </style>