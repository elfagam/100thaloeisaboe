<template>
  <div class="history-page">
    <div class="header-section text-center">
      <h1 class="font-historical text-gold-gradient">Satu Abad Perjalanan</h1>
      <p class="subtitle">Menelusuri Jejak Sejarah Pelayanan Kesehatan RSUD Prof. Dr. H. Aloei Saboe</p>
    </div>

    <!-- Timeline Section -->
    <section class="timeline-section">
      <h2 class="section-title font-historical text-gold-gradient">⏳ Milestones Satu Abad</h2>
      <div class="timeline-container">
        <div 
          v-for="(milestone, index) in milestones" 
          :key="index" 
          class="timeline-item"
          :class="{ 'left': index % 2 === 0, 'right': index % 2 !== 0 }"
        >
          <div class="timeline-dot"></div>
          <div class="timeline-card premium-card">
            <span class="timeline-year text-gold">{{ milestone.year }}</span>
            <h3 class="timeline-title font-historical">{{ milestone.title }}</h3>
            <p class="timeline-desc">{{ milestone.desc }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const milestones = ref([
  {
    year: "1926",
    title: "Zevenhoven Hospital Didirikan",
    desc: "Peletakan batu pertama oleh Pemerintah Hindia Belanda sebagai fasilitas kesehatan modern perdana di Gorontalo."
  },
  {
    year: "1945",
    title: "Nasionalisasi Rumah Sakit",
    desc: "Rebut alih manajemen rumah sakit dari sekutu menjadi milik pemerintah Republik Indonesia pasca Proklamasi."
  },
  {
    year: "1987",
    title: "Diresmikan Sebagai RSUD Aloei Saboe",
    desc: "Pemberian nama resmi dr. Aloei Saboe sebagai bentuk penghormatan abadi atas jasa pahlawan medis nasional."
  },
  {
    year: "2002",
    title: "Relokasi Kampus Utama Baru",
    desc: "Pemindahan gedung pelayanan ke lokasi baru di Jl. Prof. Dr. H. Aloei Saboe guna perluasan daya tampung rujukan."
  },
  {
    year: "2026",
    title: "Satu Abad & Transformasi Digital",
    desc: "Implementasi 100% EMR Compliance, sistem audit medis otomatis, dan peraihan predikat Akreditasi Paripurna Bintang Lima."
  }
])

onMounted(async () => {
  if (window.__stopHeartbeat) {
    window.__stopHeartbeat()
  }

  // Fetch Milestones Data
  try {
    const response = await fetch('http://localhost:8081/api/milestones')
    const data = await response.json()
    if (data && data.length > 0) {
      milestones.value = data.map(m => ({
        year: m.year,
        title: m.title,
        desc: m.description
      }))
    }
  } catch (error) {
    console.error('Failed to fetch milestones data:', error)
  }

  // Scroll Entrance Reveal Animation Observer
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible')
      }
    })
  }, { threshold: 0.1 })

  setTimeout(() => {
    document.querySelectorAll('.timeline-item').forEach(item => {
      observer.observe(item)
    })
  }, 100)
})
</script>

<style lang="scss" scoped>
.history-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 4rem 1.5rem;
  position: relative;
  overflow: hidden;
  background: radial-gradient(circle at top right, rgba(11, 48, 29, 0.4) 0%, rgba(3, 20, 10, 0.8) 100%);
  border-radius: 20px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(247, 150, 51, 0.1);
  margin-top: 2rem;
  margin-bottom: 4rem;
}

.text-center {
  text-align: center;
}

.header-section {
  margin-bottom: 4rem;
  position: relative;
  z-index: 2;

  h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
  }

  .subtitle {
    color: #94A3B8;
    font-size: 1.1rem;
    max-width: 600px;
    margin: 0 auto;
  }
}

.timeline-section {
  border-top: 1px solid rgba(247, 150, 51, 0.15);
  padding-top: 5rem;
  margin-top: 3rem;
  text-align: center;

  .section-title {
    font-size: 2rem;
    margin-bottom: 4rem;
  }
}

.timeline-container {
  position: relative;
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem 0;

  &::before {
    content: '';
    position: absolute;
    width: 2px;
    background: linear-gradient(180deg, rgba(247,150,51,0.1) 0%, rgba(247,150,51,0.8) 50%, rgba(247,150,51,0.1) 100%);
    top: 0;
    bottom: 0;
    left: 50%;
    margin-left: -1px;

    @media (max-width: 768px) {
      left: 30px;
    }
  }
}

.timeline-item {
  padding: 1rem 2rem;
  position: relative;
  width: 50%;
  text-align: left;

  opacity: 0;
  transition: opacity 0.8s cubic-bezier(0.16, 1, 0.3, 1), transform 0.8s cubic-bezier(0.16, 1, 0.3, 1);

  @media (max-width: 768px) {
    width: 100%;
    padding-left: 70px;
    padding-right: 15px;
    
    &.left, &.right {
      left: 0 !important;
      transform: translateX(-30px);
    }
  }

  &.left {
    left: 0;
    transform: translateX(-40px);
  }

  &.right {
    left: 50%;
    transform: translateX(40px);
  }

  &.visible {
    opacity: 1;
    transform: translateX(0) !important;
  }

  .timeline-dot {
    width: 16px;
    height: 16px;
    border-radius: 50%;
    background: var(--color-primary, #F79633);
    border: 3px solid var(--color-bg-dark, #03140A);
    position: absolute;
    top: 2.2rem;
    left: 50%;
    margin-left: -8px;
    z-index: 5;
    box-shadow: 0 0 10px var(--color-primary, #F79633);
    transition: background 0.3s ease, border-color 0.3s ease;

    @media (max-width: 768px) {
      left: 30px;
      margin-left: -8px;
    }
  }

  .timeline-card {
    padding: 1.5rem;
    border-radius: 12px;
    transition: transform 0.4s cubic-bezier(0.34, 1.56, 0.64, 1), box-shadow 0.4s ease, border-color 0.3s ease, background 0.3s ease;
    cursor: pointer;
    border: 1px solid rgba(247, 150, 51, 0.1);

    &:hover {
      transform: translateY(-5px) scale(1.02);
      border-color: var(--color-primary, #F79633);
      box-shadow: 0 10px 30px var(--color-status-active-shadow, rgba(0, 255, 135, 0.15));
      background: var(--color-card-bg);
      
      .timeline-year {
        text-shadow: 0 0 8px var(--color-primary);
      }
    }

    .timeline-year {
      font-size: 1.4rem;
      font-weight: 900;
      display: block;
      margin-bottom: 0.5rem;
      letter-spacing: 1px;
      transition: text-shadow 0.3s ease;
    }

    .timeline-title {
      font-size: 1.2rem;
      color: var(--color-primary, #FFC385);
      margin-bottom: 0.75rem;
    }

    .timeline-desc {
      color: var(--color-card-text, #94A3B8);
      font-size: 0.9rem;
      line-height: 1.5;
    }
  }
}
</style>
