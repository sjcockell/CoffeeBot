package reminder

import ("math/rand"
  "time"
)

var invites = []string {
  "Time for coffee",
  "Coffee Tiem!!11!",
  "Can haz coffee? KTHXBAI",
  "Dental Bites is about to close. Run! Run!",
  "You want coffee don't you? Better get a move on then.",
  "Now is the winter of our discontent. Oh, and time for coffee.",
  "Got mil^H^H^H coffee?",
  "This is your CoffeeBot. Time for coffee.",
  "Much coffee. So caffeine. Wow.",
  "ZOMBIE APOCALYPSE IMMINENT! Go get coffee.",
}

//generates a reminder
func Reminder() string {
  rand.Seed(time.Now().UTC().UnixNano())
  i := rand.Intn(len(invites))
  return invites[i]
}
