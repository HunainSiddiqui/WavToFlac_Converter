package tests

import (
	"os"
	"task-golang/controller"

	"testing"
)

func TestConvertWAVToFLAC(t *testing.T) {
	inputWAVFilePath := "test.wav"
	outputFLACFilePath := "test.flac"

	wavData, err := os.ReadFile(inputWAVFilePath)
	if err != nil {
		t.Fatalf("Failed to read WAV file: %v", err)
	}

	flacData, err := controller.ConvertWAVToFLAC(wavData)
	if err != nil {
		t.Fatalf("Failed to convert WAV to FLAC: %v", err)
	}

	err = os.WriteFile(outputFLACFilePath, flacData, 0644)
	if err != nil {
		t.Fatalf("Failed to write FLAC file: %v", err)
	}

	fileInfo, err := os.Stat(outputFLACFilePath)
	if err != nil || fileInfo.Size() == 0 {
		t.Fatal("FLAC file was not created or is empty after conversion")
	} else {
		t.Log("FLAC file was successfully created with data.")
	}

	defer os.Remove(outputFLACFilePath)
}
