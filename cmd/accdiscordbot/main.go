package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/themaluxis/accdiscordbot"
	"github.com/themaluxis/accdiscordbot/pkg/data"
	"github.com/themaluxis/accdiscordbot/pkg/discord"
	"github.com/themaluxis/accdiscordbot/pkg/utils"
)

func main() {
	var Cfg accdiscordbot.Config
	accdiscordbot.LoadConfig(&Cfg)
	dossier := Cfg.General.AccServerPath
	groupedBestLaps := make(map[string]map[int][]data.BestLapInfo)
	carList := data.GetCarList()
	trackName := data.GetTrackName()
	trackImage := data.GetTrackImage()

	dg, err := discordgo.New("Bot " + Cfg.Discord.BotToken)
	if err != nil {
		fmt.Println("Erreur lors de la création de la session Discord,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la connexion,", err)
		return
	}

	utils.MessageIDs, err = utils.LoadMessageIDsFromFile("messageIDs.json")
	if err != nil {
		fmt.Println("Error loading message IDs:", err)
		return
	}

	// Boucle pour mettre à jour les messages toutes les 10 minutes
	for {
		utils.ParcourirDossier(dossier, groupedBestLaps, carList)

		var tracks []string
		for track := range groupedBestLaps {
			tracks = append(tracks, track)
		}
		sort.Strings(tracks)

		discord.BestLapMessages(Cfg, dg, groupedBestLaps, carList, trackName, trackImage)
		discord.TrackLeaderboardMessages(Cfg, dg, groupedBestLaps, carList, trackName, trackImage)

		// Attendre 10 minutes
		time.Sleep(10 * time.Minute)
	}

}
