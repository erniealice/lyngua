package v1

// labels_staff.go — Go label structs for entity/staff
// (performance-evaluation layer 12, Phase A).
//
// Root key: "staff" (file staff.json + outsourcing/staff.json tier-3 overlay).
//
// The outsourcing tier overrides Page.Heading → "Associates" and relabels
// employment.contractor/external/subcontractor → "VA" (LBL-2).
// The general tier uses neutral "Staff" / "Contractor" wording.

// StaffLabels holds all translatable strings for the staff module.
type StaffLabels struct {
	Page         StaffPageLabels         `json:"page"`
	Buttons      StaffButtonLabels       `json:"buttons"`
	Columns      StaffColumnLabels       `json:"columns"`
	Availability StaffAvailabilityLabels `json:"availability"`
	Employment   StaffEmploymentLabels   `json:"employment"`
	Rating       StaffRatingLabels       `json:"rating"`
	Form         StaffFormLabels         `json:"form"`
	Empty        StaffEmptyLabels        `json:"empty"`
	Actions      StaffActionLabels       `json:"actions"`
	Detail       StaffDetailLabels       `json:"detail"`
	Tabs         StaffTabLabels          `json:"tabs"`
	Errors       StaffErrorLabels        `json:"errors"`
}

type StaffPageLabels struct {
	Heading         string `json:"heading"`
	Caption         string `json:"caption"`
	HeadingActive   string `json:"headingActive"`
	HeadingInactive string `json:"headingInactive"`
}

type StaffButtonLabels struct {
	AddStaff string `json:"addStaff"`
}

type StaffColumnLabels struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Status          string `json:"status"`
	EmploymentType  string `json:"employmentType"`
	Seniority       string `json:"seniority"`
	Rating          string `json:"rating"`
	EmploymentStart string `json:"employmentStart"`
}

// StaffAvailabilityLabels maps staff.status availability values.
// Mirrors available | assigned | bench | offboarded (entities.md §E1, Q-IDENTITY-1).
// The `status` field on entity/staff carries AVAILABILITY only — never employment model.
type StaffAvailabilityLabels struct {
	Available   string `json:"available"`
	Assigned    string `json:"assigned"`
	Bench       string `json:"bench"`
	Offboarded  string `json:"offboarded"`
}

// StaffEmploymentLabels maps employment_type enum values.
// The outsourcing tier-3 overlay relabels Contractor/External/Subcontractor → "VA".
// One canonical field name: employment_type (Q-EMPLOYMENT-TYPE-1, never overloaded).
type StaffEmploymentLabels struct {
	Employed      string `json:"employed"`
	Contractor    string `json:"contractor"`
	External      string `json:"external"`
	Partner       string `json:"partner"`
	Retained      string `json:"retained"`
	Subcontractor string `json:"subcontractor"`
}

type StaffRatingLabels struct {
	Label    string `json:"label"`
	NoRating string `json:"noRating"`
}

type StaffFormLabels struct {
	Status               string `json:"status"`
	StatusPlaceholder    string `json:"statusPlaceholder"`
	EmploymentType       string `json:"employmentType"`
	EmploymentTypePH     string `json:"employmentTypePlaceholder"`
	Seniority            string `json:"seniority"`
	SeniorityPlaceholder string `json:"seniorityPlaceholder"`
	EmploymentStart      string `json:"employmentStart"`
	EmploymentEnd        string `json:"employmentEnd"`
}

type StaffEmptyLabels struct {
	Title           string `json:"title"`
	Message         string `json:"message"`
	ActiveTitle     string `json:"activeTitle"`
	ActiveMessage   string `json:"activeMessage"`
	InactiveTitle   string `json:"inactiveTitle"`
	InactiveMessage string `json:"inactiveMessage"`
}

type StaffActionLabels struct {
	View   string `json:"view"`
	Edit   string `json:"edit"`
	Delete string `json:"delete"`
}

type StaffDetailLabels struct {
	PageTitle       string `json:"pageTitle"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Status          string `json:"status"`
	EmploymentType  string `json:"employmentType"`
	Seniority       string `json:"seniority"`
	EmploymentStart string `json:"employmentStart"`
	EmploymentEnd   string `json:"employmentEnd"`
	Rating          string `json:"rating"`
	CreatedDate     string `json:"createdDate"`
	ModifiedDate    string `json:"modifiedDate"`
}

type StaffTabLabels struct {
	Info        string `json:"info"`
	Evaluations string `json:"evaluations"`
	AuditTrail  string `json:"auditTrail"`
}

type StaffErrorLabels struct {
	PermissionDenied string `json:"permissionDenied"`
	InvalidFormData  string `json:"invalidFormData"`
	NotFound         string `json:"notFound"`
}

// DefaultStaffLabels returns StaffLabels with sensible English defaults.
// The outsourcing tier-3 overlay (outsourcing/staff.json) relabels
// Page.Heading → "Associates" at load time via the cascade.
func DefaultStaffLabels() StaffLabels {
	return StaffLabels{
		Page: StaffPageLabels{
			Heading:         "Staff",
			Caption:         "Manage your staff members",
			HeadingActive:   "Active Staff",
			HeadingInactive: "Inactive Staff",
		},
		Buttons: StaffButtonLabels{
			AddStaff: "Add Staff",
		},
		Columns: StaffColumnLabels{
			Name:            "Name",
			Email:           "Email",
			Status:          "Availability",
			EmploymentType:  "Employment Type",
			Seniority:       "Seniority",
			Rating:          "Rating",
			EmploymentStart: "Start Date",
		},
		Availability: StaffAvailabilityLabels{
			Available:  "Available",
			Assigned:   "Assigned",
			Bench:      "Bench",
			Offboarded: "Offboarded",
		},
		Employment: StaffEmploymentLabels{
			Employed:      "Employed",
			Contractor:    "Contractor",
			External:      "External",
			Partner:       "Partner",
			Retained:      "Retained",
			Subcontractor: "Subcontractor",
		},
		Rating: StaffRatingLabels{
			Label:    "Rating",
			NoRating: "—",
		},
		Form: StaffFormLabels{
			Status:               "Availability",
			StatusPlaceholder:    "Select availability status",
			EmploymentType:       "Employment Type",
			EmploymentTypePH:     "Select employment type",
			Seniority:            "Seniority",
			SeniorityPlaceholder: "e.g. Senior, Mid, Junior",
			EmploymentStart:      "Employment Start",
			EmploymentEnd:        "Employment End",
		},
		Empty: StaffEmptyLabels{
			Title:           "No staff found",
			Message:         "No staff members to display.",
			ActiveTitle:     "No active staff",
			ActiveMessage:   "Add your first staff member to get started.",
			InactiveTitle:   "No inactive staff",
			InactiveMessage: "Offboarded staff will appear here.",
		},
		Actions: StaffActionLabels{
			View:   "View Staff",
			Edit:   "Edit Staff",
			Delete: "Delete Staff",
		},
		Detail: StaffDetailLabels{
			PageTitle:       "Staff Details",
			Name:            "Name",
			Email:           "Email",
			Status:          "Availability",
			EmploymentType:  "Employment Type",
			Seniority:       "Seniority",
			EmploymentStart: "Employment Start",
			EmploymentEnd:   "Employment End",
			Rating:          "Latest Rating",
			CreatedDate:     "Created",
			ModifiedDate:    "Last Modified",
		},
		Tabs: StaffTabLabels{
			Info:        "Information",
			Evaluations: "Reviews",
			AuditTrail:  "Audit Trail",
		},
		Errors: StaffErrorLabels{
			PermissionDenied: "You do not have permission to perform this action",
			InvalidFormData:  "Invalid form data. Please check your inputs and try again.",
			NotFound:         "Staff member not found",
		},
	}
}
