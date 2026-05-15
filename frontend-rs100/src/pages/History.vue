<template>
  <div class="history-page">
    <div class="header-section text-center">
      <h1 class="font-historical text-gold-gradient">Satu Abad Perjalanan</h1>
      <p class="subtitle">Menelusuri Jejak Sejarah Pelayanan Kesehatan RSUD Prof. Dr. H. Aloei Saboe</p>
    </div>

    <!-- Interactive E-Book Section -->
    <section class="ebook-section">
      <h2 class="section-title font-historical text-gold-gradient">📖 E-Book Satu Abad RSAS</h2>
      <p class="desc">Gunakan navigasi atau tab di bawah untuk melihat dokumen sejarah dan panduan logo baru resmi.</p>
      
      <!-- Book Type Toggle Tabs -->
      <div class="book-tabs">
        <button 
          @click="switchBook('sejarah')" 
          class="tab-btn" 
          :class="{ 'active': activeBook === 'sejarah' }"
        >
          📖 E-Book Sejarah
        </button>
        <button 
          @click="switchBook('logo')" 
          class="tab-btn" 
          :class="{ 'active': activeBook === 'logo' }"
        >
          🎨 E-Book Panduan Logo (PDF)
        </button>
      </div>

      <div class="book-container">
        <!-- The Virtual Book Component -->
        <div class="virtual-book">
          <!-- Left Spine Shadow -->
          <div class="spine-line"></div>

          <!-- Active Page Content -->
          <div class="book-page active-page" :class="{ 'is-pdf-book': activeBook === 'logo' }" :key="activeBook + '_' + currentPage">
            <div class="page-num">{{ currentPage + 1 }} / {{ maxPages }}</div>
            
            <div class="page-content" :class="{ 'pdf-page-container': activeBook === 'logo' }">
              <!-- BOOK 1: SEJARAH -->
              <template v-if="activeBook === 'sejarah'">
                <!-- Render current page component -->
                <div v-if="currentPage === 0" class="page-cover">
                  <div class="cover-gold-border">
                    <h3 class="cover-sub">MEMORANDUM SATU ABAD</h3>
                    <h2 class="cover-main font-historical text-gold-gradient">RSUD ALOEI SABOE</h2>
                    <p class="cover-years">1926 - 2026</p>
                    <div class="cover-emblem">🏆</div>
                    <p class="cover-footer">Bakti & Dedikasi Kesehatan Masyarakat Gorontalo</p>
                  </div>
                </div>

                <div v-else class="page-body">
                  <h3 class="page-title font-historical">{{ bookPages[currentPage].title }}</h3>
                  <div class="page-text" v-html="bookPages[currentPage].content"></div>
                  <div class="page-illustration" v-if="bookPages[currentPage].image">
                    <span class="illustration-icon">{{ bookPages[currentPage].image }}</span>
                  </div>
                </div>
              </template>

              <!-- BOOK 2: LOGO GUIDELINES (PDF RENDER) -->
              <template v-else>
                <div class="pdf-image-wrapper">
                  <img 
                    :src="logoPages[currentPage]" 
                    alt="Logo Guideline Page" 
                    class="pdf-page-image" 
                    loading="lazy"
                  />
                </div>
              </template>
            </div>
          </div>
        </div>

        <!-- Book Navigation Controls -->
        <div class="book-controls">
          <button @click="prevPage" :disabled="currentPage === 0" class="control-btn">&larr; Sebelumnya</button>
          <div class="page-indicator">Halaman {{ currentPage + 1 }}</div>
          <button @click="nextPage" :disabled="currentPage === maxPages - 1" class="control-btn">Selanjutnya &rarr;</button>
        </div>
      </div>
    </section>

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
import { ref, onMounted, computed } from 'vue'
import { useActivationStore } from '../store/activation'

const store = useActivationStore()
const currentPage = ref(0)
const activeBook = ref('sejarah') // 'sejarah' or 'logo'

const logoPages = [
  '/logo-guidelines/page_1.png',
  '/logo-guidelines/page_2.png',
  '/logo-guidelines/page_3.png',
  '/logo-guidelines/page_4.png',
  '/logo-guidelines/page_5.png',
  '/logo-guidelines/page_6.png',
  '/logo-guidelines/page_7.png',
  '/logo-guidelines/page_8.png',
  '/logo-guidelines/page_9.png',
  '/logo-guidelines/page_10.png'
]

const maxPages = computed(() => {
  return activeBook.value === 'sejarah' ? bookPages.length : logoPages.length
})

const switchBook = (book) => {
  activeBook.value = book
  currentPage.value = 0
}

onMounted(() => {
  store.fetchBooks()

  // Scroll Entrance Reveal Animation Observer
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible')
      }
    })
  }, {
    threshold: 0.1,
    rootMargin: '0px 0px -60px 0px' // Triggers slightly before entering fully
  })

  // Observe all timeline elements
  setTimeout(() => {
    document.querySelectorAll('.timeline-item').forEach(item => {
      observer.observe(item)
    })
  }, 100)
})

const bookPages = [
  {
    title: "Cover Page",
    content: ""
  },
  {
    title: "KATA PENGANTAR",
    content: `
      <p>Puji dan syukur kita panjatkan ke hadirat Tuhan Yang Maha Esa atas rahmat-Nya sehingga RSUD Prof. Dr. H. Aloei Saboe dapat berdiri kokoh melayani masyarakat selama genap 100 tahun.</p>
      <p>Buku sejarah digital ini dirancang khusus untuk memetakan sejarah penting pembentukan, perjuangan, integrasi teknologi hingga kepatuhan rekam medis elektronik (EMR) modern demi pelayanan kesehatan yang berintegritas dan kredibel di Provinsi Gorontalo.</p>
      <p>Semoga karya tulis ini menjadi inspirasi bakti tanpa henti bagi generasi mendatang.</p>
      <p><strong>- Jajaran Direksi RSUD Aloei Saboe</strong></p>
    `,
    image: "📜"
  },
  {
    title: "ERA KOLONIAL: ZEVENHOVEN HOSPITAL (1926)",
    content: `
      <p>Rumah sakit ini didirikan pertama kali pada tahun 1926 oleh Pemerintah Hindia Belanda dengan nama <strong>Zevenhoven Hospital</strong>.</p>
      <p>Nama ini diambil dari nama Residen Gorontalo saat itu yang memprakarsai berdirinya pelayanan kesehatan formal terintegrasi pertama guna menanggulangi penyakit endemik di wilayah jazirah Gorontalo.</p>
      <p>Bangunan fisik awal bercorak kolonial klasik dengan kapasitas 40 tempat tidur saja, menjadi cikal bakal sejarah kesehatan modern Gorontalo.</p>
    `,
    image: "🏛️"
  },
  {
    title: "PERJUANGAN DR. ALOEI SABOE (1942-1950)",
    content: `
      <p>Memasuki masa pendudukan Jepang dan perang kemerdekaan, rumah sakit ini berganti fungsi strategis menjadi benteng medis pejuang republik.</p>
      <p><strong>dr. Aloei Saboe</strong>, seorang dokter pejuang lulusan NIAS Surabaya, mengorbankan jiwa dan raganya untuk menyelamatkan masyarakat sipil dan gerilyawan nasional.</p>
      <p>Beliau memimpin reformasi medis, mengintegrasikan obat-obatan tradisional kala pasokan terputus, dan melatih pemuda Gorontalo dalam pertolongan pertama taktis medan tempur.</p>
    `,
    image: "🩺"
  },
  {
    title: "REFORMASI MODERN & DIGITALISASI (2026)",
    content: `
      <p>Kini, di usianya yang ke-100 tahun (1926-2026), RSUD Prof. Dr. H. Aloei Saboe bertransformasi total menjadi Smart Hospital.</p>
      <p>Dengan integrasi penuh Rekam Medis Elektronik (EMR) yang memenuhi standar Kepatuhan EMR (EMR Compliance), rumah sakit ini mampu menyajikan integrasi asuransi jaminan kesehatan nasional (BPJS) dengan akurasi klaim 100%.</p>
      <p>Dilengkapi sistem antrean digital, resep elektronik, dan SLA audit medis real-time, kami menyongsong abad kedua dengan optimisme paripurna digital!</p>
    `,
    image: "💻"
  },
  {
    title: "SIMBOLISME LOGO BARU SATU ABAD",
    content: `
      <div class="logo-page">
        <p class="intro-desc">Mencerminkan nilai luhur RSUD Prof. Dr. H. Aloei Saboe yang kokoh, modern, dan penuh empati:</p>
        <ul class="symbol-list">
          <li><strong>💖 Bentuk Hati:</strong> Komitmen melayani dengan kasih sayang, empati tinggi, dan ketulusan mendalam bagi kesembuhan pasien.</li>
          <li><strong>🌀 4 Bilah Lengkung:</strong> Melambangkan 4 pilar utama rumah sakit (Sarpras, SDM Kompeten, Transformasi Layanan, & Akreditasi Paripurna) yang saling terikat membentuk siluet <em>Cross</em> (Simbol Medis).</li>
          <li><strong>👐 Tangan Terbuka:</strong> Sinergi, kolaborasi dinamis, dan komitmen seluruh elemen dalam mengelola rumah sakit secara profesional.</li>
          <li><strong>⭕ Lingkaran Sempurna:</strong> Wilayah Provinsi Gorontalo, menegaskan peran RSAS sebagai pusat rujukan utama regional terpercaya.</li>
        </ul>
      </div>
    `,
    image: "🎨"
  },
  {
    title: "PALET WARNA & TIPOGRAFI",
    content: `
      <div class="logo-page">
        <p class="intro-desc">Identitas visual dirancang harmonis menggunakan skema warna bermakna tinggi:</p>
        <div class="color-palette-container">
          <div class="color-item">
            <span class="color-circle bg-emerald"></span>
            <div class="color-desc">
              <strong>Hijau RSAS (#00A859):</strong> Melambangkan pemulihan, keseimbangan tubuh & jiwa, serta kearifan lokal kerukunan Gorontalo.
            </div>
          </div>
          <div class="color-item">
            <span class="color-circle bg-orange"></span>
            <div class="color-desc">
              <strong>Oranye (#F79633):</strong> Energi optimisme tinggi, kehangatan pelayanan, antusiasme, serta kreativitas.
            </div>
          </div>
          <div class="color-item">
            <span class="color-circle bg-white"></span>
            <div class="color-desc">
              <strong>Putih (#FFFFFF):</strong> Fondasi kebersihan mutlak, transparansi digital, ketulusan pengabdian, dan integritas.
            </div>
          </div>
        </div>
        <p class="typography-note"><strong>Tipografi Resmi:</strong> Menggunakan rupa huruf <strong>Futura Md BT</strong> yang bercorak geometris, modern, anggun, dan tegas.</p>
      </div>
    `,
    image: "📐"
  }
]

const milestones = [
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
]

const prevPage = () => {
  if (currentPage.value > 0) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < maxPages.value - 1) {
    currentPage.value++
  }
}
</script>

<style lang="scss" scoped>
.history-page {
  max-width: 1100px;
  margin: 0 auto;
  padding: 3rem 1.5rem;
}

.text-center {
  text-align: center;
}

.header-section {
  margin-bottom: 4rem;

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

// E-book Styles
.ebook-section {
  text-align: center;
  margin-bottom: 6rem;

  .section-title {
    font-size: 2rem;
    margin-bottom: 0.5rem;
  }

  .desc {
    color: #94A3B8;
    margin-bottom: 3rem;
  }
}

.book-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.virtual-book {
  width: 100%;
  max-width: 650px;
  height: 480px;
  background: #fdfdfd; // Paper feel
  border-radius: 4px 16px 16px 4px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.4), inset 5px 0 10px rgba(0,0,0,0.1);
  position: relative;
  overflow: hidden;
  display: flex;
}

.spine-line {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 12px;
  background: linear-gradient(90deg, rgba(0,0,0,0.2) 0%, rgba(0,0,0,0.05) 50%, rgba(0,0,0,0.15) 100%);
  border-right: 1px solid rgba(0,0,0,0.08);
  z-index: 10;
}

.book-page {
  width: 100%;
  height: 100%;
  padding: 3.5rem 3rem 2rem 4rem;
  color: #1e293b; // Deep readability dark grey
  text-align: left;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  animation: page-flip 0.4s ease-out;
}

@keyframes page-flip {
  0% { transform: rotateY(-5deg) scaleX(0.95); opacity: 0.5; }
  100% { transform: rotateY(0deg) scaleX(1); opacity: 1; }
}

.page-num {
  position: absolute;
  bottom: 1.5rem;
  right: 2rem;
  font-size: 0.85rem;
  font-weight: 700;
  color: #64748b;
  letter-spacing: 1px;
}

.page-content {
  height: 100%;
  display: flex;
  flex-direction: column;
}

// Cover Page design
.page-cover {
  height: 100%;
  border: 3px double #F79633; // Centennial Gold-Orange
  padding: 2rem;
  background: radial-gradient(circle, #0B301D 0%, #03140A 100%); // Deep Emerald to dark green
  color: #ffffff;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  border-radius: 8px;

  .cover-gold-border {
    width: 100%;
    height: 100%;
    border: 1px solid rgba(247, 150, 51, 0.4);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 1.5rem;
  }

  .cover-sub {
    font-size: 0.8rem;
    letter-spacing: 4px;
    color: #F79633;
    margin-bottom: 1rem;
    font-weight: 700;
  }

  .cover-main {
    font-size: 2.2rem;
    line-height: 1.2;
    margin-bottom: 1rem;
  }

  .cover-years {
    font-size: 1.2rem;
    letter-spacing: 6px;
    color: #FFC385; // Champagne Orange-Gold
    font-weight: 600;
    margin-bottom: 2rem;
  }

  .cover-emblem {
    font-size: 3.5rem;
    margin-bottom: 2rem;
  }

  .cover-footer {
    font-size: 0.8rem;
    color: #94a3b8;
    line-height: 1.5;
  }
}

// Page body text
.page-body {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;

  .page-title {
    font-size: 1.5rem;
    color: #0f172a;
    border-bottom: 1px solid rgba(0,0,0,0.1);
    padding-bottom: 0.75rem;
    margin-bottom: 1.5rem;
    font-weight: 700;
    letter-spacing: 1px;
  }

  .page-text {
    font-size: 0.95rem;
    line-height: 1.7;
    color: #334155;
    flex-grow: 1;

    p {
      margin-bottom: 1rem;
    }
    
    strong {
      color: #0f172a;
    }
  }

  .page-illustration {
    align-self: center;
    font-size: 3.5rem;
    margin-top: 1rem;
  }
}

.book-controls {
  display: flex;
  align-items: center;
  gap: 2rem;
  margin-top: 2rem;
}

.control-btn {
  background: rgba(247, 150, 51, 0.1);
  border: 1px solid rgba(247, 150, 51, 0.3);
  color: #FFC385;
  padding: 0.6rem 1.5rem;
  border-radius: 30px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s ease;

  &:hover:not(:disabled) {
    background: #F79633;
    color: #03140A; // Dark Emerald Green background equivalent
    border-color: #F79633;
  }

  &:disabled {
    opacity: 0.3;
    cursor: not-allowed;
  }
}

.page-indicator {
  font-weight: 600;
  color: #94A3B8;
  font-size: 0.9rem;
}

// Timeline styles
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

  // Scroll Entrance Fade Reveal Base State
  opacity: 0;
  transition: opacity 0.8s cubic-bezier(0.16, 1, 0.3, 1), transform 0.8s cubic-bezier(0.16, 1, 0.3, 1);

  @media (max-width: 768px) {
    width: 100%;
    padding-left: 70px;
    padding-right: 15px;
    
    &.left, &.right {
      left: 0 !important;
      transform: translateX(-30px); // Slide in from left for responsive stacked lists
    }
  }

  &.left {
    left: 0;
    transform: translateX(-40px); // Slide in from left
  }

  &.right {
    left: 50%;
    transform: translateX(40px); // Slide in from right
  }

  // Active scrolled state toggled by IntersectionObserver!
  &.visible {
    opacity: 1;
    transform: translateX(0) !important;
  }

  .timeline-dot {
    width: 16px;
    height: 16px;
    border-radius: 50%;
    background: var(--color-primary, #F79633); // Adaptive dot!
    border: 3px solid var(--color-bg-dark, #03140A); // Match theme background
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
      color: var(--color-primary, #FFC385); // Dynamic contrast title
      margin-bottom: 0.75rem;
    }

    .timeline-desc {
      color: var(--color-card-text, #94A3B8); // Adaptive text contrast
      font-size: 0.9rem;
      line-height: 1.5;
    }
  }
}

// Custom Styles for Logo Philosophy Pages
.logo-page {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  text-align: left;
  
  .intro-desc {
    font-size: 0.85rem;
    color: #475569;
    font-style: italic;
    margin-bottom: 0.25rem;
    line-height: 1.4;
  }
}

.symbol-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.4rem;

  li {
    font-size: 0.8rem;
    line-height: 1.4;
    color: #334155;
    padding-left: 0.5rem;
    border-left: 3px solid #00A859;
    margin-bottom: 0.2rem;
    
    strong {
      color: #0f172a;
    }
    
    em {
      color: #F79633;
      font-weight: 700;
      font-style: normal;
    }
  }
}

.color-palette-container {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
  margin-bottom: 0.3rem;
}

.color-item {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  background: rgba(0, 0, 0, 0.02);
  padding: 0.35rem 0.5rem;
  border-radius: 6px;
  border: 1px solid rgba(0,0,0,0.04);
}

.color-circle {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  flex-shrink: 0;
  box-shadow: 0 1.5px 3px rgba(0,0,0,0.15);
  border: 1px solid rgba(0, 0, 0, 0.1);
  
  &.bg-emerald {
    background-color: #00A859;
  }
  
  &.bg-orange {
    background-color: #F79633;
  }
  
  &.bg-white {
    background-color: #FFFFFF;
  }
}

.color-desc {
  font-size: 0.78rem;
  line-height: 1.3;
  color: #475569;
  
  strong {
    color: #0f172a;
    font-size: 0.78rem;
  }
}

.typography-note {
  font-size: 0.78rem;
  line-height: 1.3;
  color: #334155;
  background: rgba(247, 150, 51, 0.05);
  padding: 0.4rem;
  border-radius: 4px;
  border-left: 3px solid #F79633;
  margin-top: 0.1rem;
}

// Book Selector Tabs
.book-tabs {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 2.5rem;
  background: rgba(11, 48, 29, 0.6); // Emerald translucent dark background
  padding: 0.4rem;
  border-radius: 50px;
  border: 1px solid rgba(247, 150, 51, 0.2);
  max-width: 540px;
  margin-left: auto;
  margin-right: auto;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.tab-btn {
  background: transparent;
  border: none;
  color: #94A3B8;
  padding: 0.6rem 1.4rem;
  font-family: 'Inter', sans-serif;
  font-weight: 700;
  font-size: 0.85rem;
  cursor: pointer;
  border-radius: 40px;
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  letter-spacing: 0.5px;

  &:hover {
    color: #FFC385;
    background: rgba(247, 150, 51, 0.08);
  }

  &.active {
    background: linear-gradient(135deg, #00A859 0%, #F79633 100%);
    color: #FFFFFF;
    box-shadow: 0 4px 15px rgba(247, 150, 51, 0.3);
  }
}

// PDF Image Book Presentation styling
.book-page.is-pdf-book {
  padding: 0 !important; // Overrides the default padding for a border-to-border layout
}

.pdf-page-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.pdf-image-wrapper {
  width: 100%;
  height: 100%;
  background-color: #ffffff;
  display: flex;
  justify-content: center;
  align-items: center;
}

.pdf-page-image {
  width: 100%;
  height: 100%;
  object-fit: contain; // Perfectly preserve logo PDF page proportions
}
</style>
