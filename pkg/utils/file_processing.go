package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/themaluxis/accdiscordbot/pkg/data"
)

func ParcourirDossier(dossier string, groupedBestLaps map[string]map[int][]data.BestLapInfo, carList map[int]string) {
	fichiers, err := os.ReadDir(dossier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier :", err)
		return
	}

	for _, fichier := range fichiers {
		cheminFichier := filepath.Join(dossier, fichier.Name())

		if fichier.IsDir() {
			// Assuming that the 'results' directory is what you want to process
			if fichier.Name() == "results" {
				parcourirDossierResults(cheminFichier, groupedBestLaps, carList)
			} else {
				// Recursively process other directories
				ParcourirDossier(cheminFichier, groupedBestLaps, carList)
			}
		}
		// Add additional conditions here if you need to process specific files
	}
}

func parcourirDossierResults(dossierResults string, groupedBestLaps map[string]map[int][]data.BestLapInfo, carList map[int]string) {
	fichiers, err := os.ReadDir(dossierResults)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier 'results' :", err)
		return
	}

	for _, fichier := range fichiers {
		cheminFichier := filepath.Join(dossierResults, fichier.Name())

		if strings.HasSuffix(fichier.Name(), ".json") {
			file, err := os.Open(cheminFichier)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier", fichier.Name(), ":", err)
				continue
			}

			cleanedReader, err := convertToUTF8AndClean(file)
			if err != nil {
				fmt.Println("Erreur lors de la conversion et nettoyage pour le fichier", fichier.Name(), ":", err)
				file.Close()
				continue
			}

			var trackData data.TrackData
			decoder := json.NewDecoder(cleanedReader)
			if err := decoder.Decode(&trackData); err != nil {
				fmt.Println("Erreur lors du décodage du JSON pour le fichier", fichier.Name(), ":", err)
				file.Close()
				continue
			}

			file.Close()

			for _, line := range trackData.SessionResult.LeaderBoardLines {
				if line.Timing.BestLap < 600000 { // Check if lap time is less than 10 minutes
					driver := line.Car.Drivers[0]
					if !data.PlayerIDWhitelist(driver.PlayerId) {
						continue
					}

					driverID := driver.FirstName + " " + driver.LastName
					carName, exists := carList[line.Car.CarModel]
					if !exists {
						carName = "Unknown Car Model"
					}

					bestLap := data.BestLapInfo{
						TrackName: trackData.TrackName,
						CarModel:  line.Car.CarModel,
						CarName:   carName,
						BestLap:   line.Timing.BestLap,
						DriverID:  driverID,
					}

					// key := fmt.Sprintf("%s_%s_%d", trackData.TrackName, driverID, line.Car.CarModel)
					UpdateBestLap(groupedBestLaps, trackData.TrackName, bestLap)
				}
			}
		}
	}
}

func SaveMessageIDsToFile(messageIDs map[string]string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(messageIDs); err != nil {
		return err
	}

	return nil
}

func LoadMessageIDsFromFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, err
	}
	defer file.Close()

	var messageIDs map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&messageIDs); err != nil {
		return nil, err
	}

	return messageIDs, nil
}
