package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

type Handler interface {
	GetCommand() (string, error)
}

type inputHandler struct {
	infoText string
	*bufio.Reader
	validInputList []string
}

func NewHandler(info string, validInput ...string) Handler {
	return &inputHandler{
		Reader:         bufio.NewReader(os.Stdin),
		validInputList: validInput,
		infoText:       info,
	}
}

func (h *inputHandler) GetCommand() (string, error) {
	fmt.Println(h.infoText)
	inputString, err := h.ReadString('\n')
	if err != nil {
		return "", err
	}

	replacer := strings.NewReplacer("\n", "", "\r", "")

	inputString = replacer.Replace(inputString)

	if err = h.validateInput(inputString); err != nil {
		return "", err
	}

	return inputString, nil
}

func (h *inputHandler) validateInput(input string) error {
	for _, validInput := range h.validInputList {
		if strings.EqualFold(validInput, input) {
			return nil
		}
	}
	return ErrInvalidInput
}
