package main

import (
	"context"
	"db/ent"
	"db/ent/migrate"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	if err := client.Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatal("failed creating schema resources: &v", err)
	}

	userRepository := UserRepository{Client: client.User}
	tourProductRepository := TourProductRepository{Client: client.TourProduct}

	user1 := client.User.Create().SetID("jmj005").SetName("장민주").SaveX(context.TODO())
	user2 := client.User.Create().SetID("dorlltb").SetName("개발자").SaveX(context.TODO())

	client.TourProduct.Create().SetName("미국 여행").SetManager(user1).SetPrice(20000000).SaveX(context.TODO())
	client.TourProduct.Create().SetName("유럽 여행").SetManager(user2).SetPrice(40000000).SaveX(context.TODO())

	fmt.Println("전체 유저 조회")
	for _, user := range userRepository.FindAll() {
		fmt.Printf("User(id=%s, name=%s)\n", user.ID, user.Name)
	}
	fmt.Println("------------------------------------------------------")

	fmt.Println("전체 여행 상품 조회")
	for _, tour := range tourProductRepository.FindAll() {
		fmt.Printf("TourProduct(id=%d, name=%s, manager=%s)\n", tour.ID, tour.Name, tour.Edges.Manager.ID)
	}
	fmt.Println("------------------------------------------------------")

	fmt.Println(user1.ID + "가 관리하는 전체 여행 상품 조회")
	for _, tour := range tourProductRepository.FindAllByManager(user1.ID) {
		fmt.Printf("TourProduct(id=%d, name=%s, manager=%s)\n", tour.ID, tour.Name, tour.Edges.Manager.ID)
	}
}
