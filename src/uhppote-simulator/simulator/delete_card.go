package simulator

import (
	"uhppote"
)

func (s *Simulator) DeleteCard(request *uhppote.DeleteCardRequest) (interface{}, error) {
	if s.SerialNumber != request.SerialNumber {
		return nil, nil
	}

	deleted := s.Cards.Delete(request.CardNumber)
	saved := false

	if deleted {
		if err := s.Save(); err == nil {
			saved = true
		}
	}

	response := uhppote.DeleteCardResponse{
		SerialNumber: s.SerialNumber,
		Succeeded:    deleted && saved,
	}

	return &response, nil
}