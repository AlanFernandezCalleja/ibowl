<template>
  <div class="tabla-mascotas-container">
    <table class="tabla-mascotas">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Raza</th>
          <th>Edad</th>
          <th>Peso (kg)</th>
          <th>Frecuencia</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="mascota in mascotas" :key="mascota.id">
          <td>{{ mascota.id }}</td>
          <td>{{ mascota.nombre }}</td>
          <td>{{ mascota.raza }}</td>
          <td>{{ mascota.edad }}</td>
          <td>{{ mascota.peso }}</td>
          <td>{{ mascota.frecuencia }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';


export default {
  name: 'TablaMascotas',
  data() {
    return {
      mascotas:[]
    }
  },
  mounted(){
    this.obtenerMascotas();
    console.log(this.mascotas)
  },
  methods: {
    async obtenerMascotas(){
      try {
        const response = await axios.get('http://localhost:8080/mascotas');
        this.mascotas = response.data;
      } catch (error) {
        console.error('Error al obtener mascotas:', error);
      }
    }
  }
}
</script>

<style scoped>
.tabla-mascotas-container {
  width: 100%;
  overflow-x: auto;
  margin: 20px 0;
}

.tabla-mascotas {
  width: 100%;
  border-collapse: collapse;
  font-family: 'Arial', sans-serif;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.tabla-mascotas th,
.tabla-mascotas td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #e0e0e0;
}

.tabla-mascotas th {
  background-color: #f5f5f5;
  font-weight: 600;
  color: #333;
  text-transform: uppercase;
  font-size: 0.85em;
  letter-spacing: 0.5px;
}

.tabla-mascotas tbody tr:hover {
  background-color: #f9f9f9;
}

.tabla-mascotas tbody tr:nth-child(even) {
  background-color: #fafafa;
}

.tabla-mascotas tbody tr:nth-child(even):hover {
  background-color: #f5f5f5;
}
</style>