package stacks

type Action struct {
	Name string
	meta string
	Next *Action
}

type ActionHistory struct {
	Top  *Action
	Size int
}

func (actionHistory *ActionHistory) Add(newAction *Action) {
	if actionHistory.Top != nil {
		oldHistory := actionHistory.Top
		actionHistory.Top = newAction
		newAction.Next = oldHistory
	}
	actionHistory.Top = newAction
	actionHistory.Size++
}

func (history ActionHistory) Undo() *Action {
	rAction := history.Top
	if rAction == nil {
		history.Top = nil
		history.Size = 1
	} else {
		history.Top = rAction.Next
	}
	history.Size--
	return rAction
}
