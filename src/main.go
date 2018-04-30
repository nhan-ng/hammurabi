package main

import (
	"log"

	"./cmd"
)

const (
	intro = `
Congratulations, you are the newest ruler of ancient Samaria, elected for a ten year term of office. Your duties are to dispsense food, direct farming, and buy and sell land as needed to support your people. Watch out for rat infestations and the plague! Gain is the general currency, measured in bushels. The following will help you in your decisions:

- Each person needs at least 20 bushels of grain per year to survive.
- Each person can farm at most 10 acres of land.
- It takes 1 bushel of grain to farm an acre of land.
- The mark price for land fluctuates yearly.

Rule wisely and you will be showered with appreciation at the end of your term. Rule poorly and you will be kicked out of office!
	`
)

func main() {
	if err := cmd.NewHammurabiCmd().Execute(); err != nil {
		log.Println(err)
	}
}
