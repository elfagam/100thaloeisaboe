<template>
  <div class="preshow-container" @click="startAudio">
    <!-- Video Latar Belakang (Blur) -->
    <video 
      ref="bgVideoPlayer"
      :src="videoList[currentVideoIndex]"
      muted
      class="bg-video"
    ></video>

    <!-- Pemutar Video Utama -->
    <video 
      ref="videoPlayer"
      :src="videoList[currentVideoIndex]"
      @ended="playNextVideo"
      @timeupdate="updateProgress"
      @loadedmetadata="onLoadedMetadata"
      @play="syncPlay"
      @pause="syncPause"
      autoplay
      class="main-video"
    ></video>

    <!-- Suara Latar (Ambient BGM) -->
    <video 
      ref="bgMusicPlayer"
      src="/suara_latar.mp4"
      loop
      class="hidden-player"
    ></video>

    <!-- Overlay Jeda Antar Video -->
    <div v-if="isTransitioning" class="transition-overlay">
      <div class="transition-content">
        <div class="pulse-loader"></div>
        <p>Menyiapkan Video Berikutnya...</p>
        <span class="next-title">{{ getNextVideoTitle() }}</span>
      </div>
    </div>

    <!-- Overlay Klik untuk Suara (Hanya muncul di awal jika audio terblokir) -->
    <div v-if="showAudioOverlay" class="audio-overlay">
      <div class="glow-card">
        <h2>🔊 KLIK UNTUK MENGAKTIFKAN SUARA</h2>
        <p>Pastikan volume speaker Anda sudah menyala</p>
      </div>
    </div>

    <!-- Overlay Judul Kecil di Pojok -->
    <div class="brand-overlay">
      <h1>100 TAHUN RSAS</h1>
      <p>Video Ucapan Selamat</p>
    </div>

    <!-- Tombol Navigasi ke Halaman Launching -->
    <button @click="goToLaunching" class="nav-button">
      Masuk Sesi Peluncuran ➔
    </button>

    <!-- Control Panel Pengaturan Video -->
    <div class="video-controls" @click.stop>
      <div class="progress-container" @click="seekVideo">
        <div class="progress-bar" :style="{ width: videoProgress + '%' }"></div>
      </div>
      
      <div class="controls-row">
        <div class="left-controls">
          <button @click="playPrevVideo" class="control-btn" title="Video Sebelumnya">⏮</button>
          <button @click="togglePlay" class="control-btn main-play-btn" :title="isPlaying ? 'Pause' : 'Play'">
            {{ isPlaying ? '⏸' : '▶' }}
          </button>
          <button @click="playNextVideo" class="control-btn" title="Video Berikutnya">⏭</button>
          
          <div class="volume-container">
            <button @click="toggleMute" class="control-btn">
              {{ isMuted ? '🔇' : '🔊' }}
            </button>
            <input type="range" min="0" max="1" step="0.05" v-model="volumeLevel" @input="updateVolume" class="volume-slider" />
          </div>
          
          <span class="time-display">{{ formattedTime }}</span>
        </div>
        
        <div class="right-controls">
          <span class="video-title">{{ currentVideoTitle }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const videoPlayer = ref(null)
const bgVideoPlayer = ref(null)
const bgMusicPlayer = ref(null)
const showAudioOverlay = ref(true)
const isTransitioning = ref(false)

// Daftar video di folder public/ucapan
const videoList = ref([
  '/ucapan/video 1.mp4',
  '/ucapan/video 2.mp4',
  '/ucapan/video 3.mp4',
  '/ucapan/video 4.mp4',
  '/ucapan/video 5.mp4',
  '/ucapan/video 6.mp4',
  '/ucapan/video 7.mp4',
  '/ucapan/video 8.mp4',
  '/ucapan/video 9.mp4',
  '/ucapan/video 10.mp4',
  '/ucapan/video 11.mp4',
  
])

const currentVideoIndex = ref(0)
const isPlaying = ref(true)
const isMuted = ref(false)
const volumeLevel = ref(1)
const videoProgress = ref(0)
const currentTime = ref(0)
const duration = ref(0)

const currentVideoTitle = computed(() => {
  const path = videoList.value[currentVideoIndex.value]
  return path.split('/').pop().replace('.mp4', '')
})

const getNextVideoTitle = () => {
  const nextIndex = (currentVideoIndex.value + 1) % videoList.value.length
  const path = videoList.value[nextIndex]
  return path.split('/').pop().replace('.mp4', '')
}

const formattedTime = computed(() => {
  const format = (seconds) => {
    const mins = Math.floor(seconds / 60)
    const secs = Math.floor(seconds % 60)
    return `${mins}:${secs < 10 ? '0' : ''}${secs}`
  }
  return `${format(currentTime.value)} / ${format(duration.value)}`
})

const playNextVideo = () => {
  if (isTransitioning.value) return // Cegah double trigger
  
  isTransitioning.value = true
  
  // Naikkan volume suara latar saat jeda (transisi)
  if (bgMusicPlayer.value) {
    fadeAudio(bgMusicPlayer.value, 0.6, 500) // Naik ke 0.6 dalam 500ms
  }
  
  // Pause video yang sedang berjalan
  if (videoPlayer.value) videoPlayer.value.pause()
  if (bgVideoPlayer.value) bgVideoPlayer.value.pause()
  
  // Tunggu 3 detik
  setTimeout(() => {
    currentVideoIndex.value = (currentVideoIndex.value + 1) % videoList.value.length
    isTransitioning.value = false
    
    // Turunkan volume suara latar saat video mulai diputar
    if (bgMusicPlayer.value) {
      fadeAudio(bgMusicPlayer.value, 0.15, 500) // Turun ke 0.15 dalam 500ms
    }
    
    // Beri sedikit waktu untuk memuat src baru sebelum di-play
    setTimeout(() => {
      if (videoPlayer.value) videoPlayer.value.play()
    }, 200)
  }, 3000)
}

const playPrevVideo = () => {
  currentVideoIndex.value = (currentVideoIndex.value - 1 + videoList.value.length) % videoList.value.length
}

const togglePlay = () => {
  if (videoPlayer.value) {
    if (videoPlayer.value.paused) {
      videoPlayer.value.play()
    } else {
      videoPlayer.value.pause()
    }
  }
}

const syncPlay = () => {
  isPlaying.value = true
  if (bgVideoPlayer.value) {
    bgVideoPlayer.value.play()
  }
}

const syncPause = () => {
  isPlaying.value = false
  if (bgVideoPlayer.value) {
    bgVideoPlayer.value.pause()
  }
}

const toggleMute = () => {
  if (videoPlayer.value) {
    videoPlayer.value.muted = !videoPlayer.value.muted
    isMuted.value = videoPlayer.value.muted
    if (!isMuted.value && volumeLevel.value === 0) {
      volumeLevel.value = 0.5
      videoPlayer.value.volume = 0.5
    }
  }
}

const updateVolume = () => {
  if (videoPlayer.value) {
    videoPlayer.value.volume = volumeLevel.value
    videoPlayer.value.muted = volumeLevel.value === 0
    isMuted.value = videoPlayer.value.muted
  }
}

const seekVideo = (e) => {
  if (videoPlayer.value) {
    const rect = e.currentTarget.getBoundingClientRect()
    const x = e.clientX - rect.left
    const percentage = x / rect.width
    const time = percentage * videoPlayer.value.duration
    videoPlayer.value.currentTime = time
    if (bgVideoPlayer.value) {
      bgVideoPlayer.value.currentTime = time
    }
  }
}

const updateProgress = () => {
  if (videoPlayer.value) {
    currentTime.value = videoPlayer.value.currentTime
    videoProgress.value = (currentTime.value / duration.value) * 100
  }
}

const onLoadedMetadata = () => {
  if (videoPlayer.value) {
    duration.value = videoPlayer.value.duration || 0
    updateProgress()
  }
}

const startAudio = () => {
  if (videoPlayer.value) {
    videoPlayer.value.muted = false
    isMuted.value = false
    videoPlayer.value.play().catch(e => console.error('Play failed:', e))
  }
  if (bgVideoPlayer.value) {
    bgVideoPlayer.value.play().catch(e => console.error('BG Play failed:', e))
  }
  if (bgMusicPlayer.value) {
    bgMusicPlayer.value.volume = 0.15 // Mulai dengan volume redup
    bgMusicPlayer.value.play().catch(e => console.error('BG Music Play failed:', e))
  }
  showAudioOverlay.value = false
}

const goToLaunching = () => {
  router.push('/launching')
}

// Fungsi pembantu untuk membuat transisi volume halus (Fade)
const fadeAudio = (audioElement, targetVolume, durationMs) => {
  const startVolume = audioElement.volume
  const diff = targetVolume - startVolume
  const steps = 10
  const stepTime = durationMs / steps
  const stepValue = diff / steps
  
  let currentStep = 0
  const interval = setInterval(() => {
    currentStep++
    const newVolume = startVolume + (stepValue * currentStep)
    // Pastikan tidak melebihi batas 0 s.d 1
    audioElement.volume = Math.max(0, Math.min(1, newVolume))
    
    if (currentStep >= steps) {
      clearInterval(interval)
      audioElement.volume = targetVolume
    }
  }, stepTime)
}

onMounted(() => {
  if (videoPlayer.value) {
    videoPlayer.value.play().catch(e => {
      console.log('Autoplay blocked')
      showAudioOverlay.value = true
    })
  }
})
</script>

<style scoped>
.preshow-container {
  position: relative;
  width: 100vw;
  height: 100vh;
  background: #000;
  overflow: hidden;
  cursor: pointer;
}

.bg-video {
  position: absolute;
  top: -5%; left: -5%; width: 110%; height: 110%;
  object-fit: cover;
  filter: blur(50px) brightness(0.5);
  z-index: 1;
}

.main-video {
  position: relative;
  width: 100%;
  height: 100%;
  object-fit: contain;
  z-index: 2;
}

.hidden-player {
  display: none;
}

.audio-overlay {
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}

.glow-card {
  background: rgba(3, 20, 10, 0.9);
  border: 2px solid #00FF87;
  padding: 3rem;
  border-radius: 20px;
  text-align: center;
  box-shadow: 0 0 30px rgba(0, 255, 135, 0.5);
  color: white;
  
  h2 { color: #00FF87; margin-bottom: 1rem; font-size: 2rem; }
  p { color: #E2E8F0; font-size: 1.2rem; }
}

.brand-overlay {
  position: absolute;
  top: 2rem;
  left: 2rem;
  color: white;
  text-shadow: 0 2px 10px rgba(0,0,0,0.8);
  z-index: 10;
  
  h1 { font-size: 1.8rem; color: #F79633; margin: 0; font-weight: 900; }
  p { font-size: 1rem; color: #94A3B8; margin: 0; }
}

.nav-button {
  position: absolute;
  top: 2rem;
  right: 2rem;
  background: rgba(247, 150, 51, 0.9);
  color: white;
  border: none;
  padding: 1rem 2rem;
  border-radius: 50px;
  font-weight: 700;
  cursor: pointer;
  backdrop-filter: blur(5px);
  transition: all 0.3s ease;
  z-index: 10;
  box-shadow: 0 4px 15px rgba(247, 150, 51, 0.3);
}

.nav-button:hover {
  background: #F79633;
  transform: scale(1.05);
  box-shadow: 0 0 20px rgba(247, 150, 51, 0.5);
}

/* Video Controls Styles */
.video-controls {
  position: absolute;
  bottom: 0; left: 0; right: 0;
  background: linear-gradient(to top, rgba(0,0,0,0.9) 0%, rgba(0,0,0,0.5) 70%, transparent 100%);
  padding: 3rem 2rem 2rem 2rem;
  transition: opacity 0.3s ease, transform 0.3s ease;
  opacity: 0;
  transform: translateY(10px);
  z-index: 20;
}

.preshow-container:hover .video-controls {
  opacity: 1;
  transform: translateY(0);
}

.progress-container {
  width: 100%;
  height: 8px;
  background: rgba(255,255,255,0.2);
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 1.5rem;
  position: relative;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #00FF87, #F79633);
  border-radius: 4px;
  transition: width 0.1s linear;
}

.controls-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.left-controls {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.control-btn {
  background: transparent;
  border: none;
  color: white;
  font-size: 1.5rem;
  cursor: pointer;
  transition: transform 0.2s, color 0.2s;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  
  &:hover {
    color: #00FF87;
    transform: scale(1.2);
    background: rgba(255,255,255,0.1);
  }
}

.main-play-btn {
  font-size: 2rem;
  color: #F79633;
  width: 50px;
  height: 50px;
  background: rgba(255,255,255,0.05);
  
  &:hover {
    color: #00FF87;
    background: rgba(255,255,255,0.15);
  }
}

.volume-container {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  
  &:hover .volume-slider {
    width: 100px;
    opacity: 1;
  }
}

.volume-slider {
  width: 0;
  opacity: 0;
  transition: all 0.3s ease;
  accent-color: #00FF87;
  cursor: pointer;
}

.time-display {
  color: #94A3B8;
  font-family: monospace;
  font-size: 1rem;
}

.video-title {
  color: #E2E8F0;
  font-weight: 600;
  font-size: 1.1rem;
  letter-spacing: 1px;
}

/* Transition Overlay */
.transition-overlay {
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background: #000;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 50;
  animation: fadeIn 0.3s ease;
}

.transition-content {
  text-align: center;
  color: white;
  
  p { font-size: 1.5rem; color: #E2E8F0; margin-bottom: 0.5rem; }
  .next-title { font-size: 1.1rem; color: #F79633; font-weight: 700; }
}

.pulse-loader {
  width: 60px;
  height: 60px;
  border: 4px solid rgba(0, 255, 135, 0.1);
  border-left-color: #00FF87;
  border-radius: 50%;
  margin: 0 auto 2rem auto;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>
