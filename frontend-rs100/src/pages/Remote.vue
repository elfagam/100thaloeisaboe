<template>
  <div class="remote-container">
    <div class="remote-card">
      <div class="icon-wrap">
        <span class="pulse-icon">📡</span>
      </div>
      <h2>REMOT KONTROL VIP</h2>
      <p>Klik tombol di bawah ini untuk memulai prosesi peluncuran di layar proyektor utama.</p>
      
      <button 
        class="trigger-btn" 
        @click="triggerLaunch" 
        :disabled="isTriggered"
        :class="{ 'triggered': isTriggered }"
      >
        <span v-if="!isTriggered">🚀 LUNCURKAN SEKARANG</span>
        <span v-else>✅ PERINTAH DIKIRIM</span>
      </button>
      
      <div v-if="isTriggered" class="success-msg">
        Sinyal telah dikirim ke layar utama.
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isTriggered = ref(false)

const triggerLaunch = async () => {
  isTriggered.value = true
  
  // Mengirim sinyal melalui jaringan via Backend Server (Mendukung Perangkat Berbeda)
  try {
    const backendUrl = `${window.location.protocol}//${window.location.hostname}:8081`
    await fetch(`${backendUrl}/api/remote/trigger`, {
      method: 'POST'
    })
  } catch (error) {
    console.error('Gagal mengirim sinyal remot:', error)
  }
  
  // Kembalikan status tombol setelah 3 detik
  setTimeout(() => {
    isTriggered.value = false
  }, 3000)
}
</script>

<style scoped>
.remote-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: var(--color-bg-dark, #03140A);
  color: white;
}

.remote-card {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(0, 255, 135, 0.3);
  padding: 4rem;
  border-radius: 30px;
  text-align: center;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(10px);
  max-width: 500px;
  width: 90%;
}

.icon-wrap {
  font-size: 4rem;
  margin-bottom: 1rem;
  animation: float 3s ease-in-out infinite;
}

h2 {
  color: var(--color-gold, #F79633);
  font-size: 2rem;
  margin-bottom: 1rem;
  letter-spacing: 2px;
}

p {
  color: #94A3B8;
  font-size: 1.1rem;
  margin-bottom: 3rem;
  line-height: 1.6;
}

.trigger-btn {
  background: linear-gradient(135deg, #00FF87 0%, #00A859 100%);
  color: #03140A;
  border: none;
  padding: 1.5rem 3rem;
  font-size: 1.5rem;
  font-weight: 900;
  border-radius: 50px;
  cursor: pointer;
  box-shadow: 0 10px 30px rgba(0, 255, 135, 0.4);
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  width: 100%;
}

.trigger-btn:hover:not(:disabled) {
  transform: scale(1.05);
  box-shadow: 0 15px 40px rgba(0, 255, 135, 0.6);
}

.trigger-btn:active:not(:disabled) {
  transform: scale(0.95);
}

.trigger-btn.triggered {
  background: #334155;
  color: #94A3B8;
  box-shadow: none;
  cursor: not-allowed;
  transform: scale(1);
}

.success-msg {
  margin-top: 1.5rem;
  color: #00FF87;
  font-weight: 700;
  animation: fadeIn 0.5s ease;
}

@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
