/* eslint-disable no-unused-vars */
import { defineStore } from 'pinia';
export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    isAuthenticated: false,
  }),
  actions: {
    
    async login(email, password) {
      // Lógica para llamar al backend (simulada)
      this.user = { email };
      this.isAuthenticated = true;
      
      return true;
    },
    async register(name, email, password) {
      // Lógica para llamar al backend (simulada)
      this.user = { name, email };
      this.isAuthenticated = true;
      return true;
    },
  },
});