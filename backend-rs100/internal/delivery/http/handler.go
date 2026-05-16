package http

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"backend-rs100/internal/domain"
)

// HttpHandler handles HTTP requests for books and auth/activation.
type HttpHandler struct {
	bookUsecase domain.BookUsecase
	authUsecase domain.AuthUsecase
}

// NewHttpHandler creates a new HttpHandler instance.
func NewHttpHandler(bu domain.BookUsecase, au domain.AuthUsecase) *HttpHandler {
	return &HttpHandler{
		bookUsecase: bu,
		authUsecase: au,
	}
}

// RegisterRoutes registers the handler routes with the provided multiplexer.
func (h *HttpHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/books", h.handleBooks)
	mux.HandleFunc("/api/books/", h.handleBookDetail)
	mux.HandleFunc("/api/auth/register", h.handleRegister)
	mux.HandleFunc("/api/auth/activate", h.handleActivate)
	mux.HandleFunc("/api/auth/login", h.handleLogin)
	
	// Gallery routes
	mux.HandleFunc("/admin/gallery", h.handleGalleryForm)
	mux.HandleFunc("/api/gallery", h.handleGallery)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
}

// CORS Helper
func enableCORS(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

// JSON Response Helpers
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func (h *HttpHandler) handleBooks(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if r.Method == http.MethodGet {
		books, err := h.bookUsecase.Fetch(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, books)
		return
	}

	if r.Method == http.MethodPost {
		var book domain.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if err := h.bookUsecase.Store(r.Context(), &book); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, book)
	}
}

func (h *HttpHandler) handleBookDetail(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	idStr := parts[3]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	if r.Method == http.MethodGet {
		book, err := h.bookUsecase.GetByID(r.Context(), id)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, book)
		return
	}

	respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

func (h *HttpHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.authUsecase.Register(r.Context(), req.Email, req.Password); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Registration successful, activation code sent"})
}

func (h *HttpHandler) handleActivate(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.authUsecase.Activate(r.Context(), req.Code); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Activation successful! Selamat Hari Jadi Ke-100 RSUD Aloei Saboe!"})
}

func (h *HttpHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := h.authUsecase.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// Gallery Handlers

func (h *HttpHandler) handleGalleryForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Input Foto Sejarah</title>
		<style>
			body { font-family: sans-serif; padding: 2rem; background: #0B301D; color: white; }
			form { background: rgba(255,255,255,0.05); padding: 2rem; border-radius: 8px; max-width: 500px; margin: 0 auto; box-shadow: 0 4px 20px rgba(0,0,0,0.3); border: 1px solid rgba(247,150,51,0.2); backdrop-filter: blur(10px); }
			.form-group { margin-bottom: 1rem; }
			label { display: block; margin-bottom: 0.5rem; font-weight: bold; color: #FFC385; }
			input[type=text], textarea { width: 100%; padding: 0.5rem; border: 1px solid rgba(247,150,51,0.2); border-radius: 4px; box-sizing: border-box; background: rgba(0,0,0,0.2); color: white; }
			textarea { height: 100px; }
			button { background: #F79633; color: #03140A; padding: 0.75rem 1.5rem; border: none; border-radius: 4px; cursor: pointer; font-weight: bold; width: 100%; }
			button:hover { background: #e0852a; }
			h2 { text-align: center; color: #F79633; }
		</style>
	</head>
	<body>
		<form action="/api/gallery" method="POST" enctype="multipart/form-data">
			<h2>Input Foto & Narasi Sejarah</h2>
			<div class="form-group">
				<label for="title">Judul Foto</label>
				<input type="text" id="title" name="title" required placeholder="Contoh: Gedung Pertama 1926">
			</div>
			<div class="form-group">
				<label for="narration">Narasi Suara</label>
				<textarea id="narration" name="narration" required placeholder="Ketik narasi yang akan dibacakan di sini..."></textarea>
			</div>
			<div class="form-group">
				<label for="photo">File Foto</label>
				<input type="file" id="photo" name="photo" accept="image/*" required>
			</div>
			<button type="submit">Simpan Data</button>
		</form>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func (h *HttpHandler) handleGallery(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	// Simple file-based storage for demo/fallback
	dataFile := "gallery.json"

	if r.Method == http.MethodGet {
		// Read from file
		data, err := os.ReadFile(dataFile)
		if err != nil {
			// If not found, return empty array
			respondWithJSON(w, http.StatusOK, []interface{}{})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}

	if r.Method == http.MethodPost {
		// Handle file upload and text
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "File too large")
			return
		}

		title := r.FormValue("title")
		narration := r.FormValue("narration")

		file, handler, err := r.FormFile("photo")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Error retrieving file")
			return
		}
		defer file.Close()

		// Create uploads dir if not exists (just in case)
		os.MkdirAll("uploads", os.ModePerm)

		// Save file
		filename := time.Now().Format("20060102150405") + "_" + handler.Filename
		filepath := "uploads/" + filename
		
		dst, err := os.Create(filepath)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error saving file")
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error saving file content")
			return
		}

		// Save metadata to JSON
		type GalleryItem struct {
			Title     string `json:"title"`
			Narration string `json:"narration"`
			PhotoPath string `json:"photoPath"`
		}

		var items []GalleryItem
		
		// Read existing
		existingData, err := os.ReadFile(dataFile)
		if err == nil {
			json.Unmarshal(existingData, &items)
		}

		// Append new
		items = append(items, GalleryItem{
			Title:     title,
			Narration: narration,
			PhotoPath: "/uploads/" + filename,
		})

		// Write back
		updatedData, _ := json.MarshalIndent(items, "", "  ")
		os.WriteFile(dataFile, updatedData, 0644)

		// Success response (styled HTML)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<head>
				<style>
					body { font-family: sans-serif; padding: 2rem; background: #0B301D; color: white; text-align: center; }
					.success-box { background: rgba(255,255,255,0.05); padding: 2rem; border-radius: 8px; max-width: 500px; margin: 0 auto; border: 1px solid #00A859; }
					a { color: #F79633; text-decoration: none; font-weight: bold; }
				</style>
			</head>
			<body>
				<div class="success-box">
					<h2 style="color: #00A859;">Data Berhasil Disimpan!</h2>
					<p>Foto dan narasi Anda telah tersimpan di sistem.</p>
					<p><a href="/admin/gallery">← Input Lagi</a> | <a href="http://localhost:5173/history" target="_blank">Lihat Halaman Sejarah →</a></p>
				</div>
			</body>
			</html>
		`))
	}
}
