import { createRouter, createWebHistory } from 'vue-router'
import InicioPagina from '@/views/InicioPagina.vue'
import DashboardPagina from '@/views/DashboardPagina.vue'
import MascotasPagina from '@/views/MascotasPagina.vue'
import PerfilConfiguracion from '@/views/PerfilConfiguracion.vue'


const routes = [
  {
    path: '/',
    name: 'inicio',
    component: InicioPagina
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: DashboardPagina
  },
  {
    path: '/mascotas',
    name: 'mascotas',
    component: MascotasPagina
  },
  {
    path: '/configuracion',
    name: 'configuracion',
    component: PerfilConfiguracion
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
