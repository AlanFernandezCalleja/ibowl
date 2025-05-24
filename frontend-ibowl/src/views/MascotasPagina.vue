<template>
    <TablaMascotas v-if="mounted" />
</template>

<script>
import TablaMascotas from '@/components/TablaMascotas.vue';
export default {
    components: {
        TablaMascotas
    },
    data() {
        return {
            mounted: false
        }
    },
    mounted() {
        this.mounted = true;
    }
}
</script>