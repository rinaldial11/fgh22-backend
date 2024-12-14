package controllers

type Response struct {
	Succsess bool   `json:"success"`
	Message  string `json:"message"`
	Results  any    `json:"results,omitempty"`
}

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ListUsers []User

var Users = ListUsers{
	{
		Id:       1,
		Fullname: "Budiono Siregar",
		Email:    "budi@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$xYjCiAsMbo/xmH8I/4SkeQ$MvbfLtQCqyryc+p9ghXxwEniSGDZLMF2ckH7+Hzxqq0",
	},
	{
		Id:       2,
		Fullname: "Endra Prasmanan",
		Email:    "endra@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$NRaP9cPxjmhEqCvwF5+4ig$hY+X9oq8wlA3aWKEiwx0z/THUkypKplR4C0j+jv2qtA",
	},
	{
		Id:       3,
		Fullname: "Rama Ajarindong",
		Email:    "rama@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$ksC1dsVheRE4cfO9fR8odw$+hStsdH+e9w68Zxn30RMKWLjEcIFXVhhQ6EIGfbKeoU",
	},
	{
		Id:       4,
		Fullname: "Adiv Bened",
		Email:    "adiv@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$vRI56QMSk9gBXGnbLJ3XiQ$O7hU2yZt7zSTokAueXwcQAEX7lZs6ufAEyddfjc12vk",
	},
	{
		Id:       5,
		Fullname: "Nanda Brew",
		Email:    "nanda@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$gCjUU9eQbquuRJcNEoQm9g$bMrkwmI9O4bA23xLEqVEYX26uc1s0k/pKyrPkaC6J1c",
	},
	{
		Id:       6,
		Fullname: "Alice Johnson",
		Email:    "alice.johnson@example.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54",
	},
	{
		Id:       7,
		Fullname: "Bob Smith",
		Email:    "bob.smith@example.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54",
	},
	{Id: 8, Fullname: "Charlie Brown", Email: "charlie.brown@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 9, Fullname: "David Lee", Email: "david.lee@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 10, Fullname: "Eva Green", Email: "eva.green@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 11, Fullname: "Frank White", Email: "frank.white@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 12, Fullname: "Grace Blue", Email: "grace.blue@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 13, Fullname: "Hannah Scott", Email: "hannah.scott@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 14, Fullname: "Ian Turner", Email: "ian.turner@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 15, Fullname: "Jack Robinson", Email: "jack.robinson@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 16, Fullname: "Karen Harris", Email: "karen.harris@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 17, Fullname: "Liam Clark", Email: "liam.clark@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 18, Fullname: "Mona Lopez", Email: "mona.lopez@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 19, Fullname: "Nathan Allen", Email: "nathan.allen@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 20, Fullname: "Olivia King", Email: "olivia.king@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 21, Fullname: "Paul Adams", Email: "paul.adams@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 22, Fullname: "Quincy Moore", Email: "quincy.moore@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 23, Fullname: "Rachel Young", Email: "rachel.young@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 24, Fullname: "Samuel Nelson", Email: "samuel.nelson@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 25, Fullname: "Tina Adams", Email: "tina.adams@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
	{Id: 26, Fullname: "Ursula Price", Email: "ursula.price@example.com", Password: "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54"},
}
