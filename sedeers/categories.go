package sedeers

import (
	"fmt"
	"pustaka-api/config"
	"pustaka-api/src/models"
)

func CreateSedeerCategories() {
	db := config.InitConfig()
	var existingCategories []models.Category

	if err := db.Find(&existingCategories).Error; err != nil {
		fmt.Printf("Failed to check existing categories in the database: %s\n", err.Error())
		return
	}

	// Slice berisi daftar kategori dan subkategori
	categories := map[string][]string{
		"Umum":                   {"Informasi Umum", "Bibliografi", "Perpustakaan dan Informasi", "Ensiklopedia dan Buku yang Memuat Fakta-fakta", "Majalah dan Jurnal", "Asosiasi, Organisasi dan Museum", "Media Massa, Jurnalisme dan Publikasi", "Kutipan", "Manuskrip dan Buku Langka"},
		"Filsafat dan Psikologi": {"Filsafat dan Psikologi", "Metafisika", "Epistemologi", "Parapsikologi dan Okultisme", "Pemikiran Filosofis", "Psikologi", "Filosofis Logis", "Etika", "Filosofi Kuno, Zaman Pertengahan, dan Filosofi Ketimuran", "Filosofi Barat Modern"},
		"Agama":                  {"Agama"},
		"Sosial":                 {"Ilmu Sosial, Sosiologi dan Antropologi", "Statistik", "Ilmu Politik", "Ekonomi", "Hukum", "Administrasi Publik dan Ilmu Kemiliteran", "Masalah dan Layanan Sosial", "Pendidikan", "Perdagangan, Komunikasi, dan Transportasi", "Norma, Etika, dan Tradisi"},
		"Bahasa":                 {"Bahasa"},
		"Sains dan Matematika":   {"Sains", "Matematika", "Astronomi", "Fisika", "Kimia", "Ilmu Kebumian dan Geologi", "Fosil dan Kehidupan Prasejarah", "Biologi", "Tanaman", "Zoologi"},
		"Teknologi":              {"Teknologi", "Kesehatan dan Obat-obatan", "Teknik", "Pertanian", "Manajemen Rumah Tangga dan Keluarga", "Manajemen dan Hubungan dengan Publik", "Teknik Kimia", "Manufaktur", "Manufaktur untuk Keperluan Khusus", "Konstruksi"},
		"Seni dan Rekreasi":      {"Kesenian dan Rekreasi", "Perencanaan dan Arsitektur Lanskap", "Arsitektur", "Patung, Keramik, dan Seni Metal", "Seni Grafis dan Dekoratif", "Lukisan", "Percetakan", "Fotografi, Film, Video", "Musik", "Olahraga, Permainan, dan Hiburan"},
		"Literartur dan Sastra":  {"Literatur, Sastra, Retorika, dan Kritik"},
		"Sejarah dan Geografi":   {"Sejarah", "Geografi dan Perjalanan", "Biografi dan Asal-usul", "Sejarah Dunia Lama", "Asal-Usul Eropa", "Asal-Usul Asia", "Asal-Usul Afrika", "Asal-Usul Amerika Utara", "Asal-Usul Amerika Selatan", "Asal-Usul Wilayah Lain"},
	}

	if len(existingCategories) == 0 {
		// Iterasi melalui map dan slice untuk memasukkan data ke dalam database
		for _, subCategory := range categories {
			// Memasukkan subkategori ke dalam kategori
			for _, value := range subCategory {
				createCategory := models.Category{
					Category: value,
				}

				if err := db.Create(&createCategory).Error; err != nil {
					fmt.Printf("Failed to save subcategory %s to database: %s\n", value, err.Error())
				}
			}
		}

		fmt.Println("==> All categories created successfully <==")
	}
}
