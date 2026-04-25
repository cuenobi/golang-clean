---
name: unit-test-mockery-suite
description: Generate and maintain Go use case unit tests with mockery-generated repository mocks and testify suite style. Use when adding new unit tests, refactoring hand-written fakes to mockery, or standardizing tests to suite-based structure.
---

# Skill: Unit Test Mockery Suite

## Goal
Create deterministic unit tests for application use cases with `mockery` mocks and `testify/suite`.

## Standards
- Unit under test lives in `internal/application/usecase/...`.
- Mock dependencies from `internal/application/port/out`.
- Use `suite.Suite` and `suite.Run`.
- Keep one suite per use case test file.
- Cover happy path, failure path, and edge/validation path.
- Assert expectations in `TearDownTest`.

## Mock Generation
1. Find target interface in `internal/application/port/out`.
2. Generate or refresh only required mocks (avoid full repo regen unless explicitly asked):

```bash
go run github.com/vektra/mockery/v2@v2.53.6 \
  --name <InterfaceName> \
  --dir internal/application/port/out \
  --output internal/application/port/out/mocks \
  --outpkg mocks \
  --with-expecter
```

3. Reuse existing generated mocks when available.

## Suite Template
```go
type CreateUserSuite struct {
	suite.Suite
	repo *mocks.UserRepository
	uc   *usecase.UserUseCase
}

func TestCreateUserSuite(t *testing.T) {
	suite.Run(t, new(CreateUserSuite))
}

func (s *CreateUserSuite) SetupTest() {
	s.repo = mocks.NewUserRepository(s.T())
	s.uc = usecase.NewUserUseCase(s.repo, fixedClock, fixedID)
}

func (s *CreateUserSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}
```

## Test Method Style
- Name test methods `Test_<Behavior>`.
- Arrange with explicit expectations (`On(...).Return(...)`).
- Act once per test method.
- Assert result and error contract.
- Keep data deterministic: fixed IDs, fixed UTC timestamps, no network/time randomness.

## Validation
- Run focused tests first:

```bash
go test ./internal/application/usecase/<module> -run <SuiteOrCaseName>
```

- Expand to broader scope only if needed.
