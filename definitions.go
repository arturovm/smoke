package main

type RESTError struct {
	code int
	Message string `json:"message"`
}

type Session struct {
	SessionId string // `json:"session_id"`
	SessionKey string // `json:"session_key"`
}

type Service struct {
	Type string // `json:"type"`
	Name string // `json:"name"`
}

type User struct {
	Name string // `json:"name"`
	Email string // `json:"email"`
	Password string // `json:"password"`
	Services []Service // `json:"services"`
}

// Here we face a dilemma: If returning user details to end-user needed, how to protect password?
// This problem is due to rethinkgo using the json package to store and retrieve data; if Password is unexported (to avoid exposure), it's also not saved.
// TEMPORARY SOLUTION: To create a new type that's almost identical to User, but omits Password
// IDEA: To implement MarshalJSON() with a flag on password to decide if should marshal or not.
// NEW SOLUTION: TO MARSHAL A JSON MAP WITH JUST THE REQUIRED FIELDS. SHOULD BECOME STANDARD PRACTICE?

/*type UserDetails struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Services []Service `json:"services"`
}*/
