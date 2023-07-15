package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numOfSubs int
	var numOfUpdateReq int

	inStream := bufio.NewReader(os.Stdin)
	outStream := bufio.NewWriter(os.Stdout)
	defer outStream.Flush()

	_, err := fmt.Fscanln(inStream, &numOfSubs, &numOfUpdateReq)
	if err != nil {
		return
	}

	if numOfSubs <= 0 || numOfUpdateReq <= 0 {
		return
	}
	if numOfSubs > 50 {
		return
	}
	if numOfUpdateReq > 10000 {
		return
	}

	subscribers := make([]Subscriber, 0, numOfSubs)

	for i := 0; i < numOfSubs; i++ {
		message, _ := inStream.ReadString('\n')
		messageFields := strings.Fields(message)

		numOfTriggerFields := messageFields[0]
		numOfTriggerFieldsInt, _ := strconv.Atoi(numOfTriggerFields)
		triggerFields := make([]string, numOfTriggerFieldsInt)
		copy(triggerFields, messageFields[2:2+numOfTriggerFieldsInt])

		numOfShipmentFields := messageFields[1]
		numOfShipmentFieldsInt, _ := strconv.Atoi(numOfShipmentFields)
		shipmentFields := make([]string, numOfShipmentFieldsInt)
		copy(shipmentFields, messageFields[2+numOfTriggerFieldsInt:])

		sub := Subscriber{
			triggeredFields: triggerFields,
			shipmentFields:  shipmentFields,
		}

		subscribers = append(subscribers, sub)
	}

	offers := make(map[int]*Offer, numOfUpdateReq)

	for i := 0; i < numOfUpdateReq; i++ {
		str, _ := inStream.ReadBytes('\n')
		var message Message
		json.Unmarshal(str, &message)

		isChangedField := map[string]bool{"price": false, "stock_count": false, "partner_content": false, "title": false, "description": false}
		offerId, _ := strconv.Atoi(*message.Offer.Id)

		if offers[offerId] == nil {
			offers[offerId] = &Offer{PartnerContent: new(PartnerContent)}
			offers[offerId].Id = message.Offer.Id
		}

		if message.Offer.Price != nil {
			if offers[offerId].Price != nil && *offers[offerId].Price != *message.Offer.Price {
				offers[offerId].Price = message.Offer.Price
				isChangedField["price"] = true
			} else if offers[offerId].Price == nil {
				offers[offerId].Price = message.Offer.Price
				isChangedField["price"] = true
			}
		}

		if message.Offer.StockCount != nil {
			if offers[offerId].StockCount != nil && *offers[offerId].StockCount != *message.Offer.StockCount {
				offers[offerId].StockCount = message.Offer.StockCount
				isChangedField["stock_count"] = true
			} else if offers[offerId].StockCount == nil {
				offers[offerId].StockCount = message.Offer.StockCount
				isChangedField["stock_count"] = true
			}
		}

		if message.Offer.PartnerContent != nil {
			if message.Offer.PartnerContent.Title != nil {
				if offers[offerId].PartnerContent.Title != nil {
					if *offers[offerId].PartnerContent.Title != *message.Offer.PartnerContent.Title {
						offers[offerId].PartnerContent.Title = message.Offer.PartnerContent.Title
						isChangedField["title"] = true
					}
				} else if offers[offerId].PartnerContent.Title == nil {
					offers[offerId].PartnerContent.Title = message.Offer.PartnerContent.Title
					isChangedField["title"] = true
				}
				isChangedField["partner_content"] = true
			}
			if message.Offer.PartnerContent.Description != nil {
				if offers[offerId].PartnerContent.Description != nil {
					if *offers[offerId].PartnerContent.Description != *message.Offer.PartnerContent.Description {
						offers[offerId].PartnerContent.Description = message.Offer.PartnerContent.Description
						isChangedField["description"] = true
					}
				} else if offers[offerId].PartnerContent.Description == nil {
					offers[offerId].PartnerContent.Description = message.Offer.PartnerContent.Description
					isChangedField["description"] = true
				}
				isChangedField["partner_content"] = true
			}
		}

		for _, subscriber := range subscribers {
			isUpdate := false
			for _, subField := range subscriber.triggeredFields {
				if isChangedField[subField] {
					isUpdate = true
					break
				}
			}
			if isUpdate {
				var outMessage *Message
				outMessage = new(Message)
				outMessage.Offer = new(Offer)
				outMessage.TraceId = message.TraceId
				outMessage.Offer.Id = message.Offer.Id
				for _, subField := range subscriber.triggeredFields {
					switch subField {
					case "offer":
						outMessage.Offer = offers[offerId]
					case "price":
						outMessage.Offer.Price = offers[offerId].Price
					case "stock_count":
						outMessage.Offer.StockCount = offers[offerId].StockCount
					case "partner_content":
						outMessage.Offer.PartnerContent = offers[offerId].PartnerContent
					case "title":
						outMessage.Offer.PartnerContent.Title = offers[offerId].PartnerContent.Title
					case "description":
						outMessage.Offer.PartnerContent.Description = offers[offerId].PartnerContent.Description
					}
				}
				for _, shipField := range subscriber.shipmentFields {
					switch shipField {
					case "offer":
						outMessage.Offer = offers[offerId]
					case "price":
						outMessage.Offer.Price = offers[offerId].Price
					case "stock_count":
						outMessage.Offer.StockCount = offers[offerId].StockCount
					case "partner_content":
						outMessage.Offer.PartnerContent = offers[offerId].PartnerContent
					case "title":
						outMessage.Offer.PartnerContent.Title = offers[offerId].PartnerContent.Title
					case "description":
						outMessage.Offer.PartnerContent.Description =
							offers[offerId].PartnerContent.Description
					}
				}
				out, _ := json.Marshal(outMessage)
				outStream.Write(out)
				fmt.Fprintln(outStream)
			}
		}
	}
}

type Subscriber struct {
	triggeredFields []string
	shipmentFields  []string
}

type Message struct {
	TraceId *string `json:"trace_id,omitempty"`
	Offer   *Offer  `json:"offer,omitempty"`
}

type Offer struct {
	Id             *string         `json:"id,omitempty"`
	Price          *int            `json:"price,omitempty"`
	StockCount     *int            `json:"stock_count,omitempty"`
	PartnerContent *PartnerContent `json:"partner_content,omitempty"`
}

type PartnerContent struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
