package safe_slice

type UpdateFunc func(any) any

type SafeSlice interface {
	Append(any)
	At(int) any
	Len() int
	Close() []any
	Delete(int)
	Update(int, UpdateFunc)
}

type commandAction int

const (
	insert commandAction = iota
	remove
	at
	update
	end
	length
)

type commandData struct {
	action  commandAction
	index   int
	item    any
	result  chan<- any
	updater UpdateFunc
}

type safeSlice chan commandData

func (s safeSlice) Append(item any) {
	s <- commandData{action: insert, item: item}
}

func (s safeSlice) At(index int) any {
	result := make(chan any)
	s <- commandData{action: at, index: index, result: result}
	return <-result
}

func (s safeSlice) Len() int {
	result := make(chan any)
	s <- commandData{action: length, result: result}
	return (<-result).(int)
}

func (s safeSlice) Delete(index int) {
	s <- commandData{action: remove, index: index}
}

func (s safeSlice) Update(index int, updater UpdateFunc) {
	s <- commandData{action: update, index: index, updater: updater}
}

func (s safeSlice) Close() []any {
	result := make(chan any)
	s <- commandData{action: end, result: result}
	return (<-result).([]any)
}

func (s safeSlice) run() {
	var storage []any
	for command := range s {
		switch command.action {
		case insert:
			storage = append(storage, command.item)
		case remove:
			if 0 <= command.index && command.index < len(storage) {
				storage = append(storage[:command.index], storage[command.index+1:]...)
			}
		case at:
			if 0 <= command.index && command.index < len(storage) {
				command.result <- storage[command.index]
			} else {
				command.result <- nil
			}
		case length:
			command.result <- len(storage)
		case update:
			if 0 <= command.index && command.index < len(storage) {
				storage[command.index] = command.updater(storage[command.index])
			}
		case end:
			close(s)
			command.result <- storage
		}
	}
}

func New() SafeSlice {
	safeSliceInstance := make(safeSlice)
	go safeSliceInstance.run()
	return safeSliceInstance
}
