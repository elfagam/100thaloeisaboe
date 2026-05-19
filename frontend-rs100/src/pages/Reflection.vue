<template>
  <div class="reflection-page">
    <router-link to="/launching" class="back-btn">⬅ Kembali ke Peluncuran</router-link>

    <div class="reflection-panel-card">
      <video ref="reflectionVideo" class="reflection-video-bg" src="/Video-Narasi.mp4" playsinline loop muted></video>
      <video ref="reflectionAudio" src="/Video-Narasi.mp4" playsinline style="display: none;" @ended="onReflectionAudioEnded"></video>

      <!-- Floating Narrator -->
      <div 
        ref="narratorPortal" 
        class="narrator-portal" 
        :class="{ expanded: isNarratorExpanded, dragging: isDragging }"
        :style="narratorStyle"
        @mousedown="startDrag" @touchstart="startDrag" @click="toggleNarratorSize"
      >
        <video ref="narratorVisual" class="narrator-video-window" src="/Video-Narasi.mp4" playsinline muted></video>
        <div class="narrator-overlay">
          <span class="narrator-label">NARATOR</span>
          <span class="drag-handle">⠿</span>
        </div>
      </div>

      <div class="reflection-header">
        <span class="reflection-tag font-historical text-gold-gradient">📖 PENJELASAN SATU ABAD (AUTO-PLAY)</span>
        <div class="reflection-controls">
          <button @click="prevReflection" class="control-arrow-btn" :disabled="currentReflectionIndex === 0">◀</button>
          <span class="control-count">{{ currentReflectionIndex + 1 }} / {{ reflections.length }}</span>
          <button @click="nextReflection" class="control-arrow-btn" :disabled="currentReflectionIndex === reflections.length - 1">▶</button>
          <button @click="toggleReflectionAutoPlay" class="control-play-btn" :class="{ active: isReflectionPlaying }">
            <span v-if="isReflectionPlaying">⏸ Jeda</span>
            <span v-else>▶ Putar</span>
          </button>
          <button @click="stopAndResetReflection" class="control-stop-btn">⏹ Berhenti</button>
          
          <!-- 🔊 Volume Control -->
          <div class="volume-control">
            <span class="volume-icon">{{ globalVolume > 0.5 ? '🔊' : globalVolume > 0 ? '🔉' : '🔇' }}</span>
            <input type="range" v-model.number="globalVolume" min="0" max="1" step="0.05" class="volume-slider" @input="updateVolume" />
          </div>

          <button @click="showSyncSettings = !showSyncSettings" class="control-settings-btn" :class="{ active: showSyncSettings }">⚙️</button>
        </div>
      </div>

      <transition name="slide-fade" mode="out-in">
        <div v-if="showSyncSettings" class="sync-settings-panel">
          <div class="settings-header">
            <div class="header-left">
              <h5>⏱️ PENGATURAN WAKTU SLIDE</h5>
              <div class="live-timer" :class="{ pulsing: isReflectionPlaying }">
                <span class="timer-label">WAKTU SEKARANG:</span>
                <span class="timer-value">{{ formatTime(reflectionPlaybackTime) }}</span>
              </div>
            </div>
            <button @click="resetTimestamps" class="mini-reset-btn">Reset Default</button>
          </div>
          <div class="settings-grid">
            <div v-for="(ts, index) in reflectionTimestamps" :key="index" class="settings-item">
              <label>Slide {{ index + 1 }}</label>
              <input type="number" v-model.number="reflectionTimestamps[index]" @change="saveTimestamps" step="0.5" min="0" :class="{ error: index > 0 && reflectionTimestamps[index] <= reflectionTimestamps[index-1] }" />
            </div>
          </div>
        </div>
      </transition>

      <div class="reflection-body">
        <p class="reflection-text">{{ reflections[currentReflectionIndex] }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const reflectionVideo = ref(null)
const reflectionAudio = ref(null)
const narratorPortal = ref(null)
const narratorVisual = ref(null)

const reflections = [
  "100 Tahun Pemanfaatan Rumah Sakit. Satu abad perjalanan ini bukan sekadar hitungan waktu, tetapi menjadi bukti nyata pengabdian, pelayanan, perjuangan, dan dedikasi rumah sakit dalam memberikan pelayanan kesehatan kepada masyarakat.",
  "Selama seratus tahun, rumah sakit telah menjadi tempat lahirnya harapan, tempat bertumbuhnya kehidupan, serta tempat masyarakat memperoleh pelayanan, perhatian, dan kepedulian. Berbagai tantangan telah dilalui, mulai dari keterbatasan sarana di masa awal berdiri, perkembangan ilmu pengetahuan dan teknologi kesehatan, hingga menghadapi berbagai dinamika pelayanan kesehatan modern. Semua itu dijalani dengan semangat kemanusiaan dan komitmen untuk terus melayani.",
  "Kiranya peringatan 100 Tahun Pemanfaatan Rumah Sakit ini menjadi sumber motivasi bagi seluruh jajaran untuk terus berkarya, mengabdi, dan memberikan pelayanan terbaik bagi masyarakat. Semoga rumah sakit ini semakin maju, humanis, profesional, dan tetap menjadi tempat yang menghadirkan harapan serta kehidupan bagi semua."
]

const currentReflectionIndex = ref(0)
const isReflectionPlaying = ref(false)
const reflectionPlaybackTime = ref(0)
const globalVolume = ref(0.6)
const showSyncSettings = ref(false)

const defaultTimestamps = [0, 20, 60]
const reflectionTimestamps = ref([...defaultTimestamps])

onMounted(() => {
  const saved = JSON.parse(localStorage.getItem('rs100_reflection_ts'))
  reflectionTimestamps.value = (saved && saved.length === reflections.length) ? saved : [...defaultTimestamps]
  
  // Auto play when page loads for entertainment
  setTimeout(() => {
    startReflectionAutoPlay()
  }, 1000)
})

const formatTime = (seconds) => {
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  const ms = Math.floor((seconds % 1) * 10)
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}.${ms}`
}

const updateVolume = () => {
  if (reflectionAudio.value) reflectionAudio.value.volume = globalVolume.value
  if (reflectionVideo.value) reflectionVideo.value.volume = globalVolume.value
}

const saveTimestamps = () => {
  localStorage.setItem('rs100_reflection_ts', JSON.stringify(reflectionTimestamps.value))
}

const resetTimestamps = () => {
  reflectionTimestamps.value = [...defaultTimestamps]
  saveTimestamps()
}

const startReflectionAutoPlay = () => {
  isReflectionPlaying.value = true
  if (reflectionVideo.value) reflectionVideo.value.play().catch(() => {})
  if (narratorVisual.value) narratorVisual.value.play().catch(() => {})
  
  if (reflectionAudio.value) {
    reflectionAudio.value.play().catch(() => {})
    reflectionAudio.value.ontimeupdate = () => {
      const currentTime = reflectionAudio.value.currentTime
      reflectionPlaybackTime.value = currentTime
      let targetIndex = 0
      for (let i = 0; i < reflectionTimestamps.value.length; i++) {
        if (currentTime >= reflectionTimestamps.value[i]) targetIndex = i
      }
      if (targetIndex !== currentReflectionIndex.value) currentReflectionIndex.value = targetIndex
    }
  }
}

const pauseReflectionAutoPlay = () => {
  isReflectionPlaying.value = false
  if (reflectionVideo.value) reflectionVideo.value.pause()
  if (narratorVisual.value) narratorVisual.value.pause()
  if (reflectionAudio.value) reflectionAudio.value.pause()
}

const toggleReflectionAutoPlay = () => isReflectionPlaying.value ? pauseReflectionAutoPlay() : startReflectionAutoPlay()

const stopAndResetReflection = () => {
  isReflectionPlaying.value = false
  currentReflectionIndex.value = 0
  if (reflectionVideo.value) {
    reflectionVideo.value.pause()
    reflectionVideo.value.currentTime = 0
  }
  if (narratorVisual.value) {
    narratorVisual.value.pause()
    narratorVisual.value.currentTime = 0
  }
  if (reflectionAudio.value) {
    reflectionAudio.value.pause()
    reflectionAudio.value.currentTime = 0
    reflectionPlaybackTime.value = 0
  }
}

const onReflectionAudioEnded = () => {
  stopAndResetReflection()
  // Auto loop for entertainment!
  setTimeout(() => {
    startReflectionAutoPlay()
  }, 1000)
}

const nextReflection = () => {
  if (currentReflectionIndex.value < reflections.length - 1) {
    currentReflectionIndex.value++
    syncAudioToCurrentIndex()
  }
}

const prevReflection = () => {
  if (currentReflectionIndex.value > 0) {
    currentReflectionIndex.value--
    syncAudioToCurrentIndex()
  }
}

const syncAudioToCurrentIndex = () => {
  const targetTime = reflectionTimestamps.value[currentReflectionIndex.value]
  if (reflectionAudio.value) reflectionAudio.value.currentTime = targetTime
  if (reflectionVideo.value) reflectionVideo.value.currentTime = targetTime
  if (narratorVisual.value) narratorVisual.value.currentTime = targetTime
}

// Draggable Narrator Logic
const isNarratorExpanded = ref(false)
const isDragging = ref(false)
const narratorPos = ref({ x: 0, y: 0 })
const startPos = ref({ x: 0, y: 0 })

const narratorStyle = computed(() => {
  if (isNarratorExpanded.value) return {}
  return {
    transform: `translate(${narratorPos.value.x}px, ${narratorPos.value.y}px)`,
    cursor: isDragging.value ? 'grabbing' : 'grab'
  }
})

const toggleNarratorSize = (e) => {
  if (isDragging.value) return
  isNarratorExpanded.value = !isNarratorExpanded.value
}

const startDrag = (e) => {
  if (isNarratorExpanded.value) return
  isDragging.value = true
  const clientX = e.type.startsWith('touch') ? e.touches[0].clientX : e.clientX
  const clientY = e.type.startsWith('touch') ? e.touches[0].clientY : e.clientY
  startPos.value = {
    x: clientX - narratorPos.value.x,
    y: clientY - narratorPos.value.y
  }
  
  const moveHandler = (e) => {
    if (!isDragging.value) return
    const cx = e.type.startsWith('touch') ? e.touches[0].clientX : e.clientX
    const cy = e.type.startsWith('touch') ? e.touches[0].clientY : e.clientY
    narratorPos.value = {
      x: cx - startPos.value.x,
      y: cy - startPos.value.y
    }
  }
  
  const stopHandler = () => {
    isDragging.value = false
    window.removeEventListener('mousemove', moveHandler)
    window.removeEventListener('mouseup', stopHandler)
    window.removeEventListener('touchmove', moveHandler)
    window.removeEventListener('touchend', stopHandler)
  }
  
  window.addEventListener('mousemove', moveHandler)
  window.addEventListener('mouseup', stopHandler)
  window.addEventListener('touchmove', moveHandler, { passive: false })
  window.addEventListener('touchend', stopHandler)
}
</script>

<style scoped lang="scss">
.reflection-page {
  min-height: 100vh;
  background: #03140a;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  position: relative;
  overflow: hidden;
}

.back-btn {
  position: absolute;
  top: 2rem;
  left: 2rem;
  color: #00FF87;
  text-decoration: none;
  font-weight: 800;
  border: 1px solid rgba(0, 255, 137, 0.4);
  padding: 0.6rem 1.2rem;
  border-radius: 30px;
  background: rgba(11, 48, 29, 0.6);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
  z-index: 10;
  box-shadow: 0 4px 15px rgba(0,0,0,0.3);
  
  &:hover {
    background: rgba(0, 255, 135, 0.2);
    transform: translateX(-5px);
    border-color: #00FF87;
  }
}

.reflection-panel-card {
  position: relative; 
  overflow: hidden;
  background: rgba(3, 20, 10, 0.9); 
  backdrop-filter: blur(25px);
  border: 1px solid rgba(247, 150, 51, 0.3); 
  border-radius: 24px;
  padding: 2.5rem; 
  text-align: left; 
  box-shadow: 0 30px 60px rgba(0,0,0,0.5);
  width: 90%;
  max-width: 1000px;
  min-height: 500px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  .reflection-video-bg {
    position: absolute; top: 0; left: 0; width: 100%; height: 100%;
    object-fit: cover; opacity: 0.25; z-index: 0; pointer-events: none;
    filter: grayscale(0.2) contrast(1.1);
  }

  .reflection-header {
    position: relative; z-index: 5; display: flex; justify-content: space-between;
    align-items: center; border-bottom: 1px solid rgba(247, 150, 51, 0.2);
    padding-bottom: 1rem; margin-bottom: 1.5rem;
  }

  .reflection-tag { color: #F79633; font-weight: 800; letter-spacing: 2px; font-size: 0.85rem; }

  .reflection-controls { display: flex; align-items: center; gap: 1rem; }

  .control-arrow-btn {
    background: transparent; border: none; color: #F79633; cursor: pointer;
    font-size: 1.2rem; transition: transform 0.2s;
    &:hover:not(:disabled) { transform: scale(1.3); }
    &:disabled { opacity: 0.3; cursor: not-allowed; }
  }

  .control-count { color: #FFF; font-weight: 800; font-size: 0.9rem; min-width: 40px; text-align: center; }

  .control-play-btn, .control-stop-btn {
    border-radius: 20px; padding: 6px 16px; font-size: 0.75rem; font-weight: 700;
    cursor: pointer; transition: all 0.3s;
  }

  .control-play-btn {
    background: rgba(247, 150, 51, 0.2); border: 1px solid #F79633; color: #FFF;
    &.active { background: #22C55E; border-color: #22C55E; }
  }

  .control-stop-btn { background: rgba(239, 68, 68, 0.1); border: 1px solid #EF4444; color: #EF4444; }

  .volume-control {
    display: flex;
    align-items: center;
    gap: 8px;
    background: rgba(255, 255, 255, 0.05);
    padding: 4px 12px;
    border-radius: 20px;
    border: 1px solid rgba(247, 150, 51, 0.2);

    .volume-icon { font-size: 0.9rem; }
    
    .volume-slider {
      width: 80px;
      height: 4px;
      -webkit-appearance: none;
      appearance: none;
      background: rgba(247, 150, 51, 0.2);
      border-radius: 2px;
      outline: none;

      &::-webkit-slider-thumb {
        -webkit-appearance: none;
        appearance: none;
        width: 12px;
        height: 12px;
        background: #F79633;
        border-radius: 50%;
        cursor: pointer;
        box-shadow: 0 0 10px rgba(247, 150, 51, 0.5);
      }
    }
  }

  .control-settings-btn {
    background: transparent; border: none; font-size: 1.2rem; cursor: pointer;
    transition: transform 0.3s;
    &.active { transform: rotate(90deg); filter: hue-rotate(90deg); }
  }

  .reflection-body { position: relative; z-index: 5; min-height: 150px; padding-right: 120px; display: flex; align-items: center; }

  .reflection-text {
    font-size: 1.3rem; line-height: 1.9; color: #E2E8F0; margin: 0;
    text-shadow: 0 2px 10px rgba(0,0,0,0.5);
  }
}

.sync-settings-panel {
  position: relative; z-index: 10; background: rgba(0,0,0,0.85);
  border: 1px solid rgba(247, 150, 51, 0.5); border-radius: 16px;
  padding: 1.5rem; margin-bottom: 1.5rem;

  .settings-header {
    display: flex; justify-content: space-between; align-items: center;
    margin-bottom: 1rem; border-bottom: 1px solid rgba(255,255,255,0.1); padding-bottom: 0.5rem;
    
    h5 { color: #00FF87; margin: 0; font-size: 0.9rem; }
    
    .header-left { display: flex; align-items: center; gap: 1.5rem; }
    
    .live-timer {
      display: flex; align-items: center; gap: 8px;
      background: #111; padding: 4px 10px; border-radius: 6px;
      border: 1px solid #333;
      &.pulsing { border-color: #22C55E; animation: timer-pulse 1s infinite alternate; }
      .timer-label { font-size: 0.6rem; color: #94A3B8; }
      .timer-value { font-family: monospace; font-size: 1rem; color: #22C55E; font-weight: 800; }
    }

    .mini-reset-btn {
      background: transparent; border: 1px solid rgba(255,255,255,0.2);
      color: #94A3B8; font-size: 0.7rem; padding: 4px 10px; border-radius: 4px;
      cursor: pointer; &:hover { background: rgba(255,255,255,0.05); color: #FFF; }
    }
  }

  .settings-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 1rem; }

  .settings-item {
    display: flex; flex-direction: column; gap: 6px;
    label { font-size: 0.7rem; color: #94A3B8; font-weight: 700; }
    input {
      background: #111; border: 1px solid #F79633; color: #FFF;
      padding: 8px; border-radius: 8px; font-weight: 800; text-align: center;
      &.error { border-color: #EF4444; color: #EF4444; }
    }
  }
}

.narrator-portal {
  position: absolute; bottom: 2.5rem; right: 2.5rem; width: 120px; height: 120px;
  border-radius: 50%; border: 3px solid #F79633; overflow: hidden;
  box-shadow: 0 0 20px rgba(0,0,0,0.4); z-index: 10;
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  background: #000;
  
  &.expanded { width: 400px; height: 225px; border-radius: 16px; bottom: 50%; right: 50%; transform: translate(50%, 50%) !important; }
  
  video { width: 100%; height: 100%; object-fit: cover; }
  
  .narrator-overlay {
    position: absolute; bottom: 0; left: 0; width: 100%;
    background: linear-gradient(to top, rgba(0,0,0,0.8), transparent);
    padding: 8px; display: flex; justify-content: space-between; align-items: center;
    opacity: 0; transition: opacity 0.3s;
  }
  
  &:hover .narrator-overlay { opacity: 1; }
  
  .narrator-label { color: #FFF; font-size: 0.6rem; font-weight: 800; }
  .drag-handle { color: #F79633; cursor: grab; }
}

@keyframes timer-pulse {
  from { box-shadow: 0 0 0px rgba(34, 197, 94, 0); }
  to { box-shadow: 0 0 10px rgba(34, 197, 94, 0.5); }
}

.slide-fade-enter-active, .slide-fade-leave-active { transition: all 0.3s ease; }
.slide-fade-enter-from, .slide-fade-leave-to { transform: translateY(20px); opacity: 0; }
</style>
