<template>
  <div class="scroll-smooth">
    <!-- Navbar -->
    <header class=" w-full bg-white shadow z-1">
      <nav class="max-w-7xl mx-auto flex justify-between items-center py-4 px-6">
        <h1 class="text-2xl font-bold text-purple-700">iBowl</h1>
        <ul class="flex space-x-6 text-gray-700">
          <li><a href="#mision" class="hover:text-purple-700">Misión</a></li>
          <li><a href="#features" class="hover:text-purple-700">Características</a></li>
          <li><a href="#como-funciona" class="hover:text-purple-700">¿Cómo funciona?</a></li>
          <li><a href="#testimonios" class="hover:text-purple-700">Testimonios</a></li>
        </ul>
      </nav>
    </header>

    <!-- Hero -->
    <section class="h-screen flex items-center justify-center relative bg-cover bg-center" :style="`background-image: url(${imgFondo })`">
      <div class="absolute inset-0"></div>
      <div class="relative text-center text-white max-w-2xl px-6">
        <h1 class="text-5xl font-extrabold mb-4">IBOWL</h1>
        <p class="text-xl italic">La mejor experiencia en alimentación de tu mascota</p>
        <p class="mt-4 text-sm">Tecnología inteligente al servicio del bienestar animal</p>
        <div class="mt-6">
          <a href="#mision" class="bg-purple-600 px-6 py-2 rounded-full text-white hover:bg-purple-700 transition">Descubre más</a>
        </div>
      </div>
    </section>

    <!-- Misión y Visión -->
    <section id="mision" class="py-20 px-8 bg-gray-50 text-center">
      <h2 class="text-4xl font-bold mb-10 text-gray-800">Nuestra Misión y Visión</h2>
      <div class="grid md:grid-cols-2 gap-12 max-w-6xl mx-auto">
        <div class="bg-white p-8 rounded-xl shadow-md hover:shadow-xl transition">
          <h3 class="text-2xl font-semibold text-purple-700 mb-4">Misión</h3>
          <p>Ofrecer soluciones tecnológicas que mejoren la calidad de vida de las mascotas y la tranquilidad de sus dueños.</p>
        </div>
        <div class="bg-white p-8 rounded-xl shadow-md hover:shadow-xl transition">
          <h3 class="text-2xl font-semibold text-purple-700 mb-4">Visión</h3>
          <p>Convertirnos en la referencia global en tecnología para el cuidado inteligente de animales domésticos.</p>
        </div>
      </div>
    </section>

    <!-- Características / Cards -->
    <section id="features" class="py-20 px-8 bg-white">
      <h2 class="text-4xl font-bold text-center mb-12">¿Por qué elegir iBowl?</h2>
      <div class="grid md:grid-cols-3 gap-10 max-w-6xl mx-auto">
        <div class="bg-purple-100 p-6 rounded-lg text-center shadow hover:scale-105 transition-transform">
          <div class="text-4xl mb-4">🤖</div>
          <h4 class="text-xl font-semibold mb-2">Automatiza la alimentacion</h4>
          <p>Personaliza la dieta de tu mascota según hábitos y necesidades.</p>
        </div>
        <div class="bg-purple-100 p-6 rounded-lg text-center shadow hover:scale-105 transition-transform">
          <div class="text-4xl mb-4">📱</div>
          <h4 class="text-xl font-semibold mb-2">Control desde App</h4>
          <p>Gestiona los horarios, porciones y recibe alertas desde tu celular.</p>
        </div>
        <div class="bg-purple-100 p-6 rounded-lg text-center shadow hover:scale-105 transition-transform">
          <div class="text-4xl mb-4">🐾</div>
          <h4 class="text-xl font-semibold mb-2">Diseño ergonómico</h4>
          <p>Fabricado pensando en el confort y la seguridad de tu mascota.</p>
        </div>
      </div>
    </section>

    <!-- ¿Cómo funciona? -->
    <section id="como-funciona" class="py-20 px-8 bg-gray-100">
      <h2 class="text-4xl font-bold text-center mb-12">¿Cómo funciona?</h2>
      <div class="max-w-5xl mx-auto grid md:grid-cols-3 gap-8 text-center">
        <div class="bg-white p-6 rounded-lg shadow-md hover:shadow-lg transition">
          <h3 class="text-2xl font-semibold text-purple-700 mb-2">1. Configura</h3>
          <p>Registra a tu mascota y ajusta sus parámetros desde la app.</p>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-md hover:shadow-lg transition">
          <h3 class="text-2xl font-semibold text-purple-700 mb-2">2. Alimenta</h3>
          <p>El comedero dosifica automáticamente según el plan alimenticio.</p>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-md hover:shadow-lg transition">
          <h3 class="text-2xl font-semibold text-purple-700 mb-2">3. Monitorea</h3>
          <p>Recibe alertas y reportes del consumo y salud alimenticia.</p>
        </div>
      </div>
    </section>

    <!-- Testimonios -->
    <section id="testimonios" class="py-20 px-8 bg-white">
      <h2 class="text-4xl font-bold text-center mb-12">Lo que dicen nuestros usuarios</h2>
      <div class="max-w-4xl mx-auto grid md:grid-cols-2 gap-10">
        <div class="bg-gray-50 p-6 rounded-xl shadow hover:shadow-lg transition">
          <p class="italic">"¡Mi perro Max ya no se salta ninguna comida! Además recibo notificaciones si no ha comido."</p>
          <h4 class="mt-4 font-bold text-purple-600">Ana R.</h4>
        </div>
        <div class="bg-gray-50 p-6 rounded-xl shadow hover:shadow-lg transition">
          <p class="italic">"Increíble solución. La app es intuitiva y el comedero se ve muy moderno."</p>
          <h4 class="mt-4 font-bold text-purple-600">Carlos M.</h4>
        </div>
      </div>
    </section>

    
    
  </div>
</template>
<script setup>
import  imgFondo  from '@/assets/images/img_04.jpg'
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Syncopate:wght@400;700&display=swap');

/* Animaciones personalizadas */
@keyframes fade-in {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

@keyframes fade-up {
  0% {
    transform: translateY(20px);
    opacity: 0;
  }

  100% {
    transform: translateY(0px);
    opacity: 1;
  }
}

@keyframes bounce-slow {

  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-10px);
  }
}

.animate-fade-in {
  animation: fade-in 1s ease-out forwards;
}

.animate-fade-up {
  animation: fade-up 1s ease-out forwards;
}

.animate-bounce-slow {
  animation: bounce-slow 2s infinite;
}
</style>
