package v1

// ---------------------------------------------------------------------------
// PricePlan form labels
// ---------------------------------------------------------------------------

// PricePlanFormLabels holds translatable labels for the PricePlan create/edit
// drawer form. Load from price_plan.json → "price_plan.form" subtree.
type PricePlanFormLabels struct {
	// Section headers
	SectionBasic   string `json:"sectionBasic"`
	SectionPricing string `json:"sectionPricing"`

	// BillingKind enum value labels
	BillingKindLabel    string `json:"billingKindLabel"`
	BillingKindOneTime  string `json:"billingKindOneTime"`
	BillingKindRecurring string `json:"billingKindRecurring"`
	BillingKindContract  string `json:"billingKindContract"`

	// AmountBasis enum value labels
	AmountBasisLabel             string `json:"amountBasisLabel"`
	AmountBasisPerCycle          string `json:"amountBasisPerCycle"`
	AmountBasisTotalPackage      string `json:"amountBasisTotalPackage"`
	AmountBasisDerivedFromLines  string `json:"amountBasisDerivedFromLines"`

	// Billing cycle field labels
	BillingCycleLabel       string `json:"billingCycleLabel"`
	BillingCyclePlaceholder string `json:"billingCyclePlaceholder"`

	// Default term field labels
	DefaultTermLabel           string `json:"defaultTermLabel"`
	DefaultTermPlaceholder     string `json:"defaultTermPlaceholder"`
	DefaultTermOpenEndedHelp   string `json:"defaultTermOpenEndedHelp"`
}

// ---------------------------------------------------------------------------
// ProductPricePlan form labels
// ---------------------------------------------------------------------------

// ProductPricePlanFormLabels holds translatable labels for the ProductPricePlan
// create/edit drawer form. Load from product_price_plan.json → "product_price_plan.form" subtree.
type ProductPricePlanFormLabels struct {
	// BillingTreatment enum value labels and help text
	BillingTreatmentLabel               string `json:"billingTreatmentLabel"`
	BillingTreatmentRecurring           string `json:"billingTreatmentRecurring"`
	BillingTreatmentRecurringHelp       string `json:"billingTreatmentRecurringHelp"`
	BillingTreatmentOneTimeInitial      string `json:"billingTreatmentOneTimeInitial"`
	BillingTreatmentOneTimeInitialHelp  string `json:"billingTreatmentOneTimeInitialHelp"`
	BillingTreatmentUsageBased          string `json:"billingTreatmentUsageBased"`
	BillingTreatmentUsageBasedHelp      string `json:"billingTreatmentUsageBasedHelp"`

	// Product select field labels
	ProductLabel       string `json:"productLabel"`
	ProductPlaceholder string `json:"productPlaceholder"`

	// Price and currency field labels
	PriceLabel         string `json:"priceLabel"`
	PricePlaceholder   string `json:"pricePlaceholder"`
	CurrencyLabel      string `json:"currencyLabel"`
	CurrencyPlaceholder string `json:"currencyPlaceholder"`

	// Effective date range field labels
	DateStartLabel string `json:"dateStartLabel"`
	DateEndLabel   string `json:"dateEndLabel"`
}

// ---------------------------------------------------------------------------
// PriceSchedule form labels
// ---------------------------------------------------------------------------

// PriceScheduleFormLabels holds translatable labels for the PriceSchedule
// create/edit drawer form. Load from price_schedule.json → "priceSchedule.form" subtree.
type PriceScheduleFormLabels struct {
	// Section headers
	SectionScheduleDetails string `json:"sectionScheduleDetails"`
	SectionDateRange       string `json:"sectionDateRange"`
	SectionLocation        string `json:"sectionLocation"`
}
