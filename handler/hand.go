package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Depermitto/witcher-dice-poker/model"
	"io"
	"math/rand/v2"
	"net/http"
)

// RandomHand godoc
//
//	@Description	Generate random dice poker hand
//	@Tags			Hands
//	@Produce		json
//	@Success		200	{object}	model.Hand
//	@Router			/hands/random [get]
func RandomHand(w http.ResponseWriter, _ *http.Request) {
	var dice [5]uint
	for i := range dice {
		dice[i] = rand.UintN(6) + 1
	}
	hand, _ := model.MakeHand(dice)  // assume Hand always makes correctly
	jsonStr, _ := json.Marshal(hand) // assume Hand always marshals correctly
	_, _ = fmt.Fprintf(w, "%s\n", jsonStr)
}

type updateRequest struct {
	Hand     model.Hand `json:"hand"`
	Switches []uint     `json:"switches"`
}

// UpdateHand godoc
//
//	@Description	Update dice poker hand
//	@Tags			Hands
//	@Accept			json
//	@Produce		json
//	@Param			updateRequest	body		updateRequest	true	"Hand to modify along with list of dice indexes. Die at index will be switched with a new, randomly generated value. Dice indexes (1-5), array length (1-5)"
//	@Success		200				{object}	model.Hand
//	@Failure		400				{object}	int
//	@Router			/hands/switch [post]
func UpdateHand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error: unable to read req body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := updateRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "error: could not parse JSON data in req body", http.StatusBadRequest)
		return
	}

	if len(req.Switches) <= 0 || len(req.Switches) > 5 {
		http.Error(w, "error: switches array must of length [1-5]", http.StatusBadRequest)
		return
	}

	for _, s := range req.Switches {
		if s < 1 || s > 6 {
			http.Error(w, fmt.Sprintf("error: switch index %v out of range [1, 6]", s), http.StatusBadRequest)
			return
		}
		req.Hand.Dice[s-1] = rand.UintN(6) + 1
	}
	req.Hand, err = model.MakeHand(req.Hand.Dice)
	if err != nil {
		http.Error(w, "error: "+err.Error(), http.StatusBadRequest)
	}

	_, _ = fmt.Fprintln(w, req.Hand)
}

type evalRequest struct {
	Dice [5]uint `json:"dice"`
}

// EvaluateHand godoc
//
//	@Description	Evaluate dice
//	@Tags			Hands
//	@Accept			json
//	@Produce		json
//	@Param			evalRequest	body		evalRequest	true	"Raw dice to evaluate. Value range (1-6), array length (5)"
//	@Success		200			{object}	model.Hand	"Hand created from dice"
//	@Failure		400			{object}	int
//	@Router			/hands/eval [post]
func EvaluateHand(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error: unable to read req body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := evalRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "error: could not parse JSON data in req body", http.StatusBadRequest)
		return
	}
	hand, err := model.MakeHand(req.Dice)
	if err != nil {
		http.Error(w, "error: "+err.Error(), http.StatusBadRequest)
	}

	_, _ = fmt.Fprintln(w, hand)
}
