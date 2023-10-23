package http

type Success struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type Error struct {
	Success bool
	Error   any
}

func ReturnSuccess(value any) Success {
	result := Success{Success: true, Data: value}
	return result
}

func ReturnError(value any) Error {
	return Error{Success: false, Error: value}
}
