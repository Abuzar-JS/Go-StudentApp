# Go-StudentApp API Skills

## Base URL
- `GET/POST/PUT/PATCH/DELETE` APIs are under `/api/v1`

## School Routes
- `GET /api/v1/schools`
- `GET /api/v1/schools/:school_id`
- `POST /api/v1/school`
- `PUT /api/v1/schools/:school_id`
- `PATCH /api/v1/schools/:school_id`
- `DELETE /api/v1/schools/:school_id`

## Student Routes
- `GET /api/v1/schools/:school_id/students`
- `GET /api/v1/schools/:school_id/students/:student_id`
- `POST /api/v1/schools/:school_id/student`
- `POST /api/v1/schools/:school_id/students` *(alias kept for compatibility)*
- `PUT /api/v1/schools/:school_id/students/:student_id`
- `DELETE /api/v1/schools/:school_id/students/:student_id`

## Course Routes
- `GET /api/v1/schools/:school_id/students/:student_id/courses`
- `GET /api/v1/schools/:school_id/students/:student_id/courses/:course_id`
- `POST /api/v1/schools/:school_id/students/:student_id/course`
- `PUT /api/v1/schools/:school_id/students/:student_id/courses/:course_id`
- `DELETE /api/v1/schools/:school_id/students/:student_id/courses/:course_id`

## Request DTOs (current)
- Student create/update expects `name`, `class`, and `school_id` in payloads where required.
- Course create currently expects `title` and `student_id` only.

## Response payload shape patterns
- Responses are simple resource-shaped JSONs defined under each module’s `controller/response` folder.
- Most responses include resource IDs and primary fields.
