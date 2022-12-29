package main

import (
	"testing"
)

func Test_createDisplayMessageForUser(t *testing.T) {
	cards := []card{{
		name:  "2 of Hearts",
		value: 2,
	}, {
		name:  "7 of Hearts",
		value: 7,
	}}
	result := createDisplayMessageForUser(cards)
	if !(result == "Your cards are: 2 of Hearts & 7 of Hearts! \n") {
		t.Error("Did not create the correct display message for user cards!")
	}
}

func Test_createDisplayMessageForBank(t *testing.T) {
	card := card{
		name:  "2 of Hearts",
		value: 2,
	}
	result := createDisplayMessageForBank(card)
	if !(result == "The bank's card visible to you is: 2 of Hearts (2) \n") {
		t.Error("Did not create the correct display message for bank card!")
	}
}

func Test_createDisplayMessageFor(t *testing.T) {
	cards := []card{{
		name:  "2 of Hearts",
		value: 2,
	}, {
		name:  "7 of Hearts",
		value: 7,
	}}
	result := createDisplayMessageFor(cards)
	if !(result == " 2 of Hearts & 7 of Hearts!") {
		t.Error("Did not create the correct display message!")
	}
}

func Test_createEndGameMessage(t *testing.T) {
	type args struct {
		userResult int
		bankResult int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test user wins black jack",
			args: args{userResult: 21, bankResult: 20},
			want: "YOU WIN WITH BLACK JACK, YEEEESSSS!!",
		},
		{
			name: "test bank wins black jack",
			args: args{userResult: 20, bankResult: 21},
			want: "The bank wins with Black Jack!",
		},
		{
			name: "test user and bank have blackjack",
			args: args{userResult: 21, bankResult: 21},
			want: "It is a tie, nobody wins!",
		},
		{
			name: "test user wins over bank",
			args: args{userResult: 19, bankResult: 18},
			want: "You win yeeey!",
		},
		{
			name: "test bank wins over user",
			args: args{userResult: 17, bankResult: 18},
			want: "You lose!",
		},
		{
			name: "test xxxx",
			args: args{userResult: 13, bankResult: 23},
			want: "You win yeeey!",
		},
		{
			name: "test xxxx",
			args: args{userResult: 23, bankResult: 13},
			want: "You lose!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createEndGameMessage(tt.args.userResult, tt.args.bankResult); got != tt.want {
				t.Errorf("createEndGameMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
