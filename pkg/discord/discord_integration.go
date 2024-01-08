package discord

import (
	"fmt"
	"sort"

	"github.com/bwmarrin/discordgo"
	"github.com/themaluxis/accdiscordbot"
	"github.com/themaluxis/accdiscordbot/pkg/data"
	"github.com/themaluxis/accdiscordbot/pkg/utils"
)

func BestLapMessages(cfg accdiscordbot.Config, s *discordgo.Session, groupedBestLaps map[string]map[int][]data.BestLapInfo, carList map[int]string, trackName map[string]string, trackImage map[string]string, previousBestLaps map[string]int, carBestLaps map[string]map[int]int) {
	var tracks []string
	for track := range groupedBestLaps {
		tracks = append(tracks, track)
	}
	sort.Strings(tracks)

	for _, track := range tracks {
		bestLap := int(^uint(0) >> 1) // Initialize with max int value
		bestDriver := ""
		for _, laps := range groupedBestLaps[track] {
			for _, lap := range laps {
				if lap.BestLap < bestLap {
					bestLap = lap.BestLap
					bestDriver = lap.DriverID
				}
			}
		}

		trackNameFormatted := trackName[track]
		trackImageURL := trackImage[track]
		bestLapFormatted := utils.FormatTime(bestLap)

		// Check if the best lap is an improvement
		if previousBest, exists := previousBestLaps[track]; exists && bestLap < previousBest {
			newLapMessage := fmt.Sprintf("@everyone | **%s** a amélioré le meilleur temps sur **%s** ! Nouveau temps : **%s**", bestDriver, trackNameFormatted, bestLapFormatted)
			s.ChannelMessageSend(cfg.Discord.ChannelID_Feed, newLapMessage)
		}
		previousBestLaps[track] = bestLap

		for carModel, laps := range groupedBestLaps[track] {
			carBestLap := int(^uint(0) >> 1)
			var bestDriverForCar string
			for _, lap := range laps {
				if lap.BestLap < carBestLap {
					carBestLap = lap.BestLap
					bestDriverForCar = lap.DriverID
				}
			}

			if carPreviousBestLaps, exists := carBestLaps[track]; exists {
				if previousBest, ok := carPreviousBestLaps[carModel]; !ok || carBestLap < previousBest {
					carBestLapFormatted := utils.FormatTime(carBestLap)
					carName := carList[carModel]
					newLapMessage := fmt.Sprintf("**%s** a amélioré le meilleur temps pour **%s** sur **%s** ! Nouveau temps : **%s**", bestDriverForCar, carName, trackNameFormatted, carBestLapFormatted)
					s.ChannelMessageSend(cfg.Discord.ChannelID_Feed, newLapMessage)
				}
			} else {
				carBestLaps[track] = make(map[int]int)
			}

			carBestLaps[track][carModel] = carBestLap
		}

		// existing embed creation and updating code...
		embed := &discordgo.MessageEmbed{
			Title: fmt.Sprintf("Meilleurs temps sur %s", trackNameFormatted),
			Color: 15425844,
			Image: &discordgo.MessageEmbedImage{
				URL: trackImageURL,
			},
			Fields: []*discordgo.MessageEmbedField{},
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("Meilleur tour : %s en %s", bestDriver, bestLapFormatted),
			},
		}

		carFields := []*discordgo.MessageEmbedField{}
		carCount := 0

		for carModel, laps := range groupedBestLaps[track] {
			carName := carList[carModel]
			value := ""
			for _, lap := range laps {
				value += fmt.Sprintf("%s: %s\n", lap.DriverID, utils.FormatTime(lap.BestLap))
			}
			if carCount == 2 {
				carFields = append(carFields, &discordgo.MessageEmbedField{
					Name:   "\u200b",
					Value:  "\u200b",
					Inline: false,
				})
				carCount = 0
			}
			carFields = append(carFields, &discordgo.MessageEmbedField{
				Name:   carName,
				Value:  value,
				Inline: true, // Display two cars per row
			})
			carCount++
		}

		// Add carFields to the embed.Fields
		embed.Fields = append(embed.Fields, carFields...)
		fmt.Printf("Tentative d'envoi de l'embed dans le canal : %s\n", cfg.Discord.ChannelID_Chronos)
		if msgID, exists := utils.MessageIDs[track]; exists {
			fmt.Printf("Tentative de modification du message avec ID : %s\n", msgID)
			msg, err := s.ChannelMessageEditEmbed(cfg.Discord.ChannelID_Chronos, msgID, embed)
			if err != nil {
				fmt.Printf("Erreur lors de la mise à jour de l'embed. Canal : %s, Message ID : %s, Erreur : %s\n", cfg.Discord.ChannelID_Chronos, msgID, err)
			} else {
				fmt.Printf("Message mis à jour avec succès. Nouvel ID de message : %v\n", msg.ID)
			}
		} else {
			msg, err := s.ChannelMessageSendEmbed(cfg.Discord.ChannelID_Chronos, embed)
			utils.MessageIDs[track] = msg.ID
			if err != nil {
				fmt.Printf("Erreur lors de l'envoi de l'embed. Canal : %s, Erreur : %s\n", cfg.Discord.ChannelID_Chronos, err)
			} else {
				fmt.Printf("Embed envoyé avec succès. ID de message : %v\n", msg.ID)
			}
		}
	}
}

func TrackLeaderboardMessages(cfg accdiscordbot.Config, s *discordgo.Session, groupedBestLaps map[string]map[int][]data.BestLapInfo, carList map[int]string, trackName map[string]string, trackImage map[string]string) {
	var tracks []string
	for track := range groupedBestLaps {
		tracks = append(tracks, track)
	}
	sort.Strings(tracks)

	for _, track := range tracks {
		trackNameFormatted := trackName[track]
		trackImageURL := trackImage[track]

		// Aggregate best laps per driver
		bestLapsPerDriver := make(map[string]int)
		for _, lapsByCar := range groupedBestLaps[track] {
			for _, lap := range lapsByCar {
				if currentBest, exists := bestLapsPerDriver[lap.DriverID]; !exists || lap.BestLap < currentBest {
					bestLapsPerDriver[lap.DriverID] = lap.BestLap
				}
			}
		}

		// Convert to a slice and sort
		type driverLap struct {
			DriverID string
			BestLap  int
		}
		var sortedLaps []driverLap
		for driver, lap := range bestLapsPerDriver {
			sortedLaps = append(sortedLaps, driverLap{DriverID: driver, BestLap: lap})
		}
		sort.Slice(sortedLaps, func(i, j int) bool {
			return sortedLaps[i].BestLap < sortedLaps[j].BestLap
		})

		// Prepare embed fields
		var embedFields []*discordgo.MessageEmbedField
		for i, driverLap := range sortedLaps {
			formattedTime := utils.FormatTime(driverLap.BestLap)
			position := fmt.Sprintf("%d. %s", i+1, driverLap.DriverID) // Adding position
			field := &discordgo.MessageEmbedField{
				Name:   position,
				Value:  formattedTime,
				Inline: false,
			}
			embedFields = append(embedFields, field)
		}

		// Create and send embed
		embed := &discordgo.MessageEmbed{
			Title:  fmt.Sprintf("Classement sur %s", trackNameFormatted),
			Color:  3386879,
			Image:  &discordgo.MessageEmbedImage{URL: trackImageURL},
			Fields: embedFields,
		}

		if msgID, exists := utils.MessageIDs["classement_"+track]; exists {
			msg, err := s.ChannelMessageEditEmbed(cfg.Discord.ChannelID_Leaderboard, msgID, embed)
			if err != nil {
				fmt.Printf("Erreur lors de la mise à jour de l'embed (Classement : %s): %s\n", cfg.Discord.ChannelID_Leaderboard, err)
			} else {
				fmt.Printf("Mise a jour de l'embed %v\n", msg.ID)
			}
		} else {
			msg, err := s.ChannelMessageSendEmbed(cfg.Discord.ChannelID_Leaderboard, embed)
			if err != nil {
				fmt.Printf("Erreur lors de l'envoi de l'embed (Classement : %s): %s\n", cfg.Discord.ChannelID_Leaderboard, err)
			} else {
				utils.MessageIDs["classement_"+track] = msg.ID
				fmt.Printf("Embed crée avec l'ID %v\n", msg.ID)
				err := utils.SaveMessageIDsToFile(utils.MessageIDs, "messageIDs.json")
				if err != nil {
					fmt.Printf("Erreur lors de la sauvegarde de l'ID de message dans messageIDs.json")
				}
			}
		}
	}
}

func TrackGeneralLeaderboard(cfg accdiscordbot.Config, s *discordgo.Session, groupedBestLaps map[string]map[int][]data.BestLapInfo, carList map[int]string) {
	// Step 1: Track Best Position Per Track for Each Driver
	driverBestPositions := make(map[string]map[string]int) // Map: DriverID -> (Track -> Best Position)
	for track, lapsByCar := range groupedBestLaps {
		for _, laps := range lapsByCar {
			sort.SliceStable(laps, func(i, j int) bool {
				return laps[i].BestLap < laps[j].BestLap
			})
			for position, lap := range laps {
				if position < 5 { // Only consider top 5 positions
					if driverBestPositions[lap.DriverID] == nil {
						driverBestPositions[lap.DriverID] = make(map[string]int)
					}
					if pos, exists := driverBestPositions[lap.DriverID][track]; !exists || position < pos {
						driverBestPositions[lap.DriverID][track] = position
					}
				}
			}
		}
	}

	// Step 2: Update Points Calculation Logic
	driverPoints := make(map[string]int)
	for driverID, trackPositions := range driverBestPositions {
		for _, position := range trackPositions {
			switch position {
			case 0:
				driverPoints[driverID] += 10
			case 1:
				driverPoints[driverID] += 8
			case 2:
				driverPoints[driverID] += 6
			case 3:
				driverPoints[driverID] += 4
			case 4:
				driverPoints[driverID] += 3
			}
		}
	}

	// Step 2: Create General Leaderboard
	type driverPointsInfo struct {
		DriverID string
		Points   int
	}
	var leaderboard []driverPointsInfo
	for driver, points := range driverPoints {
		leaderboard = append(leaderboard, driverPointsInfo{DriverID: driver, Points: points / 8})
	}
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Points > leaderboard[j].Points
	})

	// Step 3: Prepare Embed Message
	var embedFields []*discordgo.MessageEmbedField
	for i, driverInfo := range leaderboard {
		position := fmt.Sprintf("%d. %s", i+1, driverInfo.DriverID)
		points := fmt.Sprintf("%d points", driverInfo.Points)
		field := &discordgo.MessageEmbedField{
			Name:   position,
			Value:  points,
			Inline: false,
		}
		embedFields = append(embedFields, field)
	}

	// Step 4: Send Embed Message
	embed := &discordgo.MessageEmbed{
		Title:  "Classement général",
		Color:  3386879, // Example color
		Fields: embedFields,
	}

	if msgID, exists := utils.MessageIDs["classement_general"]; exists {
		msg, err := s.ChannelMessageEditEmbed(cfg.Discord.ChannelID_LeaderboardGeneral, msgID, embed)
		if err != nil {
			fmt.Printf("Erreur lors de la mise à jour de l'embed (General : %s): %s\n", cfg.Discord.ChannelID_LeaderboardGeneral, err)
		} else {
			fmt.Printf("Mise a jour de l'embed %v\n", msg.ID)
		}
	} else {
		msg, err := s.ChannelMessageSendEmbed(cfg.Discord.ChannelID_LeaderboardGeneral, embed)
		if err != nil {
			fmt.Printf("Erreur lors de l'envoi de l'embed (General : %s): %s\n", cfg.Discord.ChannelID_LeaderboardGeneral, err)
		} else {
			utils.MessageIDs["classement_general"] = msg.ID
			fmt.Printf("Embed crée avec l'ID %v\n", msg.ID)
			err := utils.SaveMessageIDsToFile(utils.MessageIDs, "messageIDs.json")
			if err != nil {
				fmt.Printf("Erreur lors de la sauvegarde de l'ID de message dans messageIDs.json")
			}
		}
	}
}
