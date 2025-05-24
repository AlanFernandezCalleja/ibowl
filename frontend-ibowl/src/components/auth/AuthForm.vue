
<template>
    <form @submit.prevent="submitForm" class="space-y-4">
      <!-- Email -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Correo electrónico</label>
        <input
          v-model="formData.email"
          type="email"
          required
          class="input-style"
          placeholder="tucorreo@ejemplo.com"
        />
      </div>
  
      <!-- Password -->
      <div>
        <label class="block text-sm font-medium text-gray-700">Contraseña</label>
        <input
          v-model="formData.password"
          type="password"
          required
          class="input-style"
          placeholder="••••••••"
        />
      </div>
  
      <!-- Nombre (solo para registro) -->
      <div v-if="isRegister">
        <label class="block text-sm font-medium text-gray-700">Nombre</label>
        <input
          v-model="formData.name"
          type="text"
          required
          class="input-style"
          placeholder="Tu nombre"
        />
      </div>
  
      <!-- Botón -->
      <button type="submit" class="btn-primary">
        {{ isRegister ? 'Registrarse' : 'Iniciar Sesión' }}
      </button>
    </form>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  /* eslint-disable no-unused-vars */
  const props = defineProps({
    isRegister: Boolean, // Define si es formulario de registro
  });
  
  const formData = ref({
    email: '',
    password: '',
    name: '',
  });
  
  const emit = defineEmits(['submit']);
  
  const submitForm = () => {
    emit('submit', formData.value);
  };
  </script>
  
  <style scoped>
  @import "tailwindcss";
  .input-style {
    @apply w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500;
  }
  .btn-primary {
    @apply w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-cyan-600 hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-blue-500;
  }
  </style>