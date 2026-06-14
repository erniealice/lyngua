package v1

// labels_evaluation.go — Go label structs for the evaluation domain (performance-evaluation layer 12).
//
// Root keys (camelCase, matching JSON files):
//   evaluation.json        → EvaluationLabels       (root key "evaluation")
//   evaluation_response.json → EvaluationResponseLabels (root key "evaluationResponse")
//
// These types are consumed by fayna domain/operation/evaluation/ and
// domain/operation/evaluation_response/ labels.go wrappers.
// JSON tags mirror the evaluation.json / evaluation_response.json structure.

// EvaluationLabels holds all translatable strings for the evaluation module.
type EvaluationLabels struct {
	Page        EvaluationPageLabels       `json:"page"`
	Buttons     EvaluationButtonLabels     `json:"buttons"`
	Columns     EvaluationColumnLabels     `json:"columns"`
	Status      EvaluationStatusLabels     `json:"status"`
	Tabs        EvaluationTabLabels        `json:"tabs"`
	Empty       EvaluationEmptyLabels      `json:"empty"`
	Form        EvaluationFormLabels       `json:"form"`
	Dimension   EvaluationDimensionLabels  `json:"dimension"`
	Actions     EvaluationActionLabels     `json:"actions"`
	BulkActions EvaluationBulkLabels       `json:"bulkActions"`
	Detail      EvaluationDetailLabels     `json:"detail"`
	Confirm     EvaluationConfirmLabels    `json:"confirm"`
	Errors      EvaluationErrorLabels      `json:"errors"`
	Portal      EvaluationPortalLabels     `json:"portal"`
}

type EvaluationPageLabels struct {
	Heading          string `json:"heading"`
	Caption          string `json:"caption"`
	HeadingDraft     string `json:"headingDraft"`
	HeadingSubmitted string `json:"headingSubmitted"`
	HeadingSignedOff string `json:"headingSignedOff"`
	HeadingArchived  string `json:"headingArchived"`
}

type EvaluationButtonLabels struct {
	StartReview string `json:"startReview"`
	SaveDraft   string `json:"saveDraft"`
	Submit      string `json:"submit"`
	SignOff     string `json:"signOff"`
	Archive     string `json:"archive"`
	Delete      string `json:"delete"`
}

type EvaluationColumnLabels struct {
	Associate   string `json:"associate"`
	Client      string `json:"client"`
	Period      string `json:"period"`
	Type        string `json:"type"`
	Overall     string `json:"overall"`
	Status      string `json:"status"`
	SubmittedAt string `json:"submittedAt"`
}

// EvaluationStatusLabels maps EvaluationStatus enum values to display strings.
// SignedOff corresponds to EVALUATION_STATUS_SIGNED_OFF (Q-SIGNOFF-1, v1).
type EvaluationStatusLabels struct {
	Draft     string `json:"draft"`
	Submitted string `json:"submitted"`
	SignedOff string `json:"signedOff"`
	Archived  string `json:"archived"`
}

type EvaluationTabLabels struct {
	All          string `json:"all"`
	Draft        string `json:"draft"`
	Submitted    string `json:"submitted"`
	SignedOff    string `json:"signedOff"`
	Archived     string `json:"archived"`
	Info         string `json:"info"`
	Scores       string `json:"scores"`
	SignOff      string `json:"signOff"`
	AuditTrail   string `json:"auditTrail"`
}

type EvaluationEmptyLabels struct {
	Title            string `json:"title"`
	Message          string `json:"message"`
	DraftTitle       string `json:"draftTitle"`
	DraftMessage     string `json:"draftMessage"`
	SubmittedTitle   string `json:"submittedTitle"`
	SubmittedMessage string `json:"submittedMessage"`
	SignedOffTitle   string `json:"signedOffTitle"`
	SignedOffMessage string `json:"signedOffMessage"`
}

type EvaluationFormLabels struct {
	EvaluationType        string `json:"evaluationType"`
	EvaluationTypePH      string `json:"evaluationTypePlaceholder"`
	Template              string `json:"template"`
	TemplatePlaceholder   string `json:"templatePlaceholder"`
	PeriodStart           string `json:"periodStart"`
	PeriodEnd             string `json:"periodEnd"`
	Narrative             string `json:"narrative"`
	NarrativePlaceholder  string `json:"narrativePlaceholder"`
	Associate             string `json:"associate"`
	Client                string `json:"client"`
	SectionHeader         string `json:"sectionHeader"`
	SectionRubric         string `json:"sectionRubric"`
	SectionNarrative      string `json:"sectionNarrative"`
}

type EvaluationDimensionLabels struct {
	NotScored string                        `json:"notScored"`
	PassFail  EvaluationDimensionPassFail   `json:"passFail"`
	RatingBar EvaluationDimensionRatingBar  `json:"ratingBar"`
}

type EvaluationDimensionPassFail struct {
	Pass string `json:"pass"`
	Fail string `json:"fail"`
}

type EvaluationDimensionRatingBar struct {
	AriaLabel string `json:"ariaLabel"`
}

type EvaluationActionLabels struct {
	View    string `json:"view"`
	Edit    string `json:"edit"`
	SignOff string `json:"signOff"`
	Archive string `json:"archive"`
	Delete  string `json:"delete"`
}

type EvaluationBulkLabels struct {
	Archive string `json:"archive"`
	Delete  string `json:"delete"`
}

type EvaluationDetailLabels struct {
	PageTitle    string                      `json:"pageTitle"`
	Associate    string                      `json:"associate"`
	Client       string                      `json:"client"`
	Period       string                      `json:"period"`
	Type         string                      `json:"type"`
	Template     string                      `json:"template"`
	OverallScore string                      `json:"overallScore"`
	Status       string                      `json:"status"`
	SubmittedAt  string                      `json:"submittedAt"`
	SignedOffAt  string                      `json:"signedOffAt"`
	SignedOffBy  string                      `json:"signedOffBy"`
	Narrative    string                      `json:"narrative"`
	Scores       EvaluationDetailScores      `json:"scores"`
	SignOff      EvaluationDetailSignOff     `json:"signOff"`
}

type EvaluationDetailScores struct {
	Heading      string `json:"heading"`
	Criterion    string `json:"criterion"`
	Weight       string `json:"weight"`
	Score        string `json:"score"`
	Comment      string `json:"comment"`
	WeightedTotal string `json:"weightedTotal"`
	Empty        string `json:"empty"`
}

type EvaluationDetailSignOff struct {
	Heading  string `json:"heading"`
	Status   string `json:"status"`
	SignedBy string `json:"signedBy"`
	SignedAt string `json:"signedAt"`
	Pending  string `json:"pending"`
	Complete string `json:"complete"`
}

type EvaluationConfirmLabels struct {
	SignOff              string `json:"signOff"`
	SignOffMessage       string `json:"signOffMessage"`
	Archive              string `json:"archive"`
	ArchiveMessage       string `json:"archiveMessage"`
	Delete               string `json:"delete"`
	DeleteMessage        string `json:"deleteMessage"`
	BulkArchive          string `json:"bulkArchive"`
	BulkArchiveMessage   string `json:"bulkArchiveMessage"`
	BulkDelete           string `json:"bulkDelete"`
	BulkDeleteMessage    string `json:"bulkDeleteMessage"`
}

type EvaluationErrorLabels struct {
	PermissionDenied  string `json:"permissionDenied"`
	InvalidFormData   string `json:"invalidFormData"`
	NotFound          string `json:"notFound"`
	IDRequired        string `json:"idRequired"`
	NoPermission      string `json:"noPermission"`
	TemplateRequired  string `json:"templateRequired"`
	AlreadySignedOff  string `json:"alreadySignedOff"`
	NotSubmitted      string `json:"notSubmitted"`
}

type EvaluationPortalLabels struct {
	Heading        string `json:"heading"`
	Caption        string `json:"caption"`
	StartReview    string `json:"startReview"`
	ViewLastReview string `json:"viewLastReview"`
	Empty          string `json:"empty"`
	RatingLabel    string `json:"ratingLabel"`
	LastReview     string `json:"lastReview"`
	NoReview       string `json:"noReview"`
}

// DefaultEvaluationLabels returns EvaluationLabels with sensible English defaults.
func DefaultEvaluationLabels() EvaluationLabels {
	return EvaluationLabels{
		Page: EvaluationPageLabels{
			Heading:          "Reviews",
			Caption:          "Track and manage performance evaluations",
			HeadingDraft:     "Draft Reviews",
			HeadingSubmitted: "Submitted Reviews",
			HeadingSignedOff: "Signed-Off Reviews",
			HeadingArchived:  "Archived Reviews",
		},
		Buttons: EvaluationButtonLabels{
			StartReview: "Start Review",
			SaveDraft:   "Save Draft",
			Submit:      "Sign off",
			SignOff:     "Sign off",
			Archive:     "Archive",
			Delete:      "Delete",
		},
		Columns: EvaluationColumnLabels{
			Associate:   "Associate",
			Client:      "Client",
			Period:      "Period",
			Type:        "Type",
			Overall:     "Overall Score",
			Status:      "Status",
			SubmittedAt: "Submitted",
		},
		Status: EvaluationStatusLabels{
			Draft:     "Draft",
			Submitted: "Submitted",
			SignedOff: "Signed Off",
			Archived:  "Archived",
		},
		Tabs: EvaluationTabLabels{
			All:        "All",
			Draft:      "Draft",
			Submitted:  "Submitted",
			SignedOff:  "Signed Off",
			Archived:   "Archived",
			Info:       "Information",
			Scores:     "Scores",
			SignOff:    "Sign Off",
			AuditTrail: "Audit Trail",
		},
		Empty: EvaluationEmptyLabels{
			Title:            "No reviews found",
			Message:          "No evaluations to display.",
			DraftTitle:       "No draft reviews",
			DraftMessage:     "Start a new review to get started.",
			SubmittedTitle:   "No submitted reviews",
			SubmittedMessage: "Completed reviews will appear here.",
			SignedOffTitle:   "No signed-off reviews",
			SignedOffMessage: "Reviews awaiting or completed sign-off will appear here.",
		},
		Form: EvaluationFormLabels{
			EvaluationType:       "Review Type",
			EvaluationTypePH:     "Select review type",
			Template:             "Evaluation Template",
			TemplatePlaceholder:  "Select a template",
			PeriodStart:          "Period Start",
			PeriodEnd:            "Period End",
			Narrative:            "Narrative",
			NarrativePlaceholder: "Optional comments or summary...",
			Associate:            "Associate",
			Client:               "Client",
			SectionHeader:        "Evaluation",
			SectionRubric:        "Rubric",
			SectionNarrative:     "Narrative",
		},
		Dimension: EvaluationDimensionLabels{
			NotScored: "(not scored)",
			PassFail: EvaluationDimensionPassFail{
				Pass: "Pass",
				Fail: "Fail",
			},
			RatingBar: EvaluationDimensionRatingBar{
				AriaLabel: "Score for {{.Name}}: {{.Value}} of {{.Max}}",
			},
		},
		Actions: EvaluationActionLabels{
			View:    "View Review",
			Edit:    "Edit Review",
			SignOff: "Sign Off",
			Archive: "Archive",
			Delete:  "Delete",
		},
		BulkActions: EvaluationBulkLabels{
			Archive: "Archive Selected",
			Delete:  "Delete Selected",
		},
		Detail: EvaluationDetailLabels{
			PageTitle:    "Review Details",
			Associate:    "Associate",
			Client:       "Client",
			Period:       "Period",
			Type:         "Type",
			Template:     "Template",
			OverallScore: "Overall Score",
			Status:       "Status",
			SubmittedAt:  "Submitted",
			SignedOffAt:  "Signed Off",
			SignedOffBy:  "Signed Off By",
			Narrative:    "Narrative",
			Scores: EvaluationDetailScores{
				Heading:       "Scores",
				Criterion:     "Criterion",
				Weight:        "Weight",
				Score:         "Score",
				Comment:       "Comment",
				WeightedTotal: "Weighted Total",
				Empty:         "No responses recorded.",
			},
			SignOff: EvaluationDetailSignOff{
				Heading:  "Sign Off",
				Status:   "Status",
				SignedBy: "Signed By",
				SignedAt: "Signed At",
				Pending:  "Awaiting sign-off",
				Complete: "Signed off",
			},
		},
		Confirm: EvaluationConfirmLabels{
			SignOff:            "Sign Off Review",
			SignOffMessage:     "Sign off on this review for \"%s\"? This action cannot be undone.",
			Archive:            "Archive Review",
			ArchiveMessage:     "Are you sure you want to archive this review?",
			Delete:             "Delete Review",
			DeleteMessage:      "Are you sure you want to delete this review? This action cannot be undone.",
			BulkArchive:        "Archive Selected",
			BulkArchiveMessage: "Archive {count} review(s)? This action cannot be undone.",
			BulkDelete:         "Delete Selected",
			BulkDeleteMessage:  "Delete {count} review(s)? This action cannot be undone.",
		},
		Errors: EvaluationErrorLabels{
			PermissionDenied: "You do not have permission to perform this action",
			InvalidFormData:  "Invalid form data. Please check your inputs and try again.",
			NotFound:         "Review not found",
			IDRequired:       "Review ID is required",
			NoPermission:     "No permission",
			TemplateRequired: "An evaluation template is required for performance reviews",
			AlreadySignedOff: "This review has already been signed off",
			NotSubmitted:     "Only submitted reviews can be signed off",
		},
		Portal: EvaluationPortalLabels{
			Heading:        "Rate My Team",
			Caption:        "Submit performance reviews for your team members",
			StartReview:    "Start Review",
			ViewLastReview: "View Last Review",
			Empty:          "No team members to review at this time.",
			RatingLabel:    "Rating",
			LastReview:     "Last review",
			NoReview:       "No review yet",
		},
	}
}

// EvaluationResponseLabels holds all translatable strings for evaluation responses.
type EvaluationResponseLabels struct {
	Columns EvaluationResponseColumnLabels `json:"columns"`
	Empty   EvaluationResponseEmptyLabels  `json:"empty"`
	Form    EvaluationResponseFormLabels   `json:"form"`
	Detail  EvaluationResponseDetailLabels `json:"detail"`
	Errors  EvaluationResponseErrorLabels  `json:"errors"`
}

type EvaluationResponseColumnLabels struct {
	Criterion string `json:"criterion"`
	Weight    string `json:"weight"`
	Score     string `json:"score"`
	Comment   string `json:"comment"`
}

type EvaluationResponseEmptyLabels struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type EvaluationResponseFormLabels struct {
	Criterion           string                        `json:"criterion"`
	Score               string                        `json:"score"`
	Comment             string                        `json:"comment"`
	CommentPlaceholder  string                        `json:"commentPlaceholder"`
	PassFail            EvaluationResponsePassFail    `json:"passFail"`
}

type EvaluationResponsePassFail struct {
	Pass string `json:"pass"`
	Fail string `json:"fail"`
}

type EvaluationResponseDetailLabels struct {
	WeightedTotal string `json:"weightedTotal"`
	ScoreSummary  string `json:"scoreSummary"`
	NotScored     string `json:"notScored"`
}

type EvaluationResponseErrorLabels struct {
	NotFound         string `json:"notFound"`
	PermissionDenied string `json:"permissionDenied"`
	InvalidFormData  string `json:"invalidFormData"`
}

// DefaultEvaluationResponseLabels returns EvaluationResponseLabels with sensible English defaults.
func DefaultEvaluationResponseLabels() EvaluationResponseLabels {
	return EvaluationResponseLabels{
		Columns: EvaluationResponseColumnLabels{
			Criterion: "Criterion",
			Weight:    "Weight",
			Score:     "Score",
			Comment:   "Comment",
		},
		Empty: EvaluationResponseEmptyLabels{
			Title:   "No responses",
			Message: "No evaluation responses have been recorded.",
		},
		Form: EvaluationResponseFormLabels{
			Criterion:          "Criterion",
			Score:              "Score",
			Comment:            "Comment",
			CommentPlaceholder: "Optional comment...",
			PassFail: EvaluationResponsePassFail{
				Pass: "Pass",
				Fail: "Fail",
			},
		},
		Detail: EvaluationResponseDetailLabels{
			WeightedTotal: "Weighted Total",
			ScoreSummary:  "Score Summary",
			NotScored:     "(not scored)",
		},
		Errors: EvaluationResponseErrorLabels{
			NotFound:         "Evaluation response not found",
			PermissionDenied: "You do not have permission to perform this action",
			InvalidFormData:  "Invalid form data. Please check your inputs and try again.",
		},
	}
}
