package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"backend-rs100/internal/domain"
	"cloud.google.com/go/storage"
)

// HttpHandler handles HTTP requests for books and auth/activation.
type HttpHandler struct {
	bookUsecase      domain.BookUsecase
	authUsecase      domain.AuthUsecase
	galleryUsecase   domain.GalleryUsecase
	milestoneUsecase domain.MilestoneUsecase
	storageClient    *storage.Client
	bucketName       string
}

// NewHttpHandler creates a new HttpHandler instance.
func NewHttpHandler(bu domain.BookUsecase, au domain.AuthUsecase, gu domain.GalleryUsecase, mu domain.MilestoneUsecase, sc *storage.Client, bn string) *HttpHandler {
	return &HttpHandler{
		bookUsecase:      bu,
		authUsecase:      au,
		galleryUsecase:   gu,
		milestoneUsecase: mu,
		storageClient:    sc,
		bucketName:       bn,
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
	mux.HandleFunc("/admin/", h.handleDashboard)
	mux.HandleFunc("/admin/gallery", h.handleGalleryForm)
	mux.HandleFunc("/api/gallery", h.handleGallery)
	mux.HandleFunc("/api/gallery/", h.handleGalleryDetail)
	
	// Milestone routes
	mux.HandleFunc("/admin/milestones", h.handleMilestoneForm)
	mux.HandleFunc("/api/milestones", h.handleMilestones)
	mux.HandleFunc("/api/milestones/", h.handleMilestoneDetail)
	
	// Remote routes
	mux.HandleFunc("/api/remote/trigger", h.handleRemoteTrigger)
	mux.HandleFunc("/api/remote/status", h.handleRemoteStatus)
	
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

var lastRemoteTrigger int64

func (h *HttpHandler) handleRemoteTrigger(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}
	lastRemoteTrigger = time.Now().UnixNano() / int64(time.Millisecond)
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "triggered", "timestamp": lastRemoteTrigger})
}

func (h *HttpHandler) handleRemoteStatus(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"lastTrigger": lastRemoteTrigger})
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
	items, err := h.galleryUsecase.Fetch(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch gallery items: "+err.Error())
		return
	}

	// Build table rows safely
	tableRows := ""
	for _, item := range items {
		// Escape quotes for HTML attributes
		escapedTitle := strings.ReplaceAll(item.Title, "\"", "&quot;")
		escapedNarration := strings.ReplaceAll(item.Narration, "\"", "&quot;")
		
		tableRows += fmt.Sprintf(`
			<tr>
				<td><img src="%s" style="width: 50px; height: 50px; object-fit: cover; border-radius: 4px;"></td>
				<td>%s</td>
				<td>%s</td>
				<td>
					<button class="btn-edit" data-id="%s" data-title="%s" data-narration="%s" data-photo="%s" onclick="openEditModal(this)">Edit</button>
					<button class="btn-delete" onclick="deleteGallery('%s')">Hapus</button>
				</td>
			</tr>
		`, item.PhotoPath, item.Title, item.Narration, item.ID, escapedTitle, escapedNarration, item.PhotoPath, item.ID)
	}

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Kelola Galeri Foto</title>
		<style>
			body { font-family: 'Inter', sans-serif; background: #0B301D; color: white; padding: 2rem; }
			.container { max-width: 900px; margin: 0 auto; background: rgba(255,255,255,0.05); padding: 2rem; border-radius: 12px; border: 1px solid rgba(255,255,255,0.1); }
			h2 { color: #F79633; text-align: center; }
			.form-group { margin-bottom: 1.5rem; }
			label { display: block; margin-bottom: 0.5rem; color: #FFC385; font-weight: bold; }
			input[type="text"], textarea { width: 100%%; padding: 0.75rem; border-radius: 6px; border: 1px solid rgba(247,150,51,0.2); background: rgba(0,0,0,0.3); color: white; box-sizing: border-box; }
			textarea { height: 100px; resize: vertical; }
			button { background: #F79633; color: #03140A; padding: 0.75rem 1.5rem; border: none; border-radius: 6px; cursor: pointer; font-weight: bold; }
			button:hover { background: #e0852a; }
			
			table { width: 100%%; border-collapse: collapse; margin-top: 2rem; }
			th, td { padding: 0.75rem; text-align: left; border-bottom: 1px solid rgba(255,255,255,0.1); }
			th { color: #F79633; }
			.btn-edit { background: #cca43b; margin-right: 0.5rem; color: black; }
			.btn-delete { background: #c0392b; color: white; }
			
			a { color: #F79633; text-decoration: none; display: block; text-align: center; margin-top: 1rem; }

			/* Modal Styles */
			.modal { display: none; position: fixed; z-index: 1000; left: 0; top: 0; width: 100%%; height: 100%%; background-color: rgba(0,0,0,0.8); backdrop-filter: blur(5px); }
			.modal-content { background: #0B301D; margin: 10%% auto; padding: 2rem; border: 1px solid #F79633; width: 50%%; border-radius: 12px; box-shadow: 0 5px 15px rgba(0,0,0,0.5); }
			.close { color: #aaa; float: right; font-size: 28px; font-weight: bold; cursor: pointer; }
			.close:hover { color: #F79633; }
		</style>
		<script>
			function deleteGallery(id) {
				if (confirm('Apakah Anda yakin ingin menghapus data ini?')) {
					fetch('/api/gallery/' + id, { method: 'DELETE' })
					.then(res => res.json())
					.then(data => {
						alert('Berhasil dihapus!');
						window.location.reload();
					})
					.catch(err => alert('Gagal menghapus'));
				}
			}
			
			function openEditModal(btn) {
				const id = btn.getAttribute('data-id');
				const title = btn.getAttribute('data-title');
				const narration = btn.getAttribute('data-narration');
				const photo = btn.getAttribute('data-photo');
				
				document.getElementById('edit-id').value = id;
				document.getElementById('edit-title').value = title;
				document.getElementById('edit-narration').value = narration;
				document.getElementById('edit-photo-path').value = photo;
				
				document.getElementById('editModal').style.display = 'block';
			}
			
			function closeEditModal() {
				document.getElementById('editModal').style.display = 'none';
			}
			
			function saveEdit() {
				const id = document.getElementById('edit-id').value;
				const title = document.getElementById('edit-title').value;
				const narration = document.getElementById('edit-narration').value;
				const photoPath = document.getElementById('edit-photo-path').value;
				
				if (title && narration) {
					fetch('/api/gallery/' + id, {
						method: 'PUT',
						headers: { 'Content-Type': 'application/json' },
						body: JSON.stringify({ title: title, narration: narration, photoPath: photoPath })
					})
					.then(res => res.json())
					.then(data => {
						alert('Berhasil diupdate!');
						window.location.reload();
					})
					.catch(err => alert('Gagal update'));
				} else {
					alert('Judul dan Narasi harus diisi!');
				}
			}

			window.onclick = function(event) {
				const modal = document.getElementById('editModal');
				if (event.target == modal) {
					modal.style.display = 'none';
				}
			}
		</script>
	</head>
	<body>
		<div class="container">
			<h2>Kelola Galeri Foto</h2>
			
			<!-- Form Input -->
			<form action="/api/gallery" method="POST" enctype="multipart/form-data">
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
			
			<!-- Table List -->
			<table>
				<thead>
					<tr>
						<th>Foto</th>
						<th>Judul</th>
						<th>Narasi</th>
						<th>Aksi</th>
					</tr>
				</thead>
				<tbody>
					%s
				</tbody>
			</table>
			
			<a href="/admin">← Kembali ke Dashboard</a>
		</div>

		<!-- Edit Modal -->
		<div id="editModal" class="modal">
			<div class="modal-content">
				<span class="close" onclick="closeEditModal()">&times;</span>
				<h2 style="color: #F79633;">Edit Galeri</h2>
				<input type="hidden" id="edit-id">
				<div class="form-group">
					<label for="edit-title">Judul:</label>
					<input type="text" id="edit-title" required>
				</div>
				<div class="form-group">
					<label for="edit-narration">Narasi:</label>
					<textarea id="edit-narration" required></textarea>
				</div>
				<div class="form-group">
					<label for="edit-photo-path">Path Foto (URL):</label>
					<input type="text" id="edit-photo-path" readonly style="opacity: 0.5;">
					<small style="color: #aaa;">Untuk saat ini, pengeditan file foto belum didukung di modal ini.</small>
				</div>
				<button onclick="saveEdit()" style="background: #F79633; color: black;">Simpan Perubahan</button>
			</div>
		</div>
	</body>
	</html>
	`, tableRows)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func (h *HttpHandler) handleGallery(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	if r.Method == http.MethodGet {
		items, err := h.galleryUsecase.Fetch(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, items)
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

		// Upload to GCS
		filename := time.Now().Format("20060102150405") + "_" + handler.Filename
		
		bucket := h.storageClient.Bucket(h.bucketName)
		obj := bucket.Object(filename)
		
		wc := obj.NewWriter(r.Context())
		if _, err := io.Copy(wc, file); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error uploading to GCS")
			return
		}
		if err := wc.Close(); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error closing GCS writer: "+err.Error())
			return
		}

		// Get Public URL
		photoURL := "https://storage.googleapis.com/" + h.bucketName + "/" + filename

		// Save metadata to Firestore
		item := domain.GalleryItem{
			Title:     title,
			Narration: narration,
			PhotoPath: photoURL,
		}

		if err := h.galleryUsecase.Store(r.Context(), &item); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

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
					<p>Foto dan narasi Anda telah tersimpan di Firestore.</p>
					<p><a href="/admin/gallery">← Input Lagi</a> | <a href="http://localhost:5173/history" target="_blank">Lihat Halaman Sejarah →</a></p>
				</div>
			</body>
			</html>
		`))
	}
}

func (h *HttpHandler) handleGalleryDetail(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	id := parts[3]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if r.Method == http.MethodDelete {
		if err := h.galleryUsecase.Delete(r.Context(), id); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Data deleted successfully"})
		return
	}

	if r.Method == http.MethodPut {
		var req struct {
			Title     string `json:"title"`
			Narration string `json:"narration"`
			PhotoPath string `json:"photoPath"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		item := domain.GalleryItem{
			ID:        id,
			Title:     req.Title,
			Narration: req.Narration,
			PhotoPath: req.PhotoPath,
		}

		if err := h.galleryUsecase.Update(r.Context(), &item); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Data updated successfully"})
		return
	}

	respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

func (h *HttpHandler) handleMilestoneForm(w http.ResponseWriter, r *http.Request) {
	items, err := h.milestoneUsecase.Fetch(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch milestones: "+err.Error())
		return
	}

	// Build table rows safely
	tableRows := ""
	for _, item := range items {
		// Escape quotes for HTML attributes
		escapedTitle := strings.ReplaceAll(item.Title, "\"", "&quot;")
		escapedDesc := strings.ReplaceAll(item.Description, "\"", "&quot;")
		
		tableRows += fmt.Sprintf(`
			<tr>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>
					<button class="btn-edit" data-id="%s" data-year="%s" data-title="%s" data-description="%s" onclick="openEditModal(this)">Edit</button>
					<button class="btn-delete" onclick="deleteMilestone('%s')">Hapus</button>
				</td>
			</tr>
		`, item.Year, item.Title, item.Description, item.ID, item.Year, escapedTitle, escapedDesc, item.ID)
	}

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Kelola Milestone Satu Abad</title>
		<style>
			body { font-family: 'Inter', sans-serif; background: #0B301D; color: white; padding: 2rem; }
			.container { max-width: 900px; margin: 0 auto; background: rgba(255,255,255,0.05); padding: 2rem; border-radius: 12px; border: 1px solid rgba(255,255,255,0.1); }
			h2 { color: #F79633; text-align: center; }
			.form-group { margin-bottom: 1.5rem; }
			label { display: block; margin-bottom: 0.5rem; color: #00A859; font-weight: bold; }
			input[type="text"], textarea { width: 100%%; padding: 0.75rem; border-radius: 6px; border: 1px solid rgba(255,255,255,0.2); background: rgba(0,0,0,0.3); color: white; box-sizing: border-box; }
			textarea { height: 100px; resize: vertical; }
			button { background: #00A859; color: white; padding: 0.75rem 1.5rem; border: none; border-radius: 6px; cursor: pointer; font-weight: bold; }
			button:hover { background: #00c86e; }
			
			table { width: 100%%; border-collapse: collapse; margin-top: 2rem; }
			th, td { padding: 0.75rem; text-align: left; border-bottom: 1px solid rgba(255,255,255,0.1); }
			th { color: #F79633; }
			.btn-edit { background: #cca43b; margin-right: 0.5rem; color: black; }
			.btn-delete { background: #c0392b; color: white; }
			
			a { color: #F79633; text-decoration: none; display: block; text-align: center; margin-top: 1rem; }

			/* Modal Styles */
			.modal { display: none; position: fixed; z-index: 1000; left: 0; top: 0; width: 100%%; height: 100%%; background-color: rgba(0,0,0,0.8); backdrop-filter: blur(5px); }
			.modal-content { background: #0B301D; margin: 10%% auto; padding: 2rem; border: 1px solid #F79633; width: 50%%; border-radius: 12px; box-shadow: 0 5px 15px rgba(0,0,0,0.5); }
			.close { color: #aaa; float: right; font-size: 28px; font-weight: bold; cursor: pointer; }
			.close:hover { color: #F79633; }
		</style>
		<script>
			function deleteMilestone(id) {
				if (confirm('Apakah Anda yakin ingin menghapus data ini?')) {
					fetch('/api/milestones/' + id, { method: 'DELETE' })
					.then(res => res.json())
					.then(data => {
						alert('Berhasil dihapus!');
						window.location.reload();
					})
					.catch(err => alert('Gagal menghapus'));
				}
			}
			
			function openEditModal(btn) {
				const id = btn.getAttribute('data-id');
				const year = btn.getAttribute('data-year');
				const title = btn.getAttribute('data-title');
				const description = btn.getAttribute('data-description');
				
				document.getElementById('edit-id').value = id;
				document.getElementById('edit-year').value = year;
				document.getElementById('edit-title').value = title;
				document.getElementById('edit-description').value = description;
				
				document.getElementById('editModal').style.display = 'block';
			}
			
			function closeEditModal() {
				document.getElementById('editModal').style.display = 'none';
			}
			
			function saveEdit() {
				const id = document.getElementById('edit-id').value;
				const year = document.getElementById('edit-year').value;
				const title = document.getElementById('edit-title').value;
				const description = document.getElementById('edit-description').value;
				
				if (year && title && description) {
					fetch('/api/milestones/' + id, {
						method: 'PUT',
						headers: { 'Content-Type': 'application/json' },
						body: JSON.stringify({ year: year, title: title, description: description })
					})
					.then(res => res.json())
					.then(data => {
						alert('Berhasil diupdate!');
						window.location.reload();
					})
					.catch(err => alert('Gagal update'));
				} else {
					alert('Semua data harus diisi!');
				}
			}

			// Close modal when clicking outside
			window.onclick = function(event) {
				const modal = document.getElementById('editModal');
				if (event.target == modal) {
					modal.style.display = 'none';
				}
			}
		</script>
	</head>
	<body>
		<div class="container">
			<h2>Kelola Milestone Satu Abad</h2>
			
			<!-- Form Input -->
			<form action="/api/milestones" method="POST">
				<div class="form-group">
					<label for="year">Tahun:</label>
					<input type="text" id="year" name="year" placeholder="Contoh: 1926" required>
				</div>
				<div class="form-group">
					<label for="title">Judul Milestone:</label>
					<input type="text" id="title" name="title" placeholder="Contoh: Peresmian Gedung Baru" required>
				</div>
				<div class="form-group">
					<label for="description">Deskripsi:</label>
					<textarea id="description" name="description" placeholder="Ceritakan detail sejarah pada tahun tersebut..." required></textarea>
				</div>
				<button type="submit">Simpan Milestone</button>
			</form>
			
			<!-- Table List -->
			<table>
				<thead>
					<tr>
						<th>Tahun</th>
						<th>Judul</th>
						<th>Deskripsi</th>
						<th>Aksi</th>
					</tr>
				</thead>
				<tbody>
					%s
				</tbody>
			</table>
			
			<a href="/admin">← Kembali ke Dashboard</a>
		</div>

		<!-- Edit Modal -->
		<div id="editModal" class="modal">
			<div class="modal-content">
				<span class="close" onclick="closeEditModal()">&times;</span>
				<h2 style="color: #F79633;">Edit Milestone</h2>
				<input type="hidden" id="edit-id">
				<div class="form-group">
					<label for="edit-year">Tahun:</label>
					<input type="text" id="edit-year" required>
				</div>
				<div class="form-group">
					<label for="edit-title">Judul Milestone:</label>
					<input type="text" id="edit-title" required>
				</div>
				<div class="form-group">
					<label for="edit-description">Deskripsi:</label>
					<textarea id="edit-description" required></textarea>
				</div>
				<button onclick="saveEdit()" style="background: #F79633; color: black;">Simpan Perubahan</button>
			</div>
		</div>
	</body>
	</html>
	`, tableRows)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func (h *HttpHandler) handleMilestones(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	if r.Method == http.MethodGet {
		items, err := h.milestoneUsecase.Fetch(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, items)
		return
	}

	if r.Method == http.MethodPost {
		var item domain.Milestone
		
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
		} else {
			item.Year = r.FormValue("year")
			item.Title = r.FormValue("title")
			item.Description = r.FormValue("description")
		}

		if item.Year == "" || item.Title == "" {
			respondWithError(w, http.StatusBadRequest, "Year and Title are required")
			return
		}

		if err := h.milestoneUsecase.Store(r.Context(), &item); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
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
						<h2 style="color: #00A859;">Milestone Berhasil Disimpan!</h2>
						<p>Data sejarah telah tersimpan di Firestore.</p>
						<p><a href="/admin/milestones">← Input Lagi</a> | <a href="http://localhost:5173/history" target="_blank">Lihat Timeline →</a></p>
					</div>
				</body>
				</html>
			`))
			return
		}

		respondWithJSON(w, http.StatusCreated, item)
		return
	}

	respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

func (h *HttpHandler) handleMilestoneDetail(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	id := parts[3]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if r.Method == http.MethodDelete {
		if err := h.milestoneUsecase.Delete(r.Context(), id); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Milestone deleted successfully"})
		return
	}

	if r.Method == http.MethodPut {
		var req struct {
			Year        string `json:"year"`
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		item := domain.Milestone{
			ID:          id,
			Year:        req.Year,
			Title:       req.Title,
			Description: req.Description,
		}

		if err := h.milestoneUsecase.Update(r.Context(), &item); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Milestone updated successfully"})
		return
	}

	respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
}

func (h *HttpHandler) handleDashboard(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Dashboard Admin - 100 Tahun RSAS</title>
		<style>
			body { font-family: 'Inter', sans-serif; background: #0B301D; color: white; padding: 2rem; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; box-sizing: border-box; }
			.dashboard { background: rgba(255,255,255,0.05); padding: 3rem; border-radius: 16px; border: 1px solid rgba(255,255,255,0.1); width: 100%; max-width: 500px; text-align: center; box-shadow: 0 10px 30px rgba(0,0,0,0.5); }
			h1 { color: #F79633; margin-bottom: 2rem; font-size: 2rem; }
			.nav-links { display: flex; flex-direction: column; gap: 1rem; }
			.nav-btn { background: rgba(255,255,255,0.1); color: white; padding: 1rem; border-radius: 8px; text-decoration: none; font-weight: bold; transition: all 0.3s ease; border: 1px solid rgba(255,255,255,0.1); }
			.nav-btn:hover { background: #00A859; color: white; transform: translateY(-3px); box-shadow: 0 5px 15px rgba(0,168,89,0.3); border-color: #00A859; }
			.footer { margin-top: 3rem; color: rgba(255,255,255,0.5); font-size: 0.8rem; }
		</style>
	</head>
	<body>
		<div class="dashboard">
			<h1>Dashboard Admin</h1>
			<div class="nav-links">
				<a href="/admin/gallery" class="nav-btn">📸 Kelola Galeri Foto (GCS & Firestore)</a>
				<a href="/admin/milestones" class="nav-btn">⏳ Kelola Milestones (Sejarah)</a>
			</div>
			<div class="footer">
				Panel Kontrol Satu Abad RSUD Prof. Dr. H. Aloei Saboe
			</div>
		</div>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
