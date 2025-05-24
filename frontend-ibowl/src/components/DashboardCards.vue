<template>
  <section class="dashboard">
    <div class="card card-color2">
      <h3><i class="fas fa-user-tag"></i> Nombre de Usuario</h3>
      <p>{{ userName }}</p>
    </div>

    <div class="card card-color3">
      <h3><i class="fas fa-microchip"></i> Dispositivo</h3>
      <p>{{ deviceInfo }}</p>
    </div>

    <div class="card card-color4">
      <h3><i class="fas fa-dog"></i> Mascota Actual</h3>
      <p>{{ petName }}</p>
      <p>🍽️ Alimentaciones: <span>{{ feedCount }}</span></p>
    </div>

    <div class="card card-color5">
      <h3><i class="fas fa-cogs"></i> Estado del Servo</h3>
      <p>{{ servoStatus }}</p>
    </div>
  </section>
</template>

<script>
import axios from 'axios'; 

export default {
  name: 'DashboardCards',
  props: {
    userName: {
      type: String,
      default: 'Cargando...'
    },
    deviceInfo: {
      type: String,
      default: 'Cargando...'
    },
    petName: {
      type: String,
      default: 'Cargando...'
    },
    feedCount: {
      type: Number,
      default: 0
    },
    servoStatus: {
      type: String,
      default: 'Desconectado'
    }
  }
  ,

  // mounted() { // o mounted()
  //   axios.get('http://localhost:8080/mascotas')
  //     .then(response => {
  //       const mascotas = response.data;
  //       console.log(response.data);
  //     })
  //     .catch(error => {
  //       console.error('Error fetching data:', error);
  //     });
  // }

}

</script>

<style scoped>
.dashboard {
  display: flex;
  flex-wrap: wrap;
  gap: 20px; /* Espacio entre tarjetas */
  padding: 20px;
  justify-content: center; /* Centra las tarjetas si no ocupan todo el ancho */
}

.card {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 20px;
  flex: 1 1 250px; /* Permite que las tarjetas crezcan y se encojan, con un ancho base */
  min-width: 220px; /* Ancho mínimo para tarjetas pequeñas */
  max-width: 300px; /* Ancho máximo */
  color: #333;
  transition: transform 0.2s ease-in-out;
}

.card:hover {
  transform: translateY(-5px);
}

.card h3 {
  margin-top: 0;
  margin-bottom: 15px;
  font-size: 1.2em;
  display: flex;
  align-items: center;
}

.card h3 i {
  margin-right: 10px;
  font-size: 1.3em; /* Iconos un poco más grandes */
}

.card p {
  margin-bottom: 5px;
  font-size: 1em;
}

.card p:last-child {
  margin-bottom: 0;
}

/* Colores específicos (puedes ajustarlos) */
.card-color2 {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}
.card-color2 h3 i { color: #e0e0e0; }

.card-color3 {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
  color: #2c3e50;
}
.card-color3 h3 i { color: #1abc9c; }


.card-color4 {
  background: linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%);
  color: #333;
}
.card-color4 h3 i { color: #e91e63; }


.card-color5 {
  background: linear-gradient(135deg, #ffc371 0%, #ff5f6d 100%);
  color: white;
}
.card-color5 h3 i { color: #f0f0f0; }

/* Asegúrate de tener Font Awesome si usas las clases fas */
/* Puedes añadir el CDN en tu index.html:
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
*/
</style>