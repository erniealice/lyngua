# Payment Terms

Define and manage the payment terms used across your business — from standard net-day schedules to early payment discounts and immediate payment requirements.

## What payment terms are and why they matter

A payment term sets the expectation for when an invoice must be paid and whether a discount applies for early settlement. Common examples:

- **Net 30** — full amount due within 30 days of the invoice date
- **Net 60** — full amount due within 60 days
- **2/10 Net 30** — 2% discount if paid within 10 days, otherwise full amount due within 30 days
- **COD (Cash on Delivery)** — payment collected at the time of delivery
- **Due on Receipt** — payment expected immediately upon receiving the invoice

Accurate payment terms help you manage cash flow, set client and supplier expectations, and automate due-date calculations on invoices and purchase orders.

## What you can do here

- View all configured payment terms and their current status (active / inactive)
- Add new payment terms with a name, short code, type, net days, and optional early payment discount
- Edit existing terms to adjust days, discount rate, or scope
- Set a term as the default so it is pre-selected on new transactions
- Configure entity scope to limit a term to clients, suppliers, or make it available to both
- Delete terms that are no longer needed (terms in use cannot be deleted)

## How payment terms connect to clients and suppliers

Each client and supplier record can have a default payment term assigned. When a new transaction (invoice or purchase order) is created for that entity, the assigned term is pre-filled automatically.

You can always override the term on an individual transaction without changing the entity's default. This lets you handle one-off arrangements without affecting your standard setup.

## Common workflows

**To add a new payment term:** Click the Add button in the toolbar, enter a name (e.g. "Net 30"), a short code (e.g. "N30"), select the type, and set the number of net days. Save to make it available for selection.

**To set a term as the default:** Open the edit drawer for the term you want and enable the Default toggle. Only one term per entity scope can be the default at a time — setting a new default will automatically clear the previous one.

**To configure an early payment discount:** When adding or editing a term of type "discount", enter the discount percentage in basis points (e.g. 200 = 2.00%) and the number of days within which the discount applies.

**To deactivate a term without deleting it:** Edit the term and set it to inactive. Inactive terms no longer appear in transaction dropdowns but are retained for historical records.

## Tips

- **Entity scope** — use scope filtering on the list to quickly see which terms apply to clients, which apply to suppliers, and which are shared. Assign the correct scope when creating a term to avoid it appearing in the wrong context.
- **Display order** — terms are sorted alphabetically by name by default. Use clear, consistent naming (e.g. "Net 07", "Net 14", "Net 30") so they sort predictably in dropdowns.
- **Basis points for discount percent** — discount rates are stored in basis points where 100 = 1.00%. Enter 150 for 1.5%, 200 for 2%, and so on. This avoids floating-point rounding on small percentages.
- **COD and Due on Receipt** — these types typically have net days set to 0. They are useful for walk-in sales or situations where credit is not extended.
