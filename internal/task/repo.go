package task

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

var ErrNotFound = errors.New("task not found")

type Repo struct {
	mu       sync.RWMutex
	seq      int64
	items    map[int64]*Task
	filePath string
}

func NewRepo() *Repo {
	repo := &Repo{
		items:    make(map[int64]*Task),
		filePath: "tasks.json",
	}
	repo.loadFromFile()
	return repo
}

func (r *Repo) List() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		out = append(out, t)
	}
	return out
}

func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	now := time.Now()
	t := &Task{ID: r.seq, Title: title, CreatedAt: now, UpdatedAt: now, Done: false}
	r.items[t.ID] = t
	go r.saveToFile()
	return t
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()
	go r.saveToFile()
	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	go r.saveToFile()
	return nil
}

func (r *Repo) loadFromFile() {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}

	var tasks []*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return
	}

	for _, task := range tasks {
		r.items[task.ID] = task
		if task.ID > r.seq {
			r.seq = task.ID
		}
	}
}

func (r *Repo) saveToFile() error {
	tasks := func() []*Task {
		r.mu.RLock()
		defer r.mu.RUnlock()
		out := make([]*Task, 0, len(r.items))
		for _, t := range r.items {
			out = append(out, t)
		}
		return out
	}()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	tmpFile := r.filePath + ".tmp"
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, r.filePath)
}
