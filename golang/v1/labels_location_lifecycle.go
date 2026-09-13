package v1

// LocationLifecycleLabels is an opt-in vocabulary contract, not a persisted
// status enum or transition policy. Existing location consumers keep using active.
// Load with LoadPath(locale, vertical, "location_lifecycle.json", "location_lifecycle", &labels).
type LocationLifecycleLabels struct {
	Label  string                  `json:"label"`
	Help   string                  `json:"help"`
	States LocationLifecycleStates `json:"states"`
}

type LocationLifecycleStates struct {
	Planned           LocationLifecycleStateLabels `json:"planned"`
	Operational       LocationLifecycleStateLabels `json:"operational"`
	UnderMaintenance  LocationLifecycleStateLabels `json:"under_maintenance"`
	TemporarilyClosed LocationLifecycleStateLabels `json:"temporarily_closed"`
	Closed            LocationLifecycleStateLabels `json:"closed"`
}

type LocationLifecycleStateLabels struct {
	Label string `json:"label"`
	Help  string `json:"help"`
}
