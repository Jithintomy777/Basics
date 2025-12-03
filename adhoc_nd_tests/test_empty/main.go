package main

import (
	"encoding/json"
	"fmt"
)

type Price struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

type BreakupItem struct {
	Price *Price `json:"price,omitempty"`
}

type PriceBreakup struct {
	ItemID    string       `json:"item_id,omitempty"`
	Title     string       `json:"title,omitempty"`
	TitleType string       `json:"title_type,omitempty"`
	Price     *Price       `json:"price,omitempty"`
	Item      *BreakupItem `json:"item,omitempty"`
}

type OrderQuote struct {
	Breakup []PriceBreakup `json:"breakup,omitempty"`
}

type Order struct {
	Quote OrderQuote `json:"quote,omitempty"`
}

func buildBadBreakups(n int) []PriceBreakup {
	// BUG: preallocate with length n -> creates n zero-value entries
	breakups := make([]PriceBreakup, n)
	for i := 0; i < n; i++ {
		// append additional meaningful entries (will come AFTER zero entries)
		breakups = append(breakups, PriceBreakup{
			ItemID:    fmt.Sprintf("item_%d", i),
			Title:     fmt.Sprintf("Item %d", i),
			TitleType: "item",
			Price:     &Price{Currency: "INR", Value: fmt.Sprintf("%.2f", float64(100+i))},
			Item:      &BreakupItem{Price: &Price{Currency: "INR", Value: fmt.Sprintf("%.2f", float64(100+i))}},
		})
	}
	return breakups
}

func buildFixedBreakups(n int) []PriceBreakup {
	// FIX: create zero-length slice with capacity n, then append
	breakups := make([]PriceBreakup, 0, n)
	for i := 0; i < n; i++ {
		breakups = append(breakups, PriceBreakup{
			ItemID:    fmt.Sprintf("item_%d", i),
			Title:     fmt.Sprintf("Item %d", i),
			TitleType: "item",
			Price:     &Price{Currency: "INR", Value: fmt.Sprintf("%.2f", float64(100+i))},
			Item:      &BreakupItem{Price: &Price{Currency: "INR", Value: fmt.Sprintf("%.2f", float64(100+i))}},
		})
	}
	return breakups
}

func main() {
	const n = 2

	badOrder := Order{Quote: OrderQuote{Breakup: buildBadBreakups(n)}}
	fixedOrder := Order{Quote: OrderQuote{Breakup: buildFixedBreakups(n)}}

	badJSON, _ := json.MarshalIndent(badOrder, "", "  ")
	fixedJSON, _ := json.MarshalIndent(fixedOrder, "", "  ")

	fmt.Println(">> BAD (preallocated with length) output:")
	fmt.Println(string(badJSON))
	fmt.Println()
	fmt.Println(">> FIXED (make with length 0 and capacity) output:")
	fmt.Println(string(fixedJSON))
}
