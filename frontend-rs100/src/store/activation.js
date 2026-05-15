import { defineStore } from 'pinia'
import { apiService } from '../api/service'

export const useActivationStore = defineStore('activation', {
  state: () => ({
    isActivated: localStorage.getItem('rs100_activated') === 'true',
    isActivating: false,
    activationError: null,
    activationSuccessMsg: '',
    books: [],
    activeTheme: localStorage.getItem('rs100_theme') || 'light-mint',
    customColors: {
      navbarBg: localStorage.getItem('rs100_custom_navbarBg') || '#00A859',
      navbarText: localStorage.getItem('rs100_custom_navbarText') || '#FFFFFF',
      bodyBg: localStorage.getItem('rs100_custom_bodyBg') || '#E6F5EC',
      bodyText: localStorage.getItem('rs100_custom_bodyText') || '#1E293B',
      primary: localStorage.getItem('rs100_custom_primary') || '#00A859',
      cardBg: localStorage.getItem('rs100_custom_cardBg') || '#FFFFFF',
      cardText: localStorage.getItem('rs100_custom_cardText') || '#1E293B',
      shieldBg: localStorage.getItem('rs100_custom_shieldBg') || '#0B301D',
      shieldText: localStorage.getItem('rs100_custom_shieldText') || '#F79633'
    },
    brightnessModifier: localStorage.getItem('rs100_brightness') || 'normal',
    activeSoundTheme: localStorage.getItem('rs100_sound_theme') || 'futuristic',
    activeVoiceCharacter: localStorage.getItem('rs100_voice_character') || 'female-warm',
    sessionSoundConsole: {
      standby: localStorage.getItem('rs100_session_standby') || 'clinical',
      loading: localStorage.getItem('rs100_session_loading') || 'futuristic',
      success: localStorage.getItem('rs100_session_success') || 'orchestral',
      signature: localStorage.getItem('rs100_session_signature') || 'futuristic',
      narrator: localStorage.getItem('rs100_session_narrator') || 'female-warm'
    }
  }),
  actions: {
    setTheme(themeName) {
      this.activeTheme = themeName
      localStorage.setItem('rs100_theme', themeName)
      this.applyThemeToBody()
    },
    setSoundTheme(soundThemeName) {
      this.activeSoundTheme = soundThemeName
      localStorage.setItem('rs100_sound_theme', soundThemeName)
      
      // Auto-synchronise the granular Session Sound Console for maximum operator convenience!
      if (soundThemeName === 'mute') {
        this.sessionSoundConsole = {
          standby: 'mute',
          loading: 'mute',
          success: 'mute',
          signature: 'mute',
          narrator: 'mute'
        }
      } else {
        this.sessionSoundConsole = {
          standby: soundThemeName === 'clinical' ? 'clinical' : (soundThemeName === 'orchestral' ? 'clinical' : 'futuristic'),
          loading: soundThemeName,
          success: soundThemeName,
          signature: soundThemeName === 'orchestral' ? 'orchestral' : 'futuristic',
          narrator: this.activeVoiceCharacter || 'female-warm'
        }
      }

      // Persist the synchronised sessions to localStorage
      Object.keys(this.sessionSoundConsole).forEach(key => {
        localStorage.setItem(`rs100_session_${key}`, this.sessionSoundConsole[key])
      })
    },
    setVoiceCharacter(characterName) {
      this.activeVoiceCharacter = characterName
      localStorage.setItem('rs100_voice_character', characterName)
      
      // Also synchronise with the narrator session setting!
      this.sessionSoundConsole.narrator = characterName
      localStorage.setItem('rs100_session_narrator', characterName)
    },
    setSessionSound(sessionKey, soundThemeName) {
      this.sessionSoundConsole[sessionKey] = soundThemeName
      localStorage.setItem(`rs100_session_${sessionKey}`, soundThemeName)
    },
    resetSessionSoundConsole() {
      this.sessionSoundConsole = {
        standby: 'clinical',
        loading: 'futuristic',
        success: 'orchestral',
        signature: 'futuristic',
        narrator: 'female-warm'
      }
      localStorage.setItem('rs100_session_standby', 'clinical')
      localStorage.setItem('rs100_session_loading', 'futuristic')
      localStorage.setItem('rs100_session_success', 'orchestral')
      localStorage.setItem('rs100_session_signature', 'futuristic')
      localStorage.setItem('rs100_session_narrator', 'female-warm')
    },
    setBrightnessModifier(mode) {
      this.brightnessModifier = mode
      localStorage.setItem('rs100_brightness', mode)
      this.applyThemeToBody()
    },
    updateCustomColor(key, value) {
      this.customColors[key] = value
      localStorage.setItem(`rs100_custom_${key}`, value)
      if (this.activeTheme === 'custom') {
        this.applyThemeToBody()
      }
    },
    applyThemeToBody() {
      const root = document.documentElement
      
      // Handle page-level cinematic brightness modifications!
      root.classList.remove('theme-modifier-darker', 'theme-modifier-lighter')
      if (this.brightnessModifier === 'darker') {
        root.classList.add('theme-modifier-darker')
      } else if (this.brightnessModifier === 'lighter') {
        root.classList.add('theme-modifier-lighter')
      }

      root.classList.remove(
        'theme-light-mint', 
        'theme-dark-obsidian', 
        'theme-champagne-gold', 
        'theme-centennial-gold', 
        'theme-emerald-royal', 
        'theme-midnight-velvet', 
        'theme-teal-platinum', 
        'theme-custom'
      )
      
      if (this.activeTheme === 'custom') {
        root.classList.add('theme-custom')
        root.style.setProperty('--color-nav-bg', this.customColors.navbarBg)
        root.style.setProperty('--color-nav-text', this.customColors.navbarText)
        root.style.setProperty('--color-bg-dark', this.customColors.bodyBg)
        root.style.setProperty('--color-white', this.customColors.bodyText)
        root.style.setProperty('--color-primary', this.customColors.primary)
        root.style.setProperty('--color-success', this.customColors.primary)
        root.style.setProperty('--color-card-bg', this.customColors.cardBg)
        root.style.setProperty('--color-card-text', this.customColors.cardText)
        root.style.setProperty('--color-shield-bg', this.customColors.shieldBg)
        root.style.setProperty('--color-shield-text', this.customColors.shieldText)
        // Dynamic status badge translucent colors calculated from primary custom color!
        root.style.setProperty('--color-status-active-bg', `${this.customColors.primary}14`) // 8% opacity
        root.style.setProperty('--color-status-active-shadow', `${this.customColors.primary}26`) // 15% opacity
        // Set text to navbarText (guaranteed contrast) and border to primary (main brand accent)
        root.style.setProperty('--color-status-active-text', this.customColors.navbarText)
        root.style.setProperty('--color-status-active-border', this.customColors.primary)
      } else {
        root.style.removeProperty('--color-nav-bg')
        root.style.removeProperty('--color-nav-text')
        root.style.removeProperty('--color-bg-dark')
        root.style.removeProperty('--color-white')
        root.style.removeProperty('--color-primary')
        root.style.removeProperty('--color-success')
        root.style.removeProperty('--color-card-bg')
        root.style.removeProperty('--color-card-text')
        root.style.removeProperty('--color-shield-bg')
        root.style.removeProperty('--color-shield-text')
        root.style.removeProperty('--color-status-active-bg')
        root.style.removeProperty('--color-status-active-shadow')
        root.style.removeProperty('--color-status-active-text')
        root.style.removeProperty('--color-status-active-border')
        root.classList.add(`theme-${this.activeTheme}`)
      }
    },
    async activateApp(code) {
      this.isActivating = true
      this.activationError = null
      this.activationSuccessMsg = ''
      try {
        const data = await apiService.activate(code)
        this.isActivated = true
        this.activationSuccessMsg = data.message || 'Satu Abad RSUD Aloei Saboe Berhasil Diresmikan!'
        localStorage.setItem('rs100_activated', 'true')
        return { success: true }
      } catch (error) {
        this.activationError = error.response?.data?.error || 'Gagal melakukan peresmian. Periksa Kode Otorisasi Anda.'
        return { success: false, error: this.activationError }
      } finally {
        this.isActivating = false
      }
    },
    async fetchBooks() {
      try {
        const data = await apiService.getBooks()
        this.books = data
      } catch (error) {
        console.error('Gagal mengambil data buku:', error)
      }
    },
    resetActivation() {
      this.isActivated = false
      this.activationSuccessMsg = ''
      localStorage.removeItem('rs100_activated')
    }
  }
})
