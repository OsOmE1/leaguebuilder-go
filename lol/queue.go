package lol

type QueueType int

const (
	Custom        QueueType = 0
	DraftPick     QueueType = 400
	RankedSoloDue QueueType = 420
	BlindPick     QueueType = 430
	RankedFlex    QueueType = 440
)
