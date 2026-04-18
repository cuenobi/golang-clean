# Error Code Catalog

This catalog is FE-facing and must remain backward compatible once public.

## Common Codes (1000-1499)
| Code | Type | Meaning | Typical HTTP |
|---|---|---|---|
| 1000 | `internal_error` | Unhandled server failure | 500 |
| 1001 | `validation_error` | Request validation failed | 400 |
| 1002 | `bad_request` | Malformed or invalid request | 400 |
| 1401 | `unauthorized` | Missing/invalid API key | 401 |
| 1403 | `forbidden` | Permission denied | 403 |
| 1404 | `not_found` | Resource not found | 404 |
| 1409 | `conflict` | Resource state conflict | 409 |
| 1422 | `invalid_state` | Domain state transition invalid | 422 |
| 1429 | `rate_limited` | Request quota exceeded | 429 |

## User Context (2000-2999)
| Code | Type | Meaning | Typical HTTP |
|---|---|---|---|
| 2001 | `user_email_already_exists` | Duplicate email address | 409 |
| 2002 | `user_name_invalid` | Name violates business rule | 422 |
| 2003 | `user_not_active` | User cannot perform action | 422 |

## Order Context (3000-3999)
| Code | Type | Meaning | Typical HTTP |
|---|---|---|---|
| 3001 | `order_amount_invalid` | Amount/currency business rule violation | 422 |
| 3002 | `order_status_transition_invalid` | Invalid order status change | 422 |
| 3003 | `order_idempotency_conflict` | Same idempotency key with incompatible payload | 409 |

## Governance Rules
1. Never reuse an old numeric code with a different meaning.
2. Add new code before releasing related endpoint behavior.
3. Update OpenAPI examples and FE mapping table in same PR.
