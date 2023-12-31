package discord

import (
	"fmt"
	"sort"

	"github.com/bwmarrin/discordgo"
	"github.com/themaluxis/accdiscordbot"
	"github.com/themaluxis/accdiscordbot/pkg/data"
	"github.com/themaluxis/accdiscordbot/pkg/utils"
)

func BestLapMessages(cfg accdiscordbot.Config, s *discordgo.Session, groupedBestLaps map[string]map[int][]data.BestLapInfo, carList map[int]string, trackName map[string]string, trackImage map[string]string) {
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
		if msgID, exists := utils.MessageIDs[track]; exists {
			msg, err := s.ChannelMessageEditEmbed(cfg.Discord.ChannelID_Chronos, msgID, embed)
			if err != nil {
				fmt.Printf("Erreur lors de la mise à jour de l'embed: %s\n", err)
			} else {
				fmt.Printf("Mise a jour de l'embed %v\n", msg.ID)
			}
		} else {
			msg, err := s.ChannelMessageEditEmbed(cfg.Discord.ChannelID_Chronos, msgID, embed)
			if err != nil {
				fmt.Printf("Erreur lors de l'envoi de l'embed: %s\n", err)
			} else {
				utils.MessageIDs[track] = msg.ID
				fmt.Printf("Embed crée avec l'ID %v\n", msg.ID)
				err := utils.SaveMessageIDsToFile(utils.MessageIDs, "messageIDs.json")
				if err != nil {
					fmt.Printf("Erreur lors de la sauvegarde de l'ID de message dans messageIDs.json")
				}
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
				Inline: true,
			}
			embedFields = append(embedFields, field)
		}

		// Create and send embed
		embed := &discordgo.MessageEmbed{
			Title:  fmt.Sprintf("Classement sur %s", trackNameFormatted),
			Color:  15425844,
			Image:  &discordgo.MessageEmbedImage{URL: trackImageURL},
			Fields: embedFields,
		}

		if msgID, exists := utils.MessageIDs["classement_"+track]; exists {
			msg, err := s.ChannelMessageEditEmbed(cfg.Discord.ChannelID_Leaderboard, msgID, embed)
			if err != nil {
				fmt.Printf("Erreur lors de la mise à jour de l'embed: %s\n", err)
			} else {
				fmt.Printf("Mise a jour de l'embed %v\n", msg.ID)
			}
		} else {
			msg, err := s.ChannelMessageSendEmbed(cfg.Discord.ChannelID_Leaderboard, embed)
			if err != nil {
				fmt.Printf("Erreur lors de l'envoi de l'embed: %s\n", err)
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
