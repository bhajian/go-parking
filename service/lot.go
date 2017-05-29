
package service

import (
	"sync"
	"sync/atomic"
	"time"
	errors "github.com/go-openapi/errors"
	"github.com/user/go-parking/models"
	
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name TodoList --spec ../swagger.yml

type LotService struct {

	lots map[int64]*models.Lot
	lastID int64

	lotsLock sync.Mutex
	ticketSerevice TicketService
}



func (l *LotService) Init() {
	l.lots = make(map[int64]*models.Lot)
	l.ticketSerevice = TicketService{}
	l.ticketSerevice.Init()
}

func (l *LotService) newLotID() int64 {
	return atomic.AddInt64(&l.lastID, 1)
}


func (l *LotService) AddLot(lot *models.Lot) error {
	if lot == nil {
		return errors.New(500, "item must be present")
	}

	l.lotsLock.Lock()
	defer l.lotsLock.Unlock()

	newID := l.newLotID()
	lot.ID = newID
	lot.Init()
	l.lots[newID] = lot

	return nil
}

func (l *LotService) UpdateLot(id int64, lot *models.Lot) error {
	if lot == nil {
		return errors.New(500, "item must be present")
	}

	l.lotsLock.Lock()
	defer l.lotsLock.Unlock()

	_, exists := l.lots[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	lot.ID = id
	l.lots[id] = lot
	return nil
}

func (l *LotService) DeleteLot(id int64) error {
	l.lotsLock.Lock()
	defer l.lotsLock.Unlock()

	_, exists := l.lots[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(l.lots, id)
	return nil
}

func (l *LotService) AllLots(since int64, limit int32) (result []*models.Lot) {
	result = make([]*models.Lot, 0)
	for id, lot := range l.lots {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, lot)
		}
	}
	return
}

func (l *LotService) GetTicket(lotId int64, carSize string) (*models.Ticket, error) {
	var ticket = models.Ticket{}
	t := time.Now()
	ticket.LotID = lotId
	ticket.CarSize = carSize
	ticket.EntryTime = t.Format(time.RFC3339)
	var retTicket, err = l.ticketSerevice.AddTicket(&ticket) 
	if err == nil {
		var lot = l.lots[lotId]
		lot.GetTicket(*retTicket)
		return retTicket, nil
	}
	return nil, err
}

func (l *LotService) CarLeaves(lotId int64, ticketId int64) *models.Ticket {
	var lot = l.lots[lotId]
	var ticket = l.ticketSerevice.tickets[ticketId]
	return lot.CarLeaves(ticket.CarSize, ticket.SpotNumber)
}

