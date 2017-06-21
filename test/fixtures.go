package test

const (
	User1 = "testuser1"
	Pass1 = "testpass1"
	User2 = "testuser2"
	Pass2 = "testpass2"
	User3 = "testuser3"
	Pass3 = "testpass3"
	// BookTitle
	// BookIsbn
	// BookImage
	// BookUser
)

var testusers = []struct {
	Username string
	Password string
}{
	{User1, Pass1},
	{User2, Pass2},
	{User3, Pass3},
}

var testbooks = []struct {
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	ImageUrl string `json:"image_url"`
	// UserId   int64  `json:"user_id,omitempty"`
}{
	{"1Q84", "9786074213720", "https://images-na.ssl-images-amazon.com/images/I/41aXLI626BL._SX323_BO1,204,203,200_.jpg"},
	{"Artemis Fowl", "9780786848690", "http://t2.gstatic.com/images?q=tbn:ANd9GcTnSUKt1gyCI_X4amzC4lHk__we3zYxtDkjWPK73RxQa9m06-s0"},
}
