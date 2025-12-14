package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(t Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, exists := l.tasks[t.Title]; exists {
		return ErrTaskAlreadyExists
	}

	l.tasks[t.Title] = t

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, exists := l.tasks[title]
	if !exists {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (l *List) GetTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task)
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) GetCompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task)
	for k, v := range l.tasks {
		if v.Completed {
			tmp[k] = v
		}

	}
	return tmp
}

func (l *List) GetUncompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task)
	for k, v := range l.tasks {
		if !v.Completed {
			tmp[k] = v
		}
	}
	return tmp
}

func (l *List) CompleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if task, exists := l.tasks[title]; exists {
		task.Complete()
		l.tasks[title] = task
		return nil
	}
	return ErrTaskNotFound
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, exists := l.tasks[title]; exists {
		delete(l.tasks, title)
		return nil
	}
	return ErrTaskNotFound
}
