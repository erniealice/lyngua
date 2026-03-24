# Sales → Revenue Full Rename — Progress Log

**Plan:** [plan.md](./plan.md)
**Started:** 2026-03-12
**Branch:** `dev/20260312-sales-to-revenue`

---

## Phase 1: Rename translation JSON files — NOT STARTED

- [ ] Rename `retail/sale.json` → `retail/revenue.json`, root key `"sales"` → `"revenue"`
- [ ] Rename `service/sale.json` → `service/revenue.json`, root key `"sales"` → `"revenue"`
- [ ] Rename `professional/sale.json` → `professional/revenue.json`, root key `"sales"` → `"revenue"`
- [ ] Rename `"sales"` key → `"revenue"` in `professional/route.json`
- [ ] Rename `"sales"` key → `"revenue"` in `service/route.json`

---

## Phase 2: Rename Go types in centymo-golang — NOT STARTED

- [ ] `labels.go`: `Sales*Labels` → `Revenue*Labels` (15 types)
- [ ] `routes_config.go`: `SalesRoutes` → `RevenueRoutes`, `DefaultSalesRoutes` → `DefaultRevenueRoutes`
- [ ] `routes.go`: `Sales*URL` → `Revenue*URL` (25 constants)
- [ ] `views/revenue/list/page.go`: update type references
- [ ] `views/revenue/detail/page.go`: update type references
- [ ] `views/revenue/detail/line_items.go`: update type references
- [ ] `views/revenue/action/action.go`: update type references
- [ ] `views/revenue/action/payment.go`: update type references
- [ ] `views/revenue/action/invoice_download.go`: update type references
- [ ] `views/revenue/dashboard/page.go`: update type references
- [ ] `views/revenue/settings/deps.go`: update type references
- [ ] `views/revenue/settings/page.go`: update type references

---

## Phase 3: Rename presentation dirs + update app consumers — NOT STARTED

- [ ] Rename `retail-admin/internal/presentation/sales/` → `revenue/`, update package name
- [ ] Rename `service-admin/internal/presentation/sales/` → `revenue/`, update package name
- [ ] Update `retail-admin/composition/views.go`: import alias + field names
- [ ] Update `service-admin/composition/views.go`: import alias + field names
- [ ] Update `retail-admin/composition/container.go`: translation paths + type/var names
- [ ] Update `service-admin/composition/container.go`: translation paths + type/var names + route.json path
- [ ] Update `service-admin/presentation/home/module.go`: `SalesRoutes` → `RevenueRoutes`

---

## Phase 4: Build verification — NOT STARTED

- [ ] service-admin builds successfully
- [ ] retail-admin builds successfully
- [ ] service-admin E2E tests pass (61 tests)
- [ ] Display text verified: retail="Sales", service="Bookings"

---

## Summary

- **Phases complete:** 0 / 4
- **Files modified:** 0 / 24

---

## Skipped / Deferred

| Item | Reason |
|------|--------|
| — | — |

---

## How to Resume

To continue this work:
1. Read this progress file and the [plan](./plan.md)
2. Check git status for uncommitted changes
3. Start from the first incomplete phase
4. Update checkboxes as you complete steps

Key files to read first:
- `packages/centymo-golang/labels.go:298` (SalesLabels type definition)
- `packages/centymo-golang/routes_config.go:322` (SalesRoutes type definition)
- `packages/centymo-golang/routes.go:152-189` (Sales*URL constants)
- `apps/service-admin/internal/composition/container.go:218-278` (translation loading)
- `apps/retail-admin/internal/composition/container.go:266-267` (translation loading)
