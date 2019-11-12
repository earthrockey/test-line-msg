package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/callback", Callback)

	http.ListenAndServe(getPort(), nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}

func main() {
	db := DbConn()
	defer db.Close()
	HandleRequest()
}

func DbConn() (db *gorm.DB) {
	// db, err := gorm.Open("mysql", "root:e575g73wk@/gosensorproject?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:Pid1ajYUDJ@tcp(128.199.216.169:3306)/gosensorproject?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "root:Pid1ajYUDJ@tcp(172.17.0.2:3306)/gosensorproject?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func connectLineBot() *linebot.Client {
	bot, err := linebot.New("62b1c4aba0864c7f65b08d715cebbf3b", "FvOYyw81BuV4bgMxhcer9g0ebeUCp4W/jj0+EYtkep6zx08o3Owm8ptGoeiJF8SxVoOVdct1qd62S0+TgPlgS+LjilQXYDv3zqGrIGOUjnIgkm1F62zS5X+DEXGdsaIOX2VTtd8x3BWA+IJZ6SV96AdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func Callback(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	bot := connectLineBot()
	events, err := bot.ParseRequest(r)
	if err != nil {
		fmt.Println(err)
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			profile, err := bot.GetProfile(event.Source.UserID).Do()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("EventTypeMessage")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
			fmt.Println("Profile : " + profile.DisplayName)
			bot.PushMessage(userID, linebot.NewTextMessage("Display name: "+profile.DisplayName)).Do()
		}
		if event.Type == linebot.EventTypeJoin {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Println("EventTypeJoin")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
		}
		if event.Type == linebot.EventTypeLeave {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Println("EventTypeLeave")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
		}
		if event.Type == linebot.EventTypeFollow {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Println("EventTypeFollow")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
		}
		if event.Type == linebot.EventTypeUnfollow {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Println("EventTypeUnfollow")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
		}
		if event.Type == linebot.EventTypePostback {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Println("EventTypePostback")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
		}
		if event.Type == linebot.EventTypeBeacon {
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Println("EventTypeBeacon")
			fmt.Println("userID : " + userID)
			fmt.Println("groupID : " + groupID)
			fmt.Println("RoomID : " + RoomID)
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
