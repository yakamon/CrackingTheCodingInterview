package dogcat

import (
	"testing"
)

func TestDogCat(t *testing.T) {
	animals := []Animal{NewDog(), NewDog(), NewCat(), NewDog(), NewCat(), NewCat(), NewCat(), NewDog()}
	dogcat := New()
	for _, animal := range animals {
		dogcat.Enqueue(animal)
	}
	for _, expected := range animals {
		if animal, ok := dogcat.DequeueAny(); !ok || animal != expected {
			t.Errorf("DogCat.DequeueAny(): expected %v, got %v", expected, animal)
		}
	}

	dogs := []*Dog{NewDog(), NewDog(), NewDog()}
	cats := []*Cat{NewCat(), NewCat(), NewCat(), NewCat()}
	dogcat.Enqueue(dogs[0])
	dogcat.Enqueue(dogs[1])
	dogcat.Enqueue(cats[0])
	dogcat.Enqueue(cats[1])
	dogcat.Enqueue(cats[2])
	dogcat.Enqueue(dogs[2])
	dogcat.Enqueue(cats[3])
	for _, expected := range dogs {
		if dog, ok := dogcat.DequeueDog(); !ok || dog != expected {
			t.Errorf("DogCat.DequeueDog(): expected %v, got %v", expected, dog)
		}
	}
	for _, expected := range cats {
		if cat, ok := dogcat.DequeueCat(); !ok || cat != expected {
			t.Errorf("DogCat.DequeueCat(): expected %v, got %v", expected, cat)
		}
	}
}
