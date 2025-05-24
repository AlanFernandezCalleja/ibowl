<template>
  <section class="perfil-container">
    <div class="perfil-card">
      <div class="perfil-header">
        <h2><i class="fas fa-user-circle"></i> Configuración de Perfil</h2>
      </div>

      <div class="perfil-body">
        <!-- Sección de foto de perfil -->
        <div class="foto-perfil-container">
          <div class="foto-perfil" @click="triggerFileInput">
            <img v-if="usuario.foto" :src="usuario.foto" alt="Foto de perfil" class="foto-preview">
            <i v-else class="fas fa-user-circle default-avatar"></i>
            <input 
              type="file" 
              ref="fileInput" 
              @change="handleFileUpload" 
              accept="image/*" 
              style="display: none;"
            >
          </div>
          <button @click="triggerFileInput" class="btn-cambiar-foto">
            <i class="fas fa-camera"></i> Cambiar foto
          </button>
        </div>

        <!-- Formulario de perfil -->
        <form @submit.prevent="guardarPerfil" class="form-perfil">
          <div class="form-group">
            <label for="nombreUsuario"><i class="fas fa-at"></i> Nombre de usuario</label>
            <input 
              type="text" 
              id="nombreUsuario" 
              v-model="usuario.nombreUsuario" 
              placeholder="Ej: usuario123"
              required
            >
          </div>

          <div class="form-group">
            <label for="nombreCompleto"><i class="fas fa-user"></i> Nombre completo</label>
            <input 
              type="text" 
              id="nombreCompleto" 
              v-model="usuario.nombreCompleto" 
              placeholder="Ej: Juan Pérez"
              required
            >
          </div>

          <div class="form-group">
            <label for="telefono"><i class="fas fa-phone"></i> Teléfono</label>
            <input 
              type="tel" 
              id="telefono" 
              v-model="usuario.telefono" 
              placeholder="Ej: +52 55 1234 5678"
            >
          </div>

          <div class="form-group">
            <label for="email"><i class="fas fa-envelope"></i> Correo electrónico</label>
            <input 
              type="email" 
              id="email" 
              v-model="usuario.email" 
              placeholder="Ej: usuario@example.com"
              required
            >
          </div>

          <div class="form-group">
            <label for="fechaNacimiento"><i class="fas fa-birthday-cake"></i> Fecha de nacimiento</label>
            <input 
              type="date" 
              id="fechaNacimiento" 
              v-model="usuario.fechaNacimiento"
            >
          </div>

          <div class="form-actions">
            <button type="submit" class="btn-guardar">
              <i class="fas fa-save"></i> Guardar cambios
            </button>
            <button type="button" class="btn-cancelar" @click="resetForm">
              <i class="fas fa-undo"></i> Cancelar
            </button>
          </div>
        </form>

        <!-- Sección de mascotas -->
        <div class="mascotas-section">
          <h3><i class="fas fa-paw"></i> Tus Mascotas</h3>
          <div v-if="mascotas.length > 0" class="lista-mascotas">
            <div v-for="mascota in mascotas" :key="mascota.id" class="mascota-item">
              <span>{{ mascota.nombre }}</span>
              <span class="mascota-raza">{{ mascota.raza }}</span>
            </div>
          </div>
          <div v-else class="sin-mascotas">
            <p>No tienes mascotas registradas</p>
            <button class="btn-agregar-mascota">
              <i class="fas fa-plus"></i> Agregar mascota
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'PerfilConfiguracion',
  data() {
    return {
      usuario: {
        foto: null,
        nombreUsuario: 'usuarioEjemplo',
        nombreCompleto: 'Juan Pérez',
        telefono: '+52 55 1234 5678',
        email: 'usuario@example.com',
        fechaNacimiento: '1990-01-01'
      },
      mascotas: [
        { id: 1, nombre: 'Max', raza: 'Labrador' },
        { id: 2, nombre: 'Luna', raza: 'Pastor Alemán' }
      ]
    }
  },
  methods: {
    triggerFileInput() {
      this.$refs.fileInput.click();
    },
    handleFileUpload(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.usuario.foto = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    guardarPerfil() {
      // Aquí iría la lógica para guardar los cambios
      console.log('Perfil guardado:', this.usuario);
      alert('Los cambios se han guardado correctamente');
    },
    resetForm() {
      // Restablecer a los valores iniciales
      this.usuario = {
        ...this.usuario,
        foto: null,
        nombreUsuario: 'usuarioEjemplo',
        nombreCompleto: 'Juan Pérez',
        telefono: '+52 55 1234 5678',
        email: 'usuario@example.com',
        fechaNacimiento: '1990-01-01'
      };
    }
  }
}
</script>

<style scoped>
.perfil-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.perfil-card {
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.perfil-header {
  background: linear-gradient(135deg, #6e8efb, #a777e3);
  color: white;
  padding: 20px;
}

.perfil-header h2 {
  margin: 0;
  font-size: 1.5rem;
}

.perfil-body {
  padding: 25px;
}

.foto-perfil-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 25px;
}

.foto-perfil {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background-color: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  overflow: hidden;
  border: 3px solid #e0e0e0;
  margin-bottom: 10px;
}

.foto-perfil:hover {
  border-color: #a777e3;
}

.foto-preview {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.default-avatar {
  font-size: 5rem;
  color: #cccccc;
}

.btn-cambiar-foto {
  background: none;
  border: none;
  color: #6e8efb;
  cursor: pointer;
  font-size: 0.9rem;
}

.btn-cambiar-foto:hover {
  text-decoration: underline;
}

.form-perfil {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #555;
}

.form-group input {
  width: 100%;
  padding: 10px 15px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

.form-group input:focus {
  border-color: #6e8efb;
  outline: none;
  box-shadow: 0 0 0 2px rgba(110, 142, 251, 0.2);
}

.form-actions {
  grid-column: span 2;
  display: flex;
  justify-content: flex-end;
  gap: 15px;
  margin-top: 10px;
}

.btn-guardar, .btn-cancelar {
  padding: 10px 20px;
  border-radius: 6px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-guardar {
  background-color: #6e8efb;
  color: white;
  border: none;
}

.btn-guardar:hover {
  background-color: #5a7df4;
}

.btn-cancelar {
  background-color: #f5f5f5;
  color: #555;
  border: 1px solid #ddd;
}

.btn-cancelar:hover {
  background-color: #eaeaea;
}

.mascotas-section {
  border-top: 1px solid #eee;
  padding-top: 25px;
}

.mascotas-section h3 {
  color: #555;
  margin-bottom: 15px;
}

.lista-mascotas {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 15px;
}

.mascota-item {
  background: #f9f9f9;
  padding: 12px 15px;
  border-radius: 6px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-left: 4px solid #a777e3;
}

.mascota-raza {
  color: #777;
  font-size: 0.9rem;
}

.sin-mascotas {
  text-align: center;
  padding: 20px;
  color: #777;
}

.btn-agregar-mascota {
  background-color: #a777e3;
  color: white;
  border: none;
  padding: 8px 15px;
  border-radius: 6px;
  margin-top: 10px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-agregar-mascota:hover {
  background-color: #9566d1;
}

@media (max-width: 768px) {
  .form-perfil {
    grid-template-columns: 1fr;
  }
  
  .form-actions {
    grid-column: span 1;
    justify-content: center;
  }
}
</style>