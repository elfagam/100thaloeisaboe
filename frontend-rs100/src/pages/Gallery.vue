<template>
  <div class="gallery-showcase-page">
    <!-- Full Screen Background with blur effect of current image -->
    <div class="bg-blur" :style="{ backgroundImage: currentGalleryItem ? `url(${currentGalleryItem.photoPath.startsWith('http') ? currentGalleryItem.photoPath : 'http://localhost:8081' + currentGalleryItem.photoPath})` : '' }"></div>
    
    <div class="content-wrapper">
      <div class="photo-display-container">
        <transition name="fade-scale" mode="out-in">
          <img 
            v-if="currentGalleryItem && currentGalleryItem.photoPath" 
            :key="currentGalleryItem.id"
            :src="currentGalleryItem.photoPath.startsWith('http') ? currentGalleryItem.photoPath : 'http://localhost:8081' + currentGalleryItem.photoPath" 
            alt="Gallery Photo" 
            class="main-photo"
          >
          <div v-else class="photo-placeholder">
            <span>📸 Memuat Galeri Sejarah...</span>
          </div>
        </transition>

        <!-- Countdown Overlay -->
        <div class="countdown-badge" v-if="isAutoPlaying && countdown > 0">
          ⏱️ Berpindah dalam {{ countdown }}d
        </div>

        <!-- Progress Bar -->
        <div class="photo-progress-bar" v-if="isAutoPlaying && countdown > 0" :style="{ width: progressPercentage + '%' }"></div>
      </div>

      <div class="narration-container glass-panel">
        <h2 class="text-gold-gradient font-historical">{{ currentGalleryItem ? currentGalleryItem.title : 'Menelusuri Jejak Waktu' }}</h2>
        <p class="narration-text">{{ currentGalleryItem ? currentGalleryItem.narration : 'Memuat narasi...' }}</p>
        
        <div class="controls-row">
          <button v-if="!isAutoPlaying" class="control-btn play-btn" @click="speakNarration(true)">▶️ Mulai Slideshow</button>
          <button v-else class="control-btn stop-btn" @click="stopAutoplay">⏹️ Hentikan</button>
          
          <!-- Mute Button -->
          <button class="control-btn mute-btn" @click="toggleMute">
            {{ isMuted ? '🔇 Buka Suara' : '🔊 Senyap' }}
          </button>
          
          <!-- Interval Control -->
          <div class="interval-control">
            <label>⏱️ Jeda:</label>
            <select v-model="slideDuration" class="control-select">
              <option :value="3">3s</option>
              <option :value="5">5s</option>
              <option :value="10">10s</option>
              <option :value="15">15s</option>
            </select>
          </div>
          
          <button class="control-btn" @click="prevGalleryItem">◀️ Prev</button>
          <button class="control-btn" @click="nextGalleryItem">Next ▶️</button>
          
          <span class="index-indicator">{{ currentIndex + 1 }} / {{ galleryItems.length }}</span>
          
          <!-- Back to Dashboard button -->
          <router-link to="/dashboard" class="control-btn exit-btn">✖ Keluar</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

const galleryItems = ref([])
const currentIndex = ref(0)
const isAutoPlaying = ref(false)
const isMuted = ref(true)
const slideDuration = ref(5)
const countdown = ref(0)
const progressPercentage = ref(0)
let timer = null
let progressTimer = null

const currentGalleryItem = computed(() => {
  return galleryItems.value[currentIndex.value] || null
})

const fetchGallery = async () => {
  try {
    const response = await fetch('http://localhost:8081/api/gallery')
    const data = await response.json()
    if (data && data.length > 0) {
      galleryItems.value = data
      currentIndex.value = 0
    }
  } catch (error) {
    console.error('Failed to fetch gallery data:', error)
  }
}

const nextGalleryItem = () => {
  if (galleryItems.value.length === 0) return
  currentIndex.value = (currentIndex.value + 1) % galleryItems.value.length
  resetTimers()
  if (isAutoPlaying.value) {
    speakNarration()
  }
}

const prevGalleryItem = () => {
  if (galleryItems.value.length === 0) return
  currentIndex.value = (currentIndex.value - 1 + galleryItems.value.length) % galleryItems.value.length
  resetTimers()
  if (isAutoPlaying.value) {
    speakNarration()
  }
}

const toggleMute = () => {
  isMuted.value = !isMuted.value
  if (isMuted.value) {
    window.speechSynthesis.cancel()
    if (isAutoPlaying.value) {
      startCountdown(slideDuration.value) // Start countdown immediately if muted
    }
  } else if (isAutoPlaying.value) {
    speakNarration()
  }
}

const speakNarration = (startAuto = false) => {
  if (startAuto) {
    isAutoPlaying.value = true
  }
  
  const text = currentGalleryItem.value ? currentGalleryItem.value.narration : ''
  if (!text || isMuted.value) {
    startCountdown(slideDuration.value) // Just wait if muted or no text
    return
  }

  // Cancel any ongoing speech
  window.speechSynthesis.cancel()

  const utterance = new SpeechSynthesisUtterance(text)
  utterance.lang = 'id-ID' // Indonesian
  
  // Find a good voice if available
  const voices = window.speechSynthesis.getVoices()
  const indonesianVoice = voices.find(v => v.lang.includes('id'))
  if (indonesianVoice) {
    utterance.voice = indonesianVoice
  }

  utterance.onend = () => {
    if (isAutoPlaying.value) {
      startCountdown(slideDuration.value) // Pause after narration
    }
  }

  utterance.onerror = (e) => {
    console.error('SpeechSynthesis error:', e)
    if (isAutoPlaying.value) {
      startCountdown(slideDuration.value)
    }
  }

  window.speechSynthesis.speak(utterance)
}

const startCountdown = (seconds) => {
  countdown.value = seconds
  progressPercentage.value = 100
  
  const step = 100 / (seconds * 10) // 10 steps per second
  
  clearInterval(progressTimer)
  progressTimer = setInterval(() => {
    progressPercentage.value -= step
    if (progressPercentage.value <= 0) {
      clearInterval(progressTimer)
    }
  }, 100)

  clearInterval(timer)
  timer = setInterval(() => {
    countdown.value -= 1
    if (countdown.value <= 0) {
      clearInterval(timer)
      clearInterval(progressTimer)
      nextGalleryItem()
    }
  }, 1000)
}

const stopAutoplay = () => {
  isAutoPlaying.value = false
  clearInterval(timer)
  clearInterval(progressTimer)
  window.speechSynthesis.cancel()
  countdown.value = 0
  progressPercentage.value = 0
}

const resetTimers = () => {
  clearInterval(timer)
  clearInterval(progressTimer)
  countdown.value = 0
  progressPercentage.value = 0
}

onMounted(() => {
  fetchGallery()
  // Auto play on mount after a short delay
  setTimeout(() => {
    if (galleryItems.value.length > 0) {
      speakNarration(true)
    }
  }, 2000)
})

onUnmounted(() => {
  stopAutoplay()
})
</script>

<style lang="scss" scoped>
.gallery-showcase-page {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: #03140A;
  color: white;
  z-index: 2000; /* Above navbar if any */
  overflow: hidden;
}

.bg-blur {
  position: absolute;
  top: -20px; left: -20px; right: -20px; bottom: -20px;
  background-size: cover;
  background-position: center;
  filter: blur(30px) brightness(0.3);
  z-index: 1;
  transition: background-image 1s ease;
}

.content-wrapper {
  position: relative;
  z-index: 2;
  height: 100%;
  display: grid;
  grid-template-rows: 1fr auto;
  padding: 2rem;
  box-sizing: border-box;
}

.photo-display-container {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  overflow: hidden;
  border-radius: 16px;
  background: rgba(0,0,0,0.4);
  border: 1px solid rgba(247, 150, 51, 0.2);
}

.main-photo {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  box-shadow: 0 20px 50px rgba(0,0,0,0.5);
}

.photo-placeholder {
  font-size: 1.5rem;
  color: #94A3B8;
}

.countdown-badge {
  position: absolute;
  top: 20px;
  right: 20px;
  background: rgba(3, 20, 10, 0.8);
  padding: 0.8rem 1.5rem;
  border-radius: 30px;
  border: 1px solid #F79633;
  color: #FFC385;
  font-weight: 700;
}

.photo-progress-bar {
  position: absolute;
  bottom: 0;
  left: 0;
  height: 6px;
  background: #F79633;
  box-shadow: 0 0 10px #F79633;
  transition: width 0.1s linear;
}

.narration-container {
  margin-top: 2rem;
  padding: 2rem;
  border-radius: 16px;
  background: rgba(11, 48, 29, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(247, 150, 51, 0.2);
  text-align: center;

  h2 {
    font-size: 2rem;
    margin-bottom: 1rem;
  }

  .narration-text {
    font-size: 1.3rem;
    line-height: 1.8;
    color: #E2E8F0;
    max-width: 1000px;
    margin: 0 auto 2rem auto;
  }
}

.controls-row {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
}

.control-btn {
  background: rgba(247, 150, 51, 0.1);
  border: 1px solid rgba(247, 150, 51, 0.3);
  color: #FFC385;
  padding: 0.8rem 2rem;
  border-radius: 30px;
  font-weight: 700;
  cursor: pointer;
  text-decoration: none;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background: #F79633;
    color: #03140A;
    border-color: #F79633;
  }

  &.play-btn {
    background: rgba(0, 168, 89, 0.2);
    border-color: rgba(0, 168, 89, 0.4);
    color: #52ec9a;
    &:hover {
      background: #00A859;
      color: white;
    }
  }

  &.stop-btn {
    background: rgba(239, 68, 68, 0.2);
    border-color: rgba(239, 68, 68, 0.4);
    color: #fca5a5;
    &:hover {
      background: #EF4444;
      color: white;
    }
  }
  
  &.mute-btn {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.2);
    color: #FFFFFF;
    &:hover {
      background: rgba(255, 255, 255, 0.2);
    }
  }

  &.exit-btn {
    background: rgba(255, 255, 255, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.2);
    color: #FFFFFF;
    &:hover {
      background: rgba(255, 255, 255, 0.2);
    }
  }
}

.interval-control {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #FFC385;
  font-weight: 700;
  
  .control-select {
    background: rgba(247, 150, 51, 0.1);
    border: 1px solid rgba(247, 150, 51, 0.3);
    color: #FFC385;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    font-weight: 700;
    cursor: pointer;
    outline: none;
    
    option {
      background: #03140A;
      color: #FFC385;
    }
    
    &:focus {
      border-color: #F79633;
    }
  }
}

.index-indicator {
  font-weight: 700;
  color: #94A3B8;
  margin-left: 1rem;
}

/* Transitions */
.fade-scale-enter-active, .fade-scale-leave-active {
  transition: all 0.8s cubic-bezier(0.16, 1, 0.3, 1);
}
.fade-scale-enter-from {
  opacity: 0;
  transform: scale(0.95);
}
.fade-scale-leave-to {
  opacity: 0;
  transform: scale(1.05);
}
</style>
