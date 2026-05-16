<template>
  <div class="launching-page">
    <div class="interactive-container">
      
      <!-- Phase 1: Cinematic ECG Launcher -->
      <div v-if="!store.isActivated" class="launcher-panel">
        <div v-if="!isLaunching" class="instructions animate-fade">
          <h3 class="gold-subtitle">SISTEM SIAP DILUNCURKAN</h3>
          <p>Yth. Bapak Wali Kota Gorontalo,</p>
          <p class="touch-call">Silakan sentuh ikon jantung di bawah untuk memulai peresmian</p>
          
          <div v-if="errorMessage" class="error-msg" style="max-width: 500px; margin: 2rem auto 0;">
            ⚠️ {{ errorMessage }}
          </div>
        </div>

        <div class="sensor-container">
          <button v-if="!isLaunching" @click="startCinematicLaunch" class="heart-sensor">
            <div class="heart-icon-wrap">
              <svg viewBox="0 0 32 32" class="beating-heart">
                <!-- Corrected Heart Path: Converges to a single bottom point (16, 30) -->
                <path d="M16,30 C16,30 3,21 3,11 C3,6 7,2 12,2 C15,2 17,4 16,7 C15,4 17,2 20,2 C25,2 29,6 29,11 C29,21 16,30 16,30 Z" fill="url(#heart-grad)" />
                <defs>
                  <radialGradient id="heart-grad" cx="50%" cy="30%" r="50%">
                    <stop offset="0%" stop-color="#FFD1A9" />
                    <stop offset="50%" stop-color="#F79633" />
                    <stop offset="100%" stop-color="#00A859" />
                  </radialGradient>
                </defs>
              </svg>
            </div>
            <span class="heart-label">SENTUH JANTUNG</span>
            <div class="pulse-ring r1"></div>
            <div class="pulse-ring r2"></div>
            <div class="pulse-ring r3"></div>
          </button>

          <div v-else class="cinematic-canvas">
            <svg class="ecg-svg" viewBox="0 0 800 400">
              <defs>
                <filter id="gold-glow" x="-30%" y="-30%" width="160%" height="160%">
                  <feGaussianBlur stdDeviation="14" result="blur" />
                  <feMerge>
                    <feMergeNode in="blur" />
                    <feMergeNode in="SourceGraphic" />
                  </feMerge>
                </filter>
                <filter id="green-glow" x="-30%" y="-30%" width="160%" height="160%">
                  <feGaussianBlur stdDeviation="18" result="blur" />
                  <feMerge>
                    <feMergeNode in="blur" />
                    <feMergeNode in="SourceGraphic" />
                  </feMerge>
                </filter>
              </defs>
              <path class="path-element path-ecg" d="M 50,200 L 120,200 L 135,110 L 150,290 L 165,170 L 180,200 L 250,200" filter="url(#green-glow)" />
              <path class="path-element path-one" d="M 285,150 L 310,110 L 310,270" filter="url(#gold-glow)" />
              <path class="path-element path-zero-one" d="M 335,190 C 335,120 385,120 385,190 C 385,260 335,260 335,190 Z" filter="url(#gold-glow)" />
              <path class="path-element path-zero-two" d="M 410,190 C 410,120 460,120 460,190 C 460,260 410,260 410,190 Z" filter="url(#gold-glow)" />
              <path class="path-element path-building" d="M 485,200 L 515,200 L 515,160 L 545,160 L 545,110 L 580,110 L 580,70 L 615,70 L 615,110 L 650,110 L 650,160 L 680,160 L 680,200 L 750,200" filter="url(#green-glow)" />
              <path class="path-element path-cross" d="M 590,90 L 606,90 M 598,82 L 598,98" filter="url(#green-glow)" />
            </svg>
            <div class="cinematic-footer">
              <div class="loader-track">
                <span class="loader-fill" :style="{ width: progressPercentage + '%' }"></span>
              </div>
              <p class="status-msg animate-pulse">{{ activeStatusMessage }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Phase 2 & 3: Celebration & Reflection -->
      <div v-else class="celebration-panel">
        
        <!-- 🚀 Launching Components Dashboard (Fase 2) -->
        <transition name="slide-fade" mode="out-in">
          <div v-if="isActivated && !hasSigned" class="launch-dashboard-panel animate-fade">
            <div class="dashboard-card">
              <div class="dashboard-header">
                <h2 class="font-historical">PERESMIAN</h2>
                <p class="subtitle">SATU ABAD RSUD PROF. DR. H. ALOEI SABOE</p>
                
              </div>

              <div class="components-grid">
                <div class="component-item">
                  <div class="comp-icon">🏛️</div>
                  <div class="comp-info">
                    <h3>MUSEUM</h3>
                    <p>Prof.dr. H. Aloei Saboe</p>
                  </div>
                </div>
                <div class="component-item active-comp" @click="openBookPDF" style="cursor: pointer;">
                  <div class="comp-icon">
                    <img src="/cover-buku.png" alt="Cover Buku 100 Tahun" class="comp-book-img" />
                  </div>
                  <div class="comp-info">
                    <h3>BUKU</h3>
                    <p>100 Tahun RSAS</p>
                  </div>
                </div>
                <div class="component-item active-comp" @click="nextLogo" style="cursor: pointer;">
                  <div class="comp-icon">
                    <img :src="logos[activeLogoIndex]" alt="Logo RSAS 100 Tahun" class="comp-logo-img" />
                  </div>
                  <div class="comp-info">
                    <h3>LOGO</h3>
                    <p>Branding Satu Abad</p>
                  </div>
                </div>
                <div class="component-item">
                  <div class="comp-icon">🏥</div>
                  <div class="comp-info">
                    <h3>GEDUNG</h3>
                    <p>Nama Tokoh RS</p>
                  </div>
                </div>
                <div class="component-item">
                  <div class="comp-icon">💊</div>
                  <div class="comp-info">
                    <h3>INOVASI</h3>
                    <p>Safety Stock Obat</p>
                  </div>
                </div>
              </div>

              <div class="activation-hub">
                <div class="activation-trigger" @click="finalizeSignature">
                  <div class="pulse-ring"></div>
                  <div class="pulse-ring-outer"></div>
                  <div class="activation-icon">🖐️</div>
                  <span class="activation-text">SENTUH LAYAR UNTUK MERESMIKAN</span>
                </div>
              </div>
            </div>
          </div>
        </transition>

        <!-- 📖 Centennial Reflection Narrative (Fase 3) -->
        <transition name="slide-fade" mode="out-in">
          <div v-if="hasSigned" class="reflection-panel-card animate-fade">
            <video ref="reflectionVideo" class="reflection-video-bg" src="/Video-Narasi.mp4" playsinline loop muted></video>
            <video ref="reflectionAudio" src="/Video-Narasi.mp4" playsinline style="display: none;" @ended="onReflectionAudioEnded"></video>

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
        </transition>

        <div class="plaque-footer">
          <span class="badge">EST. 1926</span>
          <button @click="resetEverything" class="reset-btn">Reset</button>
        </div>
      </div>
    </div>

    <!-- Global Mute Toggle Button -->
    <button @click="toggleGlobalMute" class="global-audio-toggle" :title="isGlobalMuted ? 'Buka Suara' : 'Senyap'">
      <span v-if="isGlobalMuted" class="icon">🔇</span>
      <span v-else class="icon">🔊</span>
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed, watch } from 'vue'
import { useActivationStore } from '../store/activation'

// =========================================================
// NATIVE WEB AUDIO API SYNTHESIZER (ZERO-FILE FALLBACK)
// =========================================================
class NativeSynthesizer {
  constructor() {
    this.ctx = null
    this.riserOsc = null
    this.heartbeatInterval = null
  }

  init() {
    if (!this.ctx) {
      this.ctx = new (window.AudioContext || window.webkitAudioContext)()
    }
    if (this.ctx.state === 'suspended') {
      this.ctx.resume()
    }
  }

  playHeartbeat(theme) {
    this.init()
    this.stopHeartbeat()
    if (theme === 'mute') return

    const playThud = () => {
      const now = this.ctx.currentTime
      if (theme === 'clinical') {
        this.beep(now, 1000, 0.08)
      } else if (theme === 'orchestral') {
        this.drone(now, 110, 1.3)
      } else {
        this.thud(now, 75, 0.28)
        this.thud(now + 0.16, 60, 0.22)
      }
    }

    playThud()
    const intervalMs = theme === 'clinical' ? 1000 : theme === 'orchestral' ? 1400 : 1300
    this.heartbeatInterval = setInterval(playThud, intervalMs)
  }

  stopHeartbeat() {
    if (this.heartbeatInterval) {
      clearInterval(this.heartbeatInterval)
      this.heartbeatInterval = null
    }
  }

  playWelcomingIntro(theme) {
    this.init()
    if (theme === 'mute') return
    const now = this.ctx.currentTime

    if (theme === 'clinical') {
      const notes = [392.00, 523.25, 659.25, 783.99]
      notes.forEach((freq, index) => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'sine'
        osc.frequency.setValueAtTime(freq, now + index * 0.22)
        gain.gain.setValueAtTime(0.08, now + index * 0.22)
        gain.gain.exponentialRampToValueAtTime(0.001, now + index * 0.22 + 1.8)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now + index * 0.22)
        osc.stop(now + index * 0.22 + 1.8)
      })
    } else if (theme === 'orchestral') {
      const freqs = [196.00, 246.94, 293.66, 392.00, 493.88, 587.33]
      freqs.forEach((freq, index) => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'triangle'
        osc.frequency.setValueAtTime(freq, now)
        gain.gain.setValueAtTime(0.01, now)
        gain.gain.linearRampToValueAtTime(0.08, now + 1.0)
        gain.gain.exponentialRampToValueAtTime(0.001, now + 3.0)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now)
        osc.stop(now + 3.0)
      })
    } else {
      const notes = [523.25, 659.25, 783.99, 1046.50]
      notes.forEach((freq, index) => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'triangle'
        osc.frequency.setValueAtTime(freq, now + index * 0.12)
        gain.gain.setValueAtTime(0.06, now + index * 0.12)
        gain.gain.exponentialRampToValueAtTime(0.001, now + index * 0.12 + 1.5)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now + index * 0.12)
        osc.stop(now + index * 0.12 + 1.5)
      })
    }
  }

  beep(time, freq, duration) {
    const osc = this.ctx.createOscillator()
    const gain = this.ctx.createGain()
    osc.type = 'sine'
    osc.frequency.setValueAtTime(freq, time)
    gain.gain.setValueAtTime(0.06, time)
    gain.gain.exponentialRampToValueAtTime(0.001, time + duration)
    osc.connect(gain)
    gain.connect(this.ctx.destination)
    osc.start(time)
    osc.stop(time + duration)
  }

  drone(time, freq, duration) {
    const osc = this.ctx.createOscillator()
    const gain = this.ctx.createGain()
    osc.type = 'triangle'
    osc.frequency.setValueAtTime(freq, time)
    gain.gain.setValueAtTime(0.05, time)
    gain.gain.linearRampToValueAtTime(0.05, time + duration - 0.2)
    gain.gain.exponentialRampToValueAtTime(0.001, time + duration)
    osc.connect(gain)
    gain.connect(this.ctx.destination)
    osc.start(time)
    osc.stop(time + duration)
  }

  thud(time, freq, duration) {
    const osc = this.ctx.createOscillator()
    const gain = this.ctx.createGain()
    osc.type = 'sine'
    osc.frequency.setValueAtTime(freq, time)
    osc.frequency.exponentialRampToValueAtTime(0.01, time + duration)
    gain.gain.setValueAtTime(0.22, time)
    gain.gain.exponentialRampToValueAtTime(0.01, time + duration)
    osc.connect(gain)
    gain.connect(this.ctx.destination)
    osc.start(time)
    osc.stop(time + duration)
  }

  playLaunchStart(theme) {
    this.init()
    if (theme === 'mute') return
    const now = this.ctx.currentTime
    const duration = 0.9
    const osc = this.ctx.createOscillator()
    const gain = this.ctx.createGain()
    const filter = this.ctx.createBiquadFilter()
    if (theme === 'clinical') {
      osc.type = 'sine'
      osc.frequency.setValueAtTime(250, now)
      osc.frequency.exponentialRampToValueAtTime(500, now + duration)
      gain.gain.setValueAtTime(0.08, now)
    } else if (theme === 'orchestral') {
      osc.type = 'triangle'
      osc.frequency.setValueAtTime(55, now)
      osc.frequency.exponentialRampToValueAtTime(200, now + duration)
      gain.gain.setValueAtTime(0.18, now)
    } else {
      osc.type = 'sawtooth'
      osc.frequency.setValueAtTime(45, now)
      osc.frequency.exponentialRampToValueAtTime(320, now + duration)
      gain.gain.setValueAtTime(0.18, now)
    }
    filter.type = 'lowpass'
    filter.frequency.setValueAtTime(140, now)
    filter.frequency.exponentialRampToValueAtTime(550, now + duration)
    gain.gain.exponentialRampToValueAtTime(0.01, now + duration)
    osc.connect(filter)
    filter.connect(gain)
    gain.connect(this.ctx.destination)
    osc.start(now)
    osc.stop(now + duration)
  }

  playRiser(theme, duration) {
    this.init()
    if (theme === 'mute') return
    const now = this.ctx.currentTime
    const osc = this.ctx.createOscillator()
    const gain = this.ctx.createGain()
    if (theme === 'clinical') {
      osc.type = 'sine'
      osc.frequency.setValueAtTime(440, now)
      osc.frequency.linearRampToValueAtTime(880, now + duration)
      gain.gain.setValueAtTime(0.02, now)
      gain.gain.linearRampToValueAtTime(0.08, now + duration - 0.5)
    } else if (theme === 'orchestral') {
      osc.type = 'triangle'
      osc.frequency.setValueAtTime(80, now)
      osc.frequency.linearRampToValueAtTime(320, now + duration)
      gain.gain.setValueAtTime(0.03, now)
      gain.gain.linearRampToValueAtTime(0.14, now + duration - 0.5)
    } else {
      osc.type = 'sine'
      osc.frequency.setValueAtTime(110, now)
      osc.frequency.linearRampToValueAtTime(500, now + duration)
      gain.gain.setValueAtTime(0.03, now)
      gain.gain.linearRampToValueAtTime(0.18, now + duration - 0.5)
    }
    gain.gain.exponentialRampToValueAtTime(0.001, now + duration)
    osc.connect(gain)
    gain.connect(this.ctx.destination)
    osc.start(now)
    osc.stop(now + duration)
    this.riserOsc = osc
  }

  stopRiser() {
    if (this.riserOsc) {
      try { this.riserOsc.stop() } catch (e) {}
      this.riserOsc = null
    }
  }

  playSuccessChime(theme) {
    this.init()
    if (theme === 'mute') return
    const now = this.ctx.currentTime
    if (theme === 'clinical') {
      const notes = [392.00, 523.25, 659.25, 783.99]
      notes.forEach((freq, index) => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'sine'
        osc.frequency.setValueAtTime(freq, now + index * 0.25)
        gain.gain.setValueAtTime(0.12, now + index * 0.25)
        gain.gain.exponentialRampToValueAtTime(0.001, now + index * 0.25 + 2.0)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now + index * 0.25)
        osc.stop(now + index * 0.25 + 2.0)
      })
    } else if (theme === 'orchestral') {
      const brassFreqs = [130.81, 196.00, 261.63, 329.63, 392.00, 523.25]
      brassFreqs.forEach(freq => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'triangle'
        osc.frequency.setValueAtTime(freq, now)
        gain.gain.setValueAtTime(0.02, now)
        gain.gain.linearRampToValueAtTime(0.16, now + 0.4)
        gain.gain.exponentialRampToValueAtTime(0.001, now + 4.0)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now)
        osc.stop(now + 4.0)
      })
    } else {
      const cyberNotes = [261.63, 392.00, 523.25, 783.99, 1046.50, 1567.98]
      cyberNotes.forEach((freq, index) => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'triangle'
        osc.frequency.setValueAtTime(freq, now + index * 0.08)
        gain.gain.setValueAtTime(0.12, now + index * 0.08)
        gain.gain.exponentialRampToValueAtTime(0.001, now + index * 0.08 + 3.0)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now + index * 0.08)
        osc.stop(now + index * 0.08 + 3.0)
      })
    }
  }

  playSignatureSuccess(theme) {
    this.init()
    if (theme === 'mute') return
    const now = this.ctx.currentTime
    this.thud(now, 75, 0.3)
    this.thud(now + 0.15, 60, 0.25)
    const notes = [523.25, 659.25, 783.99, 1046.50]
    notes.forEach((freq, index) => {
      const osc = this.ctx.createOscillator()
      const gain = this.ctx.createGain()
      osc.type = 'sine'
      osc.frequency.setValueAtTime(freq, now + 0.3 + index * 0.1)
      gain.gain.setValueAtTime(0.08, now + 0.3 + index * 0.1)
      gain.gain.exponentialRampToValueAtTime(0.001, now + 0.3 + index * 0.1 + 1.5)
      osc.connect(gain)
      gain.connect(this.ctx.destination)
      osc.start(now + 0.3 + index * 0.1)
      osc.stop(now + 0.3 + index * 0.1 + 1.5)
    })
  }

  playClosingChime(theme) {
    this.init()
    if (theme === 'mute') return
    const now = this.ctx.currentTime
    if (theme === 'clinical') {
      const notes = [392.00, 261.63]
      notes.forEach((freq, index) => {
        const osc = this.ctx.createOscillator()
        const gain = this.ctx.createGain()
        osc.type = 'sine'
        osc.frequency.setValueAtTime(freq, now + index * 0.3)
        gain.gain.setValueAtTime(0.08, now + index * 0.3)
        gain.gain.exponentialRampToValueAtTime(0.001, now + index * 0.3 + 1.8)
        osc.connect(gain)
        gain.connect(this.ctx.destination)
        osc.start(now + index * 0.3)
        osc.stop(now + index * 0.3 + 1.8)
      })
    } else {
      const sweep = this.ctx.createOscillator()
      const gain = this.ctx.createGain()
      sweep.type = 'sine'
      sweep.frequency.setValueAtTime(600, now)
      sweep.frequency.exponentialRampToValueAtTime(90, now + 1.2)
      gain.gain.setValueAtTime(0.05, now)
      gain.gain.exponentialRampToValueAtTime(0.001, now + 1.2)
      sweep.connect(gain)
      gain.connect(this.ctx.destination)
      sweep.start(now)
      sweep.stop(now + 1.2)
    }
  }

  startReflectionPad() {
    this.init()
    this.stopReflectionPad()
    const now = this.ctx.currentTime
    const freqs = [130.81, 164.81, 196.00, 246.94]
    this.padOscs = []
    this.padGains = []
    freqs.forEach(freq => {
      const osc = this.ctx.createOscillator()
      const gain = this.ctx.createGain()
      osc.type = 'triangle'
      osc.frequency.setValueAtTime(freq, now)
      gain.gain.setValueAtTime(0.001, now)
      gain.gain.linearRampToValueAtTime(0.015, now + 2.0)
      osc.connect(gain)
      gain.connect(this.ctx.destination)
      osc.start(now)
      this.padOscs.push(osc)
      this.padGains.push(gain)
    })
  }

  stopReflectionPad() {
    if (this.padOscs && this.padOscs.length > 0) {
      const now = this.ctx ? this.ctx.currentTime : 0
      this.padOscs.forEach((osc, index) => {
        const gain = this.padGains[index]
        if (gain && this.ctx) {
          try {
            gain.gain.setValueAtTime(gain.gain.value, now)
            gain.gain.exponentialRampToValueAtTime(0.001, now + 1.2)
            setTimeout(() => { try { osc.stop() } catch (e) {} }, 1300)
          } catch (e) { try { osc.stop() } catch (e) {} }
        } else { try { osc.stop() } catch (e) {} }
      })
      this.padOscs = []
      this.padGains = []
    }
  }
}

const launchAudio = ref(null)
const synth = new NativeSynthesizer()
window.__stopHeartbeat = () => synth.stopHeartbeat()
const audioManager = {
  playWelcomingIntro() {
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.standby
    if (s === 'mute') return
    synth.playWelcomingIntro(s)
  },
  startHeartbeat() {
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.standby
    if (s === 'mute') return
    synth.playHeartbeat(s)
  },
  stopHeartbeat() { synth.stopHeartbeat() },
  playLaunchWhoosh() {
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.loading
    if (s === 'mute') return
    synth.playLaunchStart(s)
  },
  playRiser(durationSec) {
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.loading
    if (s === 'mute') return
    synth.playRiser(s, durationSec)
  },
  stopRiser() { synth.stopRiser() },
  playSuccess() {
    this.stopRiser()
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.success
    if (s === 'mute') return
    synth.playSuccessChime(s)
  },
  playSignatureSuccess() {
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.signature
    if (s === 'mute') return
    synth.playSignatureSuccess(s)
  },
  playClosing() {
    if (isGlobalMuted.value) return
    const s = store.sessionSoundConsole.loading
    if (s === 'mute') return
    synth.playClosingChime(s)
  },
  playReflectionBGM() {
    if (isGlobalMuted.value) return
    
    // Set initial volume to 0 for fade-in effect
    if (reflectionVideo.value) reflectionVideo.value.volume = 0
    if (reflectionAudio.value) reflectionAudio.value.volume = 0
    
    if (reflectionVideo.value) reflectionVideo.value.play().catch(() => {})
    if (narratorVisual.value) narratorVisual.value.play().catch(() => {})
    if (reflectionAudio.value) {
      reflectionAudio.value.play().catch(() => synth.startReflectionPad())
    } else {
      synth.startReflectionPad()
    }
    
    // Smoothly fade in volume over 2 seconds
    let fadeVolume = 0
    const fadeInterval = setInterval(() => {
      if (fadeVolume < globalVolume.value) {
        fadeVolume += 0.05
        if (fadeVolume > globalVolume.value) fadeVolume = globalVolume.value
        if (reflectionVideo.value) reflectionVideo.value.volume = fadeVolume
        if (reflectionAudio.value) reflectionAudio.value.volume = fadeVolume
      } else {
        clearInterval(fadeInterval)
      }
    }, 100)
  },
  stopReflectionBGM() {
    if (reflectionVideo.value) reflectionVideo.value.pause()
    if (narratorVisual.value) narratorVisual.value.pause()
    if (reflectionAudio.value) {
      reflectionAudio.value.pause()
      reflectionAudio.value.currentTime = 0
    }
    synth.stopReflectionPad()
  },
  stopAllActiveSounds() {
    this.stopHeartbeat()
    this.stopRiser()
    this.stopReflectionBGM()
    window.speechSynthesis.cancel()
    if (launchAudio.value) {
      launchAudio.value.pause()
      launchAudio.value.currentTime = 0
    }
  }
}

const store = useActivationStore()

// Watch for real-time changes in the panel for standby sound
watch(() => store.sessionSoundConsole.standby, (newVal) => {
  audioManager.stopHeartbeat()
  if (newVal !== 'mute' && !isGlobalMuted.value) {
    audioManager.startHeartbeat()
  }
})
const isLaunching = ref(false)
const progressPercentage = ref(0)

const logos = [
  '/LogoRSAS JPG/Logo_RSAS_0.jpg',
  '/LogoRSAS JPG/Logo_RSAS_1.jpg',
  '/LogoRSAS JPG/Logo_RSAS_2.jpg',
  '/LogoRSAS JPG/Logo_RSAS_3.jpg',
  '/LogoRSAS JPG/Logo_RSAS_0_bw.jpg',
  '/LogoRSAS JPG/Logo_RSAS_1_bw.jpg',
  '/LogoRSAS JPG/Logo_RSAS_2_bw.jpg',
  '/LogoRSAS JPG/Logo_RSAS_3_bw.jpg'
]
const activeLogoIndex = ref(0)
const nextLogo = () => {
  activeLogoIndex.value = (activeLogoIndex.value + 1) % logos.length
}
const activeStatusMessage = ref('')
const errorMessage = ref('')
const hasSigned = ref(false)
const narratorVisual = ref(null)
const reflectionVideo = ref(null)
const reflectionAudio = ref(null)
const isGlobalMuted = ref(false)
const isActivated = computed(() => store.isActivated)

const narratorPortal = ref(null)
const isNarratorExpanded = ref(false)
const isDragging = ref(false)
const narratorPosition = ref({ x: 0, y: 0 }) 
const dragStart = ref({ x: 0, y: 0 })

const narratorStyle = computed(() => ({
  transform: `translate(${narratorPosition.value.x}px, ${narratorPosition.value.y}px)`,
  cursor: isDragging.value ? 'grabbing' : 'grab'
}))

const startDrag = (e) => {
  const clientX = e.type === 'touchstart' ? e.touches[0].clientX : e.clientX
  const clientY = e.type === 'touchstart' ? e.touches[0].clientY : e.clientY
  dragStart.value = { x: clientX - narratorPosition.value.x, y: clientY - narratorPosition.value.y }
  isDragging.value = true
  window.addEventListener('mousemove', onDrag)
  window.addEventListener('touchmove', onDrag)
  window.addEventListener('mouseup', stopDrag)
  window.addEventListener('touchend', stopDrag)
}

const onDrag = (e) => {
  if (!isDragging.value) return
  const clientX = e.type === 'touchmove' ? e.touches[0].clientX : e.clientX
  const clientY = e.type === 'touchmove' ? e.touches[0].clientY : e.clientY
  narratorPosition.value = { x: clientX - dragStart.value.x, y: clientY - dragStart.value.y }
}

const stopDrag = () => {
  isDragging.value = false
  window.removeEventListener('mousemove', onDrag)
  window.removeEventListener('touchmove', onDrag)
  window.removeEventListener('mouseup', stopDrag)
  window.removeEventListener('touchend', stopDrag)
}

const toggleNarratorSize = () => {
  if (Math.abs(narratorPosition.value.x) < 5 && Math.abs(narratorPosition.value.y) < 5) {
     isNarratorExpanded.value = !isNarratorExpanded.value
  }
}

const showSyncSettings = ref(false)
const defaultTimestamps = [0, 39, 82] // User set Slide 1 to 39s
const reflectionTimestamps = ref([])

onMounted(() => {
  const saved = JSON.parse(localStorage.getItem('rs100_reflection_ts'))
  reflectionTimestamps.value = (saved && saved.length === reflections.length) ? saved : [...defaultTimestamps]

  // Try to start heartbeat on load
  audioManager.startHeartbeat()

  // Add click listener to bypass autoplay block
  const enableAudio = () => {
    audioManager.startHeartbeat()
    document.removeEventListener('click', enableAudio)
  }
  document.addEventListener('click', enableAudio)
})

const saveTimestamps = () => {
  reflectionTimestamps.value.sort((a, b) => a - b)
  localStorage.setItem('rs100_reflection_ts', JSON.stringify(reflectionTimestamps.value))
}

const reflections = [
  "100 Tahun Pemanfaatan Rumah Sakit. Satu abad perjalanan ini bukan sekadar hitungan waktu, tetapi menjadi bukti nyata pengabdian, pelayanan, perjuangan, dan dedikasi rumah sakit dalam memberikan pelayanan kesehatan kepada masyarakat.",
  "Selama seratus tahun, rumah sakit telah menjadi tempat lahirnya harapan, tempat bertumbuhnya kehidupan, serta tempat masyarakat memperoleh pelayanan, perhatian, dan kepedulian. Berbagai tantangan telah dilalui, mulai dari keterbatasan sarana di masa awal berdiri, perkembangan ilmu pengetahuan dan teknologi kesehatan, hingga menghadapi berbagai dinamika pelayanan kesehatan modern. Semua itu dijalani dengan semangat kemanusiaan dan komitmen untuk terus melayani.",
  "Kiranya peringatan 100 Tahun Pemanfaatan Rumah Sakit ini menjadi sumber motivasi bagi seluruh jajaran untuk terus berkarya, mengabdi, dan memberikan pelayanan terbaik bagi masyarakat. Semoga rumah sakit ini semakin maju, humanis, profesional, dan tetap menjadi tempat yang menghadirkan harapan serta kehidupan bagi semua."
]

const currentReflectionIndex = ref(0)
const isReflectionPlaying = ref(false)
const reflectionPlaybackTime = ref(0)

const formatTime = (seconds) => {
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  const ms = Math.floor((seconds % 1) * 10)
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}.${ms}`
}

const globalVolume = ref(0.6)

const updateVolume = () => {
  if (reflectionAudio.value) reflectionAudio.value.volume = globalVolume.value
  if (reflectionVideo.value) reflectionVideo.value.volume = globalVolume.value
}

const speakNarration = (text) => {
  window.speechSynthesis.cancel()
  const utterance = new SpeechSynthesisUtterance(text)
  utterance.lang = 'id-ID'
  utterance.rate = 0.9
  utterance.volume = globalVolume.value // Apply global volume to speech synthesis
  utterance.onend = () => {
    if (isReflectionPlaying.value && currentReflectionIndex.value < reflections.length - 1) {
      currentReflectionIndex.value++
      speakNarration(reflections[currentReflectionIndex.value])
    }
  }
  window.speechSynthesis.speak(utterance)
}

const startReflectionAutoPlay = () => {
  isReflectionPlaying.value = true
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
  } else {
    speakNarration(reflections[currentReflectionIndex.value])
  }
}

const stopAndResetReflection = () => {
  isReflectionPlaying.value = false
  window.speechSynthesis.cancel()
  currentReflectionIndex.value = 0
  if (reflectionAudio.value) {
    reflectionAudio.value.pause()
    reflectionAudio.value.currentTime = 0
    reflectionPlaybackTime.value = 0
  }
}

const toggleReflectionAutoPlay = () => isReflectionPlaying.value ? pauseReflectionAutoPlay() : startReflectionAutoPlay()
const pauseReflectionAutoPlay = () => {
  isReflectionPlaying.value = false
  if (reflectionAudio.value) reflectionAudio.value.pause()
}

const openBookPDF = () => {
  window.open('/BUKU RSAS 100 TAHUN.pdf', '_blank')
}

const statusMessages = ["Menerima denyut nadi VVIP...", "Menyalurkan energi EKG emas...", "Membentuk angka monumental 100...", "Memetakan siluet RSUD...", "Mengaktifkan Portal Digital..."]

const startCinematicLaunch = () => {
  if (isLaunching.value) return
  isLaunching.value = true
  progressPercentage.value = 0
  
  // Stop heartbeat sound when launch starts!
  audioManager.stopHeartbeat()
  
  const s = store.sessionSoundConsole.loading
  if (s === 'playback_6') {
    if (launchAudio.value) {
      launchAudio.value.pause()
    }
    launchAudio.value = new Audio('/playback_6.mp3')
    launchAudio.value.volume = globalVolume.value
    launchAudio.value.play().catch(e => console.error('Failed to play launch audio:', e))
  } else if (s !== 'mute') {
    audioManager.playLaunchWhoosh()
    audioManager.playRiser(6.0)
  }
  
  let steps = 60
  let currentStep = 0
  const progressInterval = setInterval(() => {
    currentStep++
    progressPercentage.value = (currentStep / steps) * 100
    activeStatusMessage.value = statusMessages[Math.floor((currentStep / steps) * statusMessages.length)]
    if (currentStep >= steps) {
      clearInterval(progressInterval)
      finalizeActivation()
    }
  }, 100)
}

const finalizeActivation = async () => {
  const result = await store.activateApp('ACT-RS100')
  if (result.success) audioManager.playSuccess()
  else { isLaunching.value = false; errorMessage.value = result.error }
}

const finalizeSignature = () => {
  hasSigned.value = true
  audioManager.playSignatureSuccess()
  setTimeout(() => {
    audioManager.playReflectionBGM()
    startReflectionAutoPlay()
  }, 1200)
}

const resetEverything = () => {
  audioManager.stopAllActiveSounds()
  store.resetActivation()
  isLaunching.value = false
  hasSigned.value = false
  stopAndResetReflection()
}

const toggleGlobalMute = () => {
  isGlobalMuted.value = !isGlobalMuted.value
  isGlobalMuted.value ? audioManager.stopAllActiveSounds() : audioManager.startHeartbeat()
}

onUnmounted(() => audioManager.stopAllActiveSounds())
</script>

<style lang="scss" scoped>
.launching-page {
  min-height: 90vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: radial-gradient(circle at center, #0B301D 0%, #03140A 100%);
  padding: 2rem;
  overflow: hidden;
  font-family: 'Inter', sans-serif;
}

.interactive-container { 
  width: 100%; 
  max-width: 1100px; 
  perspective: 1000px;
}

.launcher-panel { text-align: center; }

.gold-subtitle {
  color: #F79633;
  letter-spacing: 4px;
  font-weight: 800;
  margin-bottom: 1rem;
  text-shadow: 0 0 15px rgba(247, 150, 51, 0.4);
}

.heart-sensor {
  background: none; border: none; cursor: pointer; position: relative;
  width: 320px; height: 320px;
  display: flex; flex-direction: column; justify-content: center; align-items: center;
  transition: transform 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  margin: 0 auto;

  &:hover {
    transform: scale(1.1);
    .heart-icon-wrap { filter: drop-shadow(0 0 40px rgba(247, 150, 51, 0.7)); }
  }

  .heart-icon-wrap {
    width: 180px; height: 180px; z-index: 10;
    transition: filter 0.3s ease;
    filter: drop-shadow(0 0 30px rgba(247, 150, 51, 0.5));
  }

  .beating-heart { 
    width: 100%; height: 100%; 
    animation: heartbeat-vibe 1.2s infinite cubic-bezier(0.215, 0.610, 0.355, 1);
  }

  .heart-label {
    margin-top: 2rem; color: #F79633; font-weight: 800;
    letter-spacing: 4px; font-size: 0.9rem; text-transform: uppercase;
  }

  .pulse-ring {
    position: absolute; width: 160px; height: 160px;
    border: 2px solid rgba(0, 255, 135, 0.3); border-radius: 50%;
    &.r1 { animation: sonar 3s infinite 0s; }
    &.r2 { animation: sonar 3s infinite 1s; }
    &.r3 { animation: sonar 3s infinite 2s; }
  }
}

.cinematic-canvas {
  background: rgba(3, 20, 10, 0.8); backdrop-filter: blur(10px);
  border: 1px solid rgba(247, 150, 51, 0.2); border-radius: 24px;
  padding: 3rem; box-shadow: 0 25px 50px rgba(0,0,0,0.5);
  animation: fade-in-up 0.8s ease;

  .ecg-svg { width: 100%; height: auto; }
  .path-element {
    fill: none; stroke-linecap: round; stroke-linejoin: round;
    stroke-dasharray: 1000; stroke-dashoffset: 1000; stroke-width: 6;
    animation: draw-step 5s forwards ease-in-out;
  }
  .path-ecg, .path-building { stroke: #00FF87; }
  .path-one, .path-zero-one, .path-zero-two { stroke: #F79633; }
}

.launch-dashboard-panel {
  animation: fade-in-up 1s cubic-bezier(0.16, 1, 0.3, 1);
  perspective: 1500px;
  
  .dashboard-card {
    background: rgba(3, 20, 10, 0.7);
    backdrop-filter: blur(15px) saturate(150%);
    border-radius: 40px;
    padding: 4rem;
    border: 1px solid rgba(0, 255, 135, 0.25);
    box-shadow: 
      0 40px 100px rgba(0,0,0,0.7),
      inset 0 0 40px rgba(0, 255, 135, 0.03); /* Subtle Inner Glow */
    position: relative;
    overflow: hidden;

    &::after {
      content: '';
      position: absolute;
      top: 0; left: 0; right: 0; height: 1px;
      background: linear-gradient(90deg, transparent, rgba(247, 150, 51, 0.5), transparent);
    }
  }

  .dashboard-header {
    margin-bottom: 4rem;
    
    h2 {
      font-size: 3rem;
      font-weight: 900;
      letter-spacing: 8px;
      margin-bottom: 0.8rem;
      background: linear-gradient(135deg, #FFD1A9 0%, #F79633 50%, #B87333 100%);
      -webkit-background-clip: text;
      background-clip: text;
      -webkit-text-fill-color: transparent;
      text-transform: uppercase;
      filter: drop-shadow(0 2px 4px rgba(0,0,0,0.3));
    }

    .subtitle {
      color: rgba(226, 232, 240, 0.7); /* Subtle Contrast */
      font-weight: 600;
      letter-spacing: 3px;
      font-size: 0.95rem;
    }
  }

  .components-grid {
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    gap: 2rem;
    margin: 4rem 0;
  }

  .component-item {
    background: rgba(255, 255, 255, 0.02);
    border: 1px solid rgba(255, 255, 255, 0.08);
    padding: 1.5rem 1rem; /* Reduced padding to let image grow */
    border-radius: 24px;
    transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
    cursor: default;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    min-height: 220px;

    .comp-icon { 
      width: 100%;
      height: 120px; 
      margin-bottom: 1rem; 
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 4.5rem; /* Membuat emoji jauh lebih besar */
      filter: drop-shadow(0 0 15px rgba(0,0,0,0.3));
      transition: transform 0.4s ease;
      line-height: 1;

      .comp-logo-img {
        width: auto;
        height: 100%;
        max-width: 100%;
        object-fit: contain;
        border-radius: 12px;
        filter: brightness(1.1) drop-shadow(0 0 12px rgba(0, 255, 137, 0.3));
      }

      .comp-book-img {
        width: auto;
        height: 100%;
        max-width: 90%;
        object-fit: contain;
        border-radius: 6px;
        box-shadow: 8px 8px 20px rgba(0,0,0,0.6);
        transform: rotate(-3deg);
        transition: all 0.3s ease;
      }
    }

    h3 { 
      font-size: 0.9rem; 
      color: #F79633; 
      font-weight: 800; 
      margin-bottom: 0.6rem;
      letter-spacing: 1px;
    }

    p { 
      font-size: 0.75rem; 
      color: rgba(148, 163, 184, 0.8); 
      line-height: 1.5;
    }

    &:hover {
      background: rgba(0, 255, 135, 0.08);
      border-color: rgba(0, 255, 135, 0.4);
      transform: translateY(-12px) scale(1.02);
      box-shadow: 0 20px 40px rgba(0,0,0,0.4);
      .comp-icon { transform: scale(1.2) rotate(5deg); }
      .comp-book-img { transform: rotate(0deg) scale(1.1); }
    }

    /* Active State Glow - seolah teraliri listrik */
    &.active-comp {
      border-color: #00FF87;
      background: rgba(0, 255, 135, 0.12);
      box-shadow: 
        0 0 30px rgba(0, 255, 135, 0.3),
        inset 0 0 20px rgba(0, 255, 135, 0.1);
      h3 { color: #00FF87; text-shadow: 0 0 10px rgba(0, 255, 135, 0.5); }
    }
  }
}

.activation-trigger {
  width: 200px; 
  height: 200px;
  border-radius: 50%;
  border: 4px solid #00FF87;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  cursor: pointer;
  position: relative;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  background: radial-gradient(circle at center, rgba(0, 255, 135, 0.2) 0%, transparent 70%);
  box-shadow: 0 0 40px rgba(0, 255, 135, 0.2);
  animation: breathing-glow 3s infinite ease-in-out;

  &:hover {
    transform: scale(1.1);
    background: radial-gradient(circle at center, rgba(0, 255, 135, 0.4) 0%, transparent 70%);
    box-shadow: 0 0 60px rgba(0, 255, 135, 0.5);
  }

  .activation-icon { 
    font-size: 4.5rem; 
    filter: drop-shadow(0 0 20px rgba(0, 255, 135, 0.6));
    animation: hand-sway 2s infinite ease-in-out;
  }

  .activation-text {
    position: absolute;
    bottom: -50px;
    font-size: 0.85rem;
    color: #00FF87;
    font-weight: 800;
    letter-spacing: 3px;
    white-space: nowrap;
    text-shadow: 0 0 10px rgba(0, 255, 135, 0.4);
  }
}

@keyframes breathing-glow {
  0%, 100% { transform: scale(1); box-shadow: 0 0 30px rgba(0, 255, 135, 0.2); }
  50% { transform: scale(1.03); box-shadow: 0 0 60px rgba(0, 255, 135, 0.5); }
}

@keyframes hand-sway {
  0%, 100% { transform: rotate(0deg); }
  50% { transform: rotate(-10deg) translateY(-5px); }
}

.reflection-panel-card {
  position: relative; overflow: hidden;
  background: rgba(3, 20, 10, 0.9); backdrop-filter: blur(25px);
  border: 1px solid rgba(247, 150, 51, 0.3); border-radius: 24px;
  padding: 2.5rem; text-align: left; box-shadow: 0 30px 60px rgba(0,0,0,0.5);
  animation: fade-in-up 0.8s ease;

  .reflection-video-bg {
    position: absolute; top: 0; left: 0; width: 100%; height: 100%;
    object-fit: cover; opacity: 0.15; z-index: 0; pointer-events: none;
    filter: grayscale(0.6) contrast(1.2);
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

  .reflection-body { position: relative; z-index: 5; min-height: 150px; padding-right: 120px; }

  .reflection-text {
    font-size: 1.15rem; line-height: 1.9; color: #E2E8F0; margin: 0;
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
    h5 { color: #F79633; margin: 0; font-size: 0.8rem; }
  }

  .live-timer {
    display: flex; align-items: center; gap: 8px; background: #000;
    padding: 4px 12px; border-radius: 6px; border: 1px solid #22C55E;
    .timer-label { font-size: 0.6rem; color: #94A3B8; }
    .timer-value { font-family: monospace; font-size: 1rem; color: #22C55E; font-weight: 800; }
  }

  .settings-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 1rem; }

  .settings-item {
    display: flex; flex-direction: column; gap: 6px;
    label { font-size: 0.7rem; color: #94A3B8; font-weight: 700; }
    input {
      background: #111; border: 1px solid #F79633; color: #FFF;
      padding: 8px; border-radius: 8px; font-weight: 800; text-align: center;
    }
  }
}

.narrator-portal {
  position: absolute; top: 1.5rem; right: 1.5rem; width: 110px; height: 110px;
  border-radius: 50%; border: 3px solid #F79633; overflow: hidden;
  box-shadow: 0 0 20px rgba(0,0,0,0.4); z-index: 10;
  transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  &.expanded { width: 320px; height: 180px; border-radius: 16px; }
  video { width: 100%; height: 100%; object-fit: cover; }
}

.plaque-footer {
  margin-top: 3rem; display: flex; justify-content: center; gap: 2rem; align-items: center;
  .badge { color: #94A3B8; font-weight: 800; letter-spacing: 2px; font-size: 0.8rem; }
}

.reset-btn {
  background: rgba(255, 255, 255, 0.05); border: 1px solid rgba(255, 255, 255, 0.1);
  color: #94A3B8; padding: 0.5rem 1.5rem; border-radius: 20px; cursor: pointer;
  &:hover { background: rgba(239, 68, 68, 0.1); color: #f87171; border-color: #ef4444; }
}

@keyframes heartbeat-vibe {
  0% { transform: scale(1); }
  15% { transform: scale(1.12); }
  30% { transform: scale(1); }
  45% { transform: scale(1.15); }
  100% { transform: scale(1); }
}

@keyframes sonar {
  0% { transform: scale(0.8); opacity: 0; }
  50% { opacity: 0.6; }
  100% { transform: scale(2.2); opacity: 0; }
}

@keyframes fade-in-up {
  0% { opacity: 0; transform: translateY(40px); }
  100% { opacity: 1; transform: translateY(0); }
}

@keyframes draw-step { to { stroke-dashoffset: 0; } }

.global-audio-toggle {
  position: fixed; top: 6.5rem; right: 2rem;
  background: rgba(11, 48, 29, 0.6); backdrop-filter: blur(10px);
  border: 1px solid rgba(247, 150, 51, 0.4); color: white;
  width: 56px; height: 56px; border-radius: 50%; cursor: pointer; z-index: 1000;
  transition: all 0.3s ease;
  &:hover { transform: scale(1.1); background: rgba(247, 150, 51, 0.2); }
}
</style>
