package handlers

import (
	"student-api/internal/models"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// Hazırlık: Test verisi
	testData := []models.Student{
		{ID: 1, Name: "Test Ali", Grade: 100},
	}

	// 1. Kaydetmeyi test et
	err := models.SaveToFile(testData)
	if err != nil {
		t.Errorf("Kaydetme hatası: %v", err)
	}

	// 2. Okumayı test et
	loadedData, err := models.LoadFromFile()
	if err != nil {
		t.Errorf("Okuma hatası: %v", err)
	}

	// 3. Veri doğru mu? (Basit bir kontrol)
	if len(loadedData) != 1 || loadedData[0].Name != "Test Ali" {
		t.Error("Yüklenen veri kaydedilenle eşleşmiyor!")
	}
}
