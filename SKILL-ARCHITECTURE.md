# Go-StudentApp Architecture Skills

## Domain model
- School
  - `school/model/models.go`
  - Fields: `ID`, `Name`
- Student
  - `student/model/models.go`
  - Fields: `ID`, `Name`, `Class`, `SchoolID`
- Course
  - `course/model/models.go`
  - Fields: `ID`, `Title`, `StudentID`
  - Note: current schema enforces `StudentID` uniqueness.

## Request -> Service -> Repo Data Path

### Student create flow (example)
1. HTTP route hits student controller
2. Controller validates input and calls service
3. Service checks school existence + student validation
4. Repository persists entity via GORM

### Course create flow
1. Route receives nested school + student path params
2. Controller passes path/body to service
3. Service validates provided IDs exist
4. Repository creates mapped `Course` entity

## Responsibility split (per module)
- `controller`:
  - Parse request payload and path params
  - Build and return structured responses
- `service`:
  - Validate relationships (`school_id`, `student_id` where applicable)
  - Convert request DTO -> model mapping
  - Handle business errors
- `repository`:
  - Data access methods (`Save`, `FindById`, `FindBy...`, `Update`, `Delete`)

## Cross-module notes
- Student routes are nested under school context
- Course routes are nested under school + student context
- Existing route compatibility has been preserved by keeping aliases (e.g. student create supports both singular and plural endpoints)
