package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []*Item

func (t *Todos) Add(task string) {
	todo := &Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}

	list[index-1].CompletedAt = time.Now()
	list[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}

	*t = append(list[:index-1], list[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("file is empty")
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{
				Align: simpletable.AlignCenter,
				Text:  "#",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Task",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Done?",
			},
			{
				Align: simpletable.AlignRight,
				Text:  "CreatedAt",
			},
			{
				Align: simpletable.AlignRight,
				Text:  "CompletedAt",
			},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}

		cells = append(cells, []*simpletable.Cell{
			{
				Align: simpletable.AlignCenter,
				Text:  fmt.Sprintf("%d", idx+1),
			},
			{
				Align: simpletable.AlignCenter,
				Text:  task,
			},
			{
				Align: simpletable.AlignCenter,
				Text:  done,
			},
			{
				Align: simpletable.AlignRight,
				Text:  item.CreatedAt.Format(time.RFC822),
			},
			{
				Align: simpletable.AlignRight,
				Text:  item.CompletedAt.Format(time.RFC822),
			},
		})
	}

	table.Body = &simpletable.Body{
		Cells: cells,
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{
				Align: simpletable.AlignCenter,
				Span:  5,
				Text:  red(fmt.Sprintf("You have %d pending todos", t.CountPending())),
			},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total += 1
		}
	}
	return total
}
