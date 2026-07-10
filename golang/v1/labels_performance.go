package v1

// labels_performance.go — Go label structs for the performance panel
// (performance-evaluation layer 12, Phase E).
//
// Root key: "performance" (file performance.json).
//
// PerformancePanelLabels is the primary struct consumed by the fayna
// domain/operation/performance/ view (Surface 6 — admin performance panel).
// It is also partially consumed by the client-portal "Rate My Team" surface
// via the panel.Actions labels.

// PerformancePanelLabels holds all translatable strings for the performance panel.
type PerformancePanelLabels struct {
	Page   PerformancePageLabels  `json:"page"`
	Panel  PerformancePanelBlock  `json:"panel"`
	Errors PerformanceErrorLabels `json:"errors"`
}

type PerformancePageLabels struct {
	Heading string `json:"heading"`
	Caption string `json:"caption"`
}

// PerformancePanelBlock holds the sub-labels for the panel matrix view.
// Groups maps overdue/due/upToDate group headings.
type PerformancePanelBlock struct {
	Heading string                       `json:"heading"`
	Caption string                       `json:"caption"`
	Groups  PerformancePanelGroupLabels  `json:"groups"`
	Columns PerformancePanelColumnLabels `json:"columns"`
	Actions PerformancePanelActionLabels `json:"actions"`
	Empty   PerformancePanelEmptyLabels  `json:"empty"`
	Rating  PerformancePanelRatingLabels `json:"rating"`
	Cycle   PerformancePanelCycleLabels  `json:"cycle"`
}

// PerformancePanelGroupLabels holds the three group-heading strings for the matrix.
// These correspond to the three urgency buckets in GetPerformancePanelData output.
type PerformancePanelGroupLabels struct {
	Overdue  string `json:"overdue"`
	Due      string `json:"due"`
	UpToDate string `json:"up_to_date"`
}

type PerformancePanelColumnLabels struct {
	Associate  string `json:"associate"`
	Client     string `json:"client"`
	LastReview string `json:"last_review"`
	NextDue    string `json:"next_due"`
	Rating     string `json:"rating"`
	Status     string `json:"status"`
}

type PerformancePanelActionLabels struct {
	StartReview    string `json:"start_review"`
	ViewLastReview string `json:"view_last_review"`
}

type PerformancePanelEmptyLabels struct {
	Title         string `json:"title"`
	Message       string `json:"message"`
	OverdueEmpty  string `json:"overdue_empty"`
	DueEmpty      string `json:"due_empty"`
	UpToDateEmpty string `json:"up_to_date_empty"`
}

type PerformancePanelRatingLabels struct {
	NoRating string `json:"no_rating"`
	Label    string `json:"label"`
}

type PerformancePanelCycleLabels struct {
	Heading string `json:"heading"`
	NoCycle string `json:"no_cycle"`
}

type PerformanceErrorLabels struct {
	PermissionDenied string `json:"permission_denied"`
	LoadFailed       string `json:"load_failed"`
}

// DefaultPerformancePanelLabels returns PerformancePanelLabels with sensible English defaults.
func DefaultPerformancePanelLabels() PerformancePanelLabels {
	return PerformancePanelLabels{
		Page: PerformancePageLabels{
			Heading: "Performance",
			Caption: "Monitor review status and scores across all engagements",
		},
		Panel: PerformancePanelBlock{
			Heading: "Performance Overview",
			Caption: "Review status grouped by urgency",
			Groups: PerformancePanelGroupLabels{
				Overdue:  "Overdue",
				Due:      "Due",
				UpToDate: "Up to Date",
			},
			Columns: PerformancePanelColumnLabels{
				Associate:  "Associate",
				Client:     "Client",
				LastReview: "Last Review",
				NextDue:    "Next Due",
				Rating:     "Rating",
				Status:     "Status",
			},
			Actions: PerformancePanelActionLabels{
				StartReview:    "Start Review",
				ViewLastReview: "View Last Review",
			},
			Empty: PerformancePanelEmptyLabels{
				Title:         "No reviews due",
				Message:       "All team members are up to date.",
				OverdueEmpty:  "No overdue reviews",
				DueEmpty:      "No reviews due soon",
				UpToDateEmpty: "No reviews up to date",
			},
			Rating: PerformancePanelRatingLabels{
				NoRating: "—",
				Label:    "Rating",
			},
			Cycle: PerformancePanelCycleLabels{
				Heading: "Active Cycle",
				NoCycle: "No active cycle",
			},
		},
		Errors: PerformanceErrorLabels{
			PermissionDenied: "You do not have permission to view this panel",
			LoadFailed:       "Failed to load performance data",
		},
	}
}
