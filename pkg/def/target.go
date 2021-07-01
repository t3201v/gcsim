package def

type Target interface {
	Index() int
	SetIndex(ind int) //update the current index
	MaxHP() float64
	HP() float64
	//aura/reactions
	AuraTick() //tick this first to avoid messing with combat
	Tick()
	AddOnReactionHook(key string, fun func(ds *Snapshot))
	//attacks
	Attack(ds *Snapshot) (float64, bool)
	AddOnAttackLandedHook(fun func(ds *Snapshot), key string)
	RemoveOnAttackLandedHook(key string)
	AddDefMod(key int, val float64, dur int)
	AddResMod(key string, val ResistMod)
	DeactivateResMod(key string)
	HasResMod(key string) bool

	Delete() //gracefully deference everything so that it can be gc'd
}

type ResistMod struct {
	Key      string
	Ele      EleType
	Value    float64
	Duration int
	Expiry   int
}

type DefMod struct {
	Key    int
	Value  float64
	Expiry int
}

type ReactionBonusMod struct {
	Key    string
	Amount func(ds *Snapshot) float64
}
