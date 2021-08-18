package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/dtos"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/services"
)

type TicketController struct {
	Service services.TicketService
}

var restClient resty.Client = *resty.New()

func (c TicketController) GetTicketList(rw http.ResponseWriter, r *http.Request) {
	// Call the TicketService to GetTickets
	res, err := c.Service.GetTickets()

	if err != nil {
		http.Error(rw, http.StatusText(404), 404)
	}

	fmt.Fprintf(rw, "%+v", res)
}

func (c TicketController) CreateTicket(rw http.ResponseWriter, r *http.Request) {
	createTicketDTO := dtos.CreateTicketDTO{}

	//Decoding the Body into a CreateTicketDTO
	json.NewDecoder(r.Body).Decode(&createTicketDTO)

	// Call the TicketService to Create
	res, err := c.Service.Create(createTicketDTO)

	if err != nil {
		http.Error(rw, http.StatusText(403), 403)
	}

	fmt.Fprintf(rw, "Created Ticket: %+v", res)
}

func (c TicketController) GetUsersList(rw http.ResponseWriter, r *http.Request) {
	// Call the TicketService to GetTickets
	res, err := c.Service.GetUsers()

	if err != nil {
		http.Error(rw, http.StatusText(401), 401)
	}

	fmt.Fprintf(rw, "%+v", res)
}
