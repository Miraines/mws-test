package service

import (
	"context"
	"testing"

	"mws-test/internal/api"
	"mws-test/internal/store"
)

func TestCatService_CRUD(t *testing.T) {
	st := store.NewMemoryStore()
	svc := NewCatService(st)

	ctx := context.Background()

	// 2) Изначально ListCats должен вернуть пустой срез
	list0, err := svc.ListCats(ctx)
	if err != nil {
		t.Fatalf("ListCats returned unexpected error: %v", err)
	}
	if len(list0) != 0 {
		t.Fatalf("expected no cats initially, got %d", len(list0))
	}

	// 3) CreateCat: проверяем, что возвращается новый объект с ID=1
	newCat := api.NewCat{Name: "Whiskers", Age: 4, Color: "white"}
	created, err := svc.CreateCat(ctx, &newCat)
	if err != nil {
		t.Fatalf("CreateCat returned error: %v", err)
	}
	if created.ID == 0 {
		t.Errorf("expected non-zero ID for created Cat, got 0")
	}
	if created.Name != newCat.Name || created.Age != newCat.Age {
		t.Errorf("created Cat fields mismatch: got %+v, want %+v", created, newCat)
	}

	// 4) ListCats: теперь длина должна быть 1
	list1, err := svc.ListCats(ctx)
	if err != nil {
		t.Fatalf("ListCats returned error: %v", err)
	}
	if len(list1) != 1 {
		t.Fatalf("expected 1 cat after Create, got %d", len(list1))
	}
	if list1[0].ID != created.ID {
		t.Errorf("expected first cat ID=%d, got %d", created.ID, list1[0].ID)
	}

	// 5) GetCat: существующий ID
	getParams := api.GetCatParams{ID: created.ID}
	gotRes, err := svc.GetCat(ctx, getParams)
	if err != nil {
		t.Fatalf("GetCat returned error: %v", err)
	}
	// Поскольку cat найден, svc.GetCat возвращает *api.Cat (реализация Cat) => интерфейс типа api.GetCatRes
	gotCat, ok := gotRes.(*api.Cat)
	if !ok {
		t.Fatalf("expected *api.Cat, got %T", gotRes)
	}
	if gotCat.ID != created.ID || gotCat.Name != "Whiskers" {
		t.Errorf("unexpected cat from GetCat: got %+v, want ID=%d Name=Whiskers", gotCat, created.ID)
	}

	// 6) GetCat: несуществующий ID => возвращается api.GetCatNotFound
	_, err = svc.GetCat(ctx, api.GetCatParams{ID: created.ID + 1000})
	if err != nil {
		t.Fatalf("GetCat(non-existent) returned unexpected error: %v", err)
	}
	res404, _ := svc.GetCat(ctx, api.GetCatParams{ID: created.ID + 1000})
	if _, is404 := res404.(*api.GetCatNotFound); !is404 {
		t.Errorf("expected GetCat to return *api.GetCatNotFound for non-existent ID, got %T", res404)
	}

	// 7) UpdateCat: корректный update
	upd := api.UpdateCat{Name: "WhiskersII", Age: 5, Color: "black"}
	updatedRes, err := svc.UpdateCat(ctx, &upd, api.UpdateCatParams{ID: created.ID})
	if err != nil {
		t.Fatalf("UpdateCat returned error: %v", err)
	}
	updatedCat, ok := updatedRes.(*api.Cat)
	if !ok {
		t.Fatalf("expected *api.Cat after UpdateCat, got %T", updatedRes)
	}
	if updatedCat.Name != "WhiskersII" || updatedCat.Age != 5 {
		t.Errorf("unexpected fields after UpdateCat: got %+v, want Name=WhiskersII Age=5", updatedCat)
	}

	// 8) UpdateCat: несуществующий ID => *api.UpdateCatNotFound
	upd404 := api.UpdateCat{Name: "Nope", Age: 1, Color: "red"}
	notFoundRes, err := svc.UpdateCat(ctx, &upd404, api.UpdateCatParams{ID: created.ID + 1000})
	if err != nil {
		t.Fatalf("UpdateCat(non-existent) returned error: %v", err)
	}
	if _, isNotFound := notFoundRes.(*api.UpdateCatNotFound); !isNotFound {
		t.Errorf("expected *api.UpdateCatNotFound for update non-existent ID, got %T", notFoundRes)
	}

	// 9) DeleteCat: существующий ID => *api.DeleteCatNoContent
	delRes, err := svc.DeleteCat(ctx, api.DeleteCatParams{ID: created.ID})
	if err != nil {
		t.Fatalf("DeleteCat returned error: %v", err)
	}
	if _, isNoContent := delRes.(*api.DeleteCatNoContent); !isNoContent {
		t.Errorf("expected *api.DeleteCatNoContent, got %T", delRes)
	}

	// 10) DeleteCat: повторный вызов => *api.DeleteCatNotFound
	delRes2, err := svc.DeleteCat(ctx, api.DeleteCatParams{ID: created.ID})
	if err != nil {
		t.Fatalf("DeleteCat(second) returned error: %v", err)
	}
	if _, isNotFound := delRes2.(*api.DeleteCatNotFound); !isNotFound {
		t.Errorf("expected *api.DeleteCatNotFound for second delete, got %T", delRes2)
	}
}
