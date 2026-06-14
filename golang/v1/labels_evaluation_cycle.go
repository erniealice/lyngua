package v1

// labels_evaluation_cycle.go — Go label structs for evaluation_cycle domain
// (performance-evaluation layer 12, v1 — Q-EVAL-CYCLE-1).
//
// Root key: "evaluationCycle" (file evaluation_cycle.json).
//
// EvaluationCycleLabels includes EvaluationCycleBannerLabels for the
// "X of Y" read-projection partial (cycle-progress-banner.html), rendered
// on cycle detail, admin performance panel, and client portal.
//
// EvaluationCycleStatusLabels maps OPEN | SIGN_OFF | CLOSED values.

// EvaluationCycleLabels holds all translatable strings for the evaluation cycle module.
type EvaluationCycleLabels struct {
	Page        EvaluationCyclePageLabels    `json:"page"`
	Buttons     EvaluationCycleButtonLabels  `json:"buttons"`
	Columns     EvaluationCycleColumnLabels  `json:"columns"`
	Status      EvaluationCycleStatusLabels  `json:"status"`
	Tabs        EvaluationCycleTabLabels     `json:"tabs"`
	Banner      EvaluationCycleBannerLabels  `json:"banner"`
	Empty       EvaluationCycleEmptyLabels   `json:"empty"`
	Form        EvaluationCycleFormLabels    `json:"form"`
	Actions     EvaluationCycleActionLabels  `json:"actions"`
	Detail      EvaluationCycleDetailLabels  `json:"detail"`
	Confirm     EvaluationCycleConfirmLabels `json:"confirm"`
	Errors      EvaluationCycleErrorLabels   `json:"errors"`
}

type EvaluationCyclePageLabels struct {
	Heading          string `json:"heading"`
	Caption          string `json:"caption"`
	HeadingOpen      string `json:"headingOpen"`
	HeadingSignOff   string `json:"headingSignOff"`
	HeadingClosed    string `json:"headingClosed"`
}

type EvaluationCycleButtonLabels struct {
	AddCycle string `json:"addCycle"`
	Open     string `json:"open"`
	Close    string `json:"close"`
}

type EvaluationCycleColumnLabels struct {
	Name        string `json:"name"`
	Engagement  string `json:"engagement"`
	Period      string `json:"period"`
	Progress    string `json:"progress"`
	SignOffDue  string `json:"signOffDue"`
	Closes      string `json:"closes"`
	Status      string `json:"status"`
}

// EvaluationCycleStatusLabels maps EvaluationCycleStatus enum values.
// Mirrors OPEN | SIGN_OFF | CLOSED (entities.md §E4, v1).
type EvaluationCycleStatusLabels struct {
	Open    string `json:"open"`
	SignOff string `json:"signOff"`
	Closed  string `json:"closed"`
}

type EvaluationCycleTabLabels struct {
	All      string `json:"all"`
	Open     string `json:"open"`
	SignOff  string `json:"signOff"`
	Closed   string `json:"closed"`
	Info     string `json:"info"`
	Members  string `json:"members"`
}

// EvaluationCycleBannerLabels holds labels for the "X of Y" read-projection
// cycle-progress-banner.html partial (SR-1; NOT a form; three shared surfaces:
// cycle detail, admin performance panel, client portal).
//
// Progress contains a Go template string with {{.Complete}} and {{.Total}} placeholders.
// SignOffDue and Closes are date-field header labels (LBL-2, Phase E, v1).
type EvaluationCycleBannerLabels struct {
	SignOffDue    string `json:"signOffDue"`
	Closes        string `json:"closes"`
	Progress      string `json:"progress"`
	ProgressDetail string `json:"progressDetail"`
	AllComplete   string `json:"allComplete"`
	NoneComplete  string `json:"noneComplete"`
	MembersLabel  string `json:"membersLabel"`
	CycleOpen     string `json:"cycleOpen"`
	CycleClosed   string `json:"cycleClosed"`
	CycleSignOff  string `json:"cycleSignOff"`
}

type EvaluationCycleEmptyLabels struct {
	Title         string `json:"title"`
	Message       string `json:"message"`
	OpenTitle     string `json:"openTitle"`
	OpenMessage   string `json:"openMessage"`
	ClosedTitle   string `json:"closedTitle"`
	ClosedMessage string `json:"closedMessage"`
}

type EvaluationCycleFormLabels struct {
	Name               string `json:"name"`
	NamePlaceholder    string `json:"namePlaceholder"`
	Engagement         string `json:"engagement"`
	EngagementPH       string `json:"engagementPlaceholder"`
	EngagementInfo     string `json:"engagementInfo"`
	PeriodStart        string `json:"periodStart"`
	PeriodEnd          string `json:"periodEnd"`
	SignOffDueDate     string `json:"signOffDueDate"`
	SignOffDueDateInfo string `json:"signOffDueDateInfo"`
	CloseDate          string `json:"closeDate"`
	CloseDateInfo      string `json:"closeDateInfo"`
}

type EvaluationCycleActionLabels struct {
	View  string `json:"view"`
	Open  string `json:"open"`
	Close string `json:"close"`
}

type EvaluationCycleDetailLabels struct {
	PageTitle    string                        `json:"pageTitle"`
	Name         string                        `json:"name"`
	Engagement   string                        `json:"engagement"`
	PeriodStart  string                        `json:"periodStart"`
	PeriodEnd    string                        `json:"periodEnd"`
	SignOffDueDate string                      `json:"signOffDueDate"`
	CloseDate    string                        `json:"closeDate"`
	Status       string                        `json:"status"`
	CreatedDate  string                        `json:"createdDate"`
	ModifiedDate string                        `json:"modifiedDate"`
	Members      EvaluationCycleMembersLabels  `json:"members"`
}

type EvaluationCycleMembersLabels struct {
	Heading         string `json:"heading"`
	Empty           string `json:"empty"`
	ColumnAssociate string `json:"columnAssociate"`
	ColumnClient    string `json:"columnClient"`
	ColumnProbation string `json:"columnProbation"`
	ColumnAdded     string `json:"columnAdded"`
}

type EvaluationCycleConfirmLabels struct {
	Open         string `json:"open"`
	OpenMessage  string `json:"openMessage"`
	Close        string `json:"close"`
	CloseMessage string `json:"closeMessage"`
}

type EvaluationCycleErrorLabels struct {
	PermissionDenied    string `json:"permissionDenied"`
	InvalidFormData     string `json:"invalidFormData"`
	NotFound            string `json:"notFound"`
	IDRequired          string `json:"idRequired"`
	NoPermission        string `json:"noPermission"`
	AlreadyOpen         string `json:"alreadyOpen"`
	AlreadyClosed       string `json:"alreadyClosed"`
	EngagementRequired  string `json:"engagementRequired"`
}

// DefaultEvaluationCycleLabels returns EvaluationCycleLabels with sensible English defaults.
func DefaultEvaluationCycleLabels() EvaluationCycleLabels {
	return EvaluationCycleLabels{
		Page: EvaluationCyclePageLabels{
			Heading:        "Evaluation Cycles",
			Caption:        "Manage structured review periods across your engagements",
			HeadingOpen:    "Open Cycles",
			HeadingSignOff: "Sign-Off Cycles",
			HeadingClosed:  "Closed Cycles",
		},
		Buttons: EvaluationCycleButtonLabels{
			AddCycle: "New Cycle",
			Open:     "Open",
			Close:    "Close",
		},
		Columns: EvaluationCycleColumnLabels{
			Name:       "Name",
			Engagement: "Engagement",
			Period:     "Period",
			Progress:   "Progress",
			SignOffDue: "Sign-off Due",
			Closes:     "Closes",
			Status:     "Status",
		},
		Status: EvaluationCycleStatusLabels{
			Open:    "Open",
			SignOff: "Sign Off",
			Closed:  "Closed",
		},
		Tabs: EvaluationCycleTabLabels{
			All:     "All",
			Open:    "Open",
			SignOff: "Sign Off",
			Closed:  "Closed",
			Info:    "Information",
			Members: "Members",
		},
		Banner: EvaluationCycleBannerLabels{
			SignOffDue:     "Sign-off due",
			Closes:         "Closes",
			Progress:       "{{.Complete}} of {{.Total}} complete",
			ProgressDetail: "{{.Complete}} reviews submitted or signed off",
			AllComplete:    "All reviews complete",
			NoneComplete:   "No reviews submitted yet",
			MembersLabel:   "Members",
			CycleOpen:      "Cycle is open",
			CycleClosed:    "Cycle is closed",
			CycleSignOff:   "Awaiting sign-off",
		},
		Empty: EvaluationCycleEmptyLabels{
			Title:         "No cycles found",
			Message:       "No evaluation cycles to display.",
			OpenTitle:     "No open cycles",
			OpenMessage:   "Open a cycle to start tracking reviews.",
			ClosedTitle:   "No closed cycles",
			ClosedMessage: "Closed cycles will appear here.",
		},
		Form: EvaluationCycleFormLabels{
			Name:               "Cycle Name",
			NamePlaceholder:    "e.g. H1 2026 Performance Review",
			Engagement:         "Engagement",
			EngagementPH:       "Select engagement",
			EngagementInfo:     "The subscription (engagement) this cycle is scoped to.",
			PeriodStart:        "Period Start",
			PeriodEnd:          "Period End",
			SignOffDueDate:     "Sign-off Due Date",
			SignOffDueDateInfo: "Deadline for all sign-offs. Distinct from the close date.",
			CloseDate:          "Close Date",
			CloseDateInfo:      "When this cycle will be closed. Reviews cannot be submitted after this date.",
		},
		Actions: EvaluationCycleActionLabels{
			View:  "View Cycle",
			Open:  "Open",
			Close: "Close",
		},
		Detail: EvaluationCycleDetailLabels{
			PageTitle:      "Cycle Details",
			Name:           "Name",
			Engagement:     "Engagement",
			PeriodStart:    "Period Start",
			PeriodEnd:      "Period End",
			SignOffDueDate: "Sign-off Due",
			CloseDate:      "Close Date",
			Status:         "Status",
			CreatedDate:    "Created",
			ModifiedDate:   "Last Modified",
			Members: EvaluationCycleMembersLabels{
				Heading:         "Members",
				Empty:           "No members enrolled in this cycle.",
				ColumnAssociate: "Associate",
				ColumnClient:    "Client",
				ColumnProbation: "Probation",
				ColumnAdded:     "Added",
			},
		},
		Confirm: EvaluationCycleConfirmLabels{
			Open:         "Open Cycle",
			OpenMessage:  "Open \"%s\"? Members will be enrolled and reviews can be submitted.",
			Close:        "Close Cycle",
			CloseMessage: "Close \"%s\"? No further reviews can be submitted after closing.",
		},
		Errors: EvaluationCycleErrorLabels{
			PermissionDenied:   "You do not have permission to perform this action",
			InvalidFormData:    "Invalid form data. Please check your inputs and try again.",
			NotFound:           "Evaluation cycle not found",
			IDRequired:         "Cycle ID is required",
			NoPermission:       "No permission",
			AlreadyOpen:        "This cycle is already open",
			AlreadyClosed:      "This cycle is already closed",
			EngagementRequired: "An engagement is required",
		},
	}
}
