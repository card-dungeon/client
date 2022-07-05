package main

import (
	"context"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cardpb "github.com/card-dungeon/client/src/protos"
	"github.com/card-dungeon/client/src/setting"
)

func main() {
	rl.InitWindow(setting.WINDOW_WIDTH, setting.WINDOW_HEIGHT, "Card Dungeon")

	rl.SetTargetFPS(60)

	conn, err := grpc.Dial("card-dungeon.myroom.dev:443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("asd")
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	client := cardpb.NewCardClient(conn)
	charCard, err := client.GetCharacterCard(context.Background(), &cardpb.CardId{CardId: 1})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(charCard)

	charCardList, err := client.GetCharacterCardList(context.Background(), &cardpb.GetList{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(charCardList.CharacterCards[0].Desc)

	// for !rl.WindowShouldClose() {
	// 	rl.BeginDrawing()

	// 	rl.ClearBackground(rl.RayWhite)

	// 	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	// 	rl.EndDrawing()
	// }

	// rl.CloseWindow()
}
