<template>
  <nav class="navbar">
    <div class="logo">
      <span class="gold-text">RS100</span>
      <span class="divider">|</span>
      <span class="sub-text font-historical">100 Tahun Aloei Saboe</span>
    </div>
    <div class="nav-links">
      <router-link to="/" class="nav-link" active-class="active">Beranda</router-link>
      <router-link to="/history" class="nav-link" active-class="active">Sejarah & E-Book</router-link>
      <router-link to="/launching" class="nav-link" active-class="active">VVIP Launching</router-link>
    </div>
    <div class="status-badge" :class="{ 'activated': store.isActivated }">
      <span class="status-dot"></span>
      <span class="status-label">{{ store.isActivated ? 'Diresmikan' : 'Belum Aktif' }}</span>
    </div>
  </nav>
</template>

<script setup>
import { useActivationStore } from '../../store/activation'
const store = useActivationStore()
</script>

<style lang="scss" scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem 3rem;
  background: var(--color-nav-bg, var(--color-bg-dark, #03140A)); // Independent custom header background
  backdrop-filter: blur(15px);
  -webkit-backdrop-filter: blur(15px);
  border-bottom: 1px solid rgba(0, 168, 89, 0.1); 
  position: sticky;
  top: 0;
  z-index: 1000;
  transition: background 0.3s ease, border-bottom 0.3s ease;
}

.logo {
  display: flex;
  align-items: center;
  font-weight: 700;
  
  .gold-text {
    font-size: 1.5rem;
    font-family: 'Playfair Display', serif;
    background: linear-gradient(135deg, #FFC385 0%, #F79633 50%, #C16400 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }
  
  .divider {
    color: var(--color-primary, #00FF87);
    opacity: 0.3;
    margin: 0 1rem;
    font-size: 1.2rem;
  }

  .sub-text {
    font-size: 1rem;
    color: var(--color-nav-text, var(--color-white, #FFFFFF)); // Independent navbar text support
    letter-spacing: 1px;
    font-weight: 700;
    opacity: 0.9;
  }
}

.nav-links {
  display: flex;
  gap: 2.5rem;
}

.nav-link {
  color: var(--color-nav-text, var(--color-white, #FFFFFF)); // Independent custom navbar text color
  opacity: 0.75;
  text-decoration: none;
  font-weight: 700;
  font-size: 0.95rem;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: all 0.3s ease;
  position: relative;
  padding-bottom: 0.25rem;

  &::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    width: 0;
    height: 1px;
    background: var(--color-primary, #00FF87);
    transition: width 0.3s ease;
  }

  &:hover {
    color: var(--color-primary, #00FF87);
    opacity: 1;
    &::after {
      width: 100%;
    }
  }

  &.active {
    color: var(--color-nav-text, var(--color-white, #00FF87)); // Prevent clashing with primary/background colors
    opacity: 1;
    font-weight: 800;
    &::after {
      width: 100%;
      background: linear-gradient(90deg, var(--color-primary, #00FF87), var(--color-gold, #F79633));
      height: 2px;
    }
  }
}

.status-badge {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 0.5rem 1.2rem;
  border-radius: 50px;
  background: rgba(239, 68, 68, 0.08);
  border: 1px solid rgba(239, 68, 68, 0.25);
  color: #F87171;
  font-size: 0.8rem;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 1.2px;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
  box-shadow: 0 2px 8px rgba(239, 68, 68, 0.05);

  .status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: #EF4444;
    box-shadow: 0 0 10px rgba(239, 68, 68, 0.5);
    animation: pulse-red 2s infinite;
  }

  &.activated {
    background: var(--color-status-active-bg, rgba(16, 185, 129, 0.1)); 
    border: 1px solid var(--color-status-active-border, var(--color-primary, #00FF87));
    color: var(--color-status-active-text, var(--color-primary, #00FF87));
    box-shadow: 0 4px 15px var(--color-status-active-shadow, rgba(0, 255, 135, 0.15));

    .status-dot {
      background: var(--color-primary, #00FF87);
      box-shadow: 0 0 12px var(--color-primary, #00FF87);
      animation: pulse-green 2s infinite;
    }
  }
}

@keyframes pulse-red {
  0% { box-shadow: 0 0 0 0 rgba(239, 68, 68, 0.5); }
  70% { box-shadow: 0 0 0 6px rgba(239, 68, 68, 0); }
  100% { box-shadow: 0 0 0 0 rgba(239, 68, 68, 0); }
}

@keyframes pulse-green {
  0% { 
    box-shadow: 0 0 0 0 rgba(0, 255, 135, 0.6); 
    transform: scale(1);
  }
  50% {
    transform: scale(1.15);
  }
  70% { 
    box-shadow: 0 0 0 8px rgba(0, 255, 135, 0); 
    transform: scale(1);
  }
  100% { 
    box-shadow: 0 0 0 0 rgba(0, 255, 135, 0); 
    transform: scale(1);
  }
}
</style>
