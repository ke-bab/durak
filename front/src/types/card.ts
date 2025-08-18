export interface Card {
    suit: Suit
    rank: Rank
}

export enum Rank {
    Six = "six",
    Seven = "seven",
    Eight = "eight",
    Nine = "nine",
    Ten = "ten",
    Jack = "jack",
    Queen = "queen",
    King = "king",
    Ace = "ace",
}

export enum Suit {
    Hearts = 'hearts',
    Diamonds = 'diamonds',
    Clubs = 'clubs',
    Spades = 'spades'
}