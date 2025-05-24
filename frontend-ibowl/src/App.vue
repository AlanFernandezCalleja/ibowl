<template>
  <div id="app">
    <!-- Barra de navegación -->
    <!-- <NavBar /> -->
     <!-- Sidebar -->
     <!-- Menú lateral -->
    <Sidebar />
    <!-- Contenedor principal donde se cargarán las vistas -->
    <main class="container mx-auto p-4 min-h-[calc(100vh-8vh)]">
      <router-view></router-view>
    </main>

    <!-- Pie de página -->
    <footer class="bg-gray-800 text-white p-4 text-center h-[8vh]">
      <p>© 2023 iBowl - Sistema de alimentación para mascotas</p>
    </footer>
  </div>
</template>

<script setup>
  // import NavBar from './components/NavBar.vue';
  import Sidebar from './components/Sidebar.vue';

</script>

<style>

</style>