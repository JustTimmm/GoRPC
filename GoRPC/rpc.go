package GoRPC

import (
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"time"
)

type RPC struct {
	ClientID, Details, State      string
	LargeImageKey, LargeImageText string
	SmallImageKey, SmallImageText string
	Button1Label, Button1URL      string
	Button2Label, Button2URL      string
}

func StartRPC(rpc RPC) {
	err := client.Login(rpc.ClientID)
	if err != nil {
		panic(err)
	}

	var buttons []*client.Button
	buttons = appendIfValid(buttons, rpc.Button1Label, rpc.Button1URL)
	buttons = appendIfValid(buttons, rpc.Button2Label, rpc.Button2URL)

	now := time.Now()
	err = client.SetActivity(client.Activity{
		Details:    rpc.Details,
		State:      rpc.State,
		LargeImage: rpc.LargeImageKey,
		LargeText:  rpc.LargeImageText,
		SmallImage: rpc.SmallImageKey,
		SmallText:  rpc.SmallImageText,
		Timestamps: &client.Timestamps{Start: &now},
		Buttons:    buttons,
	})

	if err != nil {
		fmt.Println("Erreur SetActivity:", err)
	}

	time.Sleep(time.Second * 10)
}
