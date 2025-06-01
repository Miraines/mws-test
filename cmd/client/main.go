package main

import (
	"context"
	"fmt"
	"log"
	"mws-test/internal/api"
)

func main() {
	cli, err := api.NewClient("http://localhost:8000")
	if err != nil {
		log.Fatalf("create client: %v", err)
	}

	ctx := context.Background()

	newCat := api.NewCat{
		Name:  "Барсик",
		Age:   3,
		Color: "рыжий",
	}

	cat, err := cli.CreateCat(ctx, &newCat)
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	fmt.Printf("Создали: %+v\n", cat)

	cats, err := cli.ListCats(ctx)
	if err != nil {
		log.Fatalf("list: %v", err)
	}
	fmt.Printf("List -> %+v\n", cats)

	upd := api.UpdateCat{
		Name:  "Барсик II",
		Age:   4,
		Color: "белый",
	}

	updated, err := cli.UpdateCat(ctx, &upd, api.UpdateCatParams{ID: cat.ID})
	if err != nil {
		log.Fatalf("update: %v", err)
	}
	fmt.Printf("Обновили: %+v\n", updated)

	cats, _ = cli.ListCats(ctx)
	fmt.Printf("After update -> %+v\n", cats)

	res, err := cli.DeleteCat(ctx, api.DeleteCatParams{ID: cat.ID})
	if err != nil {
		log.Fatalf("delete request failed: %v", err)
	}

	switch res.(type) {
	case *api.DeleteCatNoContent:
		fmt.Println("Удалили котика (204)")
	case *api.DeleteCatNotFound:
		fmt.Println("Котик не найден (404)")
	default:
		fmt.Printf("Неизвестный ответ: %#v\n", res)
	}
}
