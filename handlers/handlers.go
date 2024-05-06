package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/secar98/converter/utils"
)

func ConvertImageHandler(w http.ResponseWriter, r *http.Request) {
	const uploadPath = "uploads/"

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileType := r.FormValue("type")
	if fileType == "" {
		http.Error(w, "No type provided", http.StatusBadRequest)
		return
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fname := uploadPath + strconv.Itoa(time.Now().Nanosecond()) + header.Filename

	log.Println("Writing file to:", fname)
	os.WriteFile(fname, bytes, 0666)

	// Define the output file path
	outputPath := uploadPath + strconv.Itoa(time.Now().Nanosecond()) + "output." + fileType // Change this to the desired output file path

	// Run the ImageMagick convert command
	cmd := exec.Command("convert", fname, outputPath)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read the converted image file
	content, err := os.ReadFile(outputPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer os.Remove(fname)
	defer os.Remove(outputPath)

	contentType := utils.HandleImageFormat(fileType)
	w.Header().Set("Content-Type", contentType)
	w.Write(content)
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	const uploadPath = "uploads/"

	convertType := r.FormValue("type")
	if convertType == "" {
		http.Error(w, "No type provided", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fname := uploadPath + strconv.Itoa(time.Now().Nanosecond()) + header.Filename

	log.Println("Writing file to:", fname)
	os.WriteFile(fname, bytes, 0666)

	cmd := exec.Command(
		"soffice",
		"--headless",
		"--convert-to",
		convertType,
		"--outdir",
		"uploads",
		fname,
	)

	log.Println("Converting file with name:", fname)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nameParts := strings.Split(fname, ".")

	content, err := os.ReadFile(nameParts[0] + "." + convertType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer os.Remove(nameParts[0] + "." + convertType)
	defer os.Remove(nameParts[0] + "." + nameParts[1])

	contentType := utils.HandleFileFormat(convertType)
	w.Header().Set("Content-Type", contentType)
	w.Write(content)
}
