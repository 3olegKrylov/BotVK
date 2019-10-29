package main

import (
	"github.com/dimonchik0036/vk-api"
	"log"
)

func main() {
	//client, err := vkapi.NewClientFromLogin("89219850412", "prioritet", vkapi.ScopeMessages)
	client, err := vkapi.NewClientFromToken("f8195384c2cd1b11c19a0d99388bd43e63db3e771bb32223f2c11047337f343726b7777e6b061901b6292")
	if err != nil {
		log.Panic(err)
	}

	client.Log(true)

	if err := client.InitLongPoll(0, 2); err != nil {
		log.Panic(err)
	}

	updates, _, err := client.GetLPUpdatesChan(100, vkapi.LPConfig{25, vkapi.LPModeAttachments})
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil || !update.IsNewMessage() || update.Message.Outbox(){
			continue
		}

		log.Printf("%s", update.Message.String())
		if update.Message.Text == "/info" {
			client.SendMessage(vkapi.NewMessage(vkapi.NewDstFromUserID(update.Message.FromID), "Привет, с вами на линии бот Олега, если хотите,(информцию обо мне /info)"))
		}
	}
}
