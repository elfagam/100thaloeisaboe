<template>
  <div class="home-page-wrapper">
    <div class="home-container">
      <!-- Hero Section -->
      <header class="hero">
        <div class="hero-bg-image"></div>
        <div class="hero-overlay"></div>
        
        <div class="hero-content">
          <h2 class="pre-title">1926 - 2026</h2>
          <h1 class="main-title font-historical">
            SATU ABAD <br />
            <span class="text-green-gradient">RSUD ALOEI SABOE</span>
          </h1>
          <p class="tagline">
            "Satu Abad Melayani: Menjadi Lebih Baik untuk Semua"
          </p>
          
          <div class="action-buttons">
            <router-link to="/launching" class="premium-btn">VVIP Peresmian</router-link>
            <router-link to="/history" class="btn-outline">Telusuri Sejarah</router-link>
          </div>
        </div>

        <!-- Hero Visuals -->
        <div class="hero-visual">
          <div class="gold-ring"></div>
          <div class="inner-shield">
            <span class="century-number text-gold-gradient">100</span>
            <span class="century-text">TAHUN</span>
          </div>
        </div>
      </header>

      <!-- Key Highlights / Stats -->
      <section class="highlights-grid">
        <div class="premium-card highlight-item">
          <div class="icon-wrap">🛡️</div>
          <h3>Akreditasi Paripurna</h3>
          <p>Standar pelayanan tertinggi nasional yang menjamin mutu keselamatan pasien.</p>
        </div>

        <div class="premium-card highlight-item">
          <div class="icon-wrap">⚡</div>
          <h3>100% EMR Compliance</h3>
          <p>Digitalisasi rekam medis elektronik terintegrasi penuh guna mendukung transparansi klaim.</p>
        </div>

        <div class="premium-card highlight-item">
          <div class="icon-wrap">🏛️</div>
          <h3>Bakti Sejarah</h3>
          <p>Transformasi panjang dari Zevenhoven Hospital 1926 hingga menjadi pusat rujukan modern.</p>
        </div>
      </section>

      <!-- Interactive Section Summary -->
      <section class="overview-section">
        <h2 class="section-title font-historical text-green-gradient">Portal Utama Satu Abad</h2>
        <div class="overview-cards">
          <div class="p-card" @click="router.push('/launching')">
            <div class="card-img-placeholder launch-bg">
              <div class="ripple-circle"></div>
            </div>
            <div class="p-card-body">
              <h3>Layar Interaktif VVIP</h3>
              <p>Sistem digital peluncuran satu abad dengan verifikasi kode otorisasi milidetik tinggi.</p>
              <span class="card-link">Buka Portal Launching &rarr;</span>
            </div>
          </div>

          <div class="p-card" @click="router.push('/history')">
            <div class="card-img-placeholder history-bg">
              <div class="book-spine"></div>
            </div>
            <div class="p-card-body">
              <h3>E-Book Sejarah & Timeline</h3>
              <p>Lembaran narasi sejarah perjuangan kesehatan gorontalo dan album dokumentasi arsip.</p>
              <span class="card-link">Baca E-Book Sejarah &rarr;</span>
            </div>
          </div>
        </div>
      </section>

      <!-- Footer -->
      <footer class="footer">
        <p>&copy; 2026 RSUD Prof. Dr. H. Aloei Saboe. Seluruh Hak Cipta Dilindungi.</p>
        <p class="muted">Dikembangkan dengan dedikasi penuh menyambut perayaan 1 Abad</p>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()
</script>

<style lang="scss" scoped>
.home-page-wrapper {
  // Sophisticated radial backlight centered exactly behind the floating Centennial Shield!
  background: radial-gradient(circle at 80% 35%, var(--color-status-active-bg, rgba(0, 255, 135, 0.08)) 0%, var(--color-bg-dark, #E6F5EC) 75%);
  color: var(--color-white, #1E293B); // Real-time adaptive body text!
  min-height: 90vh;
  width: 100%;
  transition: background 0.6s cubic-bezier(0.16, 1, 0.3, 1), color 0.3s ease;
}

.home-container {
  max-width: 1400px; // Wider for premium large screens
  margin: 0 auto;
  padding: 3rem 1.5rem 8rem 1.5rem;
  
  @media (min-width: 1920px) {
    max-width: 1800px;
    padding-top: 5rem;
  }
}

.hero {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 4rem;
  min-height: 80vh;
  padding: 4rem 3rem;
  border-radius: 30px;
  overflow: hidden;
  margin-bottom: 4rem;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.2);

  @media (min-width: 1920px) {
    min-height: 85vh;
    padding: 6rem 5rem;
    gap: 8rem;
  }

  .hero-bg-image {
    position: absolute;
    inset: 0;
    background-image: url('/rsas.jpeg');
    background-size: cover;
    background-position: center;
    transform: scale(1.05);
    filter: brightness(0.7) contrast(1.1);
    z-index: 0;
  }

  .hero-overlay {
    position: absolute;
    inset: 0;
    // Layer 1: Base Darkening
    background: linear-gradient(
      135deg,
      rgba(3, 20, 10, 0.92) 0%,
      rgba(11, 48, 29, 0.6) 50%,
      rgba(3, 20, 10, 0.8) 100%
    );
    z-index: 1;
    backdrop-filter: blur(1px);

    // Layer 2: Animated Mesh Shimmer
    &::before {
      content: '';
      position: absolute;
      inset: 0;
      background: radial-gradient(circle at 20% 30%, rgba(0, 255, 135, 0.15) 0%, transparent 50%),
                  radial-gradient(circle at 80% 70%, rgba(247, 150, 51, 0.1) 0%, transparent 50%);
      animation: mesh-float 10s ease-in-out infinite alternate;
      z-index: 1;
    }

    // Layer 3: Cinematic Light Leak
    &::after {
      content: '';
      position: absolute;
      top: -10%;
      right: -10%;
      width: 50%;
      height: 50%;
      background: radial-gradient(circle, rgba(0, 255, 135, 0.2) 0%, transparent 70%);
      filter: blur(60px);
      z-index: 2;
    }
  }

  // Glowing Border
  &::before {
    content: '';
    position: absolute;
    inset: 0;
    padding: 1px;
    border-radius: 30px;
    background: linear-gradient(135deg, rgba(0, 255, 135, 0.4), transparent, rgba(247, 150, 51, 0.4));
    -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
    mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
    -webkit-mask-composite: xor;
    mask-composite: exclude;
    pointer-events: none;
    z-index: 3;
  }

  .hero-content, .hero-visual {
    position: relative;
    z-index: 2;
  }

  @media (max-width: 968px) {
    flex-direction: column;
    text-align: center;
    justify-content: center;
    gap: 3rem;
  }
}

.hero-content {
  flex: 1.2;
  
  .pre-title {
    font-size: 1.2rem;
    color: #00FF87; // Electric mint for better contrast on image
    letter-spacing: 6px;
    margin-bottom: 1rem;
    font-weight: 700;
    text-shadow: 0 0 20px rgba(0, 255, 135, 0.5);
  }

  .main-title {
    font-size: 4rem;
    line-height: 1.1;
    margin-bottom: 1.5rem;
    font-weight: 900;
    letter-spacing: -1px;
    color: #FFFFFF;
    text-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);

    @media (min-width: 1920px) {
      font-size: 6rem;
    }

    @media (max-width: 576px) {
      font-size: 2.8rem;
    }
  }

  .text-green-gradient {
    // 5-stop ultra-premium emerald-to-gold-to-mint jewel gradient
    background: linear-gradient(
      110deg,
      var(--color-primary, #00A859) 0%,
      #00FF87 30%,                           // Radiant electric mint highlight
      #FFFFFF 48%,                           // Pure specular glossy silver reflection
      var(--color-gold, #F79633) 70%,
      var(--color-gold-dark, #C16400) 100%
    );
    background-size: 200% auto;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
    
    // Slow sweeping jewel luster animation
    animation: green-metallic-shimmer 7s linear infinite;
    
    // Soft outer neon glow to make the letters pop majestically off deep backgrounds
    filter: drop-shadow(0 2px 10px rgba(0, 255, 135, 0.25));
  }

  .tagline {
    font-size: 1.25rem;
    color: rgba(255, 255, 255, 0.9);
    margin-bottom: 2.5rem;
    line-height: 1.6;
    max-width: 600px;
    font-weight: 600;
    font-style: italic;
    text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  }
}

.action-buttons {
  display: flex;
  gap: 1.5rem;
  
  @media (max-width: 968px) {
    justify-content: center;
  }

  @media (max-width: 480px) {
    flex-direction: column;
    width: 100%;
    
    a, button {
      width: 100%;
    }
  }
}

.premium-btn {
  display: inline-block;
  text-decoration: none;
  font-family: 'Inter', sans-serif;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  background: linear-gradient(135deg, var(--color-primary, #00FF87) 0%, var(--color-gold, #00A859) 100%); // Dynamic button gradient!
  color: #FFFFFF;
  padding: 1rem 2.5rem;
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  text-align: center;
  box-shadow: 0 4px 15px rgba(0, 168, 89, 0.2);

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(0, 168, 89, 0.35);
  }
}

.btn-outline {
  display: inline-block;
  text-decoration: none;
  font-family: 'Inter', sans-serif;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  border: 2px solid var(--color-primary, #00A859); // Dynamic border!
  color: var(--color-primary, #00A859); // Dynamic text!
  padding: 1rem 2.5rem;
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.3s ease;
  text-align: center;

  &:hover {
    background: rgba(0, 168, 89, 0.08);
    color: var(--color-primary, #00A859);
    opacity: 0.9;
  }
}

// Hero Visual Ring
.hero-visual {
  flex: 0.8;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  width: 320px;
  height: 320px;

  @media (min-width: 1920px) {
    width: 480px;
    height: 480px;
  }

  .gold-ring {
    position: absolute;
    width: 100%;
    height: 100%;
    border-radius: 50%;
    
    // Solid polished metallic gold borders
    border: 3.5px solid var(--color-gold, #F79633);
    
    // Soft premium ambient glow to cradle the floating shield
    box-shadow: 
      0 0 35px var(--color-status-active-shadow, rgba(0, 255, 135, 0.2)), // Beautiful neon glow
      inset 0 0 15px var(--color-gold-dark, #C16400); // Warm internal bronze shade
      
    animation: pulse-gold 4s ease-in-out infinite; // Smooth, slow breathing animation
    z-index: 1;
  }

  .inner-shield {
    width: 250px;
    height: 250px;

    @media (min-width: 1920px) {
      width: 380px;
      height: 380px;
    }
    
    // Hyper-realistic 3D green crystal glass sphere composition (Offset studio flare + transparent glass)
    background: 
      radial-gradient(circle at 35% 20%, rgba(255, 255, 255, 0.35) 0%, rgba(255, 255, 255, 0) 45%), // Specular studio light highlight spot
      radial-gradient(circle, rgba(11, 48, 29, 0.72) 15%, rgba(0, 168, 89, 0.5) 100%); // Translucent emerald green glass core
      
    // High-refraction double-layered border to simulate physical glass bezel edges
    border: 1.5px solid rgba(255, 255, 255, 0.42);
    border-radius: 50%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    
    // Glassmorphism refraction (Blurs the moving ambient background video playing behind the shield!)
    backdrop-filter: blur(16px);
    -webkit-backdrop-filter: blur(16px);
    
    // Deep 3D curved glass shading (Refracted borders, light leaks, and shadows)
    box-shadow: 
      0 20px 45px rgba(0, 0, 0, 0.45),                               // Realistic 3D drop shadow
      inset 0 8px 18px rgba(255, 255, 255, 0.42),                    // Specular top glass bezel refraction
      inset 0 -8px 20px rgba(0, 0, 0, 0.55),                         // Ambient bottom shadow occlusion density
      inset 0 0 12px rgba(255, 255, 255, 0.25);                      // Inner refraction glass halo
      
    animation: float 6s ease-in-out infinite;
    position: relative;
    overflow: hidden;
    z-index: 2;

    // The iconic crescent glass dome highlight overlay (Draped over the rotating facets)
    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 50%; // Upper half dome
      background: linear-gradient(
        180deg,
        rgba(255, 255, 255, 0.35) 0%,
        rgba(255, 255, 255, 0.05) 75%,
        transparent 100%
      );
      border-radius: 125px 125px 0 0;
      pointer-events: none;
      z-index: 2; // Sits on top of the rotating diamond facets
    }

    // Faceted rotating conic gradient to simulate multi-faceted light reflections of a diamond
    &::after {
      content: '';
      position: absolute;
      top: -50%; // Expand boundary bounds to prevent flat corners from showing during 360 rotation
      left: -50%;
      width: 200%;
      height: 200%;
      border-radius: 50%;
      background: conic-gradient(
        from 0deg at 50% 50%,
        transparent 0deg,
        rgba(255, 255, 255, 0.16) 15deg,  // Pure Specular White (Glass Highlight)
        transparent 35deg,
        rgba(0, 168, 89, 0.25) 55deg,     // Emerald Green (RSUD Aloei Saboe Brand)
        transparent 75deg,
        rgba(247, 150, 51, 0.22) 95deg,   // Centennial Orange-Gold (100 Years Majesty)
        transparent 115deg,
        rgba(255, 255, 255, 0.16) 135deg,
        transparent 155deg,
        rgba(0, 168, 89, 0.25) 175deg,
        transparent 195deg,
        rgba(247, 150, 51, 0.22) 215deg,
        transparent 235deg,
        rgba(255, 255, 255, 0.16) 255deg,
        transparent 275deg,
        rgba(0, 168, 89, 0.25) 295deg,
        transparent 315deg,
        rgba(247, 150, 51, 0.22) 335deg,
        transparent 355deg,
        transparent 360deg
      );
      mix-blend-mode: overlay; // Integrates facets smoothly into the emerald-green radial base background
      animation: spin-gemstone 25s linear infinite; // Cinematic slow, mesmerizing rotation
      z-index: 1; // Sits behind text (z-index 3) and crescent (z-index 2)
      pointer-events: none;
    }

    .century-number {
      font-size: 5.5rem;
      font-weight: 900;
      font-family: 'Playfair Display', serif;
      line-height: 1;
      color: var(--color-shield-text, #F79633); // Real-time adaptive numeral color!
      text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3); // High contrast text shadow
      z-index: 3;

      @media (min-width: 1920px) {
        font-size: 8.5rem;
      }
    }

    .century-text {
      font-size: 1rem;
      letter-spacing: 8px;
      color: var(--color-shield-text, #FFC385); // Real-time adaptive label color!
      font-weight: 700;
      margin-top: 0.5rem;
      opacity: 0.9;
      text-shadow: 0 1px 5px rgba(0, 0, 0, 0.3);
      z-index: 3;
    }
  }
}

// Highlights Section
.highlights-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(290px, 1fr));
  gap: 2.5rem; // Expansive premium gap for better information breathing space
  margin: 6rem 0 3rem;
  padding: 1rem; // Safety padding cushion to prevent hover shadows from being clipped by overflow: hidden containers!
}

.premium-card {
  // Ultra-premium glassmorphism base
  background: var(--color-card-bg, rgba(255, 255, 255, 0.8)); 
  backdrop-filter: blur(15px);
  -webkit-backdrop-filter: blur(15px);
  
  // High-contrast reflective bezel border mimicking polished cut glass
  border: 1px solid rgba(255, 255, 255, 0.15); 
  border-radius: 16px;
  padding: 2.5rem 2rem;
  box-shadow: 0 8px 32px 0 rgba(0, 0, 0, 0.08);
  transition: transform 0.4s cubic-bezier(0.16, 1, 0.3, 1), box-shadow 0.4s ease, border-color 0.4s ease, background 0.3s ease;
  
  position: relative;
  overflow: hidden;

  // Specular Diagonal Sweep Highlight (Kilatan Cahaya Refleksi Kaca)
  &::after {
    content: '';
    position: absolute;
    top: 0;
    left: -150%;
    width: 100%;
    height: 100%;
    background: linear-gradient(
      120deg,
      transparent 0%,
      rgba(255, 255, 255, 0) 30%,
      rgba(255, 255, 255, 0.35) 50%, // Translucent glass shine beam
      rgba(255, 255, 255, 0) 70%,
      transparent 100%
    );
    transform: skewX(-20deg);
    transition: left 0.7s cubic-bezier(0.16, 1, 0.3, 1);
    pointer-events: none;
  }

  &:hover {
    transform: translateY(-6px);
    border-color: var(--color-primary, #00FF87);
    box-shadow: 0 15px 35px var(--color-status-active-shadow, rgba(0, 255, 135, 0.15));
    background: var(--color-card-bg, rgba(255, 255, 255, 0.95));
    
    // Sweep the glossy ray of light across the glass card on hover!
    &::after {
      left: 150%;
    }
  }
}

.highlight-item {
  text-align: center;

  .icon-wrap {
    font-size: 3rem;
    margin-bottom: 1.5rem;
  }

  h3 {
    font-family: 'Playfair Display', serif;
    font-size: 1.5rem;
    color: var(--color-primary, #054826); // Dynamic headings!
    margin-bottom: 1rem;
  }

  p {
    color: var(--color-card-text, #475569); // High-contrast dynamic card body text!
    opacity: 0.9;
    line-height: 1.6;
    font-size: 0.95rem;
  }
}

// Overview Portal Section
.overview-section {
  text-align: center;
  margin: 6rem 0 3rem;

  .section-title {
    font-size: 2.5rem;
    margin-bottom: 3rem;
  }
}

.overview-cards {
  display: flex;
  justify-content: center;
  gap: 3rem;
  flex-wrap: wrap;
}

.p-card {
  width: 450px;
  background: var(--color-card-bg, rgba(255, 255, 255, 0.95)); // Adaptive portal card!
  border: 1px solid var(--color-secondary, rgba(0, 255, 135, 0.2));
  border-radius: 16px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
  text-align: left;
  box-shadow: 0 10px 30px rgba(0, 168, 89, 0.04);

  &:hover {
    transform: translateY(-8px) scale(1.02);
    border-color: var(--color-primary, #00FF87);
    box-shadow: 0 20px 40px rgba(0, 168, 89, 0.12);

    .ripple-circle {
      transform: scale(1.2);
    }
    
    .book-spine {
      transform: rotateY(-10deg);
    }
  }

  .card-img-placeholder {
    height: 200px;
    background-color: var(--color-secondary, var(--color-bg-dark, #E6F5EC));
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    overflow: hidden;

    &.launch-bg {
      background: linear-gradient(135deg, var(--color-primary, #00A859) 0%, var(--color-secondary, var(--color-bg-dark, #E6F5EC)) 100%);
    }

    &.history-bg {
      background: linear-gradient(135deg, var(--color-secondary, var(--color-bg-dark, #E6F5EC)) 0%, var(--color-primary, #00A859) 100%);
    }
  }

  .ripple-circle {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    border: 2px solid rgba(255, 255, 255, 0.6);
    background: rgba(255, 255, 255, 0.15);
    transition: transform 0.6s ease;
    box-shadow: 0 0 20px rgba(0, 255, 135, 0.2);
  }

  .book-spine {
    width: 60px;
    height: 100px;
    background: linear-gradient(90deg, var(--color-gold-dark, #C16400) 0%, var(--color-gold, #F79633) 50%, var(--color-gold-dark, #C16400) 100%);
    border-radius: 4px;
    box-shadow: 5px 5px 15px rgba(0,0,0,0.2);
    transition: transform 0.6s ease;
  }

  .p-card-body {
    padding: 2rem;

    h3 {
      font-size: 1.4rem;
      color: var(--color-primary, #054826);
      margin-bottom: 0.75rem;
      font-family: 'Playfair Display', serif;
    }

    p {
      color: var(--color-card-text, #475569);
      opacity: 0.9;
      font-size: 0.95rem;
      line-height: 1.6;
      margin-bottom: 1.5rem;
    }

    .card-link {
      color: var(--color-primary, #00A859);
      font-weight: 700;
      letter-spacing: 1px;
      text-transform: uppercase;
      font-size: 0.85rem;
      transition: color 0.3s ease;
    }
  }
}

// Footer
.footer {
  margin-top: 8rem;
  border-top: 1px solid var(--color-secondary, rgba(0, 168, 89, 0.15));
  padding: 3rem 0;
  text-align: center;
  color: var(--color-white, #475569);
  opacity: 0.8;
  font-size: 0.9rem;

  p {
    margin-bottom: 0.5rem;
  }

  .muted {
    font-size: 0.8rem;
    color: var(--color-white, #64748B);
    opacity: 0.65;
  }
}

@keyframes sweep-gloss {
  0% {
    left: -150%;
  }
  25% {
    left: 150%;
  }
  100% {
    left: 150%; // Elegant resting pause to avoid constant frantic flashing
  }
}

@keyframes green-metallic-shimmer {
  0% {
    background-position: 200% center;
  }
  100% {
    background-position: -200% center;
  }
}

@keyframes spin-gemstone {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

@keyframes mesh-float {
  0% {
    transform: scale(1) translate(0, 0);
  }
  100% {
    transform: scale(1.2) translate(5%, 5%);
  }
}
</style>
