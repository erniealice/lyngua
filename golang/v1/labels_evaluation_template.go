package v1

// labels_evaluation_template.go — Go label structs for evaluation_template and
// evaluation_template_item domains (performance-evaluation layer 12).
//
// Root keys (snake_case, matching JSON files):
//   evaluation_template.json      → EvaluationTemplateLabels      (root key "evaluation_template")
//   evaluation_template_item.json → EvaluationTemplateItemLabels   (root key "evaluation_template_item")

// EvaluationTemplateLabels holds all translatable strings for the evaluation template module.
type EvaluationTemplateLabels struct {
	Page        EvaluationTemplatePageLabels    `json:"page"`
	Buttons     EvaluationTemplateButtonLabels  `json:"buttons"`
	Columns     EvaluationTemplateColumnLabels  `json:"columns"`
	Status      EvaluationTemplateStatusLabels  `json:"status"`
	Tabs        EvaluationTemplateTabLabels     `json:"tabs"`
	Empty       EvaluationTemplateEmptyLabels   `json:"empty"`
	Form        EvaluationTemplateFormLabels    `json:"form"`
	Actions     EvaluationTemplateActionLabels  `json:"actions"`
	BulkActions EvaluationTemplateBulkLabels    `json:"bulk_actions"`
	Detail      EvaluationTemplateDetailLabels  `json:"detail"`
	Confirm     EvaluationTemplateConfirmLabels `json:"confirm"`
	Errors      EvaluationTemplateErrorLabels   `json:"errors"`
}

type EvaluationTemplatePageLabels struct {
	Heading           string `json:"heading"`
	Caption           string `json:"caption"`
	HeadingDraft      string `json:"heading_draft"`
	HeadingActive     string `json:"heading_active"`
	HeadingDeprecated string `json:"heading_deprecated"`
}

type EvaluationTemplateButtonLabels struct {
	AddTemplate string `json:"add_template"`
	Activate    string `json:"activate"`
	Deprecate   string `json:"deprecate"`
	Clone       string `json:"clone"`
}

type EvaluationTemplateColumnLabels struct {
	Name             string `json:"name"`
	EvaluationType   string `json:"evaluation_type"`
	RelationshipType string `json:"relationship_type"`
	Version          string `json:"version"`
	Status           string `json:"status"`
	Visibility       string `json:"visibility"`
	ItemCount        string `json:"item_count"`
	Created          string `json:"created"`
}

// EvaluationTemplateStatusLabels maps EvaluationTemplateStatus enum values.
// Mirrors DRAFT | ACTIVE | DEPRECATED (new-entity enum convention, per entities.md).
type EvaluationTemplateStatusLabels struct {
	Draft      string `json:"draft"`
	Active     string `json:"active"`
	Deprecated string `json:"deprecated"`
}

type EvaluationTemplateTabLabels struct {
	All        string `json:"all"`
	Draft      string `json:"draft"`
	Active     string `json:"active"`
	Deprecated string `json:"deprecated"`
	Info       string `json:"info"`
	Items      string `json:"items"`
}

type EvaluationTemplateEmptyLabels struct {
	Title         string `json:"title"`
	Message       string `json:"message"`
	ActiveTitle   string `json:"active_title"`
	ActiveMessage string `json:"active_message"`
	DraftTitle    string `json:"draft_title"`
	DraftMessage  string `json:"draft_message"`
}

type EvaluationTemplateFormLabels struct {
	Name                   string `json:"name"`
	NamePlaceholder        string `json:"name_placeholder"`
	Description            string `json:"description"`
	DescriptionPlaceholder string `json:"description_placeholder"`
	EvaluationType         string `json:"evaluation_type"`
	EvaluationTypePH       string `json:"evaluation_type_placeholder"`
	RelationshipType       string `json:"relationship_type"`
	RelationshipTypePH     string `json:"relationship_type_placeholder"`
	VisibilityType         string `json:"visibility_type"`
	VisibilityTypePH       string `json:"visibility_type_placeholder"`
}

type EvaluationTemplateActionLabels struct {
	View      string `json:"view"`
	Edit      string `json:"edit"`
	Activate  string `json:"activate"`
	Deprecate string `json:"deprecate"`
	Clone     string `json:"clone"`
	Delete    string `json:"delete"`
}

type EvaluationTemplateBulkLabels struct {
	Deprecate string `json:"deprecate"`
}

type EvaluationTemplateDetailLabels struct {
	PageTitle        string                        `json:"page_title"`
	Name             string                        `json:"name"`
	Description      string                        `json:"description"`
	EvaluationType   string                        `json:"evaluation_type"`
	RelationshipType string                        `json:"relationship_type"`
	Version          string                        `json:"version"`
	Status           string                        `json:"status"`
	Visibility       string                        `json:"visibility"`
	CreatedDate      string                        `json:"created_date"`
	ModifiedDate     string                        `json:"modified_date"`
	Items            EvaluationTemplateDetailItems `json:"items"`
}

type EvaluationTemplateDetailItems struct {
	Heading string `json:"heading"`
	AddItem string `json:"add_item"`
	Reorder string `json:"reorder"`
	Empty   string `json:"empty"`
}

type EvaluationTemplateConfirmLabels struct {
	Activate             string `json:"activate"`
	ActivateMessage      string `json:"activate_message"`
	Deprecate            string `json:"deprecate"`
	DeprecateMessage     string `json:"deprecate_message"`
	Clone                string `json:"clone"`
	CloneMessage         string `json:"clone_message"`
	Delete               string `json:"delete"`
	DeleteMessage        string `json:"delete_message"`
	BulkDeprecate        string `json:"bulk_deprecate"`
	BulkDeprecateMessage string `json:"bulk_deprecate_message"`
}

type EvaluationTemplateErrorLabels struct {
	PermissionDenied        string `json:"permission_denied"`
	InvalidFormData         string `json:"invalid_form_data"`
	NotFound                string `json:"not_found"`
	IDRequired              string `json:"id_required"`
	NoPermission            string `json:"no_permission"`
	WeightedNonNumericGuard string `json:"weighted_non_numeric_guard"`
	InUse                   string `json:"in_use"`
}

// DefaultEvaluationTemplateLabels returns EvaluationTemplateLabels with sensible English defaults.
func DefaultEvaluationTemplateLabels() EvaluationTemplateLabels {
	return EvaluationTemplateLabels{
		Page: EvaluationTemplatePageLabels{
			Heading:           "Evaluation Templates",
			Caption:           "Manage reusable rubric templates for performance reviews",
			HeadingDraft:      "Draft Templates",
			HeadingActive:     "Active Templates",
			HeadingDeprecated: "Deprecated Templates",
		},
		Buttons: EvaluationTemplateButtonLabels{
			AddTemplate: "New Template",
			Activate:    "Activate",
			Deprecate:   "Deprecate",
			Clone:       "Clone",
		},
		Columns: EvaluationTemplateColumnLabels{
			Name:             "Name",
			EvaluationType:   "Eval Type",
			RelationshipType: "Relationship",
			Version:          "Version",
			Status:           "Status",
			Visibility:       "Visibility",
			ItemCount:        "Items",
			Created:          "Created",
		},
		Status: EvaluationTemplateStatusLabels{
			Draft:      "Draft",
			Active:     "Active",
			Deprecated: "Deprecated",
		},
		Tabs: EvaluationTemplateTabLabels{
			All:        "All",
			Draft:      "Draft",
			Active:     "Active",
			Deprecated: "Deprecated",
			Info:       "Information",
			Items:      "Rubric Items",
		},
		Empty: EvaluationTemplateEmptyLabels{
			Title:         "No templates found",
			Message:       "No evaluation templates to display.",
			ActiveTitle:   "No active templates",
			ActiveMessage: "Create and activate a template to make it available for reviews.",
			DraftTitle:    "No draft templates",
			DraftMessage:  "Draft templates will appear here.",
		},
		Form: EvaluationTemplateFormLabels{
			Name:                   "Template Name",
			NamePlaceholder:        "Enter template name",
			Description:            "Description",
			DescriptionPlaceholder: "Optional description...",
			EvaluationType:         "Evaluation Type",
			EvaluationTypePH:       "Select evaluation type",
			RelationshipType:       "Relationship Type",
			RelationshipTypePH:     "Select relationship type",
			VisibilityType:         "Visibility",
			VisibilityTypePH:       "Select visibility",
		},
		Actions: EvaluationTemplateActionLabels{
			View:      "View Template",
			Edit:      "Edit Template",
			Activate:  "Activate",
			Deprecate: "Deprecate",
			Clone:     "Clone",
			Delete:    "Delete",
		},
		BulkActions: EvaluationTemplateBulkLabels{
			Deprecate: "Deprecate Selected",
		},
		Detail: EvaluationTemplateDetailLabels{
			PageTitle:        "Template Details",
			Name:             "Name",
			Description:      "Description",
			EvaluationType:   "Evaluation Type",
			RelationshipType: "Relationship Type",
			Version:          "Version",
			Status:           "Status",
			Visibility:       "Visibility",
			CreatedDate:      "Created",
			ModifiedDate:     "Last Modified",
			Items: EvaluationTemplateDetailItems{
				Heading: "Rubric Items",
				AddItem: "Add Question",
				Reorder: "Drag to reorder",
				Empty:   "No items yet — add a question to build the rubric.",
			},
		},
		Confirm: EvaluationTemplateConfirmLabels{
			Activate:             "Activate Template",
			ActivateMessage:      "Activate \"%s\"? It will be available for new reviews.",
			Deprecate:            "Deprecate Template",
			DeprecateMessage:     "Deprecate \"%s\"? It will no longer be available for new reviews. Existing draft reviews can still be submitted.",
			Clone:                "Clone Template",
			CloneMessage:         "Clone \"%s\"? A new draft copy will be created.",
			Delete:               "Delete Template",
			DeleteMessage:        "Are you sure you want to delete \"%s\"? This action cannot be undone.",
			BulkDeprecate:        "Deprecate Selected",
			BulkDeprecateMessage: "Deprecate {count} template(s)?",
		},
		Errors: EvaluationTemplateErrorLabels{
			PermissionDenied:        "You do not have permission to perform this action",
			InvalidFormData:         "Invalid form data. Please check your inputs and try again.",
			NotFound:                "Evaluation template not found",
			IDRequired:              "Template ID is required",
			NoPermission:            "No permission",
			WeightedNonNumericGuard: "Cannot activate: a rubric item with a non-zero weight has a non-numeric criteria type. Remove the weight or change the criteria type.",
			InUse:                   "Cannot delete: this template is in use by existing reviews",
		},
	}
}

// EvaluationTemplateItemLabels holds all translatable strings for rubric items.
type EvaluationTemplateItemLabels struct {
	Page    EvaluationTemplateItemPageLabels   `json:"page"`
	Buttons EvaluationTemplateItemButtonLabels `json:"buttons"`
	Columns EvaluationTemplateItemColumnLabels `json:"columns"`
	Empty   EvaluationTemplateItemEmptyLabels  `json:"empty"`
	Form    EvaluationTemplateItemFormLabels   `json:"form"`
	Actions EvaluationTemplateItemActionLabels `json:"actions"`
	Errors  EvaluationTemplateItemErrorLabels  `json:"errors"`
}

type EvaluationTemplateItemPageLabels struct {
	Heading string `json:"heading"`
}

type EvaluationTemplateItemButtonLabels struct {
	AddItem    string `json:"add_item"`
	RemoveItem string `json:"remove_item"`
}

type EvaluationTemplateItemColumnLabels struct {
	Criterion    string `json:"criterion"`
	CriteriaType string `json:"criteria_type"`
	Weight       string `json:"weight"`
	Required     string `json:"required"`
	Order        string `json:"order"`
}

type EvaluationTemplateItemEmptyLabels struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type EvaluationTemplateItemFormLabels struct {
	Criterion            string `json:"criterion"`
	CriterionPlaceholder string `json:"criterion_placeholder"`
	CriterionInfo        string `json:"criterion_info"`
	QuestionLabel        string `json:"question_label"`
	QuestionLabelPH      string `json:"question_label_placeholder"`
	QuestionPrompt       string `json:"question_prompt"`
	QuestionPromptPH     string `json:"question_prompt_placeholder"`
	Weight               string `json:"weight"`
	WeightInfo           string `json:"weight_info"`
	Required             string `json:"required"`
	RequiredInfo         string `json:"required_info"`
	CriteriaType         string `json:"criteria_type"`
	CriteriaTypeReadOnly string `json:"criteria_type_read_only"`
}

type EvaluationTemplateItemActionLabels struct {
	Edit    string `json:"edit"`
	Remove  string `json:"remove"`
	Reorder string `json:"reorder"`
}

type EvaluationTemplateItemErrorLabels struct {
	PermissionDenied   string `json:"permission_denied"`
	InvalidFormData    string `json:"invalid_form_data"`
	NotFound           string `json:"not_found"`
	IDRequired         string `json:"id_required"`
	CriterionRequired  string `json:"criterion_required"`
	DuplicateCriterion string `json:"duplicate_criterion"`
}

// DefaultEvaluationTemplateItemLabels returns EvaluationTemplateItemLabels with sensible English defaults.
func DefaultEvaluationTemplateItemLabels() EvaluationTemplateItemLabels {
	return EvaluationTemplateItemLabels{
		Page: EvaluationTemplateItemPageLabels{
			Heading: "Rubric Items",
		},
		Buttons: EvaluationTemplateItemButtonLabels{
			AddItem:    "Add Question",
			RemoveItem: "Remove",
		},
		Columns: EvaluationTemplateItemColumnLabels{
			Criterion:    "Criterion",
			CriteriaType: "Type",
			Weight:       "Weight",
			Required:     "Required",
			Order:        "Order",
		},
		Empty: EvaluationTemplateItemEmptyLabels{
			Title:   "No rubric items",
			Message: "Add a question to start building the rubric.",
		},
		Form: EvaluationTemplateItemFormLabels{
			Criterion:            "Criterion",
			CriterionPlaceholder: "Search criteria...",
			CriterionInfo:        "Select the outcome criterion this item evaluates.",
			QuestionLabel:        "Question Label",
			QuestionLabelPH:      "Override criterion name (optional)",
			QuestionPrompt:       "Question Prompt",
			QuestionPromptPH:     "Instructions shown to the evaluator (optional)",
			Weight:               "Weight",
			WeightInfo:           "Relative importance when computing the overall score. Leave blank to use the criterion default.",
			Required:             "Required",
			RequiredInfo:         "Whether this question must be answered before submitting.",
			CriteriaType:         "Type",
			CriteriaTypeReadOnly: "Set by the linked criterion",
		},
		Actions: EvaluationTemplateItemActionLabels{
			Edit:    "Edit Item",
			Remove:  "Remove Item",
			Reorder: "Reorder",
		},
		Errors: EvaluationTemplateItemErrorLabels{
			PermissionDenied:   "You do not have permission to perform this action",
			InvalidFormData:    "Invalid form data. Please check your inputs and try again.",
			NotFound:           "Rubric item not found",
			IDRequired:         "Item ID is required",
			CriterionRequired:  "A criterion must be selected",
			DuplicateCriterion: "This criterion is already in the template",
		},
	}
}
