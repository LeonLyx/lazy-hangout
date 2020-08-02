package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	serviceAccountCredentials = "credentials.json"
	dateLayout                = "2006-01-02"
	conferenceSolution        = "eventHangout"
	calendarID                = "primary"
)

func newEvent() (*calendar.Event, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "| unable to generate uuid")
	}

	currDate := time.Now().Add(time.Hour * 1).Format(dateLayout)
	startDate := &calendar.EventDateTime{Date: currDate}
	endDate := &calendar.EventDateTime{Date: currDate}
	conferenceData := &calendar.ConferenceData{
		CreateRequest: &calendar.CreateConferenceRequest{
			RequestId:             uuid.String(),
			ConferenceSolutionKey: &calendar.ConferenceSolutionKey{Type: conferenceSolution},
		},
	}

	event := &calendar.Event{
		Start:          startDate,
		End:            endDate,
		ConferenceData: conferenceData,
	}
	return event, nil
}

func main() {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithCredentialsFile(serviceAccountCredentials))
	if err != nil {
		fmt.Printf("unable to initialise service, %v", err)
		os.Exit(1)
	}

	event, err := newEvent()
	if err != nil {
		fmt.Printf("unable to initialise event, %v", err)
		os.Exit(1)
	}

	newEvent, err := srv.Events.Insert(calendarID, event).ConferenceDataVersion(1).Do()
	if err != nil {
		fmt.Printf("unable to create event %v", err)
		os.Exit(1)
	}
	fmt.Printf("The following is your Google hangout link : \n%s\n", newEvent.HangoutLink)
}
