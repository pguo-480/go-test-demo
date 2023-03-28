package pkg

import (
	"errors"
	"fmt"
	"io"
)

//go:generate mockery --inpackage --name=Magic
type Magic interface {
	BuildMagicBook(message int) (*magicBook, error)
	Magic(book *magicBook) error
}

type magic struct {
	db Db
}

func NewMagic(db Db) Magic {
	return &magic{
		db: db,
	}
}

type magicBook struct {
	filed1  int
	filed2  bool
	filed3  string
	filed4  float64
	filed5  []byte
	filed6  interface{}
	filed7  []interface{}
	filed8  <-chan int
	filed9  chan<- string
	filed10 io.Reader
}

func (m *magic) BuildMagicBook(message int) (*magicBook, error) {
	if message == 1 {
		return nil, fmt.Errorf("BuildMagicBook error %w", ErrorUnrecoverable)
	}
	// 很多你无法理解的代码
	// call other functions * n
	// 要花30分钟理解的长算法
	// error cases * n
	// logging to datadog ..
	// 很多你无法理解的代码
	return &magicBook{
		filed1: message,
		// magic....
	}, nil
}

func (m *magic) Magic(book *magicBook) error {
	// magic.... * 100 lines

	err := m.db.Put(book.filed1)
	if err != nil {
		if errors.Is(err, ErrorDb409) {
			return fmt.Errorf("Magic db.Put error %w", ErrorUnrecoverable)
		}
		return fmt.Errorf("Magic db.Put error %w", ErrorRecoverable)
	}

	return nil
}
