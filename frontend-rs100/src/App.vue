<template>
  <div class="app-wrapper">
    <!-- High-Performance hardware-accelerated brightness overlays (0% CPU/GPU overhead) -->
    <div class="global-brightness-overlay" :class="activationStore.brightnessModifier"></div>

    <NavBar />
    <main class="main-content">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>

    <!-- Floating Theme Switcher Widget -->
    <div class="theme-switcher-container">
      <button class="theme-floating-btn" @click.stop="toggleMenu" title="Pilih Tema Layar">
        <span class="palette-icon">🎨</span>
        <span class="btn-ripple"></span>
      </button>

      <transition name="pop-up">
        <div v-if="isMenuOpen" class="theme-menu-card glass-panel" @click.stop>
          <div class="menu-header">
            <h4 class="menu-title font-historical">STUDIO TEMA</h4>
            <span class="menu-subtitle">Sesuaikan Visual Seremoni</span>
          </div>

          <!-- Tab Navigation -->
          <div class="theme-tabs">
            <button 
              class="tab-btn" 
              :class="{ active: activeTab === 'presets' }" 
              @click="activeTab = 'presets'"
            >
              Preset
            </button>
            <button 
              class="tab-btn" 
              :class="{ active: activeTab === 'designer' }" 
              @click="switchToDesigner"
            >
              Desainer 🛠️
            </button>
          </div>
          
          <!-- Tab 1: Presets -->
          <div v-if="activeTab === 'presets'" class="theme-options">
            <button 
              class="theme-opt-btn light-mint-opt" 
              :class="{ active: activationStore.activeTheme === 'light-mint' }"
              @click="selectTheme('light-mint')"
            >
              <span class="color-preview preview-light-mint"></span>
              <span class="opt-label">Light Mint</span>
            </button>
            <button 
              class="theme-opt-btn dark-obsidian-opt" 
              :class="{ active: activationStore.activeTheme === 'dark-obsidian' }"
              @click="selectTheme('dark-obsidian')"
            >
              <span class="color-preview preview-dark-obsidian"></span>
              <span class="opt-label">Dark Obsidian</span>
            </button>
            <button 
              class="theme-opt-btn champagne-gold-opt" 
              :class="{ active: activationStore.activeTheme === 'champagne-gold' }"
              @click="selectTheme('champagne-gold')"
            >
              <span class="color-preview preview-champagne-gold"></span>
              <span class="opt-label">Royal Gold</span>
            </button>
            <button 
              class="theme-opt-btn centennial-gold-opt" 
              :class="{ active: activationStore.activeTheme === 'centennial-gold' }"
              @click="selectTheme('centennial-gold')"
            >
              <span class="color-preview preview-centennial-gold"></span>
              <span class="opt-label">Centennial Gold</span>
            </button>
            <button 
              class="theme-opt-btn emerald-royal-opt" 
              :class="{ active: activationStore.activeTheme === 'emerald-royal' }"
              @click="selectTheme('emerald-royal')"
            >
              <span class="color-preview preview-emerald-royal"></span>
              <span class="opt-label">Emerald Royal</span>
            </button>
            <button 
              class="theme-opt-btn midnight-velvet-opt" 
              :class="{ active: activationStore.activeTheme === 'midnight-velvet' }"
              @click="selectTheme('midnight-velvet')"
            >
              <span class="color-preview preview-midnight-velvet"></span>
              <span class="opt-label">Midnight Velvet</span>
            </button>
            <button 
              class="theme-opt-btn teal-platinum-opt" 
              :class="{ active: activationStore.activeTheme === 'teal-platinum' }"
              @click="selectTheme('teal-platinum')"
            >
              <span class="color-preview preview-teal-platinum"></span>
              <span class="opt-label">Teal Platinum</span>
            </button>

            <!-- Dynamic Brightness Fine-Tuning controls -->
            <div class="brightness-divider"></div>
            <div class="brightness-section">
              <span class="brightness-label">Kecerahan Preset</span>
              <div class="brightness-btn-group">
                <button 
                  class="brightness-btn darker-btn" 
                  :class="{ active: activationStore.brightnessModifier === 'darker' }"
                  @click="activationStore.setBrightnessModifier('darker')"
                  title="Mode Bioskop Lebih Gelap"
                >
                  🌙 Gelap
                </button>
                <button 
                  class="brightness-btn normal-btn" 
                  :class="{ active: activationStore.brightnessModifier === 'normal' }"
                  @click="activationStore.setBrightnessModifier('normal')"
                  title="Reset Normal"
                >
                  ⚙️ Normal
                </button>
                <button 
                  class="brightness-btn lighter-btn" 
                  :class="{ active: activationStore.brightnessModifier === 'lighter' }"
                  @click="activationStore.setBrightnessModifier('lighter')"
                  title="Kontras Lebih Terang"
                >
                  ☀️ Terang
                </button>
              </div>
            </div>

            <!-- Dynamic Audio / Sound Effect controls -->
            <div class="brightness-divider"></div>
            <div class="audio-section">
              <span class="brightness-label">Efek Suara Seremoni</span>
              <div class="audio-btn-group">
                <button 
                  v-for="theme in soundThemes" 
                  :key="theme.id"
                  class="brightness-btn" 
                  :class="{ active: activationStore.activeSoundTheme === theme.id }"
                  @click="activationStore.setSoundTheme(theme.id)"
                  :title="theme.desc"
                >
                  {{ theme.icon }} {{ theme.name }}
                </button>
              </div>
            </div>

            <!-- Dynamic Voice Character controls -->
            <div class="brightness-divider"></div>
            <div class="audio-section">
              <span class="brightness-label">Karakter Asisten Suara</span>
              <div class="audio-btn-group">
                <button 
                  v-for="char in voiceCharacters" 
                  :key="char.id"
                  class="brightness-btn" 
                  :class="{ active: activationStore.activeVoiceCharacter === char.id }"
                  @click="activationStore.setVoiceCharacter(char.id)"
                  :title="char.desc"
                >
                  {{ char.icon }} {{ char.name }}
                </button>
              </div>
            </div>
            <!-- Advanced Collapsible Session Sound Mixer Console -->
            <div class="brightness-divider"></div>
            <div class="advanced-mixer">
              <button class="mixer-toggle-btn" @click="isMixerOpen = !isMixerOpen">
                <span>🎛️ Konsol Suara Per Sesi (Advanced)</span>
                <span>{{ isMixerOpen ? '▲' : '▼' }}</span>
              </button>
              
              <transition name="slide-fade">
                <div v-if="isMixerOpen" class="mixer-content animate-fade">
                  <div v-for="sess in sessionsList" :key="sess.key" class="mixer-row">
                    <span class="mixer-row-label">{{ sess.icon }} {{ sess.name }}</span>
                    <select 
                      :value="activationStore.sessionSoundConsole[sess.key]" 
                      @change="activationStore.setSessionSound(sess.key, $event.target.value)"
                      class="mixer-select"
                    >
                      <option v-for="opt in sess.options" :key="opt.id" :value="opt.id">
                        {{ opt.icon }} {{ opt.name }}
                      </option>
                    </select>
                  </div>
                  <button class="reset-mixer-btn" @click="activationStore.resetSessionSoundConsole">
                    🔄 Reset Setelan Mixer
                  </button>
                </div>
              </transition>
            </div>
          </div>

          <!-- Tab 2: Custom Designer -->
          <div v-else class="designer-options">
            <div class="custom-picker-row">
              <span class="picker-label">Latar Konten</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.bodyBg || '#E6F5EC'" 
                  @input="updateColor('bodyBg', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.bodyBg || '#E6F5EC').toUpperCase()" 
                  @input="updateColorFromText('bodyBg', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Latar Navbar</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.navbarBg || '#00A859'" 
                  @input="updateColor('navbarBg', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.navbarBg || '#00A859').toUpperCase()" 
                  @input="updateColorFromText('navbarBg', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Teks Konten</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.bodyText || '#1E293B'" 
                  @input="updateColor('bodyText', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.bodyText || '#1E293B').toUpperCase()" 
                  @input="updateColorFromText('bodyText', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Teks Navbar</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.navbarText || '#FFFFFF'" 
                  @input="updateColor('navbarText', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.navbarText || '#FFFFFF').toUpperCase()" 
                  @input="updateColorFromText('navbarText', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Warna Aksen</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.primary || '#00A859'" 
                  @input="updateColor('primary', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.primary || '#00A859').toUpperCase()" 
                  @input="updateColorFromText('primary', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Latar Kartu</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.cardBg || '#FFFFFF'" 
                  @input="updateColor('cardBg', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.cardBg || '#FFFFFF').toUpperCase()" 
                  @input="updateColorFromText('cardBg', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Teks Kartu</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.cardText || '#1E293B'" 
                  @input="updateColor('cardText', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.cardText || '#1E293B').toUpperCase()" 
                  @input="updateColorFromText('cardText', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Latar Perisai</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.shieldBg || '#0B301D'" 
                  @input="updateColor('shieldBg', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.shieldBg || '#0B301D').toUpperCase()" 
                  @input="updateColorFromText('shieldBg', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>

            <div class="custom-picker-row">
              <span class="picker-label">Teks Perisai</span>
              <div class="picker-wrapper">
                <input 
                  type="color" 
                  :value="activationStore.customColors.shieldText || '#F79633'" 
                  @input="updateColor('shieldText', $event.target.value)" 
                  class="color-input" 
                />
                <input 
                  type="text" 
                  :value="(activationStore.customColors.shieldText || '#F79633').toUpperCase()" 
                  @input="updateColorFromText('shieldText', $event.target.value)" 
                  class="hex-input" 
                  maxlength="7"
                />
              </div>
            </div>
            
            <button class="reset-designer-btn" @click="resetToDefaultCustom">
              Kembalikan Default
            </button>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import NavBar from './components/shared/NavBar.vue'
import { useActivationStore } from './store/activation'

const activationStore = useActivationStore()
const isMenuOpen = ref(false)
const activeTab = ref(activationStore.activeTheme === 'custom' ? 'designer' : 'presets')

const soundThemes = [
  { id: 'futuristic', name: 'Futuristik', icon: '⚡', desc: 'Sintesis frekuensi tinggi & cyber scans' },
  { id: 'orchestral', name: 'Orkestra', icon: '🎻', desc: 'Dengung orkestra tebal & Gong agung' },
  { id: 'clinical', name: 'Medis', icon: '🏥', desc: 'Monitor EKG rumah sakit asli & Chimes' },
  { id: 'mute', name: 'Hening', icon: '🔇', desc: 'Sistem peresmian berjalan tanpa suara' }
]

const voiceCharacters = [
  { id: 'female-warm', name: 'Ibu Hangat', icon: '👩', desc: 'Suara wanita lembut, teduh & menyentuh kalbu' },
  { id: 'male-protocol', name: 'Pak Protokol', icon: '🤵', desc: 'Suara pria gagah, berat & sangat khidmat' },
  { id: 'cyber-assistant', name: 'Asisten Siber', icon: '🤖', desc: 'Suara wanita cerdas, tangkas & presisi digital' }
]

const isMixerOpen = ref(false)

const sessionsList = [
  {
    key: 'standby',
    name: 'Sesi Standby',
    icon: '💓',
    options: [
      { id: 'clinical', name: 'Medis (Detak Jantung)', icon: '💓' },
      { id: 'futuristic', name: 'Futuristik (Drone)', icon: '🌌' },
      { id: 'mute', name: 'Hening', icon: '🔇' }
    ]
  },
  {
    key: 'loading',
    name: 'Sesi Riser',
    icon: '🚀',
    options: [
      { id: 'futuristic', name: 'Futuristik (Siber Whoosh)', icon: '⚡' },
      { id: 'orchestral', name: 'Orkestra (Dengung Biola)', icon: '🎻' },
      { id: 'clinical', name: 'Medis (Detak Cepat)', icon: '🏥' },
      { id: 'mute', name: 'Hening', icon: '🔇' }
    ]
  },
  {
    key: 'success',
    name: 'Sesi Fanfare',
    icon: '🎺',
    options: [
      { id: 'orchestral', name: 'Orkestra (Kuningan Brass)', icon: '🎺' },
      { id: 'clinical', name: 'Medis (Klimaks EKG & Bell)', icon: '🔔' },
      { id: 'futuristic', name: 'Futuristik (Siber Chime)', icon: '⚡' },
      { id: 'mute', name: 'Hening', icon: '🔇' }
    ]
  },
  {
    key: 'signature',
    name: 'Sesi Tanda Tangan',
    icon: '🖋️',
    options: [
      { id: 'futuristic', name: 'Tembakan Laser Energi', icon: '💥' },
      { id: 'orchestral', name: 'Lonceng Katedral + Gong', icon: '🔔' },
      { id: 'mute', name: 'Hening', icon: '🔇' }
    ]
  },
  {
    key: 'narrator',
    name: 'Karakter Narator',
    icon: '🎙️',
    options: [
      { id: 'female-warm', name: 'Ibu Hangat (Wanita)', icon: '👩' },
      { id: 'male-protocol', name: 'Pak Protokol (Pria)', icon: '🤵' },
      { id: 'cyber-assistant', name: 'Asisten Siber (Wanita)', icon: '🤖' },
      { id: 'mute', name: 'Hening', icon: '🔇' }
    ]
  }
]

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const selectTheme = (themeName) => {
  activationStore.setTheme(themeName)
  activeTab.value = 'presets'
}

const switchToDesigner = () => {
  activeTab.value = 'designer'
  activationStore.setTheme('custom')
}

const updateColor = (key, value) => {
  activationStore.updateCustomColor(key, value)
}

const updateColorFromText = (key, value) => {
  let formatted = value.trim()
  if (formatted.length > 0 && !formatted.startsWith('#')) {
    formatted = '#' + formatted
  }
  // Validate standard 7-character hex string (#RRGGBB)
  if (/^#[0-9A-Fa-f]{6}$/.test(formatted)) {
    activationStore.updateCustomColor(key, formatted)
  }
}

const resetToDefaultCustom = () => {
  activationStore.updateCustomColor('bodyBg', '#E6F5EC')
  activationStore.updateCustomColor('navbarBg', '#00A859')
  activationStore.updateCustomColor('bodyText', '#1E293B')
  activationStore.updateCustomColor('navbarText', '#FFFFFF')
  activationStore.updateCustomColor('primary', '#00A859')
  activationStore.updateCustomColor('cardBg', '#FFFFFF')
  activationStore.updateCustomColor('cardText', '#1E293B')
  activationStore.updateCustomColor('shieldBg', '#0B301D')
  activationStore.updateCustomColor('shieldText', '#F79633')
}

const closeMenu = (event) => {
  const container = document.querySelector('.theme-switcher-container')
  if (container && !container.contains(event.target)) {
    isMenuOpen.value = false
  }
}

onMounted(() => {
  activationStore.applyThemeToBody()
  document.addEventListener('click', closeMenu)
})

onUnmounted(() => {
  document.removeEventListener('click', closeMenu)
})
</script>

<style lang="scss">
.app-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.global-brightness-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  pointer-events: none; // Completely click-through
  z-index: 999999; // Stays on top of everything
  transition: background 0.4s cubic-bezier(0.16, 1, 0.3, 1);
  background: transparent;
  will-change: background;

  &.darker {
    background: rgba(0, 0, 0, 0.22); // Cinematic theater darkness (zero GPU calculation overhead!)
  }

  &.lighter {
    background: rgba(255, 255, 255, 0.12); // High ambient contrast luminance
  }
}

.main-content {
  flex-grow: 1;
}

// Page Transition
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

// ==========================================================================
// Floating Theme Switcher Styles
// ==========================================================================

.theme-switcher-container {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  z-index: 9999;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 1rem;
}

.theme-floating-btn {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: var(--color-primary, #00FF87);
  border: 2px solid var(--color-gold, #F79633);
  box-shadow: 0 10px 25px rgba(0, 255, 135, 0.4);
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);

  .palette-icon {
    font-size: 1.8rem;
    z-index: 2;
  }

  .btn-ripple {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: 50%;
    border: 2px solid var(--color-primary, #00FF87);
    animation: floating-ripple 2s infinite ease-out;
    opacity: 0.8;
  }

  &:hover {
    transform: scale(1.1) rotate(15deg);
    box-shadow: 0 15px 30px rgba(0, 255, 135, 0.6);
  }

  &:active {
    transform: scale(0.95);
  }
}

.theme-menu-card {
  width: 300px;
  background: rgba(255, 255, 255, 0.9);
  border: 1px solid rgba(0, 255, 135, 0.3);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  display: flex;
  flex-direction: column;
  gap: 1.1rem;
  transform-origin: bottom right;

  .menu-header {
    border-bottom: 1px solid rgba(0, 168, 89, 0.1);
    padding-bottom: 0.5rem;
    text-align: center;
  }

  .menu-title {
    font-size: 1.1rem;
    font-weight: 800;
    color: #054826;
    letter-spacing: 2px;
    margin: 0;
  }

  .menu-subtitle {
    font-size: 0.75rem;
    color: #64748B;
    font-weight: 600;
  }
}

// Tab Styles
.theme-tabs {
  display: flex;
  background: rgba(0, 0, 0, 0.05);
  border-radius: 8px;
  padding: 0.25rem;
  gap: 0.25rem;

  .tab-btn {
    flex: 1;
    border: none;
    background: transparent;
    padding: 0.5rem;
    font-size: 0.8rem;
    font-weight: 700;
    color: #475569;
    cursor: pointer;
    border-radius: 6px;
    transition: all 0.2s ease;

    &.active {
      background: #FFFFFF;
      color: #00A859;
      box-shadow: 0 2px 6px rgba(0,0,0,0.08);
    }
  }
}

// Theme Options & Designer Overrides
.theme-dark-obsidian {
  .theme-menu-card {
    background: rgba(3, 20, 10, 0.95);
    border-color: rgba(0, 255, 135, 0.4);
    box-shadow: 0 20px 40px rgba(0,0,0,0.5);

    .menu-title { color: #00FF87; }
    .menu-subtitle { color: #94A3B8; }
    .menu-header { border-color: rgba(0, 255, 135, 0.2); }
    .theme-tabs { 
      background: rgba(255,255,255,0.08); 
      .tab-btn {
        color: #94A3B8;
        &.active {
          background: rgba(0, 255, 135, 0.15);
          color: #00FF87;
        }
      }
    }

    .theme-options, .designer-options {
      &::-webkit-scrollbar-thumb {
        background: rgba(0, 255, 135, 0.2);
      }
      &:hover::-webkit-scrollbar-thumb {
        background: rgba(0, 255, 135, 0.5);
      }
    }
  }
}

.theme-champagne-gold {
  .theme-menu-card {
    background: rgba(250, 246, 238, 0.95);
    border-color: rgba(193, 100, 0, 0.4);
    box-shadow: 0 20px 40px rgba(193, 100, 0, 0.15);

    .menu-title { color: #C16400; }
    .menu-subtitle { color: #78350F; }
    .menu-header { border-color: rgba(193, 100, 0, 0.15); }
    .theme-tabs { 
      background: rgba(0,0,0,0.04); 
      .tab-btn {
        color: #78350F;
        &.active {
          background: #FFFFFF;
          color: #C16400;
        }
      }
    }

    .theme-options, .designer-options {
      &::-webkit-scrollbar-thumb {
        background: rgba(193, 100, 0, 0.2);
      }
      &:hover::-webkit-scrollbar-thumb {
        background: rgba(193, 100, 0, 0.5);
      }
    }
  }
}

.theme-custom {
  .theme-menu-card {
    background: var(--color-bg-dark, #E6F5EC);
    border-color: var(--color-primary, #00A859);
    box-shadow: 0 20px 40px rgba(0,0,0,0.15);

    .menu-title { color: var(--color-primary); }
    .menu-subtitle { color: var(--color-white); opacity: 0.8; }
    .menu-header { border-color: rgba(0,0,0,0.1); }
    .theme-tabs { 
      background: rgba(0,0,0,0.05); 
      .tab-btn {
        color: var(--color-white);
        opacity: 0.7;
        &.active {
          background: #FFFFFF;
          color: #00A859;
          opacity: 1;
        }
      }
    }
  }
}

.theme-options {
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
  max-height: calc(100vh - 16rem);
  overflow-y: auto;
  overflow-x: hidden;
  padding-right: 0.4rem;

  /* Custom Webkit Scrollbar */
  &::-webkit-scrollbar {
    width: 6px;
  }
  
  &::-webkit-scrollbar-track {
    background: transparent;
  }
  
  &::-webkit-scrollbar-thumb {
    background: rgba(0, 168, 89, 0.2);
    border-radius: 10px;
    transition: all 0.3s ease;
  }
  
  &:hover::-webkit-scrollbar-thumb {
    background: rgba(0, 168, 89, 0.4);
  }
}

.theme-opt-btn {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 0.7rem 1rem;
  border-radius: 12px;
  border: 1px solid rgba(0, 168, 89, 0.1);
  background: rgba(255, 255, 255, 0.5);
  cursor: pointer;
  transition: all 0.2s ease;
  font-family: 'Inter', sans-serif;
  font-weight: 700;
  color: #1E293B;

  &:hover {
    background: rgba(0, 168, 89, 0.05);
    transform: translateX(-3px);
  }

  &.active {
    border-color: var(--color-primary, #00FF87);
    border-width: 2px;
    background: rgba(0, 255, 135, 0.08);
  }
}

// Color option previews
.color-preview {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 1px solid rgba(0,0,0,0.1);
  flex-shrink: 0;
}

.preview-light-mint { background: linear-gradient(135deg, #E6F5EC 50%, #00A859 50%); }
.preview-dark-obsidian { background: linear-gradient(135deg, #03140A 50%, #00FF87 50%); }
.preview-champagne-gold { background: linear-gradient(135deg, #FAF6EE 50%, #C16400 50%); }
.preview-centennial-gold { background: linear-gradient(135deg, #0F172A 50%, #F59E0B 50%); }
.preview-emerald-royal { background: linear-gradient(135deg, #062C17 50%, #10B981 50%); }
.preview-midnight-velvet { background: linear-gradient(135deg, #090514 50%, #A855F7 50%); }
.preview-teal-platinum { background: linear-gradient(135deg, #F8FAFC 50%, #0EA5E9 50%); }

.opt-label {
  font-size: 0.85rem;
}

// Cinematic Brightness Tuning
.brightness-divider {
  height: 1px;
  background: rgba(0, 168, 89, 0.15);
  margin: 1.2rem 0 0.8rem;
}

.brightness-section {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;

  .brightness-label {
    font-size: 0.8rem;
    font-weight: 800;
    color: #475569;
    letter-spacing: 0.5px;
    text-transform: uppercase;
  }

  .brightness-btn-group {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 0.4rem;
  }

  .brightness-btn {
    padding: 0.5rem 0.25rem;
    font-size: 0.75rem;
    font-weight: 700;
    border-radius: 8px;
    border: 1px solid rgba(0, 168, 89, 0.1);
    background: rgba(255, 255, 255, 0.6);
    color: #1E293B;
    cursor: pointer;
    transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1);
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.2rem;
    font-family: 'Inter', sans-serif;

    &:hover {
      background: rgba(0, 168, 89, 0.05);
      border-color: rgba(0, 168, 89, 0.3);
    }

    &.active {
      background: #00A859;
      color: #FFFFFF;
      border-color: #00A859;
      box-shadow: 0 4px 12px rgba(0, 168, 89, 0.2);
    }
  }
}

.audio-section {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-top: 0.4rem;

  .audio-btn-group {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 0.4rem;
  }
}

.advanced-mixer {
  margin-top: 0.6rem;
  
  .mixer-toggle-btn {
    width: 100%;
    background: rgba(247, 150, 51, 0.05);
    border: 1px dashed rgba(247, 150, 51, 0.3);
    color: #FFC385;
    padding: 0.6rem 0.75rem;
    border-radius: 8px;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 0.8rem;
    font-weight: 700;
    letter-spacing: 1px;
    text-transform: uppercase;
    transition: all 0.3s ease;

    &:hover {
      background: rgba(247, 150, 51, 0.1);
      border-color: #F79633;
    }
  }

  .mixer-content {
    background: rgba(3, 20, 10, 0.9);
    border: 1px solid rgba(247, 150, 51, 0.2);
    border-radius: 8px;
    padding: 0.85rem;
    margin-top: 0.4rem;
    display: flex;
    flex-direction: column;
    gap: 0.6rem;
    box-shadow: inset 0 0 10px rgba(0,0,0,0.5);
  }

  .mixer-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 0.5rem;
  }

  .mixer-row-label {
    font-size: 0.75rem;
    color: #94A3B8;
    font-weight: 600;
  }

  .mixer-select {
    background: rgba(11, 48, 29, 0.8);
    border: 1px solid rgba(247, 150, 51, 0.25);
    color: #FFC385;
    border-radius: 4px;
    padding: 0.25rem 0.4rem;
    font-size: 0.75rem;
    outline: none;
    cursor: pointer;
    min-width: 135px;
    max-width: 155px;
    transition: all 0.3s ease;

    &:focus, &:hover {
      border-color: #F79633;
      background: rgba(11, 48, 29, 1);
    }
    
    option {
      background: #03140A;
      color: #FFC385;
    }
  }

  .reset-mixer-btn {
    margin-top: 0.25rem;
    background: transparent;
    border: 1px solid rgba(239, 68, 68, 0.25);
    color: #FCA5A5;
    font-size: 0.7rem;
    padding: 0.35rem;
    border-radius: 4px;
    cursor: pointer;
    font-weight: 600;
    transition: all 0.3s ease;

    &:hover {
      background: rgba(239, 68, 68, 0.1);
      border-color: #ef4444;
      color: #ffffff;
    }
  }
}

// Custom Designer Styles
.designer-options {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  padding: 0.25rem 0;
  max-height: calc(100vh - 16rem);
  overflow-y: auto;
  overflow-x: hidden;
  padding-right: 0.4rem;

  /* Custom Webkit Scrollbar */
  &::-webkit-scrollbar {
    width: 6px;
  }
  
  &::-webkit-scrollbar-track {
    background: transparent;
  }
  
  &::-webkit-scrollbar-thumb {
    background: rgba(0, 168, 89, 0.2);
    border-radius: 10px;
    transition: all 0.3s ease;
  }
  
  &:hover::-webkit-scrollbar-thumb {
    background: rgba(0, 168, 89, 0.4);
  }
}

.custom-picker-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;

  .picker-label {
    font-size: 0.8rem;
    font-weight: 700;
    color: var(--color-white, #1E293B);
    opacity: 0.9;
  }
}

.picker-wrapper {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(0,0,0,0.1);
  padding: 0.25rem 0.5rem;
  border-radius: 8px;
  box-shadow: inset 0 1px 3px rgba(0,0,0,0.05);

  .color-input {
    -webkit-appearance: none;
    appearance: none;
    border: none;
    width: 24px;
    height: 24px;
    border-radius: 50%;
    cursor: pointer;
    background: none;
    flex-shrink: 0;

    &::-webkit-color-swatch-wrapper {
      padding: 0;
    }
    &::-webkit-color-swatch {
      border: 1px solid rgba(0,0,0,0.15);
      border-radius: 50%;
      padding: 0;
    }
  }

  .hex-input {
    font-size: 0.75rem;
    font-family: monospace;
    font-weight: 700;
    color: #475569;
    background: transparent;
    border: none;
    width: 65px;
    padding: 0;
    text-transform: uppercase;
    outline: none;
    cursor: text;
  }
}

// Override wrappers for inputs inside dark designer card
.theme-dark-obsidian, .theme-custom {
  .picker-wrapper {
    background: rgba(0, 0, 0, 0.25);
    border-color: rgba(255,255,255,0.15);

    .hex-input {
      color: #CBD5E1;
    }
  }
}

.reset-designer-btn {
  margin-top: 0.5rem;
  background: transparent;
  border: 1px dashed var(--color-primary, #00A859);
  color: var(--color-primary, #00A859);
  padding: 0.5rem;
  border-radius: 8px;
  font-family: 'Inter', sans-serif;
  font-size: 0.75rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    background: rgba(0, 255, 135, 0.08);
    transform: translateY(-1px);
  }
}

// Switcher Animations
.pop-up-enter-active,
.pop-up-leave-active {
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

.pop-up-enter-from,
.pop-up-leave-to {
  opacity: 0;
  transform: scale(0.8) translateY(20px);
}

@keyframes floating-ripple {
  0% { transform: scale(1); opacity: 0.8; }
  100% { transform: scale(1.4); opacity: 0; }
}
</style>
