package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	truecall "vinay/truecaller"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/nyaruka/phonenumbers"
)

type Response struct {
	Data []struct {
		Name    string `json:"name,omitempty"`
		Image   string `json:"image,omitempty"`
		Address []struct {
			City string `json:"city,omitempty"`
		} `json:"addresses,omitempty"`

		Phone []struct {
			Mobile  string `json:"e164Format,omitempty"`
			Carrier string `json:"carrier,omitempty"`
		} `json:"phones,omitempty"`

		INTaddress []struct {
			Email String `json:"id,omitempty"`
		} `json:"internetaddresses,omitempty"`
	} `json:"data,omitempty"`
}

type String struct {
	IsDefined bool
	Value     string
}

// This method will be automatically invoked by json.Unmarshal
// but only for values that were provided in the json, regardless
// of whether they were null or not.
func (s *String) UnmarshalJSON(d []byte) error {
	s.IsDefined = true
	if string(d) != "null" {
		return json.Unmarshal(d, &s.Value)
	}
	return nil
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {

	re := regexp.MustCompile(`^[0-9]{10}$`)
	token_bot := goDotEnvVariable("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token_bot)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = true
	log.Printf("authorized on account %s\n", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg_from_user := update.Message.Command()
		// msg_from_user_arg := update.Message.CommandArguments()
		msg_text := update.Message.Text
		bot_user := update.Message.Chat
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if msg_from_user == "help" {
			msg.Text = "Commands: \n/start to start the bot\n/help to know how to use the bot\n/author to know who created this bot\n/me to know your own details\n"
		} else if msg_from_user == "start" {
			msg.Text = "This is a telegram bot written Golang. \nIt fetches Moile number details from Truecaller!\n\nEnter mobile number in 10 Digits format to get details\n\nRun /help to know how to use this Bot\n\nMade By Vinay Kumar \nhttps://github.com/Linuxinet/\n"
		} else if msg_from_user == "me" {
			msg.Text = fmt.Sprintf("\nFirst Name: %s\nLast Name: %s\nUsername: %s\nChat ID: %d\n", bot_user.FirstName, bot_user.LastName, bot_user.UserName, bot_user.ID)
		} else if msg_from_user == "author" {
			msg.Text = "This Bot is Created By \nhttps://github.com/Linuxinet/"
		} else if msg_text != "" {
			// regex check mobile number
			match := re.MatchString(msg_text)
			if match {
				numb, err := phonenumbers.Parse(msg_text, "IN")
				if err != nil {
					msg.Text = "Enter mobile number in 10 digit format"
				} else {
					formattedNum := phonenumbers.Format(numb, phonenumbers.NATIONAL)
					mobilenumber := strings.ReplaceAll(formattedNum, " ", "")
					body := truecall.Search_num(mobilenumber)
					var jsonbody Response
					if err := json.Unmarshal(body, &jsonbody); err != nil {
						log.Println(err)
					}

					for i, p := range jsonbody.Data {
						a := p.Name
						b := p.Phone[i].Mobile
						var c string
						if len(p.INTaddress) == 0 {
							c = ""
						} else {
							c = p.INTaddress[i].Email.Value
						}
						d := p.Phone[i].Carrier
						e := p.Address[i].City
						var f string
						if p.Image != "" {
							f = p.Image
						} else {
							f = ""
						}

						// checks the key value exist or not

						if a != "" && d != "" && e != "" {
							output := fmt.Sprintf("Name : %s\nMobile : %s\nEmail : %v\nCarrier : %s\nCity : %s\nImage : %v\n", a, b, c, d, e, f)
							msg.Text = output
						} else if a != "" && d != "" && e == "" {
							output := fmt.Sprintf("Name : %s\nMobile : %s\nEmail : %v\nCarrier : %s\n", a, b, c, d)
							msg.Text = output
						} else if a != "" && e != "" && d == "" {
							output := fmt.Sprintf("Name : %s\nMobile : %s\nEmail : %v\nCity : %s\n", a, b, c, e)
							msg.Text = output
						} else if a != "" && d == "" && e == "" {
							output := fmt.Sprintf("Name : %s\nMobile : %s\n", a, b)
							msg.Text = output
						} else {
							output := "Details not found"
							msg.Text = output
						}
					}
				}
			} else {
				msg.Text = "Enter mobile number in 10 digit format"
			}
		} else {
			msg.Text = "I dont know the command"
		}
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
