<template>
  <div class="dashboard-page">
    <div class="header-section text-center">
      <h1 class="font-historical text-gold-gradient">Dashboard E-Book</h1>
      <p class="subtitle">Akses Buku Sejarah dan Panduan Logo RSUD Prof. Dr. H. Aloei Saboe</p>
    </div>

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

    <div class="history-content-grid">
      <!-- Full Width Column: E-Book -->
      <div class="ebook-column full-width">
        <div class="virtual-book large-book">
          <!-- Left Spine Shadow -->
          <div class="spine-line"></div>

          <!-- Active Page Content -->
          <div class="book-page active-page is-pdf-book" :key="activeBook + '_' + currentPage">
            <div class="page-num" v-if="activeBook !== 'sejarah'">{{ currentPage + 1 }} / {{ maxPages }}</div>
            
            <div class="page-content pdf-page-container">
              <!-- BOOK 1: SEJARAH (PDF IFRAME) -->
              <template v-if="activeBook === 'sejarah'">
                <div class="pdf-iframe-wrapper">
                  <iframe 
                    src="/BUKU RSAS 100 TAHUN.pdf" 
                    width="100%" 
                    height="100%" 
                    style="border: none;"
                  ></iframe>
                </div>
              </template>

              <!-- BOOK 2: LOGO GUIDELINES (IMAGE RENDER) -->
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
        <div class="book-controls" v-if="activeBook !== 'sejarah'">
          <button @click="prevPage" :disabled="currentPage === 0" class="control-btn">&larr; Sebelumnya</button>
          <div class="page-indicator">Halaman {{ currentPage + 1 }}</div>
          <button @click="nextPage" :disabled="currentPage === maxPages - 1" class="control-btn">Selanjutnya &rarr;</button>
        </div>
      </div>
    </div>
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
  return activeBook.value === 'sejarah' ? 23 : logoPages.length
})

const switchBook = (book) => {
  activeBook.value = book
  currentPage.value = 0
}

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

onMounted(async () => {
  store.fetchBooks()
  if (window.__stopHeartbeat) {
    window.__stopHeartbeat()
  }
})
</script>

<style lang="scss" scoped>
.dashboard-page {
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

.book-tabs {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 2.5rem;
  background: rgba(11, 48, 29, 0.6);
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

.history-content-grid {
  display: flex;
  justify-content: center;
  margin-top: 2rem;
  margin-bottom: 4rem;
}

.ebook-column {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  
  &.full-width {
    max-width: 1000px; /* Wider for better reading */
  }
}

.virtual-book {
  width: 100%;
  height: 700px; /* Taller for better PDF reading */
  background: linear-gradient(to right, #fdfdfd 0%, #f4f4f4 100%);
  border-radius: 4px 16px 16px 4px;
  box-shadow: 0 30px 60px rgba(0, 0, 0, 0.4), inset 5px 0 10px rgba(0,0,0,0.1);
  position: relative;
  overflow: hidden;
  display: flex;
  border: 1px solid rgba(247, 150, 51, 0.2);
}

.spine-line {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 20px;
  background: linear-gradient(90deg, rgba(0,0,0,0.3) 0%, rgba(0,0,0,0.05) 50%, rgba(0,0,0,0.2) 100%);
  border-right: 1px solid rgba(0,0,0,0.15);
  z-index: 10;
}

.book-page {
  width: 100%;
  height: 100%;
  padding: 0 !important; /* No padding for PDF */
  color: #1e293b;
  text-align: left;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.page-num {
  position: absolute;
  bottom: 1.5rem;
  right: 2rem;
  font-size: 0.85rem;
  font-weight: 700;
  color: #64748b;
  letter-spacing: 1px;
  z-index: 10;
}

.page-content {
  width: 100%;
  height: 100%;
}

.pdf-page-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.pdf-iframe-wrapper {
  width: 100%;
  height: 100%;
}

.pdf-image-wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.pdf-page-image {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
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
    color: #03140A;
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
</style>
