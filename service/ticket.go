
package service

import (
	"sync"
	"sync/atomic"

	errors "github.com/go-openapi/errors"
	"github.com/user/todo/models"
)

type TicketService struct {
	tickets map[int64]*models.Ticket
	lastID int64
	ticketsLock sync.Mutex
}

func (l *TicketService) newTicketID() int64 {
	return atomic.AddInt64(&l.lastID, 1)
}

func (l *TicketService) AddTicket(ticket *models.Ticket) (*models.Ticket, error) {
	if ticket == nil {
		return nil, errors.New(500, "item must be present")
	}

	l.ticketsLock.Lock()
	defer l.ticketsLock.Unlock()

	newID := l.newTicketID()
	ticket.ID = newID
	l.tickets[newID] = ticket

	return ticket, nil
}
