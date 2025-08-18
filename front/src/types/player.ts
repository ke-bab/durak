import type {Card} from "./card.ts";

export interface Player {
    id: number
    hand: Card[]
}