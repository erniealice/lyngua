# Sales → Revenue Full Rename — Design Plan

**Date:** 2026-03-12
**Branch:** `dev/20260312-sales-to-revenue`
**Status:** Draft
**App/Package:** lyngua, centymo-golang, service-admin, retail-admin

---

## Overview

Full rename of the "sales" concept to "revenue" across translation files, Go types, route constants, and presentation directories. "Sales" is a business-type-specific label (retail → "Sales", service → "Bookings") — the canonical generic term should be **"revenue"**, aligning with the proto layer (`revenuepb`) and existing view directories (`views/revenue/`).

---

## Motivation

The codebase has a split identity:
- **Already "revenue"**: proto (`revenuepb`), centymo views (`views/revenue/`), use case funcs (`ListRevenues`, `CreateRevenue`)
- **Still "sales"**: translation files (`sale.json`), Go types (`SalesLabels`, `SalesRoutes`), route constants (`SalesDashboardURL`), presentation dirs (`presentation/sales/`), route.json keys (`"sales"`)

This rename completes the alignment so the entire stack uses "revenue" as the generic concept.

---

## Architecture

### Translation cascade (display text unchanged)

```
retail/sale.json     →  revenue.json, root key "revenue", display text stays "Sales"
service/sale.json    →  revenue.json, root key "revenue", display text stays "Bookings"
professional/sale.json → revenue.json, root key "revenue", display text stays "Bookings"
```

### Route JSON cascade

```
professional/route.json  →  key "sales" → "revenue" (URLs already /revenue/)
service/route.json       →  key "sales" → "revenue" (URLs /bookings/)
```

### Go type renames

```
centymo.SalesLabels         → centymo.RevenueLabels
centymo.SalesRoutes         → centymo.RevenueRoutes
centymo.DefaultSalesRoutes  → centymo.DefaultRevenueRoutes
centymo.Sales*Labels (14 sub-types) → centymo.Revenue*Labels
Sales*URL constants (25)    → Revenue*URL
```

### Directory renames

```
apps/retail-admin/internal/presentation/sales/   → revenue/
apps/service-admin/internal/presentation/sales/  → revenue/
```

---

## Implementation Steps

### Phase 1: Rename translation JSON files (lyngua)

1. **Rename `sale.json` → `revenue.json`** in 3 tiers + change root key `"sales"` → `"revenue"`:
   - `packages/lyngua/translations/en/retail/sale.json`
   - `packages/lyngua/translations/en/service/sale.json`
   - `packages/lyngua/translations/en/professional/sale.json`

2. **Rename `"sales"` key → `"revenue"` in route.json** files:
   - `packages/lyngua/translations/en/professional/route.json`
   - `packages/lyngua/translations/en/service/route.json`

### Phase 2: Rename Go types in centymo-golang

3. **Rename label types** in `packages/centymo-golang/labels.go`:
   - `SalesLabels` → `RevenueLabels`
   - `SalesPageLabels` → `RevenuePageLabels`
   - `SalesButtonLabels` → `RevenueButtonLabels`
   - `SalesColumnLabels` → `RevenueColumnLabels`
   - `SalesEmptyLabels` → `RevenueEmptyLabels`
   - `SalesFormLabels` → `RevenueFormLabels`
   - `SalesActionLabels` → `RevenueActionLabels`
   - `SalesBulkLabels` → `RevenueBulkLabels`
   - `SalesDetailLabels` → `RevenueDetailLabels`
   - `SalesConfirmLabels` → `RevenueConfirmLabels`
   - `SalesErrorLabels` → `RevenueErrorLabels`
   - `SalesDashboardLabels` → `RevenueDashboardLabels`
   - `SalesSettingsLabels` → `RevenueSettingsLabels`
   - Plus `SalesLineItemLabels` and `SalesPaymentLabels` if they exist

4. **Rename route type + constructor** in `packages/centymo-golang/routes_config.go`:
   - `SalesRoutes` → `RevenueRoutes`
   - `DefaultSalesRoutes()` → `DefaultRevenueRoutes()`
   - `(SalesRoutes) RouteMap()` → `(RevenueRoutes) RouteMap()`

5. **Rename route constants** in `packages/centymo-golang/routes.go` (25 constants):
   - `SalesDashboardURL` → `RevenueDashboardURL`
   - `SalesListURL` → `RevenueListURL`
   - `SalesDetailURL` → `RevenueDetailURL`
   - ... (all 25 `Sales*URL` → `Revenue*URL`)
   - Comment: `// Sales (revenue) routes` → `// Revenue routes`

6. **Update all centymo view references** (already under `views/revenue/`):
   - `packages/centymo-golang/views/revenue/list/page.go` — `centymo.SalesLabels` → `centymo.RevenueLabels`, `centymo.SalesRoutes` → `centymo.RevenueRoutes`
   - `packages/centymo-golang/views/revenue/detail/page.go` — same
   - `packages/centymo-golang/views/revenue/detail/line_items.go` — same
   - `packages/centymo-golang/views/revenue/action/action.go` — same
   - `packages/centymo-golang/views/revenue/action/payment.go` — same
   - `packages/centymo-golang/views/revenue/action/invoice_download.go` — same
   - `packages/centymo-golang/views/revenue/dashboard/page.go` — same
   - `packages/centymo-golang/views/revenue/settings/deps.go` — same
   - `packages/centymo-golang/views/revenue/settings/page.go` — `SalesSettingsLabels` → `RevenueSettingsLabels`

### Phase 3: Rename presentation directories + update app consumers

7. **Rename presentation directories**:
   - `apps/retail-admin/internal/presentation/sales/` → `revenue/`
   - `apps/service-admin/internal/presentation/sales/` → `revenue/`
   - Update `package sales` → `package revenue` in both `module.go` files
   - Update all `centymo.SalesLabels` → `centymo.RevenueLabels` and `centymo.SalesRoutes` → `centymo.RevenueRoutes`

8. **Update import aliases in views.go**:
   - `apps/retail-admin/internal/composition/views.go`: `salesmod "ichizen.leapfor.xyz/.../sales"` → `revenuemod "ichizen.leapfor.xyz/.../revenue"`, update all `salesmod.` → `revenuemod.`
   - `apps/service-admin/internal/composition/views.go`: same

9. **Update container.go in both apps**:
   - Translation loading: `"sale.json", "sales"` → `"revenue.json", "revenue"` and `"route.json", "sales"` → `"route.json", "revenue"`
   - Type references: `centymo.SalesLabels` → `centymo.RevenueLabels`, `centymo.SalesRoutes` → `centymo.RevenueRoutes`
   - `DefaultSalesRoutes()` → `DefaultRevenueRoutes()`
   - `ValidateLabels("SalesLabels", ...)` → `ValidateLabels("RevenueLabels", ...)`
   - Struct field names: `SalesLabels` → `RevenueLabels`, `SalesRoutes` → `RevenueRoutes` (in ViewsDeps, TransactionRoutes, etc.)
   - Variable names: `salesLabels` → `revenueLabels`, `salesRoutes` → `revenueRoutes`

10. **Update service-admin home module** (`apps/service-admin/internal/presentation/home/module.go`):
    - `SalesRoutes centymo.SalesRoutes` → `RevenueRoutes centymo.RevenueRoutes`
    - Field access: `m.deps.SalesRoutes.DashboardURL` → `m.deps.RevenueRoutes.DashboardURL`

### Phase 4: Build verification

11. Build both apps and verify no compilation errors
12. Run service-admin E2E tests (61 tests)

---

## File References

| File | Change | Phase |
|------|--------|-------|
| `packages/lyngua/translations/en/retail/sale.json` | Rename to `revenue.json`, root key `"sales"` → `"revenue"` | 1 |
| `packages/lyngua/translations/en/service/sale.json` | Rename to `revenue.json`, root key `"sales"` → `"revenue"` | 1 |
| `packages/lyngua/translations/en/professional/sale.json` | Rename to `revenue.json`, root key `"sales"` → `"revenue"` | 1 |
| `packages/lyngua/translations/en/professional/route.json` | `"sales"` key → `"revenue"` | 1 |
| `packages/lyngua/translations/en/service/route.json` | `"sales"` key → `"revenue"` | 1 |
| `packages/centymo-golang/labels.go` | `Sales*Labels` → `Revenue*Labels` (15 types) | 2 |
| `packages/centymo-golang/routes_config.go` | `SalesRoutes` → `RevenueRoutes`, `DefaultSalesRoutes` → `DefaultRevenueRoutes` (27 refs) | 2 |
| `packages/centymo-golang/routes.go` | `Sales*URL` → `Revenue*URL` (25 constants) | 2 |
| `packages/centymo-golang/views/revenue/list/page.go` | `centymo.Sales{Labels,Routes}` → `centymo.Revenue{Labels,Routes}` | 2 |
| `packages/centymo-golang/views/revenue/detail/page.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/detail/line_items.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/action/action.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/action/payment.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/action/invoice_download.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/dashboard/page.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/settings/deps.go` | Same | 2 |
| `packages/centymo-golang/views/revenue/settings/page.go` | Same | 2 |
| `apps/retail-admin/internal/presentation/sales/module.go` | **Move to** `revenue/module.go`, `package revenue`, type refs | 3 |
| `apps/service-admin/internal/presentation/sales/module.go` | **Move to** `revenue/module.go`, `package revenue`, type refs | 3 |
| `apps/retail-admin/internal/composition/views.go` | Import + alias `salesmod` → `revenuemod`, field names | 3 |
| `apps/service-admin/internal/composition/views.go` | Import + alias `salesmod` → `revenuemod`, field names | 3 |
| `apps/retail-admin/internal/composition/container.go` | `"sale.json"/"sales"` → `"revenue.json"/"revenue"`, type + var names | 3 |
| `apps/service-admin/internal/composition/container.go` | Same + route.json path update | 3 |
| `apps/service-admin/internal/presentation/home/module.go` | `SalesRoutes` → `RevenueRoutes` field | 3 |

---

## Context & Sub-Agent Strategy

**Estimated files to read:** 24
**Estimated files to modify:** 24
**Estimated context usage:** Medium (24 files, but changes are mechanical find-replace)

**Sub-agent plan:**
- Phase 2 (centymo type renames) can run as a single agent — all changes are in one package
- Phase 3 (app consumers) can run in parallel for retail-admin and service-admin
- Or: single session with systematic replace_all edits per file

---

## Risk & Dependencies

| Risk | Impact | Mitigation |
|------|--------|------------|
| Missed reference causes build failure | Medium — caught at compile time | Build both apps after Phase 2 and Phase 3 |
| `SalesSummaryURL` referenced in views.go | Low — only 1 ref in service-admin | Rename to `RevenueSummaryURL` |
| `TransactionRoutes.Sales` field rename | Medium — used in sidebar config | Rename to `TransactionRoutes.Revenue` in both apps |
| Air hot-reload won't pick up dir rename | Low — just restart server | Manual restart after Phase 3 |

**Dependencies:**
- Phase 2 depends on Phase 1 (JSON files must exist)
- Phase 3 depends on Phase 2 (Go types must be renamed first)
- Phase 4 depends on all prior phases

---

## Acceptance Criteria

- [ ] `sale.json` no longer exists in any tier
- [ ] `revenue.json` exists in all 3 tiers with root key `"revenue"`
- [ ] route.json files use `"revenue"` key
- [ ] No `SalesLabels`, `SalesRoutes`, `DefaultSalesRoutes`, `Sales*URL` in codebase
- [ ] No `presentation/sales/` directories remain
- [ ] service-admin builds and runs without errors
- [ ] retail-admin builds without errors
- [ ] Display text unchanged: retail="Sales", service="Bookings"
- [ ] service-admin E2E tests pass (61 tests)

---

## Design Decisions

**Why rename Go types now instead of deferring?**
The types are named `Sales*` but live in `views/revenue/` directories, receive `revenuepb` protos, and are loaded from `revenue.json`. Leaving them as `Sales*` would make the mismatch worse after the translation rename. Doing it now keeps the codebase consistent.

**Why not rename the route URL paths (e.g., `/app/sales/` → `/app/revenue/`)?**
URL paths are user-facing and bookmarkable. Retail should keep `/app/sales/` as its natural URL. Service already overrides to `/app/bookings/` via route.json. The URL paths are already correct per business type; only the internal naming was wrong.

**Why rename presentation dirs from `sales/` to `revenue/`?**
The Go package name becomes the import identity. Having `package sales` when all the types are `centymo.Revenue*` is confusing. The centymo views already live under `views/revenue/` — the app-level modules should match.
