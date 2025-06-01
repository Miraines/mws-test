package store

import (
	"testing"

	"mws-test/internal/api"
)

func TestMemoryStore_CRUD(t *testing.T) {
	st := NewMemoryStore()

	// 1) Изначально List() должен вернуть пустой срез
	initial := st.List()
	if len(initial) != 0 {
		t.Fatalf("expected empty list, got %d items", len(initial))
	}

	// 2) Create: создаём новую сущность
	newCat := api.NewCat{
		Name:  "Kitty",
		Age:   2,
		Color: "gray",
	}
	created := st.Create(newCat)
	if created.ID == 0 {
		t.Errorf("expected non-zero ID after Create, got 0")
	}
	if created.Name != newCat.Name || created.Age != newCat.Age || created.Color != newCat.Color {
		t.Errorf("created fields mismatch: got %+v, want %+v", created, newCat)
	}

	// 3) После Create у нас должен быть ровно один элемент в List()
	listAfterCreate := st.List()
	if len(listAfterCreate) != 1 {
		t.Fatalf("expected 1 cat after Create, got %d", len(listAfterCreate))
	}

	// 4) Get: проверяем, что можно достать по ID
	got, ok := st.Get(created.ID)
	if !ok {
		t.Fatalf("expected to find cat with ID %d, but OK=false", created.ID)
	}
	if got.ID != created.ID || got.Name != "Kitty" {
		t.Errorf("unexpected cat returned by Get: got %+v, want ID=%d Name=Kitty", got, created.ID)
	}

	// 5) Update: изменяем поля
	upd := api.UpdateCat{
		Name:  "KittyUpdated",
		Age:   3,
		Color: "black",
	}
	updated, ok := st.Update(created.ID, upd)
	if !ok {
		t.Fatalf("expected Update to succeed for ID=%d", created.ID)
	}
	if updated.ID != created.ID || updated.Name != "KittyUpdated" || updated.Age != 3 || updated.Color != "black" {
		t.Errorf("unexpected cat after Update: got %+v, want ID=%d Name=KittyUpdated Age=3 Color=black",
			updated, created.ID)
	}

	// 6) Попробуем Update несуществующего ID
	if _, ok := st.Update(created.ID+1000, upd); ok {
		t.Errorf("expected Update to fail for non-existent ID, but ok=true")
	}

	// 7) Delete: удаляем созданный элемент
	deleted := st.Delete(created.ID)
	if !deleted {
		t.Fatalf("expected Delete to return true for existing ID=%d, got false", created.ID)
	}
	// после удаления должен быть пустой список
	afterDel := st.List()
	if len(afterDel) != 0 {
		t.Errorf("expected 0 cats after Delete, got %d", len(afterDel))
	}

	// 8) Delete ещё раз того же ID — должна вернуться false
	if st.Delete(created.ID) {
		t.Errorf("expected Delete to return false for already-deleted ID, got true")
	}

	// 9) Get для удалённого ID — OK=false
	if _, ok := st.Get(created.ID); ok {
		t.Errorf("expected Get to return ok=false for deleted ID, but ok=true")
	}
}
